package opfile

import (
	"os"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/model"
)

// Get gets the validated, deserialized representation of an "op.yml" file
func Get(
	opPath string,
) (
	*model.OpSpec,
	error,
) {
	filePath := filepath.Join(opPath, FileName)

	opFileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return Unmarshal(filePath, opFileBytes)
}
