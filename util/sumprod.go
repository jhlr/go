package util

func Product(iterable interface{}) interface{} {
	return fold(iterable, Mul)
}

func Sum(iterable interface{}) interface{} {
	return fold(iterable, Add)
}

func fold(iterable interface{}, foo func(interface{}, interface{}) (interface{}, bool)) interface{} {
	var result interface{}
	var ok bool
	ok = For(iterable, func(_, v interface{}) {
		result, ok = foo(result, v)
		if !ok {
			// this stops the loop
			panic(nil)
		}
	})
	if !ok {
		panic(errTypeNotSupported)
	}
	return result
}
