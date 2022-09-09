package node

import (
	"context"
	"io"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	containerRuntimeFakes "github.com/opctl/opctl/sdks/go/node/containerruntime/fakes"
	. "github.com/opctl/opctl/sdks/go/node/internal/fakes"
)

var _ = Context("parallelCaller", func() {
	Context("newParallelCaller", func() {
		It("should return parallelCaller", func() {
			/* arrange/act/assert */
			Expect(newParallelCaller(
				new(FakeCaller),
			)).To(Not(BeNil()))
		})
	})
	Context("Call", func() {

		Context("caller errors", func() {

			It("should return expected results", func() {
				/* arrange */
				dbDir, err := os.MkdirTemp("", "")
				if err != nil {
					panic(err)
				}

				objectUnderTest := _parallelCaller{
					caller: newCaller(
						newContainerCaller(
							new(containerRuntimeFakes.FakeContainerRuntime),
							false,
						),
						dbDir,
					),
				}

				/* act */
				_, actualErr := objectUnderTest.Call(
					context.Background(),
					"callID",
					map[string]*model.Value{},
					"rootCallID",
					"opPath",
					[]*model.CallSpec{
						{
							// intentionally invalid
							Container: &model.ContainerCallSpec{},
						},
					},
				)

				/* assert */
				Expect(actualErr.Error()).To(Equal("child call failed"))
			})
		})

		It("should start each child as expected", func() {

			/* arrange */
			dbDir, err := os.MkdirTemp("", "")
			if err != nil {
				panic(err)
			}

			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			providedOpRef := "providedOpRef"
			providedParentID := "providedParentID"
			providedRootID := "providedRootID"
			childOpRef := filepath.Join(wd, "testdata/parallelCaller")
			input1Key := "input1"
			childOp1Path := filepath.Join(childOpRef, "op1")
			childOp2Path := filepath.Join(childOpRef, "op2")

			ctx := context.Background()

			fakeContainerRuntime := new(containerRuntimeFakes.FakeContainerRuntime)
			fakeContainerRuntime.RunContainerStub = func(
				ctx context.Context,
				eventChannel chan model.Event,
				req *model.ContainerCall,
				rootCallID string,
				stdOut io.WriteCloser,
				stdErr io.WriteCloser,
			) (*int64, error) {

				stdErr.Close()
				stdOut.Close()

				return nil, nil
			}

			input1Value := "input1Value"
			providedInboundScope := map[string]*model.Value{
				input1Key: {String: &input1Value},
			}

			objectUnderTest := _parallelCaller{
				caller: newCaller(
					newContainerCaller(
						fakeContainerRuntime,
					),
					dbDir,
				),
			}

			/* act */
			_, actualErr := objectUnderTest.Call(
				ctx,
				providedParentID,
				providedInboundScope,
				providedRootID,
				providedOpRef,
				[]*model.CallSpec{
					{
						Op: &model.OpCallSpec{
							Ref: childOp1Path,
							Inputs: map[string]interface{}{
								input1Key: nil,
							},
						},
					},
					{
						Op: &model.OpCallSpec{
							Ref: childOp2Path,
							Inputs: map[string]interface{}{
								input1Key: nil,
							},
						},
					},
				},
			)

			/* assert */
			Expect(actualErr).To(BeNil())

			actualChildCalls := []model.CallStarted{}
			go func() {
				for event := range eventChannel {
					if event.CallStarted != nil && event.CallStarted.Call.Op != nil {
						// ignore props we can't readily assert
						event.CallStarted.Call.Op.ChildCallCallSpec = nil
						event.CallStarted.Call.Op.ChildCallID = ""
						event.CallStarted.Call.Op.OpID = ""
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
								Op: &model.OpCall{
									BaseCall: model.BaseCall{
										OpPath: childOp1Path,
									},
									Inputs:            providedInboundScope,
									ChildCallCallSpec: nil,
								},
								ParentID: &providedParentID,
								RootID:   providedRootID,
							},
							Ref: providedOpRef,
						},
						{
							Call: model.Call{
								Op: &model.OpCall{
									BaseCall: model.BaseCall{
										OpPath: childOp2Path,
									},
									Inputs:            providedInboundScope,
									ChildCallCallSpec: nil,
								},
								ParentID: &providedParentID,
								RootID:   providedRootID,
							},
							Ref: providedOpRef,
						},
					},
				),
			)
		})
	})
})
