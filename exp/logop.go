package exp

import (
	"fmt"
)

type LogicOp []bool

const (
	Not   = "tf"
	True  = "t"
	False = "f"
	And = "ffft"
	Or = "fttt"
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
	if len(op) == 0 {
		return i > 0
	} else if i >= len(op) {
		panic(fmt.Sprintf("exp: (%s) received too many args", op.String()))
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
