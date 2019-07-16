package helper

import "reflect"

func IsEmpty(v interface{}) bool {
	if v == nil || v == "" || v == false || v == 0 {
		return true
	}
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		if value.Len() == 0 {
			return true
		}
	}
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			value := field.Interface()
			if !(value == nil || value == "" || value == false || value == 0) {
				return false
			}
		}
		return true
	}
	return false
}
