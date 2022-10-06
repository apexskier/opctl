package op

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/data/git"
	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/inputs"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/dir"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/reference"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

// Interpret interprets an OpCallSpec into a OpCall
func Interpret(
	ctx context.Context,
	scope map[string]*model.Value,
	opCallSpec *model.OpCallSpec,
	opID string,
	parentOpPath string,
	gitOpsDir string,
	scratchDirPath string,
) (*model.OpCall, error) {
	var opPath string
	if reference.ReferenceRegexp.MatchString(opCallSpec.Ref) {
		// attempt to process as a variable reference since its variable reference like.
		dirValue, err := dir.Interpret(
			scope,
			opCallSpec.Ref,
			scratchDirPath,
			false,
		)
		if err != nil {
			return nil, fmt.Errorf("error encountered interpreting image src: %w", err)
		}
		opPath = *dirValue.Dir
	} else {
		opHandle, err := data.Resolve(
			ctx,
			opCallSpec.Ref,
			fs.New(parentOpPath, filepath.Dir(parentOpPath)),
			git.New(gitOpsDir),
		)
		if err != nil {
			return nil, err
		}
		opPath = *opHandle.Path()
	}

	opFile, err := opfile.Get(opPath)
	if err != nil {
		return nil, err
	}

	childCallID, err := uniquestring.Construct()
	if err != nil {
		return nil, err
	}

	opCall := &model.OpCall{
		BaseCall: model.BaseCall{
			OpPath: opPath,
		},
		ChildCallID:       childCallID,
		ChildCallCallSpec: opFile.Run,
		OpID:              opID,
	}

	opCall.Inputs, err = inputs.Interpret(
		opCallSpec.Inputs,
		opFile.Inputs,
		opPath,
		scope,
		scratchDirPath,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to interpret call to %v: %w", opCallSpec.Ref, err)
	}

	return opCall, nil
}
