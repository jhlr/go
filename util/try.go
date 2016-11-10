package util

import "fmt"

type Error struct {
	i interface{}
}

func (e Error) Panic() {
	if e.i != nil {
		panic(e.i)
	}
}

func (e Error) String() string {
	return fmt.Sprint(e.i)
}

func Try(try func()) Error {
	var e Error
	if try != nil {
		defer func() {
			e.i = recover()
		}()
		try()
	}
	return e
}

func (e Error) React(react func()) {
	if e.i != nil && react != nil {
		react()
	}
}

func (e Error) Recover(rec func(interface{})) {
	if e.i != nil && rec != nil {
		rec(e.i)
	}
}
