package node

import (
	"context"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/internal/fakes"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewCaller(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(
		newCaller(
			new(FakeContainerCaller),
			"dummyDataDir",
		),
	).To(Not(BeNil()))
}

func TestNilCallSpec(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeContainerCaller := new(FakeContainerCaller)
	dataDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	/* act */
	objectUnderTest := _caller{
		containerCaller: fakeContainerCaller,
		gitOpsDir:       dataDir,
	}

	/* assert */
	_, err = objectUnderTest.Call(
		context.Background(),
		make(chan model.Event, 10),
		"dummyCallID",
		map[string]*model.Value{},
		nil,
		"dummyOpPath",
		nil,
		"dummyRootCallID",
		dataDir,
	)

	g.Expect(err).To(BeNil())
}

func TestInterpretFalseIf(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	providedCallID := "dummyCallID"
	providedOpPath := "testdata/startOp"
	providedRootCallID := "dummyRootCallID"

	predicateSpec := []interface{}{
		true,
		true,
	}

	ifSpec := []*model.PredicateSpec{
		{
			Eq: &predicateSpec,
		},
	}

	expectedIf := true
	expectedEvent := model.Event{
		CallStarted: &model.CallStarted{
			Call: model.Call{
				ID:     providedCallID,
				If:     &expectedIf,
				RootID: providedRootCallID,
				Serial: []*model.CallSpec{},
			},
			OpRef: providedOpPath,
		},
	}

	fakeSerialCaller := new(FakeSerialCaller)

	dataDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	eventChannel := make(chan model.Event, 10)

	objectUnderTest := _caller{
		containerCaller: new(FakeContainerCaller),
		gitOpsDir:       dataDir,
		serialCaller:    fakeSerialCaller,
	}

	/* act */
	objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		map[string]*model.Value{},
		&model.CallSpec{
			If:     &ifSpec,
			Serial: &[]*model.CallSpec{},
		},
		providedOpPath,
		nil,
		providedRootCallID,
		dataDir,
	)

	/* assert */
	actualEvent := <-eventChannel

	// @TODO: implement/use VTime (similar to IOS & VFS) so we don't need custom assertions on temporal fields
	g.Expect(actualEvent.Timestamp).To(BeTemporally("~", time.Now().UTC(), 5*time.Second))
	// set temporal fields to expected vals since they're already asserted
	actualEvent.Timestamp = expectedEvent.Timestamp

	g.Expect(actualEvent).To(Equal(expectedEvent))
}

func TestCall(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	providedCallID := "providedCallID"
	providedOpPath := "providedOpPath"
	fakeContainerCaller := new(FakeContainerCaller)

	providedScope := map[string]*model.Value{}
	imageSpec := &model.ContainerCallImageSpec{
		Ref: "docker.io/library/ref",
	}
	providedCallSpec := &model.CallSpec{
		Container: &model.ContainerCallSpec{
			Image: imageSpec,
		},
	}
	providedRootCallID := "providedRootCallID"

	expectedCall := &model.Call{
		Container: &model.ContainerCall{
			BaseCall: model.BaseCall{
				OpPath: providedOpPath,
			},
			ContainerID: providedCallID,
			Cmd:         []string{},
			Dirs:        map[string]string{},
			Files:       map[string]string{},
			Image: &model.ContainerCallImage{
				Ref: &imageSpec.Ref,
			},
			Sockets: map[string]string{},
		},
	}

	dataDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	eventChannel := make(chan model.Event, 10)

	objectUnderTest := _caller{
		containerCaller: fakeContainerCaller,
		gitOpsDir:       dataDir,
	}

	/* act */
	_, actualErr := objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		nil,
		providedRootCallID,
		dataDir,
	)

	/* assert */
	g.Expect(actualErr).To(BeNil())
	_,
		_,
		actualContainerCall,
		actualCallSpec := fakeContainerCaller.CallArgsForCall(0)

	g.Expect(actualContainerCall).To(Equal(expectedCall.Container))
	g.Expect(actualCallSpec).To(Equal(providedCallSpec.Container))
}

func TestCall2(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeOpCaller := new(FakeOpCaller)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	providedOpPath := filepath.Join(wd, "testdata/caller")

	providedCallID := "dummyCallID"
	providedScope := map[string]*model.Value{}
	providedCallSpec := &model.CallSpec{
		Op: &model.OpCallSpec{
			Ref: providedOpPath,
		},
	}
	providedParentID := "providedParentID"
	providedRootCallID := "dummyRootCallID"

	expectedCall := &model.Call{
		Op: &model.OpCall{
			BaseCall: model.BaseCall{
				OpPath: providedOpPath,
			},
			OpID:   providedCallID,
			Inputs: map[string]*model.Value{},
		},
	}

	dataDir, err := os.MkdirTemp("", "")
	if err != nil {
		panic(err)
	}

	eventChannel := make(chan model.Event, 10)

	objectUnderTest := _caller{
		gitOpsDir: dataDir,
		opCaller:  fakeOpCaller,
	}

	/* act */
	_, actualErr := objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		&providedParentID,
		providedRootCallID,
		dataDir,
	)

	/* assert */
	g.Expect(actualErr).To(BeNil())
	_,
		_,
		actualOpCall,
		actualParentID,
		actualRootCallID,
		actualCallSpec := fakeOpCaller.CallArgsForCall(0)

	expectedCall.Op.ChildCallID = actualOpCall.ChildCallID
	g.Expect(actualOpCall).To(Equal(*expectedCall.Op))
	g.Expect(actualParentID).To(Equal(providedParentID))
	g.Expect(actualRootCallID).To(Equal(providedRootCallID))
	g.Expect(actualCallSpec).To(Equal(providedCallSpec.Op))
}

