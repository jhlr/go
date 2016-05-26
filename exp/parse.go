package exp

import (
	"fmt"
	"unicode"
)

func (e *Expression) String() string {
	if e == nil {
		return "()"
	}
	result := "(" + e.op.String()
	if v, ok := e.Var(); ok {
		vname := fmt.Sprintf("%d", v)
		if len(e.op) == 0 {
			return vname
		}
		result += " " + vname
	}
	for i := range e.args {
		result += " " + e.args[i].String()
	}
	result += ")"
	return result
}

func (e *Expression) parse(str string) bool {
	token := uint(0)
	numArgs := uint(0)
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			if token == 0 {
				// print("BEGIN\n")
				token++
				continue
			} else if token/2 > numArgs {
				return false
			}
			j := i + 1
			for count := 1; j < len(str); j++ {
				if str[j] == '(' {
					count++
				} else if str[j] == ')' {
					count--
					if count == 0 {
						break
					}
				}
			}
			if j == len(str) {
				return false
			}
			if token%2 == 1 {
				token++
				var newe Expression
				ok := newe.parse(str[i : j+1])
				e.args = append(e.args, newe)
				i = j
				if !ok {
					return false
				}
			} else {
				return false
			}
		} else if str[i] == ')' {
			// print("END\n")
			if token == 0 {
				return false
			} else if token == 1 {
				// vacuous truth
				e.op = LogicOpString(True)
			}
			return true
		} else if unicode.IsSpace(rune(str[i])) {
			// print("SPACE\n")
			if token > 0 && token%2 == 0 {
				token++
			}
		} else if unicode.IsLetter(rune(str[i])) {
			// print("OPERATION\n")
			if token == 1 {
				token++
			} else {
				return false
			}
			for j := i + 1; ; j++ {
				if !unicode.IsLetter(rune(str[j])) {
					e.op = LogicOpString(str[i:j])
					i = j - 1
					break
				}
			}
			numArgs = floorLog(2, uint(len(e.op)))
		} else if unicode.IsNumber(rune(str[i])) {
			// print("VARIABLE\n")
			variable := uint(0)
			fmt.Sscanf(str[i:], "%d", &variable)
			i += int(floorLog(10, variable))
			if token == 0 {
				e.variable = variable
				return true
			} else if token%2 == 1 && token/2 <= numArgs {
				token++
				if numArgs <= 1 {
					e.variable = variable
				} else {
					e.args = append(e.args, Expression{
						variable: variable,
					})
				}
			} else {
				return false
			}
		}
	}
	return false
}
