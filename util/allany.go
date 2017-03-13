package util

import "reflect"

// All returns true if cond(e) was always true
//
// cond = func(k?,e)bool
func All(iterable, cond interface{}) bool {
	return allany(true, iterable, cond)
}

// Any returns false if cond(e) was always false
//
// cond = func(k?,e)bool
func Any(iterable, cond interface{}) bool {
	return allany(false, iterable, cond)
}

func allany(allany bool, iterable, cond interface{}) bool {
	v := valueOf(iterable)
	t := v.Type()
	fv := valueOf(cond)

	in := []reflect.Type{t.Key(), t.Elem()}
	out := []reflect.Type{reflect.TypeOf(true)}
	switch {
	case CheckFunc(fv, in[1:], out):
	case CheckFunc(fv, in, out):
	default:
		panic(TypeError{})
	}

	err := For(v, func(_, ei interface{}) {
		in := []reflect.Value{valueOf(ei)}
		b := fv.Call(in)[0].Bool()
		if b != allany {
			allany = !allany
			Break()
		}
	})
	if err != nil {
		panic(err)
	}
	return allany
}
