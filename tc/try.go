package tc

import "fmt"

type Error struct {
	inte interface{}
}

func (e Error) Println() {
	if e.inte != nil {
		fmt.Println(e.inte)
	}
}

func (e Error) Panic() {
	if e.inte != nil {
		panic(e.inte)
	}
}

func (e Error) String() string {
	return fmt.Sprint(e.inte)
}

func Try(try func()) (e Error) {
	if try != nil {
		defer func() {
			e.inte = recover()
		}()
		try()
	}
	return e
}

func (e Error) React(react func()) {
	if e.inte != nil && react != nil {
		react()
	}
}

func (e Error) Catch(catch func(interface{})) {
	if e.inte != nil && catch != nil {
		catch(e.inte)
	}
}
