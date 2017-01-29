package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/opctl/pkg/containerengine/engines/fake"
	"github.com/opspec-io/opctl/util/eventbus"
	"github.com/opspec-io/opctl/util/pathnormalizer"
	"github.com/opspec-io/opctl/util/uniquestring"
	"github.com/opspec-io/sdk-golang/pkg/model"
	"time"
)

var _ = Context("core", func() {
	Context("StartOp", func() {
		It("should call opCaller.Call w/ expected args", func() {
			/* arrange */
			providedReq := model.StartOpReq{
				Args: map[string]*model.Data{
					"dummyArg1Name": {String: "dummyArg1Value"},
					"dummyArg2Name": {Dir: "dummyArg2Value"},
					"dummyArg3Name": {Dir: "dummyArg3Value"},
					"dummyArg4Name": {Dir: "dummyArg4Value"},
				},
				OpRef: "dummyOpRef",
			}

			expectedOpId := "dummyOpId"

			fakeOpCaller := new(fakeOpCaller)

			fakeUniqueStringFactory := new(uniquestring.Fake)
			fakeUniqueStringFactory.ConstructReturns(expectedOpId)

			objectUnderTest := _core{
				containerEngine:     new(fake.ContainerEngine),
				eventBus:            new(eventbus.Fake),
				opCaller:            fakeOpCaller,
				pathNormalizer:      pathnormalizer.NewPathNormalizer(),
				dcgNodeRepo:         new(fakeDcgNodeRepo),
				uniqueStringFactory: fakeUniqueStringFactory,
			}

			/* act */
			objectUnderTest.StartOp(providedReq)

			/* assert */
			// Call happens in go routine; wait 500ms to allow it to occur
			time.Sleep(time.Millisecond * 500)
			inboundScope,
				opId,
				opRef,
				opGraphId := fakeOpCaller.CallArgsForCall(0)

			Expect(inboundScope).To(Equal(providedReq.Args))
			Expect(opId).To(Equal(expectedOpId))
			Expect(opRef).To(Equal(providedReq.OpRef))
			Expect(opGraphId).To(Equal(opId))
		})
	})
})
