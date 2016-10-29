package iter

func Product(iterable interface{}) interface{} {
	return fold(iterable, Mul)
}

func Sum(iterable interface{}) interface{} {
	return fold(iterable, Add)
}

func fold(iterable interface{}, foo func(interface{}, interface{}) (interface{}, bool)) interface{} {
	var result interface{}
	var ok bool
	ok = ForEach(iterable, func(_, v interface{}) bool {
		result, ok = foo(result, v)
		return ok
	})
	if !ok {
		panic(errTypeNotSupported)
	}
	return result
}
