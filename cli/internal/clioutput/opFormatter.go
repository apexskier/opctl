package clioutput

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/opctl/opctl/sdks/go/opspec"
)

// OpFormatter formats an op ref in some way
type OpFormatter interface {
	FormatOpRef(opRef string) string
}

// CliOpFormatter formats an op ref in the context of a CLI run
type CliOpFormatter struct {
	workingDirPath string
	dataDirPath    string
}

var localOpPrefix = "." + string(os.PathSeparator) + opspec.DotOpspecDirName + string(os.PathSeparator)

// NewCliOpFormatter creates a new CliOpFormatter
func NewCliOpFormatter(workingDirPath, dataDirPath string) CliOpFormatter {
	return CliOpFormatter{workingDirPath, dataDirPath}
}

// FormatOpRef gives a more appropriate description of an op's reference
// Local ops will be formatted as paths relative to the working directory or
// home directory, installed ops will be formatted as url-like op refs
func (of CliOpFormatter) FormatOpRef(opRef string) string {
	if path.IsAbs(opRef) {
		home, err := os.UserHomeDir()
		if err != nil {
			return opRef
		}
		dataDirPath := of.dataDirPath
		if strings.HasPrefix(opRef, dataDirPath) {
			return opRef[len(filepath.Join(dataDirPath, "ops")+string(os.PathSeparator)):]
		}
		if strings.HasPrefix(opRef, of.workingDirPath) {
			opRef = "." + opRef[len(of.workingDirPath):]
			if strings.HasPrefix(opRef, localOpPrefix) {
				return opRef[len(localOpPrefix):]
			}
			return opRef
		}
		if strings.HasPrefix(opRef, home) {
			return "~" + opRef[len(home):]
		}
	}
	return opRef
}

// SimpleOpFormatter just mirrors the op ref as is
type SimpleOpFormatter struct{}

// FormatOpRef returns the op ref
func (SimpleOpFormatter) FormatOpRef(opRef string) string {
	return opRef
}
