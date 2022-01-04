// Package node defines the core interface for an opctl node
package node

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/node/containerruntime"
)

// New returns a new Node
func New(
	ctx context.Context,
	containerRuntime containerruntime.ContainerRuntime,
	dataDirPath string,
) (Node, error) {
	caller := newCaller(
		newContainerCaller(
			containerRuntime,
		),
		dataDirPath,
	)

	return core{
		caller:           caller,
		containerRuntime: containerRuntime,
		dataCachePath:    filepath.Join(dataDirPath, "ops"),
		opCaller: newOpCaller(
			caller,
			dataDirPath,
		),
	}, nil
}

// core is an Node that supports running ops directly on the host
type core struct {
	caller           caller
	containerRuntime containerruntime.ContainerRuntime
	dataCachePath    string
	opCaller         opCaller
}
