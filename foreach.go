package iter

import (
	"reflect"
)

// ForEach passes the key and the value of each element
// to the given function
func ForEach(iterable interface{},
	foo func(interface{}, interface{}) bool) (ok bool) {
	defer recoverBool(&ok)
	v := valueOf(iterable)
	var length int
	var keys []reflect.Value
	var typ reflect.Type

	k := v.Kind()
	switch k {
	case reflect.Map:
		keys = v.MapKeys()
		length = len(keys)
	case reflect.Struct:
		length = v.NumField()
		typ = v.Type()
	case reflect.Slice,
		reflect.Array,
		reflect.String:
		length = v.Len()
	case reflect.Chan:
		for i := 0; ; i++ {
			value, ok := v.Recv()
			if !ok {
				return true
			}
			if !foo(i, value.Interface()) {
				return false
			}
		}
	default:
		return false
	}

	var key interface{}
	var value interface{}
	for i := 0; i < length; i++ {
		switch k {
		case reflect.Map:
			key, value = keys[i].Interface(),
				v.MapIndex(keys[i]).Interface()
		case reflect.Struct:
			key, value = typ.Field(i).Name,
				v.Field(i).Interface()
		case reflect.Slice,
			reflect.Array,
			reflect.String:
			key, value = i, v.Index(i)
		}
		if !foo(key, value) {
			return false
		}
	}
	return true
}
