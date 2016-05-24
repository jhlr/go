package exp

import (
	"fmt"
)

type LogicOp []bool

const (
	Not = "tf"
	True = "t"
	False = "f"
)

func LogicOpString(str string) LogicOp {
	op := make(LogicOp, len(str))
	for i := range str {
		switch str[i] {
		case 'f', 'F':
			op[i] = false
		case 't', 'T':
			op[i] = true
		default:
			panic(fmt.Sprintf("exp: invalid char %c", str[i]))
		}
	}
	return op
}

func (op LogicOp) Apply(args ...bool) bool {
	i := 0
	inc := 1
	for _, a := range args {
		if a {
			i += inc
		}
		inc *= 2
	}
	if i >= len(op){
		fmt.Println(op.String())
		panic("exp: too many args to LogicOp")
	}
	return op[i]
}

func (op LogicOp) String() string {
	result := ""
	for i := range op {
		if op[i] {
			result += "t"
		} else {
			result += "f"
		}
	}
	return result
}
