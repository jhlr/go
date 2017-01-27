package util

type Looper func() interface{}

func Loop(res interface{}) interface{} {
	for {
		lp, ok := res.(Looper)
		if ok {
			res = lp()
		} else {
			return res
		}
	}
}

func (lp Looper) Loop() interface{} {
	return Loop(lp)
}

/*
// change the return type of your function to interface{}
func Factorial(n int, acc int) util.Flooper {
	if n <= 0 {
		return util.Looper(func() interface{} {
			return acc
		})
	}

	// add this wrapper to your tail call
	return util.Looper(func() interface{} {
		return Factorial(n-1, n*acc)
	})
}

func main() {
	// call with Loop
	Factorial(10000, 1).Loop().(int)
}
*/
