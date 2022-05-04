package node

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/opctl/opctl/sdks/go/node/containerruntime/fakes"
)

var _ = Context("core", func() {
	Context("New", func() {
		It("should return Core", func() {
			/* arrange */
			dataDir, err := os.MkdirTemp("", "")
			if err != nil {
				panic(err)
			}

			/* act/assert */
			Expect(
				New(
					new(FakeContainerRuntime),
					dataDir,
					false,
				),
			).To(Not(BeNil()))
		})
	})
})
