package util

import (
	"reflect"
	"sort"
)

func Add(a, b interface{}) (result interface{}) {
	defer recoverInterface(&result)
	va := valueOf(a)
	vb := valueOf(b)
	a = va.Interface()
	b = vb.Interface()
	switch v := a.(type) {
	case bool:
		result = v || b.(bool)
	case int:
		result = v + b.(int)
	case int8:
		result = v + b.(int8)
	case int16:
		result = v + b.(int16)
	case int32:
		result = v + b.(int32)
	case int64:
		result = v + vb.Int()
	case uint:
		result = v + b.(uint)
	case uint8:
		result = v + b.(uint8)
	case uint16:
		result = v + b.(uint16)
	case uint32:
		result = v + b.(uint32)
	case uint64:
		result = v + vb.Uint()
	case float32:
		result = v + b.(float32)
	case float64:
		result = v + vb.Float()
	case complex64:
		result = v + b.(complex64)
	case complex128:
		result = v + vb.Complex()
	case string:
		result = v + b.(string)
	}
	if result != nil {
		return result
	}
	ta := va.Type()
	tb := vb.Type()
	if ta != tb {
		return nil
	}
	switch ta.Kind() {
	case reflect.Map:
		res := reflect.MakeMap(ta)
		it := func(k, v interface{}) {
			res.SetMapIndex(valueOf(k), valueOf(v))
		}
		For(va, it)
		For(vb, it)
		result = res.Interface()
	case reflect.Slice, reflect.Array:
		la, lb := va.Len(), vb.Len()
		r := reflect.MakeSlice(ta, la, la+lb)
		reflect.Copy(r, va)
		result = reflect.AppendSlice(r, vb).Interface()
	}
	return result
}

func Mul(a, b interface{}) (result interface{}) {
	defer recoverInterface(&result)
	va := valueOf(a)
	vb := valueOf(b)
	a = va.Interface()
	b = vb.Interface()
	switch v := a.(type) {
	case bool:
		result = v && b.(bool)
	case int:
		result = v * b.(int)
	case int8:
		result = v * b.(int8)
	case int16:
		result = v * b.(int16)
	case int32:
		result = v * b.(int32)
	case int64:
		result = v * vb.Int()
	case uint:
		result = v * b.(uint)
	case uint8:
		result = v * b.(uint8)
	case uint16:
		result = v * b.(uint16)
	case uint32:
		result = v * b.(uint32)
	case uint64:
		result = v * vb.Uint()
	case float32:
		result = v * b.(float32)
	case float64:
		result = v * vb.Float()
	case complex64:
		result = v * b.(complex64)
	case complex128:
		result = v * vb.Complex()
	default:
		return nil
	}
	return result
}

func Less(a, b interface{}) (result bool, ok bool) {
	defer recoverBool(&ok)
	va := valueOf(a)
	vb := valueOf(b)
	a = va.Interface()
	b = vb.Interface()
	switch v := a.(type) {
	case bool:
		result = !v && b.(bool)
	case int:
		result = v < b.(int)
	case int8:
		result = v < b.(int8)
	case int16:
		result = v < b.(int16)
	case int32:
		result = v < b.(int32)
	case int64:
		result = v < vb.Int()
	case uint:
		result = v < b.(uint)
	case uint8:
		result = v < b.(uint8)
	case uint16:
		result = v < b.(uint16)
	case uint32:
		result = v < b.(uint32)
	case uint64:
		result = v < vb.Uint()
	case float32:
		result = v < b.(float32)
	case float64:
		result = v < vb.Float()
	case string:
		ss := []string{v, b.(string)}
		result = sort.StringSlice(ss).Less(0, 1)
	default:
		if a == nil && b != nil {
			return true, true
		}
		return va.Len() < vb.Len(), true
	}
	return result, true
}
