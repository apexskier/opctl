package main

import (
	"context"
	"fmt"
	"github.com/opctl/opctl/sdks/go/node"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/cli/internal/cliparamsatisfier"
	"github.com/opctl/opctl/cli/internal/opgraph"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
	"github.com/pkg/errors"
)

// RunOpts are options to run a given op through the CLI
type RunOpts struct {
	ArgFile    string
	Args       []string
	OpRef      string
	NoProgress bool
}

type runResults struct {
	outputs map[string]*model.Value
	err     error
}

func run(
	ctx context.Context,
	cliOutput clioutput.CliOutput,
	cliParamSatisfier cliparamsatisfier.CLIParamSatisfier,
	opFormatter clioutput.OpFormatter,
	node node.Node,
	opts *RunOpts,
) (map[string]*model.Value, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	aSigIntWasReceivedAlready := false
	sigIntChannel := make(chan os.Signal, 1)
	defer close(sigIntChannel)
	signal.Notify(
		sigIntChannel,
		syscall.SIGINT,
	)

	sigTermChannel := make(chan os.Signal, 1)
	defer close(sigTermChannel)
	signal.Notify(
		sigTermChannel,
		syscall.SIGTERM,
	)

	sigInfoChannel := make(chan os.Signal, 1)
	defer close(sigInfoChannel)
	signal.Notify(
		sigInfoChannel,
		syscall.Signal(0x1d), // portable version of syscall.SIGINFO
	)

	// "request animation frame" like loop to force refresh of display loading spinners
	animationFrame := make(chan bool)
	if !opts.NoProgress {
		go func() {
			for {
				time.Sleep(time.Second / 10)
				animationFrame <- true
			}
		}()
	}

	var state opgraph.CallGraph
	var loadingSpinner opgraph.DotLoadingSpinner
	output := opgraph.NewOutputManager()

	defer func() {
		_ = output.Print(state.String(loadingSpinner, opFormatter, time.Now(), false))
		fmt.Println()
	}()

	clearGraph := func() {
		if !opts.NoProgress {
			_ = output.Clear()
		}
	}

	displayGraph := func() {
		if !opts.NoProgress {
			_ = output.Print(state.String(loadingSpinner, opFormatter, time.Now(), true))
		}
	}

	opHandle, err := node.Resolve(
		ctx,
		opts.OpRef,
	)
	if err != nil {
		return nil, err
	}

	opFileReader, err := opHandle.GetContent(
		ctx,
		opfile.FileName,
	)
	if err != nil {
		return nil, err
	}

	opFileBytes, err := ioutil.ReadAll(opFileReader)
	if nil != err {
		return nil, err
	}

	opFile, err := opfile.Unmarshal(
		filepath.Join(opHandle.Ref(), opfile.FileName),
		opFileBytes,
	)
	if err != nil {
		return nil, err
	}

	ymlFileInputSrc, err := cliParamSatisfier.NewYMLFileInputSrc(opts.ArgFile)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to load arg file at '%v'", opts.ArgFile))
	}

	cliPromptInputSrc := cliParamSatisfier.NewCliPromptInputSrc(opFile.Inputs)
	if err != nil {
		return nil, err
	}
	argsMap, err := cliParamSatisfier.Satisfy(
		cliparamsatisfier.NewInputSourcer(
			cliParamSatisfier.NewSliceInputSrc(opts.Args, "="),
			ymlFileInputSrc,
			cliParamSatisfier.NewEnvVarInputSrc(),
			cliParamSatisfier.NewParamDefaultInputSrc(opFile.Inputs),
			cliPromptInputSrc,
		),
		opFile.Inputs,
	)
	if err != nil {
		return nil, err
	}

	eventChannel := make(chan model.Event)

	// listen for op end on a channel, to not block output streaming
	done := make(chan runResults, 1)
	go func() {
		outputs, err := node.StartOp(
			ctx,
			eventChannel,
			model.StartOpReq{
				Args: argsMap,
				Op: model.StartOpReqOp{
					Ref: opHandle.Ref(),
				},
			},
		)
		done <- runResults{
			outputs: outputs,
			err:     err,
		}
	}()

	for {
		select {
		case <-sigIntChannel:
			clearGraph()
			if !aSigIntWasReceivedAlready {
				cliOutput.Warning("Gracefully stopping... (signal Control-C again to force)")
				aSigIntWasReceivedAlready = true
				// events will continue to stream in, make sure we continue to display the graph while this happens
				displayGraph()
				cancel()
			} else {
				return nil, &RunError{
					ExitCode: 130,
					message:  "Terminated by Control-C",
				}
			}

		case <-sigInfoChannel:
			clearGraph()
			// clear two more lines
			fmt.Print("\033[1A\033[K\033[1A\033[K")
			fmt.Println(state.String(opgraph.StaticLoadingSpinner{}, opFormatter, time.Now(), false))
			displayGraph()

		case <-sigTermChannel:
			clearGraph()
			cliOutput.Error("Gracefully stopping...")
			displayGraph()
			cancel()

		case results := <-done:
			clearGraph()
			if results.err != nil && !errors.Is(results.err, context.Canceled) {
				return nil, results.err
			}
			return results.outputs, nil

		case event, isEventChannelOpen := <-eventChannel:
			clearGraph()
			if !isEventChannelOpen {
				return nil, errors.New("event channel closed unexpectedly")
			}

			if err := state.HandleEvent(&event); err != nil {
				cliOutput.Error(fmt.Sprintf("%v", err))
			}

			cliOutput.Event(&event)
			displayGraph()
		case <-animationFrame:
			clearGraph()
			displayGraph()
		}
	}
}
