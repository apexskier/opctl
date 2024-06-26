package node

import (
	"context"

	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
	"github.com/opctl/opctl/sdks/go/model"
)

//counterfeiter:generate -o internal/fakes/serialCaller.go . serialCaller
type serialCaller interface {
	// Executes a serial call
	Call(
		ctx context.Context,
		eventChannel chan model.Event,
		callID string,
		inboundScope map[string]*model.Value,
		rootCallID string,
		opPath string,
		callSpecSerialCall []*model.CallSpec,
		scratchPath string,
	) (
		map[string]*model.Value,
		error,
	)
}

func newSerialCaller(caller caller) serialCaller {
	return _serialCaller{
		caller: caller,
	}
}

type _serialCaller struct {
	caller caller
}

func (sc _serialCaller) Call(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	inboundScope map[string]*model.Value,
	rootCallID string,
	opPath string,
	callSpecSerialCall []*model.CallSpec,
	scratchPath string,
) (
	map[string]*model.Value,
	error,
) {
	scope := inboundScope

	for _, callSpecCall := range callSpecSerialCall {
		var childCallID string
		childCallID, err := uniquestring.Construct()
		if err != nil {
			// end run immediately on any error
			return nil, err
		}

		outputs, err := sc.caller.Call(
			ctx,
			eventChannel,
			childCallID,
			scope,
			callSpecCall,
			opPath,
			&callID,
			rootCallID,
			scratchPath,
		)
		if err != nil {
			return nil, err
		}

		// check to see if this has been cancelled
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		for name, value := range outputs {
			scope[name] = value
		}
	}

	return scope, nil
}
