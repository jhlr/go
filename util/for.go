package util

import (
	"reflect"
	"unicode"
)

// For ...
//   worst of 200 times slower than a normal for-range
// foo = func(k,e)e?
//   if iterable is a chan, k=int
//   return value will be re-assigned
//   may be called at random
//   should receive key and elem
//   may call break
func For(iterable, foo interface{}) (err interface{}) {
	defer func() {
		if err != nil {
			return
		}
		err = recover()
		if err == nil {
			return
		}
		_, ok := err.(BreakSignal)
		if ok {
			err = nil
		}
	}()

	v := valueOf(iterable)
	t := v.Type()
	fv := valueOf(foo)

	var et reflect.Type
	if Catch(func() {
		et = t.Elem()
	}) != nil {
		return TypeError{}
	}

	checkWith := func(k reflect.Type) {
		in := []reflect.Type{k, et}
		switch {
		case CheckFunc(fv, in, nil):
		case CheckFunc(fv, in, in[1:]):
			if !v.CanSet() {
				panic(TypeError{})
			}
		default:
			panic(TypeError{})
		}
	}
	call := func(k, v reflect.Value) []reflect.Value {
		return fv.Call([]reflect.Value{k, v})
	}

	kind := t.Kind()
	switch kind {
	case reflect.Map:
		checkWith(t.Key())
		keys := v.MapKeys()
		length := len(keys)
		for i := 0; i < length; i++ {
			k := keys[i]
			rets := call(k, v.MapIndex(k))
			if len(rets) == 1 {
				v.SetMapIndex(k, rets[0])
			}
		}
	case reflect.Struct:
		checkWith(reflect.TypeOf(""))
		length := v.NumField()
		typ := v.Type()
		for i := 0; i < length; i++ {
			field := typ.Field(i)
			if unicode.IsLower(rune(field.Name[0])) {
				continue
			}
			k := reflect.ValueOf(field.Name)
			rets := call(k, v.Field(i))
			if len(rets) == 1 {
				v.FieldByName(field.Name).Set(rets[0])
			}
		}
	case reflect.Chan:
		checkWith(reflect.TypeOf(int(0)))
		for i := 0; ; i++ {
			e, ok := v.Recv()
			if !ok {
				break
			}
			rets := call(reflect.ValueOf(i), e)
			if len(rets) == 1 {
				v.Send(rets[0])
			}
		}
	case reflect.Slice, reflect.Array,
		reflect.String:
		checkWith(reflect.TypeOf(int(0)))
		length := v.Len()
		for i := 0; i < length; i++ {
			k := reflect.ValueOf(i)
			rets := call(k, v.Index(i))
			if len(rets) == 1 {
				v.Index(i).Set(rets[0])
			}
		}
	default:
		return TypeError{}
	}
	return nil
}
