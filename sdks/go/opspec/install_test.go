package opspec

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dataFakes "github.com/opctl/opctl/sdks/go/data/fakes"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Install", func() {
	providedCtx := context.Background()

	It("should call handle.ListDescendants w/ expected args", func() {
		/* arrange */
		fakeHandle := new(dataFakes.FakeDataHandle)

		/* act */
		Install(providedCtx, "", fakeHandle)

		/* assert */
		Expect(fakeHandle.ListDescendantsArgsForCall(0)).To(Equal(providedCtx))
	})
	Context("handle.ListDescendants errs", func() {
		It("should return error", func() {
			/* arrange */
			expectedError := errors.New("dummyError")

			fakeHandle := new(dataFakes.FakeDataHandle)
			fakeHandle.ListDescendantsReturns(nil, expectedError)

			/* act */
			actualError := Install(providedCtx, "", fakeHandle)

			/* assert */
			Expect(actualError).To(MatchError(expectedError))
		})
	})
	Context("handle.ListDescendants doesn't err", func() {
		It("should call handle.GetContent w/ expected args", func() {
			/* arrange */
			fakeHandle := new(dataFakes.FakeDataHandle)
			contentsList := []*model.DirEntry{
				{
					Path: "dirEntry1PathGetContents",
				},
			}

			fakeHandle.ListDescendantsReturns(
				contentsList,
				nil,
			)

			dataDir, err := os.MkdirTemp("", "")
			if err != nil {
				panic(err)
			}

			// error to trigger immediate return
			fakeHandle.GetContentReturns(nil, expectedErr)

			/* act */
			err = Install(providedCtx, dataDir, fakeHandle)

			/* assert */
			actualContext,
				actualPath := fakeHandle.GetContentArgsForCall(0)
			Expect(err).To(MatchError(expectedErr))
			Expect(actualContext).To(Equal(providedCtx))
			Expect(actualPath).To(Equal(contentsList[0].Path))
		})
		Context("handle.GetContent errs", func() {
			It("should return error", func() {
				/* arrange */
				expectedError := errors.New("dummyError")

				fakeHandle := new(dataFakes.FakeDataHandle)
				fakeHandle.ListDescendantsReturns([]*model.DirEntry{{}}, expectedError)

				fakeHandle.GetContentReturns(nil, expectedError)

				/* act */
				actualError := Install(providedCtx, "", fakeHandle)

				/* assert */
				Expect(actualError).To(MatchError(expectedError))
			})
		})
		Context("handle.GetContent doesn't err", func() {
			Context("os.MkdirAll doesn't err", func() {
				Context("os.Create doesn't err", func() {
					Context("os.Chmod doesn't err", func() {
						It("should copy content", func() {
							/* arrange */
							fsDataSource := fs.New("")
							ref := "testdata/testop"
							handle, err := fsDataSource.Resolve(providedCtx, ref)
							if err != nil {
								panic(err)
							}

							expectedContent, err := os.ReadFile(filepath.Join(ref, "op.yml"))
							if err != nil {
								panic(err)
							}

							// create tmpfile to use as dst
							tmpDir, err := os.MkdirTemp("", "")
							if err != nil {
								panic(err)
							}

							/* act */
							Install(providedCtx, tmpDir, handle)

							/* assert */
							actualContent, err := os.ReadFile(filepath.Join(tmpDir, "op.yml"))
							if err != nil {
								panic(err)
							}

							Expect(actualContent).To(Equal(expectedContent))
						})
					})
				})
			})
		})
	})
})
