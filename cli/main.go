package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/opctl/opctl/cli/internal/clicolorer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := newCli(ctx)
	if err != nil {
		clicolorer.New().Error(fmt.Sprintf("failed to start up: %v", err.Error()))
		os.Exit(1)
	}

	defer func() {
		if panic := recover(); panic != nil {
			clicolorer.New().Error(fmt.Sprintf("panic: %v", panic))
			fmt.Printf("%s\n%s\n", panic, debug.Stack())
		}
	}()

	_ = cli.Run(os.Args)
}
