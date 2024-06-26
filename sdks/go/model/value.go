package model

import (
	"fmt"
	"strconv"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// Value represents a typed value
type Value struct {
	Array   *[]interface{}          `json:"array,omitempty"`
	Boolean *bool                   `json:"boolean,omitempty"`
	Dir     *string                 `json:"dir,omitempty"`
	File    *string                 `json:"file,omitempty"`
	Number  *float64                `json:"number,omitempty"`
	Object  *map[string]interface{} `json:"object,omitempty"`
	Socket  *string                 `json:"socket,omitempty"`
	String  *string                 `json:"string,omitempty"`
}

// Unbox unboxes a Value into a native go type
func (value Value) Unbox() (interface{}, error) {
	if value.Array != nil {
		nativeArray := []interface{}{}
		for itemKey, itemValue := range *value.Array {
			switch typedItemValue := itemValue.(type) {
			case Value:
				nativeItem, err := typedItemValue.Unbox()
				if err != nil {
					return nil, fmt.Errorf("unable to unbox item '%v' from array: %w", itemKey, err)
				}

				nativeArray = append(nativeArray, nativeItem)
			default:
				nativeArray = append(nativeArray, itemValue)
			}
		}
		return nativeArray, nil
	} else if value.Boolean != nil {
		return *value.Boolean, nil
	} else if value.Dir != nil {
		return *value.Dir, nil
	} else if value.File != nil {
		return *value.File, nil
	} else if value.Number != nil {
		return *value.Number, nil
	} else if value.Object != nil {
		nativeObject := map[string]interface{}{}
		for propKey, propValue := range *value.Object {
			switch typedPropValue := propValue.(type) {
			case Value:
				var err error
				if nativeObject[propKey], err = typedPropValue.Unbox(); err != nil {
					return nil, fmt.Errorf("unable to unbox property '%v' from object: %w", propKey, err)
				}
			default:
				nativeObject[propKey] = propValue
			}
		}
		return nativeObject, nil
	} else if value.Socket != nil {
		return *value.Socket, nil
	} else if value.String != nil {
		return *value.String, nil
	}
	return nil, fmt.Errorf("unable to unbox value '%+v'", value)
}

func (value Value) format() (interface{}, error) {
	if nil != value.Array {
		formattedArray := []interface{}{}
		for itemKey, itemValue := range *value.Array {
			switch typedItemValue := itemValue.(type) {
			case Value:
				formattedValue, err := typedItemValue.format()
				if nil != err {
					return "", errors.Wrap(err, fmt.Sprintf("unable to stringify item '%v' from array", itemKey))
				}

				formattedArray = append(formattedArray, formattedValue)
			default:
				formattedArray = append(formattedArray, itemValue)
			}
		}
		return formattedArray, nil
	} else if nil != value.Boolean {
		return strconv.FormatBool(*value.Boolean), nil
	} else if nil != value.Dir {
		return *value.Dir, nil
	} else if nil != value.File {
		return *value.File, nil
	} else if nil != value.Number {
		return fmt.Sprintf("%f", *value.Number), nil
	} else if nil != value.Object {
		formattedMap := map[string]interface{}{}
		for propKey, propValue := range *value.Object {
			switch typedPropValue := propValue.(type) {
			case Value:
				var err error
				if formattedMap[propKey], err = typedPropValue.format(); nil != err {
					return "", errors.Wrap(err, fmt.Sprintf("unable to stringify property '%v' from object", propKey))
				}
			default:
				formattedMap[propKey] = propValue
			}
		}
		return formattedMap, nil
	} else if nil != value.Socket {
		return *value.Socket, nil
	} else if nil != value.String {
		return *value.String, nil
	}
	return "", fmt.Errorf("unable to stringify value '%+v'", value)
}

func FormatValueMap(valueMap map[string]*Value) (string, error) {
	formattedValues := map[string]interface{}{}
	for name, value := range valueMap {
		description, err := value.format()
		if err != nil {
			return "", err
		}
		formattedValues[name] = description
	}
	bytes, err := yaml.Marshal(formattedValues)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
