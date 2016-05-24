package iter

// Max returns the key of the max element
func Max(iterable interface{}) interface{} {
	return minmax(true, iterable)
}

// Min returns the key of the min element
func Min(iterable interface{}) interface{} {
	return minmax(false, iterable)
}

func minmax(op bool, iterable interface{}) interface{} {
	var key, minmax interface{}
	empty := true
	ok := ForEach(iterable, func(k, v interface{}) bool {
		if empty {
			minmax = v
			key = k
			return true
		}
		l, ok := Less(minmax, v)
		if !ok {
			return false
		} else if op == l {
			minmax = v
			key = k
		}
		return true
	})
	if !ok {
		panic(errTypeNotSupported)
	}
	return key
}
