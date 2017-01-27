package util

func Product(iterable interface{}) interface{} {
	return fold(iterable, Mul)
}

func Sum(iterable interface{}) interface{} {
	return fold(iterable, Add)
}

func fold(iterable interface{}, foo func(interface{}, interface{}) interface{}) interface{} {
	var result interface{}
	var ok bool
	ok = For(iterable, func(_, v interface{}) {
		result = foo(result, v)
		if result == nil {
			// this stops the loop
			result = v
			Break()
		}
	})
	if !ok {
		panic(ErrTypeNotSupported{result})
	}
	return result
}
