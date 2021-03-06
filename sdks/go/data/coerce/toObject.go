package coerce

import (
	"encoding/json"
	"fmt"
	"github.com/opctl/opctl/sdks/go/model"
	"io/ioutil"
)

// ToObject coerces a value to an object value
func ToObject(
	value *model.Value,
) (*model.Value, error) {
	switch {
	case nil == value:
		return nil, nil
	case nil != value.Array:
		return nil, fmt.Errorf("unable to coerce array to object; incompatible types")
	case nil != value.Dir:
		return nil, fmt.Errorf("unable to coerce dir '%v' to object; incompatible types", *value.Dir)
	case nil != value.File:
		fileBytes, err := ioutil.ReadFile(*value.File)
		if nil != err {
			return nil, fmt.Errorf("unable to coerce file to object; error was %v", err.Error())
		}
		valueMap := &map[string]interface{}{}
		err = json.Unmarshal([]byte(fileBytes), valueMap)
		if nil != err {
			return nil, fmt.Errorf("unable to coerce file to object; error was %v", err.Error())
		}
		return &model.Value{Object: valueMap}, nil
	case nil != value.Number:
		return nil, fmt.Errorf("unable to coerce number '%v' to object; incompatible types", *value.Number)
	case nil != value.Object:
		return value, nil
	case nil != value.String:
		valueMap := &map[string]interface{}{}
		err := json.Unmarshal([]byte(*value.String), valueMap)
		if nil != err {
			return nil, fmt.Errorf("unable to coerce string to object; error was %v", err.Error())
		}
		return &model.Value{Object: valueMap}, nil
	default:
		return nil, fmt.Errorf("unable to coerce '%+v' to object", value)
	}
}
