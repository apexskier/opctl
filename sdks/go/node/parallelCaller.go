package node

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec"
)

//counterfeiter:generate -o internal/fakes/parallelCaller.go . parallelCaller
type parallelCaller interface {
	// Executes a parallel call
	Call(
		parentCtx context.Context,
		eventChannel chan model.Event,
		callID string,
		inboundScope map[string]*model.Value,
		rootCallID string,
		opPath string,
		callSpecParallelCall []*model.CallSpec,
		scratchPath string,
	) (
		map[string]*model.Value,
		error,
	)
}

func newParallelCaller(caller caller) parallelCaller {
	return _parallelCaller{
		caller: caller,
	}
}

type _parallelCaller struct {
	caller caller
}

func (pc _parallelCaller) Call(
	parentCtx context.Context,
	eventChannel chan model.Event,
	callID string,
	inboundScope map[string]*model.Value,
	rootCallID string,
	opPath string,
	callSpecParallelCall []*model.CallSpec,
	scratchPath string,
) (
	map[string]*model.Value,
	error,
) {
	// setup cancellation
	parallelCtx, cancelParallel := context.WithCancel(parentCtx)
	defer cancelParallel()

	childCallNeededCountByName := map[string]int{}
	for _, childCall := range callSpecParallelCall {
		// increment needed by counts for any needs
		for _, neededCallRef := range childCall.Needs {
			childCallNeededCountByName[opspec.RefToName(neededCallRef)]++
		}
	}

	childCallIndexByID := map[string]int{}
	childCallIndexByName := map[string]int{}
	childCallCancellationByIndex := map[int]context.CancelFunc{}
	childCallOutputsByIndex := map[int]map[string]*model.Value{}

	type childResult struct {
		CallID  string
		Err     error
		Outputs map[string]*model.Value
	}
	childResults := make(chan childResult, len(callSpecParallelCall))

	// This waitgroup ensures all child goroutines are allowed to clean up
	var wg sync.WaitGroup
	defer wg.Wait()

	// perform calls in parallel w/ cancellation
	for childCallIndex, childCall := range callSpecParallelCall {
		childCallID, err := uniquestring.Construct()
		if err != nil {
			// end run immediately on any error
			return nil, err
		}

		childCallIndexByID[childCallID] = childCallIndex

		var childCtx context.Context
		if childCall.Name != nil {
			childCallIndexByName[*childCall.Name] = childCallIndex
			childCtx, childCallCancellationByIndex[childCallIndex] = context.WithCancel(parallelCtx)
		} else {
			childCtx = parallelCtx
		}

		wg.Add(1)
		go func(childCtx context.Context, childCallID string, childCall *model.CallSpec) {
			defer wg.Done()
			outputs, err := pc.caller.Call(
				childCtx,
				eventChannel,
				childCallID,
				inboundScope,
				childCall,
				opPath,
				&callID,
				rootCallID,
				scratchPath,
			)
			if childCtx.Err() != nil {
				// context has been cancelled, so skip reporting results
				return
			}
			childResults <- childResult{
				CallID:  childCallID,
				Err:     err,
				Outputs: outputs,
			}
		}(childCtx, childCallID, childCall)
	}

	outboundScope := inboundScope

	for {
		select {
		case <-parallelCtx.Done():
			return nil, parallelCtx.Err()

		case result := <-childResults:
			if result.Err != nil {
				// cancel all children on any error
				cancelParallel()
				close(childResults)
				return nil, result.Err
			}

			if childCallIndex, isChildCallEnded := childCallIndexByID[result.CallID]; isChildCallEnded {
				childCallOutputsByIndex[childCallIndex] = result.Outputs

				// decrement needed by counts for any needs
				for _, neededCallRef := range callSpecParallelCall[childCallIndex].Needs {
					childCallNeededCountByName[opspec.RefToName(neededCallRef)]--
				}

				for neededCallName, neededCount := range childCallNeededCountByName {
					if 1 > neededCount {
						neededCallIndex := childCallIndexByName[neededCallName]
						if cancel, ok := childCallCancellationByIndex[neededCallIndex]; ok {
							cancel()
							// cancelled "needed" calls do not produce outputs, but need need to
							// record outputs to allow final call ended count to pass
							childCallOutputsByIndex[neededCallIndex] = map[string]*model.Value{}
						}
					}
				}
			}

			if len(childCallOutputsByIndex) == len(childCallIndexByID) {
				// all calls have ended

				// construct parallel outputs
				for i := 0; i < len(callSpecParallelCall); i++ {
					callOutputs := childCallOutputsByIndex[i]
					for varName, varData := range callOutputs {
						outboundScope[varName] = varData
					}
				}

				return outboundScope, nil
			}
		}
	}
}
