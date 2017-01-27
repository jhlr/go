package util

import "fmt"

type ErrTypeNotSupported struct {
	obj interface{}
}

func (e ErrTypeNotSupported) Error() string {
	format := "util: %v has unsupported type for this function"
	return fmt.Sprintf(format, e.obj)
}

type errBreak struct{}

func (e errBreak) Error() string {
	return "util: misuse of the Break function"
}

func recoverBool(ok *bool) {
	r := recover()
	*ok = r == nil
}

func recoverInterface(i *interface{}) {
	r := recover()
	if r != nil {
		*i = nil
	}
}
