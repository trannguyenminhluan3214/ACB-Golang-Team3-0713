package helper

import "reflect"

func Map(list interface{}, iterateeFunc interface{}) interface{} {
	items := reflect.ValueOf(list)
	listType := items.Type()

	value := reflect.ValueOf(iterateeFunc)
	typeOfResult := reflect.SliceOf(value.Type().Out(0))
	result := reflect.MakeSlice(typeOfResult, 0, 0)

	listKind := listType.Kind()
	if listKind == reflect.Slice || listKind == reflect.Array {
		for i := 0; i < items.Len(); i++ {
			elem := items.Index(i)
			in := []reflect.Value{elem}
			out := value.Call(in)[0]
			result = reflect.Append(result, out)
		}

	}
	return result.Interface()
}
