package node

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	containerRuntimeFakes "github.com/opctl/opctl/sdks/go/node/containerruntime/fakes"
	. "github.com/opctl/opctl/sdks/go/node/internal/fakes"
)

var _ = Context("serialCaller", func() {
	Context("newSerialCaller", func() {
		It("should return serialCaller", func() {
			/* arrange/act/assert */
			Expect(newSerialCaller(
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

				expectedErr := errors.New("expectedErr")

				fakeCaller := new(FakeCaller)
				fakeCaller.CallReturns(nil, expectedErr)

				objectUnderTest := _serialCaller{
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
					make(chan model.Event, 10),
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
					"",
				)

				/* assert */
				Expect(actualErr).To(MatchError("image required"))
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
			childOpRef := filepath.Join(wd, "testdata/serialCaller")
			input1Key := "input1"
			childOp1Path := filepath.Join(childOpRef, "op1")
			childOp2Path := filepath.Join(childOpRef, "op2")
			input2Key := "input2"
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

			input1Value := "input1Value"
			providedInboundScope := map[string]*model.Value{
				input1Key: {String: &input1Value},
			}

			input2Value := "input2Value"

			objectUnderTest := _serialCaller{
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
							Outputs: map[string]string{
								input2Key: "",
							},
						},
					},
					{
						Op: &model.OpCallSpec{
							Ref: childOp2Path,
							Inputs: map[string]interface{}{
								input2Key: nil,
							},
						},
					},
				},
				"",
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
							OpRef: providedOpRef,
						},
						{
							Call: model.Call{
								Op: &model.OpCall{
									BaseCall: model.BaseCall{
										OpPath: childOp2Path,
									},
									Inputs: map[string]*model.Value{
										input2Key: {String: &input2Value},
									},
									ChildCallCallSpec: nil,
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
