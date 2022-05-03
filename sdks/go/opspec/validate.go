package opspec

import (
	"github.com/opctl/opctl/sdks/go/opspec/opfile"
)

// Validate an op
func Validate(opPath string) error {
	_, err := opfile.Get(opPath)
	return err
}