func TestParallelCall(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeParallelCaller := new(FakeParallelCaller)

	providedCallID := "dummyCallID"
	providedScope := map[string]*model.Value{}
	providedCallSpec := &model.CallSpec{
		Parallel: &[]*model.CallSpec{
			{Container: &model.ContainerCallSpec{}},
		},
	}
	providedOpPath := "providedOpPath"
	providedRootCallID := "dummyRootCallID"

	eventChannel := make(chan model.Event, 2)
	objectUnderTest := _caller{
		parallelCaller: fakeParallelCaller,
	}

	/* act */
	objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		nil,
		providedRootCallID,
		"",
	)

	/* assert */
	_,
		_,
		actualCallID,
		actualScope,
		actualRootCallID,
		actualOpPath,
		actualCallSpec,
		_ := fakeParallelCaller.CallArgsForCall(0)

	g.Expect(actualCallID).To(Equal(providedCallID))
	g.Expect(actualScope).To(Equal(providedScope))
	g.Expect(actualRootCallID).To(Equal(providedRootCallID))
	g.Expect(actualOpPath).To(Equal(providedOpPath))
	g.Expect(actualCallSpec).To(Equal(*providedCallSpec.Parallel))
}

func TestParallelLoop(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeParallelLoopCaller := new(FakeParallelLoopCaller)

	providedCallID := "dummyCallID"
	providedScope := map[string]*model.Value{}
	providedCallSpec := &model.CallSpec{
		ParallelLoop: &model.ParallelLoopCallSpec{},
	}
	providedOpPath := "providedOpPath"
	providedRootCallID := "dummyRootCallID"
	providedParentID := "providedParentID"

	eventChannel := make(chan model.Event, 2)
	objectUnderTest := _caller{
		parallelLoopCaller: fakeParallelLoopCaller,
	}

	/* act */
	objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		&providedParentID,
		providedRootCallID,
		"",
	)

	/* assert */
	_,
		_,
		actualScope,
		actualParallelLoopCallSpec,
		actualOpPath,
		actualParentID,
		actualRootCallID,
		_ := fakeParallelLoopCaller.CallArgsForCall(0)

	g.Expect(actualScope).To(Equal(providedScope))
	g.Expect(actualParallelLoopCallSpec).To(Equal(*providedCallSpec.ParallelLoop))
	g.Expect(actualOpPath).To(Equal(providedOpPath))
	g.Expect(*actualParentID).To(Equal(providedParentID))
	g.Expect(actualRootCallID).To(Equal(providedRootCallID))
}

func TestSerial(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeSerialCaller := new(FakeSerialCaller)

	providedCallID := "dummyCallID"
	providedScope := map[string]*model.Value{}
	providedCallSpec := &model.CallSpec{
		Serial: &[]*model.CallSpec{
			{Container: &model.ContainerCallSpec{}},
		},
	}
	providedOpPath := "providedOpPath"
	providedRootCallID := "dummyRootCallID"

	eventChannel := make(chan model.Event, 2)
	objectUnderTest := _caller{
		containerCaller: new(FakeContainerCaller),
		serialCaller:    fakeSerialCaller,
	}

	/* act */
	objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		nil,
		providedRootCallID,
		"",
	)

	/* assert */
	_,
		_,
		actualCallID,
		actualScope,
		actualRootCallID,
		actualOpPath,
		actualCallSpec,
		_ := fakeSerialCaller.CallArgsForCall(0)

	g.Expect(actualCallID).To(Equal(providedCallID))
	g.Expect(actualScope).To(Equal(providedScope))
	g.Expect(actualRootCallID).To(Equal(providedRootCallID))
	g.Expect(actualOpPath).To(Equal(providedOpPath))
	g.Expect(actualCallSpec).To(Equal(*providedCallSpec.Serial))
}

func TestSerialLoop(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	fakeSerialLoopCaller := new(FakeSerialLoopCaller)

	providedCallID := "dummyCallID"
	providedScope := map[string]*model.Value{}
	providedCallSpec := &model.CallSpec{
		SerialLoop: &model.SerialLoopCallSpec{
			Range: []interface{}{},
		},
	}
	providedOpPath := "providedOpPath"
	providedRootCallID := "dummyRootCallID"
	providedParentID := "providedParentID"

	eventChannel := make(chan model.Event, 2)
	objectUnderTest := _caller{
		serialLoopCaller: fakeSerialLoopCaller,
	}

	/* act */
	objectUnderTest.Call(
		context.Background(),
		eventChannel,
		providedCallID,
		providedScope,
		providedCallSpec,
		providedOpPath,
		&providedParentID,
		providedRootCallID,
		"",
	)

	/* assert */
	_,
		_,
		actualScope,
		actualSerialLoopCallSpec,
		actualOpPath,
		actualParentID,
		actualRootCallID,
		_ := fakeSerialLoopCaller.CallArgsForCall(0)

	g.Expect(actualScope).To(Equal(providedScope))
	g.Expect(actualSerialLoopCallSpec).To(Equal(*providedCallSpec.SerialLoop))
	g.Expect(actualOpPath).To(Equal(providedOpPath))
	g.Expect(*actualParentID).To(Equal(providedParentID))
	g.Expect(actualRootCallID).To(Equal(providedRootCallID))
}
