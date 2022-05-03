package node

import (
	"context"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/outputs"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/reference"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

//counterfeiter:generate -o internal/fakes/opCaller.go . opCaller
type opCaller interface {
	// Call executes an op call
	Call(
		ctx context.Context,
		eventChannel chan model.Event,
		opCall *model.OpCall,
		rootCallID string,
		opCallSpec *model.OpCallSpec,
		scratchPath string,
	) (
		map[string]*model.Value,
		error,
	)
}

func newOpCaller(
	caller caller,
) opCaller {
	return _opCaller{
		caller: caller,
	}
}

type _opCaller struct {
	caller caller
}

func (oc _opCaller) Call(
	ctx context.Context,
	eventChannel chan model.Event,
	opCall *model.OpCall,
	rootCallID string,
	opCallSpec *model.OpCallSpec,
	scratchPath string,
) (
	map[string]*model.Value,
	error,
) {
	// form scope for op call by combining defined inputs & op dir
	opCallScope := map[string]*model.Value{}
	for varName, varData := range opCall.Inputs {
		opCallScope[varName] = varData
	}
	// add deprecated absolute path to scope
	opCallScope["/"] = &model.Value{
		Dir: &opCall.OpPath,
	}
	// add current directory to scope
	opCallScope["./"] = &model.Value{
		Dir: &opCall.OpPath,
	}

	// add parent directory to scope
	parentDirPath := filepath.Dir(opCall.OpPath)
	opCallScope["../"] = &model.Value{
		Dir: &parentDirPath,
	}

	opOutputs, err := oc.caller.Call(
		ctx,
		eventChannel,
		opCall.ChildCallID,
		opCallScope,
		opCall.ChildCallCallSpec,
		opCall.OpPath,
		&opCall.OpID,
		rootCallID,
		scratchPath,
	)
	if err != nil {
		return nil, err
	}

	var opFile *model.OpSpec
	opFile, err = opfile.Get(opCall.OpPath)
	if err != nil {
		return nil, err
	}
	opOutputs, err = outputs.Interpret(
		opOutputs,
		opFile.Outputs,
		opCallSpec.Outputs,
		opCall.OpPath,
		filepath.Join(scratchPath, "call", opCall.OpID),
	)

	outboundScope := map[string]*model.Value{}

	// filter op outboundScope to bound call outboundScope
	for boundName, boundValue := range opCallSpec.Outputs {
		// return bound outboundScope
		if boundValue == "" {
			// implicit value
			boundValue = boundName
		} else if !reference.ReferenceRegexp.MatchString(boundValue) {
			// handle obsolete syntax by swapping order
			prevBoundName := boundName
			boundName = boundValue
			boundValue = prevBoundName
		} else {
			boundValue = opspec.RefToName(boundValue)
		}
		for opOutputName, opOutputValue := range opOutputs {
			if boundName == opOutputName {
				outboundScope[boundValue] = opOutputValue
			}
		}
	}

	return outboundScope, err
}
