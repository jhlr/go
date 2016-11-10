package util

import (
	"reflect"
	"sort"
)

func Add(a, b interface{}) (result interface{}, ok bool) {
	defer recoverBool(&ok)
	result = nil
	ok = true
	va := valueOf(a)
	vb := valueOf(b)
	ta := va.Type()
	tb := vb.Type()
	a = va.Interface()
	b = vb.Interface()
	switch va.Kind() {
	case reflect.Bool:
		result = a.(bool) || b.(bool)
	case reflect.Int:
		result = a.(int) + b.(int)
	case reflect.Int8:
		result = a.(int8) + b.(int8)
	case reflect.Int16:
		result = a.(int16) + b.(int16)
	case reflect.Int32:
		result = a.(int32) + b.(int32)
	case reflect.Int64:
		result = a.(int64) + vb.Int()
	case reflect.Uint:
		result = a.(uint) + b.(uint)
	case reflect.Uint8:
		result = a.(uint8) + b.(uint8)
	case reflect.Uint16:
		result = a.(uint16) + b.(uint16)
	case reflect.Uint32:
		result = a.(uint32) + b.(uint32)
	case reflect.Uint64:
		result = a.(uint64) + vb.Uint()
	case reflect.Float32:
		result = a.(float32) + b.(float32)
	case reflect.Float64:
		result = a.(float64) + vb.Float()
	case reflect.Complex64:
		result = a.(complex64) + b.(complex64)
	case reflect.Complex128:
		result = a.(complex128) + vb.Complex()
	case reflect.String:
		result = a.(string) + b.(string)
	case reflect.Map:
		if ta == tb {
			res := reflect.MakeMap(ta)
			it := MakeSetter(res)
			For(va, it)
			For(vb, it)
			result = res.Interface()
		} else {
			ok = false
		}
	case reflect.Slice,
		reflect.Array:
		if ta == tb {
			la, lb := va.Len(), vb.Len()
			va = va.Slice(0, la)
			vb = vb.Slice(0, lb)
			r := reflect.MakeSlice(va.Type(), la, la+lb)
			reflect.Copy(r, va)
			result = reflect.AppendSlice(r, vb).Interface()
		} else if ta == reflect.SliceOf(tb) {
			result = reflect.Append(va, vb).Interface()
		} else if tb == reflect.SliceOf(ta) {
			result = reflect.Append(vb, va).Interface()
		} else {
			ok = false
		}
	}
	return result, ok
}

func Mul(a, b interface{}) (result interface{}, ok bool) {
	defer recoverBool(&ok)
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
		return nil, false
	}
	return result, true
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
	case int, int8, int16,
		int32, int64:
		result = va.Int() < vb.Int()
	case uint, uint8, uint16,
		uint32, uint64:
		result = va.Uint() < vb.Uint()
	case float32, float64:
		result = va.Float() < vb.Float()
	case string:
		ss := sort.StringSlice([]string{v, b.(string)})
		result = ss.Less(0, 1)
	default:
		if a == nil && b != nil {
			return true, true
		}
		return va.Len() < vb.Len(), true
	}
	return result, true
}
