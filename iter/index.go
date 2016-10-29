package iter

import (
	"errors"
	"reflect"
)

var (
	// panic if the type is not supported
	errTypeNotSupported = errors.New("type not supported")
	errNotSlice         = errors.New("interface is not an iterable")
)

func recoverBool(ok *bool) {
	r := recover()
	*ok = r == nil
}

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

// Index returns the index of the given element
func KeyOf(iterable interface{}, elem interface{}) (interface{}, bool) {
	key := interface{}(nil)
	found := !ForEach(iterable, func(k, v interface{}) bool {
		if v == elem {
			key = k
			return false
		}
		return true
	})
	return key, found
}
