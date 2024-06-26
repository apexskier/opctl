package opfile

import (
	"encoding/json"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Unmarshal", func() {
	Context("Validate returns errors", func() {
		It("should return the expected error", func() {
			/* arrange */
			/* act */
			_, actualError := Unmarshal("opRef", []byte("&"))

			/* assert */
			Expect(actualError).To(MatchError("opspec syntax error:\nopRef\n- error converting YAML to JSON: yaml: did not find expected alphabetic or numeric character"))
		})
	})
	Context("Validator.Validate doesn't return errors", func() {

		It("should return expected opFile", func() {

			/* arrange */
			paramDefault := "dummyDefault"
			dummyParams := map[string]*model.ParamSpec{
				"dummyName": {
					String: &model.StringParamSpec{
						Constraints: map[string]interface{}{
							"minLength": 0,
							"maxLength": 1000,
							"pattern":   "dummyPattern",
							"format":    "date-time",
							"enum":      []interface{}{"dummyEnumItem1"},
						},
						Default:     &paramDefault,
						Description: "dummyDescription",
						IsSecret:    true,
					},
				},
			}

			expectedOpFile := model.OpSpec{
				Description: "dummyDescription",
				Inputs:      dummyParams,
				Outputs:     dummyParams,
				Run: &model.CallSpec{
					Op: &model.OpCallSpec{
						Ref: "dummyOpRef",
					},
				},
			}
			providedBytes, err := yaml.Marshal(&expectedOpFile)
			if err != nil {
				panic(err.Error())
			}

			/* act */
			actualOpFile, actualErr := Unmarshal(providedBytes)

			/* assert */
			Expect(actualErr).To(BeNil())

			// compare as JSON; otherwise we encounter pointer inequalities
			actualBytes, err := json.Marshal(actualOpFile)
			if err != nil {
				panic(err)
			}

			expectedBytes, err := json.Marshal(expectedOpFile)
			if err != nil {
				panic(err)
			}
			Expect(string(actualBytes)).To(Equal(string(expectedBytes)))
		})
	})
})
