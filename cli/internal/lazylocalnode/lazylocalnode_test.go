package lazylocalnode

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	nodeFakes "github.com/opctl/opctl/cli/internal/nodeprovider/fakes"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node"
	apiClientFakes "github.com/opctl/opctl/sdks/go/node/api/client/fakes"
)

var _ = Context("lazylocalnode", func() {
	It("conforms to the OpNode interfaces", func() {
		lln := New(new(nodeFakes.FakeNodeProvider))
		test := func(n node.OpNode) {
			Expect(n).NotTo(BeNil())
		}
		test(lln)
	})
	Context("wraps an APIClient, first ensuring the remote node is initialized", func() {
		It("for AddAuth", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), model.AddAuthReq{}

			lln.AddAuth(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.AddAuthArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))
		})
		It("for GetEventStream", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), &model.GetEventStreamReq{}

			lln.GetEventStream(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.GetEventStreamArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))
		})
		It("for KillOp", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), model.KillOpReq{}

			lln.KillOp(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.KillOpArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))
		})
		It("for StartOp", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), model.StartOpReq{}

			lln.StartOp(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.StartOpArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))
		})
		It("for GetData", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), model.GetDataReq{}

			lln.GetData(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.GetDataArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))

		})
		It("for ListDescendants", func() {
			fakeNodeProvider := new(nodeFakes.FakeNodeProvider)
			fakeNodeHandle := new(nodeFakes.FakeNodeHandle)
			fakeAPIClient := new(apiClientFakes.FakeAPIClient)
			fakeNodeHandle.APIClientReturns(fakeAPIClient)
			fakeNodeProvider.CreateNodeIfNotExistsReturns(fakeNodeHandle, nil)
			lln := New(fakeNodeProvider)
			arg1, arg2 := context.Background(), model.ListDescendantsReq{}

			lln.ListDescendants(arg1, arg2)

			Expect(fakeNodeProvider.CreateNodeIfNotExistsCallCount()).To(Equal(1))
			Expect(fakeNodeHandle.APIClientCallCount()).To(Equal(1))
			aArg1, aArg2 := fakeAPIClient.ListDescendantsArgsForCall(0)
			Expect(aArg1).To(Equal(arg1))
			Expect(aArg2).To(Equal(arg2))
		})
	})
})