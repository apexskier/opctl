package call

import (
	"path/filepath"
	"errors"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/container"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op"
)

var _ = Context("Interpret", func() {
	Context("callSpec.If not nil", func() {
		Context("predicates returns err", func() {
			It("should return expected result", func() {
				/* arrange */
				predicateSpec := []*model.PredicateSpec{
					&model.PredicateSpec{},
				}

				/* act */
				_, actualError := Interpret(
					map[string]*model.Value{},
					&model.CallSpec{
						If: &predicateSpec,
					},
					"providedID",
					"dummyOpPath",
					nil,
					"providedRootCallID",
					os.TempDir(),
				)

				/* assert */
				Expect(actualError).To(Equal(errors.New("unable to interpret predicate; predicate was unexpected type &{Eq:<nil> Exists:<nil> Ne:<nil> NotExists:<nil>}")))
			})
		})
	})
	Context("callSpec.Container not nil", func() {
		It("should return expected result", func() {
			/* arrange */
			providedScope := map[string]*model.Value{}
			providedID := "providedID"
			providedOpPath := "providedOpPath"
			providedParentIDValue := "providedParentID"
			providedParentID := &providedParentIDValue
			providedRootCallID := "providedRootCallID"
			providedDataDirPath := os.TempDir()

			containerSpec := model.ContainerCallSpec{
				Image: &model.ContainerCallImageSpec{
					Ref: "ref",
				},
			}

			expectedContainer, err := container.Interpret(
				providedScope,
				&containerSpec,
				providedID,
				providedOpPath,
				providedDataDirPath,
			)
			if nil != err {
				panic(err)
			}

			expectedCall := &model.Call{
				Container: expectedContainer,
				ID:        providedID,
				ParentID:  providedParentID,
				RootID:    providedRootCallID,
			}

			/* act */
			actualCall, actualError := Interpret(
				providedScope,
				&model.CallSpec{
					Container: &containerSpec,
				},
				providedID,
				providedOpPath,
				providedParentID,
				providedRootCallID,
				providedDataDirPath,
			)

			/* assert */
			Expect(actualError).To(BeNil())
			Expect(actualCall).To(Equal(expectedCall))

		})
	})
	Context("callSpec.Op not nil", func() {
		It("should return expected result", func() {
			/* arrange */
			providedScope := map[string]*model.Value{}
			providedID := "providedID"
			providedOpPath := "providedOpPath"
			providedParentIDValue := "providedParentID"
			providedParentID := &providedParentIDValue
			providedRootCallID := "providedRootCallID"
			providedDataDirPath := os.TempDir()

			wd, err := os.Getwd()
			if nil != err {
				panic(err)
			}
			opRef := filepath.Join(wd, "testdata/testop")

			opSpec := model.OpCallSpec{
				Ref: opRef,
			}

			expectedOp, err := op.Interpret(
				providedScope,
				&opSpec,
				providedID,
				providedOpPath,
				providedDataDirPath,
			)
			if nil != err {
				panic(err)
			}

			expectedCall := &model.Call{
				Op:       expectedOp,
				ID:       providedID,
				ParentID: providedParentID,
				RootID:   providedRootCallID,
			}

			/* act */
			actualCall, actualError := Interpret(
				providedScope,
				&model.CallSpec{
					Op: &opSpec,
				},
				providedID,
				providedOpPath,
				providedParentID,
				providedRootCallID,
				providedDataDirPath,
			)

			/* assert */
			Expect(actualError).To(BeNil())
			// ignore Op.ChildCallID since it's a generated UUID
			actualCall.Op.ChildCallID = expectedCall.Op.ChildCallID
			Expect(*actualCall).To(Equal(*expectedCall))

		})
	})
	Context("callSpec.Parallel not empty", func() {
		It("should return expected result", func() {
			/* arrange */
			providedScope := map[string]*model.Value{}
			providedID := "providedID"
			providedOpPath := "providedOpPath"
			providedParentIDValue := "providedParentID"
			providedParentID := &providedParentIDValue
			providedRootCallID := "providedRootCallID"
			providedDataDirPath := os.TempDir()

			parallelSpec := []*model.CallSpec{}

			expectedCall := &model.Call{
				Parallel: parallelSpec,
				ID:       providedID,
				ParentID: providedParentID,
				RootID:   providedRootCallID,
			}

			/* act */
			actualCall, actualError := Interpret(
				providedScope,
				&model.CallSpec{
					Parallel: &parallelSpec,
				},
				providedID,
				providedOpPath,
				providedParentID,
				providedRootCallID,
				providedDataDirPath,
			)

			/* assert */
			Expect(actualError).To(BeNil())
			Expect(actualCall).To(Equal(expectedCall))

		})
	})
	Context("callSpec.Serial not empty", func() {
		It("should return expected result", func() {
			/* arrange */
			providedScope := map[string]*model.Value{}
			providedID := "providedID"
			providedOpPath := "providedOpPath"
			providedParentIDValue := "providedParentID"
			providedParentID := &providedParentIDValue
			providedRootCallID := "providedRootCallID"
			providedDataDirPath := os.TempDir()

			serialSpec := []*model.CallSpec{}

			expectedCall := &model.Call{
				Serial: serialSpec,
				ID:       providedID,
				ParentID: providedParentID,
				RootID:   providedRootCallID,
			}

			/* act */
			actualCall, actualError := Interpret(
				providedScope,
				&model.CallSpec{
					Serial: &serialSpec,
				},
				providedID,
				providedOpPath,
				providedParentID,
				providedRootCallID,
				providedDataDirPath,
			)

			/* assert */
			Expect(actualError).To(BeNil())
			Expect(*actualCall).To(Equal(*expectedCall))

		})
	})
})
