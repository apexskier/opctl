package dataresolver

import (
	"context"
	"github.com/pkg/errors"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	nodeFakes "github.com/opctl/opctl/sdks/go/node/fakes"
)

var _ = Context("dataResolver", func() {
	It("Can be constructed", func() {
		Expect(New(new(nodeFakes.FakeNode), "")).NotTo(BeNil())
	})
	Context("Resolve", func() {
		Context("data.Resolve errs", func() {
			Context("not data.ErrAuthenticationFailed", func() {
				It("should return expected error", func() {
					/* arrange */
					providedDataRef := "dummyDataRef"

					fakeCore := new(nodeFakes.FakeNode)
					fakeCore.LabelReturns("opctl node")
					fakeCore.ResolveReturns(nil, errors.New("expectedErr"))

					objectUnderTest := _dataResolver{
						dataProvider: fakeCore,
						basePath:     "/base",
					}

					/* act */
					response, err := objectUnderTest.Resolve(context.Background(), providedDataRef)

					/* assert */
					Expect(response).To(BeNil())

					Expect(err).To(MatchError(`unable to resolve op 'dummyDataRef':` + " " + `
- filesystem:` + " " + `
  - path /base/.opspec/dummyDataRef not found
  - path /base/dummyDataRef not found
- opctl node: expectedErr`))
				})
			})
		})
		Context("data.Resolve doesn't err", func() {
			It("should return expected result", func() {
				/* arrange */
				wd, err := os.Getwd()
				if err != nil {
					panic(err)
				}

				fakeCore := new(nodeFakes.FakeNode)
				providedDataRef := "testdata/dummy-op"

				objectUnderTest := _dataResolver{
					dataProvider: fakeCore,
					basePath:     wd,
				}

				/* act */
				actualPkgHandle, err := objectUnderTest.Resolve(
					context.Background(),
					providedDataRef,
				)

				/* assert */
				Expect(err).To(BeNil())
				Expect(actualPkgHandle.Ref()).To(Equal(path.Join(wd, providedDataRef)))
			})
		})
	})
})
