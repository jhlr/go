package gtc

import "fmt"

type Error struct {
	Interface interface{}
}

func (e Error) Println() {
	if e.Interface != nil {
		fmt.Println(e.Interface)
	}
}

func (e Error) Panic() {
	if e.Interface != nil {
		panic(e.Interface)
	}
}

func (e Error) String() string {
	return fmt.Sprint(e.Interface)
}

func Try(try func()) (e Error) {
	if try != nil {
		defer func() {
			e.Interface = recover()
		}()
		try()
	}
	return e
}

func (e Error) React(react func()) {
	if e.Interface != nil && react != nil {
		react()
	}
}

func (e Error) Catch(catch func(interface{})) {
	if e.Interface != nil && catch != nil {
		catch(e.Interface)
	}
}
