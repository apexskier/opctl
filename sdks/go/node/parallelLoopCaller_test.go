package node

import (
	"context"
	"io"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	containerRuntimeFakes "github.com/opctl/opctl/sdks/go/node/containerruntime/fakes"
	. "github.com/opctl/opctl/sdks/go/node/internal/fakes"
)

var _ = Context("parallelLoopCaller", func() {
	Context("newParallelLoopCaller", func() {
		It("should return parallelLoopCaller", func() {
			/* arrange/act/assert */
			Expect(newParallelLoopCaller(
				new(FakeCaller),
			)).To(Not(BeNil()))
		})
	})

	Context("Call", func() {
		Context("initial callParallelLoop.Range empty", func() {
			It("should not call caller.Call", func() {
				/* arrange */
				fakeCaller := new(FakeCaller)

				_, _ = _parallelLoopCaller{
					caller: fakeCaller,
				}.Call(
					context.Background(),
					make(chan model.Event, 10),
					map[string]*model.Value{},
					model.ParallelLoopCallSpec{
						Range: []interface{}{},
					},
					"dummyOpPath",
					nil,
					"rootCallID",
					"",
				)

				/* assert */
				Expect(fakeCaller.CallCallCount()).To(Equal(0))
			})
		})

		Context("iteration spec invalid", func() {

			It("should return expected results", func() {
				/* arrange */
				dbDir, err := os.MkdirTemp("", "")
				if err != nil {
					panic(err)
				}

				providedCtx := context.Background()
				providedScope := map[string]*model.Value{}

				caller := newCaller(
					newContainerCaller(
						new(containerRuntimeFakes.FakeContainerRuntime),
						false,
					),
					dbDir,
				)

				objectUnderTest := _parallelLoopCaller{
					caller: caller,
				}

				/* act */
				actualOutputs, actualErr := objectUnderTest.Call(
					providedCtx,
					make(chan model.Event, 10),
					providedScope,
					model.ParallelLoopCallSpec{
						Range: model.Value{
							Array: &[]interface{}{0},
						},
						Run: model.CallSpec{
							Container: &model.ContainerCallSpec{},
						},
					},
					"opPath",
					new(string),
					"rootCallID",
					"",
				)

				/* assert */
				Expect(actualErr.Error()).To(Equal("child call failed"))
				Expect(actualOutputs).To(BeNil())
			})
		})

		It("should start each child as expected", func() {
			/* arrange */
			dbDir, err := os.MkdirTemp("", "")
			if err != nil {
				panic(err)
			}

			providedOpRef := "providedOpRef"
			providedParentID := "providedParentID"
			providedRootID := "providedRootID"
			imageRef := "docker.io/library/alpine"
			eventChannel := make(chan model.Event, 10)

			ctx := context.Background()

			fakeContainerRuntime := new(containerRuntimeFakes.FakeContainerRuntime)
			fakeContainerRuntime.RunContainerStub = func(
				ctx context.Context,
				eventChannel chan model.Event,
				req *model.ContainerCall,
				stdOut io.WriteCloser,
				stdErr io.WriteCloser,
				privileged bool,
			) (*int64, error) {

				stdErr.Close()
				stdOut.Close()

				return nil, nil
			}

			objectUnderTest := _parallelLoopCaller{
				caller: newCaller(
					newContainerCaller(
						fakeContainerRuntime,
						false,
					),
					dbDir,
				),
			}

			/* act */
			_, actualErr := objectUnderTest.Call(
				ctx,
				eventChannel,
				map[string]*model.Value{},
				model.ParallelLoopCallSpec{
					Range: model.Value{
						Array: &[]interface{}{0, 1},
					},
					Run: model.CallSpec{
						Container: &model.ContainerCallSpec{
							Image: &model.ContainerCallImageSpec{
								Ref: imageRef,
							},
						},
					},
				},
				providedOpRef,
				&providedParentID,
				providedRootID,
				"",
			)

			/* assert */
			Expect(actualErr).To(BeNil())

			actualChildCalls := []model.CallStarted{}
			go func() {
				for event := range eventChannel {
					if event.CallStarted != nil && event.CallStarted.Call.Container != nil {
						// ignore props we can't readily assert
						event.CallStarted.Call.Container.ContainerID = ""
						event.CallStarted.Call.ID = ""

						actualChildCalls = append(actualChildCalls, *event.CallStarted)
					}
				}
			}()

			Eventually(
				func() []model.CallStarted { return actualChildCalls },
			).Should(
				ContainElements(
					[]model.CallStarted{
						{
							Call: model.Call{
								Container: &model.ContainerCall{
									BaseCall: model.BaseCall{
										OpPath: providedOpRef,
									},
									Cmd:   []string{},
									Dirs:  map[string]string{},
									Files: map[string]string{},
									Image: &model.ContainerCallImage{
										Ref: &imageRef,
									},
									Sockets: map[string]string{},
								},
								ParentID: &providedParentID,
								RootID:   providedRootID,
							},
							OpRef: providedOpRef,
						},
						{
							Call: model.Call{
								Container: &model.ContainerCall{
									BaseCall: model.BaseCall{
										OpPath: providedOpRef,
									},
									Cmd:   []string{},
									Dirs:  map[string]string{},
									Files: map[string]string{},
									Image: &model.ContainerCallImage{
										Ref: &imageRef,
									},
									Sockets: map[string]string{},
								},
								ParentID: &providedParentID,
								RootID:   providedRootID,
							},
							OpRef: providedOpRef,
						},
					},
				),
			)
		})
	})
})
