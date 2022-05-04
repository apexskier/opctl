package data

import (
	"context"
	"fmt"

	aggregateError "github.com/opctl/opctl/sdks/go/internal/aggregate_error"
)

// Resolve "dataRef" from "providers" in order
func Resolve(
	ctx context.Context,
	dataRef string,
	providers ...DataProvider,
) (
	DataHandle,
	error,
) {
	var agg aggregateError.ErrAggregate

	for _, src := range providers {
		handle, err := src.Resolve(ctx, dataRef)
		if err != nil {
			agg.AddError(fmt.Errorf("%s: %w", src.Label(), err))
		} else if handle != nil {
			return handle, nil
		}
	}

	return nil, fmt.Errorf("unable to resolve op '%s': %w", dataRef, agg)
}
