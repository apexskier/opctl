package docker

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate -o internal/fakes/commonAPIClient.go github.com/docker/docker/client.CommonAPIClient

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	dockerClientPkg "github.com/docker/docker/client"
	"github.com/opctl/opctl/sdks/go/node/containerruntime"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

const (
	containerIDLabel = "miniopctl_container_id"
)

func New(
	ctx context.Context,
	networkName,
	host,
	dockerConfigPath string,
) (containerruntime.ContainerRuntime, error) {
	dockerClient, err := dockerClientPkg.NewClientWithOpts(dockerClientPkg.FromEnv, dockerClientPkg.WithHost(host))
	if err != nil {
		return nil, err
	}

	// degrade client version to version of server
	dockerClient.NegotiateAPIVersion(ctx)

	return _containerRuntime{
		networkName:  networkName,
		runContainer: newRunContainer(ctx, networkName, dockerClient, dockerConfigPath),
		dockerClient: dockerClient,
	}, nil
}

type _containerRuntime struct {
	networkName string
	runContainer
	dockerClient dockerClientPkg.CommonAPIClient
}

func (cr _containerRuntime) Delete(
	ctx context.Context,
) error {
	containers, err := cr.dockerClient.ContainerList(
		ctx,
		types.ContainerListOptions{
			Filters: filters.NewArgs(
				filters.KeyValuePair{
					Key:   "label",
					Value: containerIDLabel,
				},
				filters.KeyValuePair{
					Key:   "network",
					Value: cr.networkName,
				},
			),
		},
	)
	if err != nil {
		return err
	}

	errGroup, _ := errgroup.WithContext(ctx)
	for _, container := range containers {
		errGroup.Go(func() error { return cr.stopAndCleanup(ctx, container.ID) })
	}

	return errGroup.Wait()
}
