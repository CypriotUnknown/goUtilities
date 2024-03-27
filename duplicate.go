package utilities

import "reflect"

func MakeDuplicate(object any) any {
	elementType := reflect.TypeOf(object).Elem()

	return reflect.New(elementType).Interface()
}
