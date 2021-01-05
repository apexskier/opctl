package cliparamsatisfier

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	clioutputFakes "github.com/opctl/opctl/cli/internal/clioutput/fakes"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Describe("inputSrcFactory", func() {
	wd, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	argsYmlTestDataPath := filepath.Join(wd, "inputsrc/ymlfile/testdata/args.yml")
	Context("NewCLIPromptInputSrc()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(_inputSrcFactory{}.NewCliPromptInputSrc(new(clioutputFakes.FakeCliOutput), nil)).To(Not(BeNil()))
		})
	})
	Context("NewEnvVarInputSrc()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(_inputSrcFactory{}.NewEnvVarInputSrc()).To(Not(BeNil()))
		})
	})
	Context("NewParamDefaultInputSrc()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(_inputSrcFactory{}.NewParamDefaultInputSrc(
				map[string]*model.Param{},
			)).To(Not(BeNil()))
		})
	})
	Context("NewSliceInputSrc()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(_inputSrcFactory{}.NewSliceInputSrc([]string{}, "")).To(Not(BeNil()))
		})
	})
	Context("NewYMLFileInputSrc()", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(_inputSrcFactory{}.NewYMLFileInputSrc(argsYmlTestDataPath)).To(Not(BeNil()))
		})
	})
})
