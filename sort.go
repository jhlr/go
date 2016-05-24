package iter

import (
	"reflect"
	"sort"
)

type sortInterface struct {
	value reflect.Value
}

func (s *sortInterface) Swap(i, j int) {
	x := s.value.Index(i).Interface()
	y := s.value.Index(j).Interface()
	s.value.Index(i).Set(valueOf(y))
	s.value.Index(j).Set(valueOf(x))
}

func (s *sortInterface) Len() int {
	return s.value.Len()
}

func (s *sortInterface) Less(i, j int) bool {
	ii := s.value.Index(i).Interface()
	ij := s.value.Index(j).Interface()
	l, ok := Less(ii, ij)
	if !ok {
		panic(errTypeNotSupported)
	}
	return l
}

// SortInterface makes a sort.Interface using reflect
func SortInterface(slice interface{}) sort.Interface {
	switch s := slice.(type) {
	case sort.Interface:
		return s
	case []int:
		return sort.IntSlice(s)
	case []float64:
		return sort.Float64Slice(s)
	case []string:
		return sort.StringSlice(s)
	}
	si := new(sortInterface)
	si.value = valueOf(slice)
	if si.value.Kind() == reflect.Array {
		panic(errNotSlice)
	}
	si.Less(0, 0) // checks the type
	return si
}

// Sort uses reflect to sort the slice
func Sort(slice interface{}) {
	si := SortInterface(slice)
	sort.Sort(si)
}
