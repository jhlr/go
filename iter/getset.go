package iter

import (
	"reflect"
)

type GetFunc func(interface{}) interface{}
type SetFunc func(interface{}, interface{}) bool

func MakeGetter(iterable interface{}) GetFunc {
	recoverInterface := func(i *interface{}) {
		r := recover()
		if r != nil {
			*i = nil
		}
	}
	value := valueOf(iterable)
	Int := reflect.TypeOf(int(0))
	kind := value.Kind()
	if kind == reflect.Ptr {
		value = value.Elem()
		kind = value.Kind()
	}
	switch kind {
	case reflect.Map:
		return func(k interface{}) (r interface{}) {
			defer recoverInterface(&r)
			vk := valueOf(k)
			return value.MapIndex(vk).Interface()
		}
	case reflect.Struct:
		return func(k interface{}) (r interface{}) {
			defer recoverInterface(&r)
			switch rvk := k.(type) {
			case string:
				return value.FieldByName(rvk).Interface()
			case []int:
				return value.FieldByIndex(rvk).Interface()
			default:
				vk := valueOf(k).Convert(Int)
				return value.Field(vk.Interface().(int)).Interface()
			}
		}
	case reflect.Slice,
		reflect.Array,
		reflect.String:
		return func(k interface{}) (r interface{}) {
			defer recoverInterface(&r)
			vk := valueOf(k).Convert(Int)
			return value.Index(vk.Interface().(int)).Interface()
		}
	case reflect.Chan:
		return func(_ interface{}) (r interface{}) {
			defer recoverInterface(&r)
			vr, ok := value.Recv()
			if !ok {
				return nil
			}
			return vr.Interface()
		}
	}
	return nil
}

func MakeSetter(iterable interface{}) SetFunc {
	value := valueOf(iterable)
	Int := reflect.TypeOf(int(0))
	kind := value.Kind()
	if kind == reflect.Ptr {
		value = value.Elem()
		kind = value.Kind()
	}
	switch kind {
	case reflect.Map:
		return func(k, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			value.SetMapIndex(valueOf(k), valueOf(v))
			return true
		}
	case reflect.Struct:
		return func(k, v interface{}) (ok bool) {
			defer recoverBool(&ok)
			vv := valueOf(v)
			switch rvk := k.(type) {
			case string:
				value.FieldByName(rvk).Set(vv)
			case []int:
				value.FieldByIndex(rvk).Set(vv)
			default:
				vk := valueOf(k).Convert(Int)
				value.Field(vk.Interface().(int)).Set(vv)
			}
			return true
		}
	case reflect.String,
		reflect.Array:
		return nil
	case reflect.Slice:
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
