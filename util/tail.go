package util

// Tail should be the type of the function
// that wraps every return statement of
// tailcall optimized functions
// the function will have to be called with Loop
type Tail func() interface{}

// Loop will make the iterations of the
// given function until the result is ready.
// It is much faster than normal recursion
func Loop(res interface{}) interface{} {
	for {
		tl, ok := res.(Tail)
		if ok {
			res = tl()
		} else {
			return res
		}
	}
}

// Loop is a syntactic sugar for the Loop function
func (tl Tail) Loop() interface{} {
	return Loop(tl)
}

/*
func Factorial(n int, acc int) util.Tail {
	if n <= 0 {
		return func() interface{} {
			return acc
		}
	}

	// add this wrapper to your tail call
	return func() interface{} {
		return Factorial(n-1, n*acc)
	}
}

func main() {
	// call with Loop
	_ = Factorial(10000, 1).Loop().(int)
}
*/
