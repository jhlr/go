package util

import "fmt"

type Error struct {
	i interface{}
}

func (e Error) Interface() interface{} {
	return e.i
}

func (e Error) Nil() bool {
	return e.i == nil
}

func (e Error) Error() string {
	return fmt.Sprint(e.Interface())
}

func Try(try func()) (e Error) {
	if try != nil {
		defer func() {
			e.i = recover()
		}()
		try()
	}
	return
}

func (e Error) React(react func()) {
	if !e.Nil() && react != nil {
		react()
	}
}

func (e Error) Recover(rec func(interface{})) {
	if !e.Nil() && rec != nil {
		rec(e.Interface())
	}
}
