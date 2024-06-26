package params

import (
	"bytes"
	"fmt"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/opspec/interpreter/call/op/params/param"
)

// Validate validates values for/against params
func Validate(
	values map[string]*model.Value,
	params map[string]*model.ParamSpec,
) error {

	paramErrMap := map[string][]error{}
	for paramName, paramValue := range params {
		errs := param.Validate(values[paramName], paramValue)
		if len(errs) > 0 {
			fmt.Println("params.validator", paramName, paramValue, errs)
			paramErrMap[paramName] = errs
		}
	}

	if len(paramErrMap) > 0 {
		// return error w/ fancy formatted msg
		messageBuffer := bytes.NewBufferString("validation error")
		if len(paramErrMap) != 1 {
			messageBuffer.WriteString("s")
		}
		messageBuffer.WriteString(":")
		for paramName, errs := range paramErrMap {
			for _, err := range errs {
				messageBuffer.WriteString(fmt.Sprintf("\n- %v: %v", paramName, err.Error()))
			}
		}
		return fmt.Errorf(messageBuffer.String())
	}

	return nil
}
