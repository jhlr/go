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

func Getter(iterable interface{}) func(interface{}) interface{} {
	v := valueOf(iterable)
	t := v.Type()
	switch v.Kind() {
	case reflect.Map:
		return func(k interface{}) interface{} {
			vk := valueOf(k)
			if vk.Type() == t.Key() {
				return v.MapIndex(vk).Interface()
			}
			return nil
		}
	case reflect.Struct:
		return func(k interface{}) (r interface{}) {
			defer func() {
				_ = recover()
				r = nil
			}()
			vk := valueOf(k)
			switch rvk := k.(type) {
			case uint, uint8, uint16,
				uint32, uint64:
				return v.Field(int(vk.Uint())).Interface()
			case int, int8, int16,
				int32, int64:
				return v.Field(int(vk.Int())).Interface()
			case string:
				return v.FieldByName(rvk).Interface()
			case []int:
				return v.FieldByIndex(rvk).Interface()
			}
			return nil
		}
	case reflect.Slice,
		reflect.Array,
		reflect.String:
		return func(k interface{}) interface{} {
			vk := valueOf(k)
			switch k.(type) {
			case uint, uint8, uint16,
				uint32, uint64:
				return v.Index(int(vk.Uint())).Interface()
			case int, int8, int16,
				int32, int64:
				return v.Index(int(vk.Int())).Interface()
			}
			return nil
		}
	case reflect.Chan:
		return func(_ interface{}) interface{} {
			vr, ok := v.Recv()
			if !ok {
				return nil
			}
			return vr.Interface()
		}
	}
	return nil
}

func Setter(iterable interface{}) func(interface{}, interface{}) bool {
	value := valueOf(iterable)
	Int := reflect.TypeOf(int(0))
	switch value.Kind() {
	case reflect.Map:
		return func(k, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			value.SetMapIndex(valueOf(k), valueOf(v))
			return true
		}
	case reflect.Struct:
		return func(k, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			vk := valueOf(k)
			vv := valueOf(v)
			switch rvk := k.(type) {
			case string:
				value.FieldByName(rvk).Set(vv)
			case []int:
				value.FieldByIndex(rvk).Set(vv)
			default:
				vk = vk.Convert(Int)
				value.Field(vk.Interface().(int)).Set(vv)
			}
			return true
		}
	case reflect.Slice,
		reflect.Array,
		reflect.String:
		return func(k, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			vk := valueOf(k).Convert(Int)
			vv := valueOf(v)
			value.Index(vk.Interface().(int)).Set(vv)
			return true
		}
	case reflect.Chan:
		return func(_, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			value.Send(valueOf(v))
			return true
		}
	}
	return nil
}
