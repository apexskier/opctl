package node

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/internal/fakes"
)

var _ = Context("caller", func() {
	Context("newCaller", func() {
		It("should return caller", func() {
			/* arrange/act/assert */
			Expect(
				newCaller(
					new(FakeContainerCaller),
					"dummyDataDir",
				),
			).To(Not(BeNil()))
		})
	})
	Context("Call", func() {
		Context("Nil CallSpec", func() {
			It("should not throw", func() {
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
					make(chan model.Event),
					"dummyCallID",
					map[string]*model.Value{},
					nil,
					"dummyOpPath",
					nil,
					"dummyRootCallID",
					dataDir,
				)

				Expect(err).To(BeNil())
			})
		})

		Context("callInterpreter.Interpret result.If falsy", func() {
			It("should emit events", func() {
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

				eventChannel := make(chan model.Event, 1)

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
				Expect(actualEvent.Timestamp).To(BeTemporally("~", time.Now().UTC(), 5*time.Second))
				// set temporal fields to expected vals since they're already asserted
				actualEvent.Timestamp = expectedEvent.Timestamp

				Expect(actualEvent).To(Equal(expectedEvent))
			})
		})

		Context("Container CallSpec", func() {
			It("should call containerCaller.Call w/ expected args", func() {
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

				eventChannel := make(chan model.Event)

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
				Expect(actualErr).To(BeNil())
				_,
					actualContainerCall,
					actualScope,
					actualCallSpec := fakeContainerCaller.CallArgsForCall(0)

				Expect(actualContainerCall).To(Equal(expectedCall.Container))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualCallSpec).To(Equal(providedCallSpec.Container))
			})
		})

		Context("Op CallSpec", func() {
			It("should call opCaller.Call w/ expected args", func() {
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

				eventChannel := make(chan model.Event)

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
				Expect(actualErr).To(BeNil())
				_,
					actualOpCall,
					actualScope,
					actualParentID,
					actualRootCallID,
					actualCallSpec := fakeOpCaller.CallArgsForCall(0)

				Expect(actualOpCall).To(Equal(*expectedCall.Op))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualParentID).To(Equal(providedParentID))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
				Expect(actualCallSpec).To(Equal(providedCallSpec.Op))
			})
		})

		Context("Parallel CallSpec", func() {
			It("should call parallelCaller.Call w/ expected args", func() {
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

				Expect(actualCallID).To(Equal(providedCallID))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
				Expect(actualOpPath).To(Equal(providedOpPath))
				Expect(actualCallSpec).To(Equal(*providedCallSpec.Parallel))
			})
		})

		Context("ParallelLoop CallSpec", func() {
			It("should call parallelLoopCaller.Call w/ expected args", func() {
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

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualParallelLoopCallSpec).To(Equal(*providedCallSpec.ParallelLoop))
				Expect(actualOpPath).To(Equal(providedOpPath))
				Expect(*actualParentID).To(Equal(providedParentID))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
			})
		})

		Context("Serial CallSpec", func() {

			It("should call serialCaller.Call w/ expected args", func() {
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

				Expect(actualCallID).To(Equal(providedCallID))
				Expect(actualScope).To(Equal(providedScope))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
				Expect(actualOpPath).To(Equal(providedOpPath))
				Expect(actualCallSpec).To(Equal(*providedCallSpec.Serial))
			})
		})

		Context("SerialLoop CallSpec", func() {
			It("should call serialLoopCaller.Call w/ expected args", func() {
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

				Expect(actualScope).To(Equal(providedScope))
				Expect(actualSerialLoopCallSpec).To(Equal(*providedCallSpec.SerialLoop))
				Expect(actualOpPath).To(Equal(providedOpPath))
				Expect(*actualParentID).To(Equal(providedParentID))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
			})
		})
	})
})
