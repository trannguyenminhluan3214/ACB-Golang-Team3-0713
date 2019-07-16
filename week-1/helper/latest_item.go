package helper

import "reflect"

func Last(lst interface{}) interface{} {
	typeOfList := reflect.TypeOf(lst)
	if typeOfList.Kind() != reflect.Slice && typeOfList.Kind() != reflect.Array {
		panic("list nên loại Array hoặc Slice")
	}
	valueOfList := reflect.ValueOf(lst)
	return valueOfList.Index(valueOfList.Len() - 1).Interface()
}
