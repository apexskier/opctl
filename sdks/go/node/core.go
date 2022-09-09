// Package node defines the core interface for an opctl node
package node

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"github.com/opctl/opctl/sdks/go/node/containerruntime"
	"path/filepath"
)

// New returns a new Node
func New(
	containerRuntime containerruntime.ContainerRuntime,
	dataDirPath string,
	privileged bool,
) (Node, error) {
	gitOpsDir := filepath.Join(dataDirPath, "ops")

	caller := newCaller(
		newContainerCaller(
			containerRuntime,
			privileged,
		),
		gitOpsDir,
	)

	return core{
		caller:      caller,
		gitOpsDir:   gitOpsDir,
		dataDirPath: dataDirPath,
	}, nil
}

// core is a Node that supports running ops directly on the host
type core struct {
	caller      caller
	gitOpsDir   string
	dataDirPath string
}
