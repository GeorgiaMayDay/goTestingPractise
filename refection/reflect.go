package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
		return
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
		return
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
