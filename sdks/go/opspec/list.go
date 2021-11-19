package opspec

import (
	"context"
	"io"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
	"github.com/pkg/errors"
)

// List ops recursively within a directory, returning discovered op files by ref. Invalid ops will be returned in the error map.
func List(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	dirHandle model.DataHandle,
) (map[string]*model.OpSpec, map[string]error, error) {
	contents, err := dirHandle.ListDescendants(ctx, eventChannel, callID)
	if err != nil {
		return nil, nil, err
	}

	opsByPath := map[string]*model.OpSpec{}
	erroringOpsByPath := map[string]error{}
	for _, content := range contents {
		if filepath.Base(content.Path) == opfile.FileName {
			opFileName := filepath.Join(dirHandle.Ref(), content.Path)
			ref := filepath.Dir(opFileName)

			opFileReader, err := dirHandle.GetContent(ctx, eventChannel, callID, content.Path)
			if err != nil {
				erroringOpsByPath[ref] = errors.Wrapf(err, "error opening %s", opFileName)
				continue
			}

			opFileBytes, err := io.ReadAll(opFileReader)
			opFileReader.Close()
			if err != nil {
				erroringOpsByPath[ref] = errors.Wrapf(err, "error reading %s", opFileName)
				continue
			}

			opFile, err := opfile.Unmarshal(opFileName, opFileBytes)
			if err != nil {
				erroringOpsByPath[ref] = errors.Wrapf(err, "error unmarshalling %s", opFileName)
				continue
			}

			opsByPath[ref] = opFile
		}

	}

	return opsByPath, erroringOpsByPath, nil
}
