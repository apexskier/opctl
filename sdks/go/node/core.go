// Package node defines the core interface for an opctl node
package node

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/containerruntime"
)

// New returns a new LocalCore initialized with the given options
func New(
	ctx context.Context,
	containerRuntime containerruntime.ContainerRuntime,
	dataDirPath string,
) (Core, error) {
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

func (c core) Liveness(
	ctx context.Context,
) error {
	return nil
}

//counterfeiter:generate -o fakes/core.go . Core

// Core is an Node that supports running ops directly on the current machine
type Core interface {
	Node

	// Resolve attempts to resolve data via local filesystem or git
	// nil pullCreds will be ignored
	//
	// expected errs:
	//  - ErrDataProviderAuthentication on authentication failure
	//  - ErrDataProviderAuthorization on authorization failure
	//  - ErrDataRefResolution on resolution failure
	ResolveData(
		ctx context.Context,
		eventChannel chan model.Event,
		callID string,
		dataRef string,
	) (
		model.DataHandle,
		error,
	)
}