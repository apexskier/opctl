package node

import (
	"context"
	"errors"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/containerruntime/fakes"
	"io"
)

var _ = Context("containerCaller", func() {
	Context("newContainerCaller", func() {
		It("should return containerCaller", func() {
			/* arrange/act/assert */
			Expect(newContainerCaller(
				new(FakeContainerRuntime),
				false,
			)).To(Not(BeNil()))
		})
	})
	Context("Call", func() {
		It("should call containerRuntime.RunContainer w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedContainerCall := &model.ContainerCall{
				BaseCall: model.BaseCall{},
				Image:    &model.ContainerCallImage{},
			}
			providedRootCallID := "providedRootCallID"
			fakeContainerRuntime := new(FakeContainerRuntime)

			objectUnderTest := _containerCaller{
				containerRuntime: fakeContainerRuntime,
			}

			/* act */
			objectUnderTest.Call(
				providedCtx,
				make(chan model.Event, 10),
				&model.ContainerCall{},
				&model.ContainerCallSpec{},
			)

			/* assert */
			_,
				actualContainerCall,
				actualRootCallID,
				_,
				_,
				_ := fakeContainerRuntime.RunContainerArgsForCall(0)
			Expect(actualContainerCall).To(Equal(providedContainerCall))
			Expect(actualRootCallID).To(Equal(providedRootCallID))
		})
		Context("containerRuntime.RunContainer errors", func() {
			It("should publish expected ContainerExited", func() {
				/* arrange */
				expectedErrorMessage := "expectedErrorMessage"

				fakeContainerRuntime := new(FakeContainerRuntime)

				fakeContainerRuntime.RunContainerStub = func(
					ctx context.Context,
					c chan model.Event,
					req *model.ContainerCall,
					stdOut io.WriteCloser,
					stdErr io.WriteCloser,
					privileged bool,
				) (*int64, error) {

					stdErr.Close()
					stdOut.Close()

					return nil, errors.New(expectedErrorMessage)
				}

				objectUnderTest := _containerCaller{
					containerRuntime: fakeContainerRuntime,
				}

				/* act */
				actualOutputs, actualErr := objectUnderTest.Call(
					context.Background(),
					make(chan model.Event, 10),
					&model.ContainerCall{
						BaseCall: model.BaseCall{},
						Image:    &model.ContainerCallImage{},
					},
					&model.ContainerCallSpec{},
				)

				/* assert */
				Expect(actualOutputs).To(Equal(map[string]*model.Value{}))
				Expect(actualErr).To(MatchError(expectedErrorMessage))
			})
		})
	})

	It("should return expected results", func() {
		/* arrange */
		providedOpPath := "providedOpPath"
		providedContainerCall := &model.ContainerCall{
			BaseCall: model.BaseCall{
				OpPath: providedOpPath,
			},
			ContainerID: "providedContainerID",
			Image:       &model.ContainerCallImage{},
		}
		providedContainerCallSpec := &model.ContainerCallSpec{}

		fakeContainerRuntime := new(FakeContainerRuntime)

		expectedErr := errors.New("io: read/write on closed pipe")
		fakeContainerRuntime.RunContainerStub = func(
			ctx context.Context,
			c chan model.Event,
			req *model.ContainerCall,
			stdOut io.WriteCloser,
			stdErr io.WriteCloser,
			privileged bool,
		) (*int64, error) {

			stdErr.Close()
			stdOut.Close()

			return nil, expectedErr
		}

		objectUnderTest := _containerCaller{
			containerRuntime: new(FakeContainerRuntime),
		}

		/* act */
		actualOutputs, actualErr := objectUnderTest.Call(
			context.Background(),
			make(chan model.Event, 10),
			providedContainerCall,
			providedContainerCallSpec,
		)

		/* assert */
		Expect(actualOutputs).To(Equal(map[string]*model.Value{}))
		Expect(actualErr).To(Equal(expectedErr))
	})
})
