package util

import "reflect"

func valueOf(i interface{}) reflect.Value {
	v, ok := i.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(i)
	}
	return v
}

func IsIterable(i interface{}) bool {
	v := valueOf(i)
	switch v.Kind() {
	case reflect.Struct,
		reflect.Map,
		reflect.Slice,
		reflect.Array,
		reflect.String,
		reflect.Chan:
		return true
	}
	return false
}

// KeyOf returns the index of the given element
func KeyOf(iterable interface{}, elem interface{}) (interface{}, bool) {
	var key interface{}
	found := !For(iterable, func(k, v interface{}) {
		if v == elem {
			key = k
			Break()
		}
	})
	return key, found
}
