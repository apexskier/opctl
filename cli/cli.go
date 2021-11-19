package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/appdataspec/sdk-golang/appdatapath"
	mow "github.com/jawher/mow.cli"
	"github.com/opctl/opctl/cli/internal/clicolorer"
	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/cli/internal/cliparamsatisfier"
	"github.com/opctl/opctl/cli/internal/dataresolver"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/core"
	"github.com/opctl/opctl/sdks/go/node/core/containerruntime"
	"github.com/opctl/opctl/sdks/go/node/core/containerruntime/docker"
	"github.com/opctl/opctl/sdks/go/node/core/containerruntime/k8s"
	"github.com/opctl/opctl/sdks/go/opspec"
	"golang.org/x/term"
)

var testModeEnvVar = "OPCTL_TEST_MODE"

type cli interface {
	Run(args []string) error
}

func newCli(
	ctx context.Context,
) (cli, error) {
	cli := mow.App(
		"opctl",
		"Opctl is a free and open source distributed operation control system.",
	)
	cli.Version("v version", version)

	perUserAppDataPath, err := appdatapath.New().PerUser()
	if nil != err {
		return nil, err
	}

	datadirPath := cli.String(
		mow.StringOpt{
			Desc:   "Path of dir used to store opctl data",
			EnvVar: "OPCTL_DATA_DIR",
			Name:   "data-dir",
			Value:  filepath.Join(perUserAppDataPath, "miniopctl"),
		},
	)

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	opFormatter := clioutput.NewCliOpFormatter(cwd, *datadirPath)

	cliOutput, err := clioutput.New(clicolorer.New(), opFormatter, os.Stderr, os.Stdout)
	if err != nil {
		return nil, err
	}

	exitWith := func(successMessage string, err error) {
		if err != nil {
			msg := err.Error()
			if msg != "" {
				cliOutput.Error(msg)
			}
			if re, ok := err.(*RunError); ok {
				mow.Exit(re.ExitCode)
			} else {
				mow.Exit(1)
			}
		}

		if successMessage != "" {
			cliOutput.Success(successMessage)
		}
		// Don't exit here with .Exit to allow container cleanup to happen
	}

	cliParamSatisfier := cliparamsatisfier.New(cliOutput)

	containerRuntime := cli.String(
		mow.StringOpt{
			Desc:   "Runtime for opctl containers",
			EnvVar: "OPCTL_CONTAINER_RUNTIME",
			Name:   "container-runtime",
			Value:  "docker",
		},
	)

	noColor := cli.BoolOpt("nc no-color", false, "Disable output coloring")

	cli.Before = func() {
		if *noColor {
			cliOutput.DisableColor()
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	cli.After = func() {
		cancel()
	}

	var cr containerruntime.ContainerRuntime
	if "k8s" == *containerRuntime {
		cr, err = k8s.New()
	} else {
		dockerConfigPath := cli.String(
			mow.StringOpt{
				Desc:   "Docker configuration file path, if using the docker container runtime",
				EnvVar: "OPCTL_DOCKER_CONFIG",
				Name:   "docker-config",
			},
		)
		cr, err = docker.New(ctx, *dockerConfigPath)
	}
	if nil != err {
		return nil, err
	}

	eventChannel := make(chan model.Event)

	opNode, err := core.New(ctx, cr, *datadirPath)
	if err != nil {
		return nil, err
	}

	cli.Command("ls", "List operations", func(lsCmd *mow.Cmd) {
		const dirRefArgName = "DIR_REF"
		lsCmd.Spec = fmt.Sprintf("[%v]", dirRefArgName)
		dirRef := lsCmd.StringArg(
			dirRefArgName,
			cwd,
			"Reference to dir ops will be listed from",
		)

		opFormatter := clioutput.NewCliOpFormatter(*dirRef, *datadirPath)

		lsCmd.Action = func() {
			exitWith("", ls(ctx, opFormatter, cliParamSatisfier, opNode, *dirRef))
		}
	})

	cli.Command("op", "Manage ops", func(opCmd *mow.Cmd) {
		dataResolver := dataresolver.New(
			cliParamSatisfier,
			opNode,
		)

		opCmd.Command("install", "Install an op", func(installCmd *mow.Cmd) {
			path := installCmd.StringOpt("path", opspec.DotOpspecDirName, "Path the op will be installed at")
			opRef := installCmd.StringArg("OP_REF", "", "Op reference (either `relative/path`, `/absolute/path`, `host/path/repo#tag`, or `host/path/repo#tag/path`)")

			installCmd.Action = func() {
				exitWith(
					"",
					opInstall(
						ctx,
						dataResolver,
						*opRef,
						*path,
					),
				)
			}
		})

		opCmd.Command("validate", "Validate an op", func(validateCmd *mow.Cmd) {
			opRef := validateCmd.StringArg("OP_REF", "", "Op reference (either `relative/path`, `/absolute/path`, `host/path/repo#tag`, or `host/path/repo#tag/path`)")

			validateCmd.Action = func() {
				exitWith(
					fmt.Sprintf("%v is valid", *opRef),
					opValidate(
						ctx,
						dataResolver,
						*opRef,
					),
				)
			}
		})
	})

	cli.Command("run", "Start and wait on an op", func(runCmd *mow.Cmd) {
		args := runCmd.StringsOpt("a", []string{}, "Explicitly pass args to op in format `-a NAME1=VALUE1 -a NAME2=VALUE2`")
		argFile := runCmd.StringOpt("arg-file", filepath.Join(opspec.DotOpspecDirName, "args.yml"), "Read in a file of args in yml format")
		noProgress := runCmd.BoolOpt("no-progress", !term.IsTerminal(int(os.Stdout.Fd())), "Disable live call graph for the op")
		opRef := runCmd.StringArg("OP_REF", "", "Op reference (either `relative/path`, `/absolute/path`, `host/path/repo#tag`, or `host/path/repo#tag/path`)")

		runCmd.Action = func() {
			outputs, err := run(
				ctx,
				cliOutput,
				cliParamSatisfier,
				opFormatter,
				eventChannel,
				opNode,
				*opRef,
				&RunOpts{Args: *args, ArgFile: *argFile},
				*noProgress,
			)
			if err != nil {
				exitWith("", err)
			} else if len(outputs) > 0 {
				exitWith(model.FormatValueMap(outputs))
			} else {
				exitWith("", nil)
			}
		}
	})

	cli.Command("self-update", "Update opctl", func(selfUpdateCmd *mow.Cmd) {
		selfUpdateCmd.Action = func() {
			exitWith(selfUpdate())
		}
	})

	return cli, nil
}
