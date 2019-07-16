package helper

import (
	"reflect"
	"strings"
)

func Contains(list interface{}, v interface{}) bool {
	items := reflect.ValueOf(list)
	value := reflect.ValueOf(v)
	listType := items.Type()
	result := false
	switch listType.Kind() {

	case reflect.String:
		result = strings.Contains(items.String(), value.String())
		break
	case reflect.Slice:
		for i := 0; i < items.Len(); i++ {
			if reflect.DeepEqual(items.Index(i).Interface(), v) {
				result = true
				break
			}
		}
		break
	}
	return result
}
