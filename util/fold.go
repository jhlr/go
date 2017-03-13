package util

import (
	"reflect"
)

// Fold ...
//   uses the callback on all the elements
//   in pairs and returns the result
// foo = func(e,e)e
//   should receive 2 elem
//   should return 1 elem
//   may call break
func Fold(iterable, foo interface{}) interface{} {
	fv := valueOf(foo)
	var result reflect.Value
	empty := true
	v := valueOf(iterable)
	t := v.Type()

	var et reflect.Type
	if Catch(func() {
		et = t.Elem()
	}) != nil {
		panic(TypeError{})
	}

	in := []reflect.Type{et, et}
	switch {
	case CheckFunc(fv, in, in[1:]):
	default:
		panic(TypeError{})
	}

	err := For(iterable, func(_, ei interface{}) {
		e := valueOf(ei)
		if empty {
			result = e
			empty = false
		} else {
			result = fv.Call([]reflect.Value{result, e})[0]
		}
	})
	if err != nil {
		panic(err)
	}
	return result.Interface()
}
