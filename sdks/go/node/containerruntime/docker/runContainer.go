package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	dockerClientPkg "github.com/docker/docker/client"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/pkg/errors"
)

func newRunContainer(
	networkName string,
	dockerClient dockerClientPkg.CommonAPIClient,
	dockerConfigPath string,
) runContainer {
	return runContainer{
		containerStdErrStreamer: newContainerStdErrStreamer(dockerClient),
		containerStdOutStreamer: newContainerStdOutStreamer(dockerClient),
		dockerClient:            dockerClient,
		ensureNetworkExistser:   newEnsureNetworkExistser(dockerClient),
		imagePuller:             newImagePuller(dockerClient, dockerConfigPath),
		imagePusher:             newImagePusher(),
		networkName:             networkName,
	}
}

type runContainer struct {
	containerStdErrStreamer containerLogStreamer
	containerStdOutStreamer containerLogStreamer
	dockerClient            dockerClientPkg.CommonAPIClient
	ensureNetworkExistser   ensureNetworkExistser
	imagePuller             imagePuller
	imagePusher             imagePusher
	networkName             string
}

// stopAndCleanup stops and cleans up a single docker container
func (cr runContainer) stopAndCleanup(
	ctx context.Context,
	container string, // docker container name/ID
) error {
	// try to stop the container gracefully prior to deletion
	stopTimeout := 10 * time.Second
	err := cr.dockerClient.ContainerStop(ctx, container, &stopTimeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to stop container: %v", err)
	}

	// now delete the container post-termination
	err = cr.dockerClient.ContainerRemove(
		ctx,
		container,
		types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to delete container: %v", err)
	}

	return nil
}

func (cr runContainer) RunContainer(
	ctx context.Context,
	eventChannel chan model.Event,
	req *model.ContainerCall,
	stdout io.WriteCloser,
	stderr io.WriteCloser,
	privileged bool,
) (*int64, error) {
	defer stdout.Close()
	defer stderr.Close()

	// ensure user defined network exists to allow inter container resolution via name
	// @TODO: remove when socket outputs supported
	if err := cr.ensureNetworkExistser.EnsureNetworkExists(
		ctx,
		cr.networkName,
	); err != nil {
		return nil, err
	}

	var imageErr error
	if req.Image.Src != nil {
		imageRef := fmt.Sprintf("%s:latest", req.ContainerID)
		req.Image.Ref = &imageRef

		imageErr = cr.imagePusher.Push(
			ctx,
			imageRef,
			req.Image.Src,
		)
	} else {
		imageErr = cr.imagePuller.Pull(
			ctx,
			req,
			eventChannel,
		)
		// don't err yet; image might be cached. We allow this to support offline use
	}

	portBindings, err := constructPortBindings(
		req.Ports,
	)
	if err != nil {
		return nil, err
	}

	// construct networking config
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			cr.networkName: {},
		},
	}
	if req.Name != nil {
		networkingConfig.EndpointsConfig[cr.networkName].Aliases = []string{
			*req.Name,
		}
	}

	// for docker, we prefix name in order to allow external tools to know it's an opctl managed container
	nameParts := []string{"miniopctl"}
	if req.Name != nil {
		nameParts = append(nameParts, *req.Name)
	}
	nameParts = append(nameParts, req.ContainerID)
	containerName := strings.Join(nameParts, "_")

	// create container
	containerCreatedResponse, err := cr.dockerClient.ContainerCreate(
		ctx,
		constructContainerConfig(
			req.Cmd,
			req.EnvVars,
			*req.Image.Ref,
			portBindings,
			req.WorkDir,
			req.ContainerID,
		),
		constructHostConfig(
			req.Dirs,
			req.Files,
			req.Sockets,
			portBindings,
			privileged,
		),
		networkingConfig,
		// platform requires API v1.41 so set to nil to avoid version errors
		nil,
		containerName,
	)

	if err != nil {
		if imageErr == nil {
			return nil, err
		}
		// if imageErr occurred prior; combine errors
		return nil, errors.New(strings.Join([]string{imageErr.Error(), err.Error()}, ", "))
	}

	defer func() {
		newCtx := context.Background() // always use a fresh context, to clean up after cancellation
		cr.stopAndCleanup(newCtx, containerCreatedResponse.ID)
	}()

	// start container
	if err := cr.dockerClient.ContainerStart(
		ctx,
		containerCreatedResponse.ID,
		types.ContainerStartOptions{},
	); err != nil {
		return nil, err
	}

	var waitGroup sync.WaitGroup
	errChan := make(chan error, 3)
	waitGroup.Add(2)

	go func() {
		if err := cr.containerStdErrStreamer.Stream(
			ctx,
			containerName,
			stderr,
		); err != nil {
			errChan <- err
		}
		waitGroup.Done()
	}()

	go func() {
		if err := cr.containerStdOutStreamer.Stream(
			ctx,
			containerName,
			stdout,
		); err != nil {
			errChan <- err
		}
		waitGroup.Done()
	}()

	var exitCode int64
	waitOkChan, waitErrChan := cr.dockerClient.ContainerWait(
		ctx,
		containerName,
		container.WaitConditionNotRunning,
	)
	select {
	case waitOk := <-waitOkChan:
		exitCode = waitOk.StatusCode
	case waitErr := <-waitErrChan:
		err = fmt.Errorf("error waiting on container: %w", waitErr)
	}

	// ensure stdout, and stderr all read before returning
	waitGroup.Wait()

	if err != nil && len(errChan) > 0 {
		// non-destructively set err
		err = <-errChan
	}
	return &exitCode, err
}
