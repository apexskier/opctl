package node

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opctl/opctl/sdks/go/internal/readchunks"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/containerruntime"
	"github.com/opctl/opctl/sdks/go/opspec"
)

//counterfeiter:generate -o internal/fakes/containerCaller.go . containerCaller
type containerCaller interface {
	// Executes a container call
	Call(
		ctx context.Context,
		eventChannel chan model.Event,
		containerCall *model.ContainerCall,
		containerCallSpec *model.ContainerCallSpec,
	) (
		map[string]*model.Value,
		error,
	)
}

func newContainerCaller(
	containerRuntime containerruntime.ContainerRuntime,
	privileged bool,
) containerCaller {
	return _containerCaller{
		containerRuntime: containerRuntime,
		privileged:       privileged,
	}
}

type _containerCaller struct {
	privileged       bool
	containerRuntime containerruntime.ContainerRuntime
}

func (cc _containerCaller) Call(
	ctx context.Context,
	eventChannel chan model.Event,
	containerCall *model.ContainerCall,
	containerCallSpec *model.ContainerCallSpec,
) (
	map[string]*model.Value,
	error,
) {
	outputs := map[string]*model.Value{}
	var exitCode int64

	logStdOutPR, logStdOutPW := io.Pipe()
	logStdErrPR, logStdErrPW := io.Pipe()

	// interpret logs
	logChan := make(chan error, 1)
	go func() {
		logChan <- cc.interpretLogs(
			eventChannel,
			logStdOutPR,
			logStdErrPR,
			containerCall,
		)
	}()

	outputs = cc.interpretOutputs(
		containerCallSpec,
		containerCall,
	)

	rawExitCode, err := cc.containerRuntime.RunContainer(
		ctx,
		eventChannel,
		containerCall,
		logStdOutPW,
		logStdErrPW,
		cc.privileged,
	)

	// @TODO: handle no exit code
	if rawExitCode != nil {
		exitCode = *rawExitCode
	}

	if exitCode != 0 {
		err = fmt.Errorf("container exited with %d", exitCode)
	}

	// wait on logChan
	if logChanErr := <-logChan; err == nil {
		// non-destructively set err
		err = logChanErr
	}

	return outputs, err
}

func (this _containerCaller) interpretLogs(
	eventChannel chan model.Event,
	stdOutReader io.Reader,
	stdErrReader io.Reader,
	containerCall *model.ContainerCall,
) error {
	stdOutLogChan := make(chan error, 1)
	go func() {
		// interpret stdOut
		stdOutLogChan <- readchunks.ReadChunks(
			stdOutReader,
			func(chunk []byte) {
				eventChannel <- model.Event{
					Timestamp: time.Now().UTC(),
					ContainerStdOutWrittenTo: &model.ContainerStdOutWrittenTo{
						Data:        chunk,
						ContainerID: containerCall.ContainerID,
						OpRef:       containerCall.OpPath,
					},
				}
			})
	}()

	stdErrLogChan := make(chan error, 1)
	go func() {
		// interpret stdErr
		stdErrLogChan <- readchunks.ReadChunks(
			stdErrReader,
			func(chunk []byte) {
				eventChannel <- model.Event{
					Timestamp: time.Now().UTC(),
					ContainerStdErrWrittenTo: &model.ContainerStdErrWrittenTo{
						Data:        chunk,
						ContainerID: containerCall.ContainerID,
						OpRef:       containerCall.OpPath,
					},
				}
			})
	}()

	// wait on logs
	stdOutLogErr := <-stdOutLogChan
	stdErrLogErr := <-stdErrLogChan

	// return errs
	if stdOutLogErr != nil {
		return stdOutLogErr
	}
	if stdErrLogErr != nil {
		return stdErrLogErr
	}

	return nil
}

func (this _containerCaller) interpretOutputs(
	containerCallSpec *model.ContainerCallSpec,
	containerCall *model.ContainerCall,
) map[string]*model.Value {
	outputs := map[string]*model.Value{}

	for socketAddr, name := range containerCallSpec.Sockets {
		// add socket outputs
		if "0.0.0.0" == socketAddr {
			outputs[name] = &model.Value{Socket: &containerCall.ContainerID}
		}
	}
	for callSpecContainerFilePath, mountSrc := range containerCallSpec.Files {
		mountSrcStr, ok := mountSrc.(string)
		if !ok {
			continue
		}

		if mountSrcStr == "" {
			// skip embedded files
			continue
		}

		// add file outputs
		for callContainerFilePath, callHostFilePath := range containerCall.Files {
			if callSpecContainerFilePath == callContainerFilePath {
				// copy callHostFilePath before taking address; range vars have same address for every iteration
				value := callHostFilePath
				outputs[opspec.RefToName(mountSrcStr)] = &model.Value{File: &value}
			}
		}
	}
	for callSpecContainerDirPath, mountSrc := range containerCallSpec.Dirs {
		mountSrcStr, ok := mountSrc.(string)
		if !ok {
			continue
		}

		if mountSrcStr == "" {
			// skip embedded dirs
			continue
		}

		// add dir outputs
		for callContainerDirPath, callHostDirPath := range containerCall.Dirs {
			if callSpecContainerDirPath == callContainerDirPath {
				// copy callHostDirPath before taking address; range vars have same address for every iteration
				value := callHostDirPath
				outputs[opspec.RefToName(mountSrcStr)] = &model.Value{Dir: &value}
			}
		}
	}

	return outputs
}
