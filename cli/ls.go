package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/opctl/opctl/cli/internal/cliparamsatisfier"
	"github.com/opctl/opctl/cli/internal/dataresolver"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/core"
	"github.com/opctl/opctl/sdks/go/opspec"
)

// ls implements "ls" command
func ls(
	ctx context.Context,
	cliParamSatisfier cliparamsatisfier.CLIParamSatisfier,
	node core.Core,
	dirRef string,
) error {
	dataResolver := dataresolver.New(
		cliParamSatisfier,
		node,
	)

	tw := new(tabwriter.Writer)
	defer tw.Flush()
	tw.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(tw, "REF\tDESCRIPTION")

	eventChannel := make(chan model.Event)
	callID := ""

	dirHandle, err := dataResolver.Resolve(
		ctx,
		eventChannel,
		callID,
		dirRef,
	)
	if err != nil {
		return err
	}

	opsByPath, err := opspec.List(
		ctx,
		eventChannel,
		callID,
		dirHandle,
	)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	type SortableOp struct {
		path string
		ref  string
	}

	// ensure runnable ops are
	sortableOps := make([]SortableOp, 0, len(opsByPath))
	for path := range opsByPath {
		ref := filepath.Join(dirHandle.Ref(), path)
		if filepath.IsAbs(ref) {
			// make absolute paths relative
			relOpRef, err := filepath.Rel(cwd, ref)
			if err != nil {
				return err
			}
			ref = strings.TrimPrefix(relOpRef, ".opspec/")
		}
		sortableOps = append(sortableOps, SortableOp{
			path: path,
			ref:  ref,
		})
	}

	sort.Slice(sortableOps, func(i, j int) bool {
		return strings.Compare(sortableOps[i].ref, sortableOps[j].ref) < 0
	})

	for _, r := range sortableOps {
		opDescription := opsByPath[r.path].Description
		scanner := bufio.NewScanner(strings.NewReader(opDescription))
		scanner.Scan()
		// first line of description, add the op ref
		fmt.Fprintf(tw, "%v\t%v", r.ref, scanner.Text())
		for scanner.Scan() {
			// subsequent lines, don't add the op ref but let the description span multiple lines
			fmt.Fprintf(tw, "\n\t%v", scanner.Text())
		}
		fmt.Fprintln(tw)
	}

	return nil
}
