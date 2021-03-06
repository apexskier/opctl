package core

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uniquestringFakes "github.com/opctl/opctl/sdks/go/internal/uniquestring/fakes"
	"github.com/opctl/opctl/sdks/go/model"
	. "github.com/opctl/opctl/sdks/go/node/core/internal/fakes"
	. "github.com/opctl/opctl/sdks/go/pubsub/fakes"
)

var _ = Context("core", func() {
	Context("StartOp", func() {
		Context("data.Resolve errs", func() {
			It("should return expected result", func() {

				/* arrange */
				providedCtx := context.Background()
				providedStartOpReq := model.StartOpReq{
					Op: model.StartOpReqOp{
						Ref: "dummyOpRef",
					},
				}

				objectUnderTest := _core{}

				/* act */
				_, actualErr := objectUnderTest.StartOp(
					providedCtx,
					providedStartOpReq,
				)

				/* assert */
				Expect(actualErr.Error()).To(Equal(`Get "https://dummyOpRef/info/refs?service=git-upload-pack": dial tcp: lookup dummyOpRef on 127.0.0.11:53: no such host`))
			})
		})
		Context("data.Resolve doesn't err", func() {
			Context("opfile.Get doesn't err", func() {
				It("should call caller.Call w/ expected args", func() {
					/* arrange */
					providedCtx := context.Background()
					providedArg1String := "dummyArg1Value"
					providedArg2Dir := "/"
					providedArg3Dir := "/"
					providedArg4Dir := "/"

					// use local op
					wd, err := os.Getwd()
					if nil != err {
						panic(err)
					}
					providedOpPath := filepath.Join(wd, "testdata/startOp")
					providedReq := model.StartOpReq{
						Args: map[string]*model.Value{
							"dummyArg1Name": {String: &providedArg1String},
							"dummyArg2Name": {Dir: &providedArg2Dir},
							"dummyArg3Name": {Dir: &providedArg3Dir},
							"dummyArg4Name": {Dir: &providedArg4Dir},
						},
						Op: model.StartOpReqOp{
							Ref: providedOpPath,
						},
					}

					opFile := &model.OpSpec{
						Outputs: map[string]*model.Param{
							"dummyOutput1": {String: &model.StringParam{}},
							"dummyOutput2": {String: &model.StringParam{}},
						},
					}

					expectedOpCallSpec := &model.OpCallSpec{
						Ref:     providedOpPath,
						Inputs:  map[string]interface{}{},
						Outputs: map[string]string{},
					}
					for name := range providedReq.Args {
						expectedOpCallSpec.Inputs[name] = ""
					}
					for name := range opFile.Outputs {
						expectedOpCallSpec.Outputs[name] = ""
					}

					expectedID := "expectedID"
					fakeUniqueStringFactory := new(uniquestringFakes.FakeUniqueStringFactory)
					fakeUniqueStringFactory.ConstructReturns(expectedID, nil)

					fakeCaller := new(FakeCaller)
					dataCachePath := os.TempDir()

					objectUnderTest := _core{
						caller:              fakeCaller,
						dataCachePath:       dataCachePath,
						pubSub:              new(FakePubSub),
						uniqueStringFactory: fakeUniqueStringFactory,
					}

					/* act */
					objectUnderTest.StartOp(
						providedCtx,
						providedReq,
					)

					/* assert */
					// Call happens in go routine; wait 500ms to allow it to occur
					time.Sleep(time.Millisecond * 500)
					_,
						actualOpID,
						actualScope,
						actualCallSpec,
						actualOpPath,
						_,
						actualRootID := fakeCaller.CallArgsForCall(0)

					Expect(actualOpID).To(Equal(expectedID))
					Expect(actualScope).To(Equal(providedReq.Args))
					Expect(*actualCallSpec).To(BeEquivalentTo(model.CallSpec{
						Op: expectedOpCallSpec,
					}))
					Expect(actualOpPath).To(Equal(providedOpPath))
					Expect(actualRootID).To(Equal(expectedID))
				})
			})
		})
	})
})
