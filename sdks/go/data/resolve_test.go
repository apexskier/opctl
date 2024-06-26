package data

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/data/fs"
	aggregateError "github.com/opctl/opctl/sdks/go/internal/aggregate_error"
)

var _ = Context("Resolve", func() {
	Context("providers[0].Resolve errs", func() {
		It("should return error", func() {
			/* arrange */
			provider0 := fs.New()
			providedProviders := []DataProvider{provider0}
			dataRef := "\\not/exist"

			/* act */
			_, actualErr := Resolve(
				context.Background(),
				dataRef,
				providedProviders...,
			)

			/* assert */
			var expected aggregateError.ErrAggregate
			expected.AddError(fmt.Errorf("%s: %w", provider0.Label(), errors.New("skipped")))
			Expect(actualErr).To(MatchError(fmt.Errorf("unable to resolve op '\\not/exist': %w", expected)))
		})
	})
	Context("providers[0].Resolve doesn't err", func() {
		It("should return expected results", func() {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			opRef := filepath.Join(wd, "testdata/testop")
			provider0 := fs.New(filepath.Dir(opRef))

			providedProviders := []DataProvider{provider0}

			/* act */
			actualHandle, actualErr := Resolve(
				context.Background(),
				opRef,
				providedProviders...,
			)

			/* assert */
			Expect(actualErr).To(BeNil())
			Expect(actualHandle.Ref()).To(Equal(opRef))
		})
	})
})
