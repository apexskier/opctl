package node

import (
	"context"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("core", func() {
	Context("resolveData", func() {
		It("should call data.Resolve w/ expected args", func() {
			/* arrange */
			gitOpsDir, err := os.MkdirTemp("", "")
			if err != nil {
				panic(err)
			}

			providedCtx := context.Background()
			// some public repo that's relatively small
			providedOpRef := "github.com/opspec-pkgs/_.op.create#3.3.1"

			objectUnderTest := core{
				gitOpsDir: gitOpsDir,
			}

			/* act */
			actualOp, actualErr := objectUnderTest.resolveData(
				providedCtx,
				providedOpRef,
			)

			/* assert */
			Expect(actualErr).To(BeNil())
			Expect(*actualOp.Path()).To(Equal(filepath.Join(gitOpsDir, providedOpRef)))
		})
	})
})
