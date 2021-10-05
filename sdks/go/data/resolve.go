package data

import (
	"context"
	"fmt"

	aggregateError "github.com/opctl/opctl/sdks/go/internal/aggregate_error"
	"github.com/opctl/opctl/sdks/go/model"
)

// Resolve "dataRef" from "providers" in order
//
// expected errs:
//  - ErrDataProviderAuthentication on authentication failure
//  - ErrDataProviderAuthorization on authorization failure
func Resolve(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	dataRef string,
	providers ...model.DataProvider,
) (
	model.DataHandle,
	error,
) {
	var agg aggregateError.ErrAggregate

	for _, src := range providers {
		handle, err := src.TryResolve(ctx, eventChannel, callID, dataRef)
		if err != nil {
			agg.AddError(fmt.Errorf("%s: %w", src.Label(), err))
		} else if handle != nil {
			return handle, nil
		}
	}

	return nil, fmt.Errorf("unable to resolve op '%s': %w", dataRef, agg)
}
