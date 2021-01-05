package core

import (
	"context"
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uniquestringFakes "github.com/opctl/opctl/sdks/go/internal/uniquestring/fakes"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/core/internal/fakes"
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
		It("should call caller for every serialCall w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedCallID := "providedCallID"
			providedInboundScope := map[string]*model.Value{}
			providedRootCallID := "providedRootCallID"
			providedOpPath := "providedOpPath"
			providedCallSpecSerialCalls := []*model.CallSpec{
				{
					Container: &model.ContainerCallSpec{},
				},
				{
					Op: &model.OpCallSpec{},
				},
				{
					Parallel: &[]*model.CallSpec{},
				},
				{
					Serial: &[]*model.CallSpec{},
				},
			}

			fakeCaller := new(FakeCaller)

			fakeUniqueStringFactory := new(uniquestringFakes.FakeUniqueStringFactory)
			uniqueStringCallIndex := 0
			fakeUniqueStringFactory.ConstructStub = func() (string, error) {
				defer func() {
					uniqueStringCallIndex++
				}()
				return fmt.Sprintf("%v", uniqueStringCallIndex), nil
			}

			objectUnderTest := _serialCaller{
				caller:              fakeCaller,
				uniqueStringFactory: fakeUniqueStringFactory,
			}

			/* act */
			objectUnderTest.Call(
				providedCtx,
				providedCallID,
				providedInboundScope,
				providedRootCallID,
				providedOpPath,
				providedCallSpecSerialCalls,
			)

			/* assert */
			for expectedCallSpecIndex, expectedCallSpec := range providedCallSpecSerialCalls {
				actualCtx,
					actualNodeID,
					actualChildOutboundScope,
					actualCallSpec,
					actualOpPath,
					actualParentCallID,
					actualRootCallID := fakeCaller.CallArgsForCall(expectedCallSpecIndex)

				Expect(actualCtx).To(Equal(actualCtx))
				Expect(actualNodeID).To(Equal(fmt.Sprintf("%v", expectedCallSpecIndex)))
				Expect(actualChildOutboundScope).To(Equal(providedInboundScope))
				Expect(actualCallSpec).To(Equal(expectedCallSpec))
				Expect(actualOpPath).To(Equal(providedOpPath))
				Expect(actualParentCallID).To(Equal(&providedCallID))
				Expect(actualRootCallID).To(Equal(providedRootCallID))
			}
		})
		Context("caller errors", func() {
			It("should return expected results", func() {
				/* arrange */
				providedCallID := "dummyCallID"
				providedInboundScope := map[string]*model.Value{}
				providedRootCallID := "dummyRootCallID"
				providedCallSpecSerialCalls := []*model.CallSpec{
					{
						Container: &model.ContainerCallSpec{},
					},
				}

				callID := "callID"

				fakeUniqueStringFactory := new(uniquestringFakes.FakeUniqueStringFactory)
				fakeUniqueStringFactory.ConstructReturns(callID, nil)

				expectedErr := errors.New("expectedErr")

				fakeCaller := new(FakeCaller)
				fakeCaller.CallReturns(nil, expectedErr)

				objectUnderTest := _serialCaller{
					caller:              fakeCaller,
					uniqueStringFactory: fakeUniqueStringFactory,
				}

				/* act */
				actualOutputs, actualErr := objectUnderTest.Call(
					context.Background(),
					providedCallID,
					providedInboundScope,
					providedRootCallID,
					"dummyOpPath",
					providedCallSpecSerialCalls,
				)

				/* assert */
				Expect(actualOutputs).To(BeNil())
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("caller doesn't error", func() {
			Context("childOutboundScope empty", func() {
				It("should call secondChild w/ inboundScope", func() {
					/* arrange */
					providedCallID := "dummyCallID"
					providedScopeName1String := "dummyParentVar1Data"
					providedScopeName2Dir := "dummyParentVar2Data"
					providedInboundScope := map[string]*model.Value{
						"dummyVar1Name": {String: &providedScopeName1String},
						"dummyVar2Name": {Dir: &providedScopeName2Dir},
					}
					expectedInboundScopeToSecondChild := providedInboundScope
					providedRootCallID := "dummyRootCallID"
					providedOpPath := "providedOpPath"
					providedCallSpecSerialCalls := []*model.CallSpec{
						{
							Container: &model.ContainerCallSpec{},
						},
						{
							Container: &model.ContainerCallSpec{},
						},
					}

					fakeCaller := new(FakeCaller)

					fakeUniqueStringFactory := new(uniquestringFakes.FakeUniqueStringFactory)
					uniqueStringCallIndex := 0
					fakeUniqueStringFactory.ConstructStub = func() (string, error) {
						defer func() {
							uniqueStringCallIndex++
						}()
						return fmt.Sprintf("%v", uniqueStringCallIndex), nil
					}

					objectUnderTest := _serialCaller{
						caller:              fakeCaller,
						uniqueStringFactory: fakeUniqueStringFactory,
					}

					/* act */
					objectUnderTest.Call(
						context.Background(),
						providedCallID,
						providedInboundScope,
						providedRootCallID,
						providedOpPath,
						providedCallSpecSerialCalls,
					)

					/* assert */
					_, _, actualInboundScopeToSecondChild, _, _, _, _ := fakeCaller.CallArgsForCall(1)
					Expect(actualInboundScopeToSecondChild).To(Equal(expectedInboundScopeToSecondChild))
				})
			})
			Context("childOutboundScope not empty", func() {
				It("should call secondChild w/ firstChildOutputs overlaying inboundScope", func() {
					/* arrange */
					providedCallID := "dummyCallID"

					providedInboundVar1String := "dummyParentVar1Data"
					providedInboundVar2Dir := "dummyParentVar2Data"
					providedInboundVar3File := "dummyParentVar3Data"
					providedInboundScope := map[string]*model.Value{
						"dummyVar1Name": {String: &providedInboundVar1String},
						"dummyVar2Name": {Dir: &providedInboundVar2Dir},
						"dummyVar3Name": {File: &providedInboundVar3File},
					}

					firstChildOutput1String := "dummyFirstChildVar1Data"
					firstChildOutput2String := "dummyFirstChildVar2Data"
					firstChildOutputs := map[string]*model.Value{
						"dummyVar1Name": {String: &firstChildOutput1String},
						"dummyVar2Name": {Dir: &firstChildOutput2String},
					}

					expectedInboundScopeToSecondChild := map[string]*model.Value{
						"dummyVar1Name": firstChildOutputs["dummyVar1Name"],
						"dummyVar2Name": firstChildOutputs["dummyVar2Name"],
						"dummyVar3Name": providedInboundScope["dummyVar3Name"],
					}
					providedRootCallID := "dummyRootCallID"
					providedCallSpecSerialCalls := []*model.CallSpec{
						{
							Container: &model.ContainerCallSpec{},
						},
						{
							Container: &model.ContainerCallSpec{},
						},
					}

					fakeCaller := new(FakeCaller)
					fakeCaller.CallReturnsOnCall(0, firstChildOutputs, nil)

					fakeUniqueStringFactory := new(uniquestringFakes.FakeUniqueStringFactory)
					uniqueStringCallIndex := 0
					fakeUniqueStringFactory.ConstructStub = func() (string, error) {
						defer func() {
							uniqueStringCallIndex++
						}()
						return fmt.Sprintf("%v", uniqueStringCallIndex), nil
					}

					objectUnderTest := _serialCaller{
						caller:              fakeCaller,
						uniqueStringFactory: fakeUniqueStringFactory,
					}

					/* act */
					outputs, err := objectUnderTest.Call(
						context.Background(),
						providedCallID,
						providedInboundScope,
						providedRootCallID,
						"dummyOpPath",
						providedCallSpecSerialCalls,
					)

					/* assert */
					Expect(err).To(BeNil())
					_, _, actualInboundScopeToSecondChild, _, _, _, _ := fakeCaller.CallArgsForCall(1)
					Expect(actualInboundScopeToSecondChild).To(Equal(expectedInboundScopeToSecondChild))

					Expect(outputs).To(Equal(map[string]*model.Value{
						"dummyVar1Name": {String: &firstChildOutput1String},
						"dummyVar2Name": {Dir: &firstChildOutput2String},
						"dummyVar3Name": {
							Array:   nil,
							Boolean: nil,
							Dir:     nil,
							File:    &providedInboundVar3File,
							Number:  nil,
							Object:  nil,
							Socket:  nil,
							String:  nil,
						},
					}))
				})
			})
		})
	})
})
