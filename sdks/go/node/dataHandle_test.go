package node

import (
	"context"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/data"
)

var _ = Context("DataHandle Resolution", func() {
	Context("GetContent", func() {
		Context("req.DataRef empty", func() {
			It("should return expected result", func() {
				/* arrange */
				ctx := context.Background()
				c := core{}
				objectUnderTest, err := c.Resolve(ctx, "test")
				Expect(err).To(BeNil())

				/* act */
				actualData, actualErr := objectUnderTest.GetContent(ctx, "")

				/* assert */
				Expect(actualData).To(BeNil())
				Expect(actualErr).To(MatchError(`"" not a valid data ref`))
			})
		})
	})
	Context("ListDescendants", func() {
		Context("req.DataRef empty", func() {
			It("should return expected result", func() {
				/* arrange */
				ctx := context.Background()
				c := core{}
				objectUnderTest, err := c.Resolve(ctx, "")
				Expect(err).To(BeNil())

				/* act */
				actualDescendants, actualErr := objectUnderTest.ListDescendants(
					context.Background(),
				)

				/* assert */
				Expect(actualDescendants).To(BeEmpty())
				Expect(actualErr).To(MatchError(`"" not a valid data ref`))
			})
		})
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		Context("req.DataRef absolute path", func() {
			It("should return expected result", func() {
				/* arrange */
				providedDataRef := path.Join(wd, "testdata/listDescendants")
				ctx := context.Background()
				c := core{}
				objectUnderTest, err := c.Resolve(ctx, providedDataRef)
				Expect(err).To(BeNil())

				expectedDescendants := []*data.DirEntry{
					{Path: "/empty.txt", Size: 0, Mode: 420},
				}

				/* act */
				actualDescendants, actualErr := objectUnderTest.ListDescendants(
					context.Background(),
				)

				/* assert */
				Expect(actualDescendants).To(ConsistOf(expectedDescendants))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
