package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/opctl/opctl/cli/internal/clioutput"
	"github.com/opctl/opctl/cli/internal/dataresolver"
	"github.com/opctl/opctl/sdks/go/opspec"
)

// ls implements "ls" command
func ls(
	ctx context.Context,
	opFormatter clioutput.OpFormatter,
	dataResolver dataresolver.DataResolver,
	dirRef string,
) error {
	tw := new(tabwriter.Writer)
	defer tw.Flush()
	tw.Init(os.Stdout, 0, 8, 1, '\t', 0)

	fmt.Fprintln(tw, "REF\tDESCRIPTION")

	dirHandle, err := dataResolver.Resolve(
		ctx,
		dirRef,
	)
	if err != nil {
		return err
	}

	opsByRef, erroringOpsByRef, err := opspec.List(
		ctx,
		dirHandle,
	)
	if err != nil {
		return err
	}

	type SortableRef struct {
		ref          string
		formattedRef string
	}

	// ensure runnable ops are sorted, otherwise order will change
	sortableOps := make([]SortableRef, 0, len(opsByRef))
	for ref := range opsByRef {
		sortableOps = append(sortableOps, SortableRef{
			ref:          ref,
			formattedRef: opFormatter.FormatOpRef(ref),
		})
	}
	sort.Slice(sortableOps, func(i, j int) bool {
		return strings.Compare(sortableOps[i].ref, sortableOps[j].ref) < 0
	})
	for _, op := range sortableOps {
		opDescription := opsByRef[op.ref].Description
		scanner := bufio.NewScanner(strings.NewReader(opDescription))
		scanner.Scan()
		// first line of description, add the op ref
		fmt.Fprintf(tw, "%v\t%v", op.formattedRef, scanner.Text())
		for scanner.Scan() {
			// subsequent lines, don't add the op ref but let the description span multiple lines
			fmt.Fprintf(tw, "\n\t%v", scanner.Text())
		}
		fmt.Fprintln(tw)
	}

	if len(erroringOpsByRef) > 0 {
		sortableErroringOpPaths := make([]SortableRef, 0, len(erroringOpsByRef))
		for ref := range erroringOpsByRef {
			sortableErroringOpPaths = append(sortableErroringOpPaths, SortableRef{
				ref:          ref,
				formattedRef: opFormatter.FormatOpRef(ref),
			})
		}
		sort.Slice(sortableErroringOpPaths, func(i, j int) bool {
			return strings.Compare(sortableErroringOpPaths[i].ref, sortableErroringOpPaths[j].ref) < 0
		})
		fmt.Fprintln(tw)
		fmt.Fprintln(tw, "INVALID OPS")
		for _, op := range sortableErroringOpPaths {
			fmt.Fprintln(tw, op.formattedRef)
		}
	}

	return nil
}
