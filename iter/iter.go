package iter

import (
	"reflect"
)

func Cycle(iterable interface{}, n uint) interface{} {
	v := valueOf(iterable)
	typ := v.Type()
	l := v.Len()
	result := reflect.MakeSlice(typ, int(n), int(n))
	for i := result.Len() - 1; i >= 0; i-- {
		result.Index(i).Set(v.Index(i % l))
	}
	return result.Interface()
}

func Repeat(elem interface{}, n uint) interface{} {
	v := valueOf(elem)
	typ := reflect.SliceOf(v.Type())
	result := reflect.MakeSlice(typ, int(n), int(n))
	for i := result.Len() - 1; i >= 0; i-- {
		result.Index(i).Set(v)
	}
	return result.Interface()
}
