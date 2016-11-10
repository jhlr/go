package util

type Flooper func() interface{}

func Loop(foo interface{}) interface{} {
	for {
		f, ok := foo.(Flooper)
		if ok {
			foo = f()
		} else {
			return foo
		}
	}
}

/*
// change the return type of your function to interface{}
func Factorial(n int, acc int) interface{} {
	if n <= 0 {
		return acc
	}

	// add this wrapper to your tail call
	return util.Flooper(func() interface{} {
		return Factorial(n-1, n*acc)
	})
}

func main() {
	// call with Loop
	util.Loop(Factorial(10000, 1)).(int)
}
*/
