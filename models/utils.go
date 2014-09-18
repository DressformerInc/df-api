package models

import (
	"reflect"
)

func T(i interface{}) interface{} {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return reflect.New(reflect.TypeOf(i).Elem()).Interface()
	}

	val := reflect.New(reflect.TypeOf(i))

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val.Interface()
}
