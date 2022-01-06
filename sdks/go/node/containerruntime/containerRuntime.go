// Package containerruntime defines an interface abstracting container runtime interactions.
// A fake implementation is included to allow faking said interactions.
package containerruntime

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"io"

	"github.com/opctl/opctl/sdks/go/model"
)

// ContainerRuntime defines the interface container runtimes must implement to be supported by opctl
//counterfeiter:generate -o fakes/containerRuntime.go . ContainerRuntime
type ContainerRuntime interface {
	// Delete deletes opctl managed resources from the container runtime
	Delete(
		ctx context.Context,
	) error

	DeleteContainerIfExists(
		ctx context.Context,
		containerID string,
	) error

	// Kill stops/kills opctl managed resources within the container runtime
	Kill(
		ctx context.Context,
	) error

	// RunContainer creates, starts, and waits on a container. ExitCode &/Or an error will be returned
	RunContainer(
		ctx context.Context,
		eventChannel chan model.Event,
		req *model.ContainerCall,
		stdout io.WriteCloser,
		stderr io.WriteCloser,
		privileged bool,
	) (*int64, error)
}
