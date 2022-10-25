package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	switch x.(type) {
	case string:
		fn(val.String())
		break
	default:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fn(field.String())
		}
	}
}
