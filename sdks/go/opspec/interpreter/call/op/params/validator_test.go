package params

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

var _ = Context("Validate", func() {
	It("should return expected result", func() {
		/* arrange */

		providedValues := map[string]*model.Value{
			"expectedName1": {
				String: new(string),
			},
		}

		providedParams := map[string]*model.ParamSpec{
			"expectedName1": {
				String: &model.StringParamSpec{
					Constraints: map[string]interface{}{
						"minLength": 10,
					},
				},
			},
		}

		/* act */
		actualErr := Validate(
			providedValues,
			providedParams,
		)

		/* assert */
		Expect(actualErr).To(MatchError("validation error(s):\n- expectedName1: String length must be greater than or equal to 10"))
	})
})
