package util

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
	ok := For(iterable, func(k, v interface{}) {
		if empty {
			minmax = v
			key = k
			return
		}
		l, ok := Less(minmax, v)
		if !ok {
			minmax = v
			Break()
		} else if op == l {
			minmax = v
			key = k
		}
	})
	if !ok {
		panic(ErrTypeNotSupported{minmax})
	}
	return key
}
