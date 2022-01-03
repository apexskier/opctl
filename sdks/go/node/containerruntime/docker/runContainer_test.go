package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/containerruntime/docker/internal/fakes"
)

var _ = Context("RunContainer", func() {
	closedContainerWaitOkBodyChan := make(chan container.ContainerWaitOKBody)
	close(closedContainerWaitOkBodyChan)

	dbDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	It("should call dockerClient.ContainerRemove w/ expected args", func() {
		/* arrange */
		providedReq := &model.ContainerCall{
			BaseCall:    model.BaseCall{},
			ContainerID: "containerID",
			Image:       &model.ContainerCallImage{Ref: new(string)},
			// invalid to trigger early return
			Ports: map[string]string{"*": "&"},
		}

		expectedContainerRemoveOptions := types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		}

		fakeDockerClient := new(FakeCommonAPIClient)
		fakeDockerClient.ContainerWaitReturns(closedContainerWaitOkBodyChan, nil)

		objectUnderTest := _runContainer{
			containerStdErrStreamer: new(FakeContainerLogStreamer),
			containerStdOutStreamer: new(FakeContainerLogStreamer),
			dockerClient:            fakeDockerClient,
			imagePuller:             new(FakeImagePuller),
			ensureNetworkExistser:   new(FakeEnsureNetworkExistser),
		}

		/* act */
		objectUnderTest.RunContainer(
			context.Background(),
			make(chan model.Event),
			providedReq,
			"rootCallID",
			nopWriteCloser{io.Discard},
			nopWriteCloser{io.Discard},
		)

		/* assert */
		_, actualContainerName, _ := fakeDockerClient.ContainerStopArgsForCall(0)
		Expect(actualContainerName).To(Equal(fmt.Sprintf("opctl_%s", providedReq.ContainerID)))

		_, actualContainerName, actualContainerRemoveOptions := fakeDockerClient.ContainerRemoveArgsForCall(0)
		Expect(actualContainerName).To(Equal(fmt.Sprintf("opctl_%s", providedReq.ContainerID)))
		Expect(actualContainerRemoveOptions).To(Equal(expectedContainerRemoveOptions))

	})
	Context("portBindingsFactory.Construct errs", func() {
		It("should return expected result", func() {
			/* arrange */

			objectUnderTest := _runContainer{
				containerStdErrStreamer: new(FakeContainerLogStreamer),
				containerStdOutStreamer: new(FakeContainerLogStreamer),
				dockerClient:            new(FakeCommonAPIClient),
				ensureNetworkExistser:   new(FakeEnsureNetworkExistser),
				imagePuller:             new(FakeImagePuller),
			}

			/* act */
			_, actualErr := objectUnderTest.RunContainer(
				context.Background(),
				make(chan model.Event),
				&model.ContainerCall{
					Image: &model.ContainerCallImage{Ref: new(string)},
					Ports: map[string]string{
						"*": "&",
					},
				},
				"rootCallID",
				nopWriteCloser{io.Discard},
				nopWriteCloser{io.Discard},
			)

			/* assert */
			Expect(actualErr).To(MatchError("Invalid containerPort: *"))
		})
	})
	Context("constructPortBindings doesn't err", func() {

		It("should call imagePuller.Pull w/ expected args", func() {

			/* arrange */
			providedCtx := context.Background()
			providedReq := &model.ContainerCall{
				BaseCall:    model.BaseCall{},
				ContainerID: "dummyContainerID",
				Image:       &model.ContainerCallImage{Ref: new(string)},
			}
			providedRootCallID := "providedRootCallID"

			providedEventChannel := make(chan model.Event)

			fakeImagePuller := new(FakeImagePuller)

			fakeDockerClient := new(FakeCommonAPIClient)
			fakeDockerClient.ContainerWaitReturns(closedContainerWaitOkBodyChan, nil)

			objectUnderTest := _runContainer{
				containerStdErrStreamer: new(FakeContainerLogStreamer),
				containerStdOutStreamer: new(FakeContainerLogStreamer),
				dockerClient:            fakeDockerClient,
				ensureNetworkExistser:   new(FakeEnsureNetworkExistser),
				imagePuller:             fakeImagePuller,
			}

			/* act */
			objectUnderTest.RunContainer(
				providedCtx,
				make(chan model.Event),
				providedReq,
				providedRootCallID,
				nopWriteCloser{io.Discard},
				nopWriteCloser{io.Discard},
			)

			/* assert */
			actualCtx,
				actualReq,
				actualRootCallID,
				actualEventPublisher := fakeImagePuller.PullArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualReq).To(Equal(providedReq))
			Expect(actualRootCallID).To(Equal(providedRootCallID))
			Expect(actualEventPublisher).To(Equal(providedEventChannel))
		})

		It("should call dockerClient.ContainerCreate w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedReq := &model.ContainerCall{
				BaseCall:    model.BaseCall{},
				ContainerID: "dummyContainerID",
				Dirs: map[string]string{
					"dir1ContainerPath": "dir1HostPath",
				},
				Files: map[string]string{
					"file1ContainerPath": "file1HostPath",
				},
				Image: &model.ContainerCallImage{Ref: new(string)},
				Name:  new(string),
				Sockets: map[string]string{
					"/unixSocket1ContainerAddress": "/unixSocket1HostAddress",
				},
				Ports: map[string]string{
					"80": "80",
				},
			}

			expectedPortBindings, err := constructPortBindings(providedReq.Ports)
			if err != nil {
				panic(err)
			}

			expectedContainerConfig := constructContainerConfig(
				providedReq.Cmd,
				providedReq.EnvVars,
				*providedReq.Image.Ref,
				expectedPortBindings,
				providedReq.WorkDir,
				"dummyContainerID",
				"rootCallID",
			)

			expectedHostConfig := &container.HostConfig{
				Mounts: []mount.Mount{
					{
						Type:          "bind",
						Target:        "file1ContainerPath",
						Source:        "file1HostPath",
						Consistency:   "cached",
						ReadOnly:      false,
						BindOptions:   nil,
						VolumeOptions: nil,
						TmpfsOptions:  nil,
					},
					{
						Type:          "bind",
						Source:        "dir1HostPath",
						Target:        "dir1ContainerPath",
						Consistency:   "cached",
						ReadOnly:      false,
						BindOptions:   nil,
						TmpfsOptions:  nil,
						VolumeOptions: nil,
					},
					{
						Type:          "bind",
						Source:        "/unixSocket1HostAddress",
						Target:        "/unixSocket1ContainerAddress",
						ReadOnly:      false,
						Consistency:   "",
						BindOptions:   nil,
						VolumeOptions: nil,
						TmpfsOptions:  nil,
					},
				},
				PortBindings: nat.PortMap{
					"80/tcp": []nat.PortBinding{
						{
							HostPort: "80",
						},
					},
				},
				Privileged: true,
			}

			expectedNetworkingConfig := &network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					networkName: {
						Aliases: []string{
							*providedReq.Name,
						},
					},
				},
			}

			fakeDockerClient := new(FakeCommonAPIClient)
			fakeDockerClient.ContainerWaitReturns(closedContainerWaitOkBodyChan, nil)

			objectUnderTest := _runContainer{
				containerStdErrStreamer: new(FakeContainerLogStreamer),
				containerStdOutStreamer: new(FakeContainerLogStreamer),
				dockerClient:            fakeDockerClient,
				ensureNetworkExistser:   new(FakeEnsureNetworkExistser),
				imagePuller:             new(FakeImagePuller),
			}

			/* act */
			objectUnderTest.RunContainer(
				providedCtx,
				make(chan model.Event),
				providedReq,
				"rootCallID",
				nopWriteCloser{io.Discard},
				nopWriteCloser{io.Discard},
			)

			/* assert */
			actualCtx,
				actualContainerConfig,
				actualHostConfig,
				actualNetworkingConfig,
				actualPlatformConfig,
				actualContainerName := fakeDockerClient.ContainerCreateArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualContainerConfig).To(Equal(expectedContainerConfig))
			Expect(*actualHostConfig).To(Equal(*expectedHostConfig))
			Expect(actualNetworkingConfig).To(Equal(expectedNetworkingConfig))
			Expect(actualPlatformConfig).To(BeNil())
			Expect(actualContainerName).To(Equal(fmt.Sprintf("opctl_%s", providedReq.ContainerID)))
		})
	})
})