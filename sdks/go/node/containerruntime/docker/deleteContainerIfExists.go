package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func (cr _containerRuntime) DeleteContainerIfExists(
	ctx context.Context,
	containerID string, // opctl container ID
) error {
	containers, err := cr.dockerClient.ContainerList(
		ctx,
		types.ContainerListOptions{
			Filters: filters.NewArgs(
				filters.KeyValuePair{
					Key:   "label",
					Value: fmt.Sprintf("%s=%s", containerIDLabel, containerID),
				},
			),
		},
	)
	if err != nil {
		return err
	}

	errGroup, _ := errgroup.WithContext(ctx)
	for _, container := range containers {
		errGroup.Go(func() error { return cr.runContainer.stopAndCleanup(ctx, container.ID) })
	}

	return errGroup.Wait()
}
