package node

import (
	"context"
	"fmt"
	"time"

	"github.com/opctl/opctl/sdks/go/model"
	callpkg "github.com/opctl/opctl/sdks/go/opspec/interpreter/call"
)

//counterfeiter:generate -o internal/fakes/caller.go . caller
type caller interface {
	// Call executes a call
	Call(
		ctx context.Context,
		eventChannel chan model.Event,
		id string,
		scope map[string]*model.Value,
		callSpec *model.CallSpec,
		opPath string,
		parentCallID *string,
		rootCallID string,
		scratchPath string,
	) (
		map[string]*model.Value,
		error,
	)
}

func newCaller(
	containerCaller containerCaller,
	gitOpsDir string,
) caller {
	instance := &_caller{
		containerCaller: containerCaller,
		gitOpsDir:       gitOpsDir,
	}
	instance.opCaller = newOpCaller(instance)
	instance.parallelCaller = newParallelCaller(instance)
	instance.parallelLoopCaller = newParallelLoopCaller(instance)
	instance.serialCaller = newSerialCaller(instance)
	instance.serialLoopCaller = newSerialLoopCaller(instance)

	return instance
}

type _caller struct {
	containerCaller    containerCaller
	opCaller           opCaller
	parallelCaller     parallelCaller
	parallelLoopCaller parallelLoopCaller
	serialCaller       serialCaller
	serialLoopCaller   serialLoopCaller

	gitOpsDir string
}

func (clr _caller) Call(
	ctx context.Context,
	eventChannel chan model.Event,
	id string,
	scope map[string]*model.Value,
	callSpec *model.CallSpec,
	opPath string,
	parentCallID *string,
	rootCallID string,
	scratchPath string,
) (
	map[string]*model.Value,
	error,
) {
	callCtx, cancelCall := context.WithCancel(ctx)
	defer cancelCall()
	var err error
	var outputs map[string]*model.Value
	var call *model.Call
	callStartTime := time.Now().UTC()

	if callCtx.Err() != nil {
		// if context done NOOP
		return nil, nil
	}

	// emit a call ended event after this call is complete
	defer func() {
		// defer must be defined before conditional return statements so it always runs

		if call == nil {
			call = &model.Call{
				ID:       id,
				RootID:   rootCallID,
				ParentID: parentCallID,
			}
		}

		event := model.Event{
			CallEnded: &model.CallEnded{
				Call:  *call,
				OpRef: opPath,
			},
			Timestamp: time.Now().UTC(),
		}

		if ctx.Err() != nil {
			// this call or parent call killed/cancelled
			event.CallEnded.Outcome = model.OpOutcomeKilled
			event.CallEnded.Error = &model.CallEndedError{
				Message: ctx.Err().Error(),
			}
		} else if err != nil {
			event.CallEnded.Outcome = model.OpOutcomeFailed
			event.CallEnded.Error = &model.CallEndedError{
				Message: err.Error(),
			}
		} else {
			event.CallEnded.Outcome = model.OpOutcomeSucceeded
		}

		eventChannel <- event
	}()

	if callSpec == nil {
		// NOOP
		return outputs, err
	}

	call, err = callpkg.Interpret(
		ctx,
		scope,
		callSpec,
		id,
		opPath,
		parentCallID,
		rootCallID,
		clr.gitOpsDir,
		scratchPath,
	)
	if err != nil {
		return nil, err
	}

	if call.If != nil && !*call.If {
		return outputs, err
	}

	// Ensure this is emitted just after the deferred operation to emit the end
	// event is set up, so we always have a matching start and end event
	eventChannel <- model.Event{
		Timestamp: callStartTime,
		CallStarted: &model.CallStarted{
			Call:  *call,
			OpRef: opPath,
		},
	}

	switch {
	case callSpec.Container != nil:
		// note: scope is effectively passed via files/dirs/env, etc
		outputs, err = clr.containerCaller.Call(
			callCtx,
			eventChannel,
			call.Container,
			callSpec.Container,
		)
	case callSpec.Op != nil:
		// note: scope is passed via call inputs instead of the scope object
		outputs, err = clr.opCaller.Call(
			callCtx,
			eventChannel,
			call.Op,
			rootCallID,
			callSpec.Op,
			scratchPath,
		)
	case callSpec.Parallel != nil:
		outputs, err = clr.parallelCaller.Call(
			callCtx,
			eventChannel,
			id,
			scope,
			rootCallID,
			opPath,
			*callSpec.Parallel,
			scratchPath,
		)
	case callSpec.ParallelLoop != nil:
		outputs, err = clr.parallelLoopCaller.Call(
			callCtx,
			eventChannel,
			scope,
			*callSpec.ParallelLoop,
			opPath,
			parentCallID,
			rootCallID,
			scratchPath,
		)
	case callSpec.Serial != nil:
		outputs, err = clr.serialCaller.Call(
			callCtx,
			eventChannel,
			id,
			scope,
			rootCallID,
			opPath,
			*callSpec.Serial,
			scratchPath,
		)
	case callSpec.SerialLoop != nil:
		outputs, err = clr.serialLoopCaller.Call(
			callCtx,
			eventChannel,
			scope,
			*callSpec.SerialLoop,
			opPath,
			parentCallID,
			rootCallID,
			scratchPath,
		)
	default:
		err = fmt.Errorf("invalid call graph '%+v'", callSpec)
	}

	return outputs, err
}
