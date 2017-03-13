package util

import (
	"reflect"
)

// Map ...
//   will allocate a copy of the iterable
//   the elems will be modified based on Foo
//   returned object will be of the same type
//     but arrays will be turned into slices
// foo = func(ae)(re,bool?)
//   ae = argument type elem
//   re = return type elem
//   may return a bool for include option
//   may call break
func Map(iterable, foo interface{}) interface{} {
	v := valueOf(iterable)
	t := v.Type()

	fv := valueOf(foo)
	ft := fv.Type()

	var res reflect.Value
	var set func(interface{}, reflect.Value)
	sliceSetFunc := func(k interface{}, v reflect.Value) {
		res.Index(k.(int)).Set(v)
	}

	var et reflect.Type
	if Catch(func() {
		et = t.Elem()
	}) != nil {
		panic(TypeError{})
	}

	in := []reflect.Type{et}
	out := []reflect.Type{
		reflect.TypeOf(nil),
		reflect.TypeOf(true)}
	switch {
	case CheckFunc(fv, in, out):
	case CheckFunc(fv, in, out[:1]):
		panic(TypeError{})
	}

	switch t.Kind() {
	case reflect.Struct:
		res = reflect.New(t).Elem()
		set = func(k interface{}, v reflect.Value) {
			res.FieldByName(k.(string)).Set(v)
		}
	case reflect.Map:
		rt := reflect.MapOf(t.Key(), ft.Out(0))
		res = reflect.MakeMap(rt)
		set = func(k interface{}, v reflect.Value) {
			res.SetMapIndex(valueOf(k), v)
		}
	case reflect.Chan:
		rt := reflect.ChanOf(reflect.BothDir, ft.Out(0))
		res = reflect.MakeChan(rt, v.Cap())
		set = func(k interface{}, v reflect.Value) {
			res.Send(v)
		}
	case reflect.Slice:
		rt := reflect.SliceOf(ft.Out(0))
		res = reflect.MakeSlice(rt, v.Len(), v.Len())
		set = sliceSetFunc
	case reflect.Array:
		rt := reflect.SliceOf(ft.Out(0))
		res = reflect.MakeSlice(rt, v.Len(), v.Len())
		set = sliceSetFunc
	case reflect.String:
		rt := reflect.TypeOf([]byte{})
		res = reflect.MakeSlice(rt, v.Len(), v.Len())
		set = sliceSetFunc
	default:
		panic(TypeError{})
	}

	if err := For(v, func(k, ei interface{}) {
		e := valueOf(ei)
		rets := fv.Call([]reflect.Value{e})
		if len(rets) != 2 || rets[1].Bool() {
			set(k, rets[0])
		}
	}); err != nil {
		panic(err)
	}

	if t.Kind() == reflect.String {
		return res.Convert(t).Interface()
	}
	return res.Interface()
}
