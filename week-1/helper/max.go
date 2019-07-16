package helper

import "reflect"

func Max(list interface{}) interface{} {
	items := reflect.ValueOf(list)
	listType := items.Type()
	listKind := listType.Kind()

	if listKind != reflect.Array && listKind != reflect.Slice {
		panic("list nên loại Array hoặc Slice")
	}
	elementType := listType.Elem()
	elementKind := elementType.Kind()

	max := items.Index(0)
	for i := 1; i < items.Len(); i++ {
		switch elementKind {
		case reflect.Int:
			if max.Int() < items.Index(i).Int() {
				max = items.Index(i)
			}
		case reflect.Int32:
			if max.Int() < items.Index(i).Int() {
				max = items.Index(i)
			}
		}

	}
	return max.Interface()
}
