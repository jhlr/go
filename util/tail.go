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
