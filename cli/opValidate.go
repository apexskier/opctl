package main

import (
	"context"

	"github.com/opctl/opctl/cli/internal/dataresolver"
	"github.com/opctl/opctl/sdks/go/opspec"
)

func opValidate(
	ctx context.Context,
	dataResolver dataresolver.DataResolver,
	opRef string,
) error {
	opDirHandle, err := dataResolver.Resolve(
		ctx,
		opRef,
	)
	if err != nil {
		return err
	}

	return opspec.Validate(*opDirHandle.Path())
}
