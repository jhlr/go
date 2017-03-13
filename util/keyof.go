package util

// KeyOf returns the key of the given elem
// or nil if not found
func KeyOf(iterable, elem interface{}) interface{} {
	var key interface{}
	err := For(iterable, func(k, v interface{}) {
		if v == elem {
			key = k
			Break()
		}
	})
	if err != nil {
		panic(err)
	}
	return key
}
