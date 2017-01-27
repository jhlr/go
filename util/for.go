package util

import (
	"reflect"
)

// Break will stop the For without panic
func Break() {
	panic(errForBreak{})
}

// For passes the key and the value of each element to the given function
// will return true only if all the element were used
func For(iterable interface{}, callback func(interface{}, interface{})) (all bool) {
	defer func() {
		r := recover()
		if r != nil {
			var ok bool
			_, ok = r.(errForBreak)
			if !ok {
				panic(r)
			} else {
				all = false
			}
		}
	}()
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
			callback(i, value.Interface())
		}
	default:
		panic(ErrTypeNotSupported{iterable})
	}
	var key interface{}
	var value interface{}
	for i := 0; i < length; i++ {
		switch k {
		case reflect.Map:
			key = keys[i].Interface()
			value = v.MapIndex(keys[i]).Interface()
		case reflect.Struct:
			key = typ.Field(i).Name
			value = v.Field(i).Interface()
		case reflect.Slice,
			reflect.Array,
			reflect.String:
			key, value = i, v.Index(i)
		}
		callback(key, value)
	}
	return true
}
