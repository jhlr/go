package exp

import (
	"fmt"
	"unicode"
)

type Expression struct {
	variable uint
	op       LogicOp
	args     []Expression
}

func (e *Expression) Var() (uint, bool) {
	if e == nil {
		return 0, false
	}
	return e.variable,
		e.args == nil &&
		e.op == nil
}

func (e *Expression) Rename(v, newv uint) bool {
	if e == nil {
		return false
	}
	changed := false
	for i := range e.args {
		changed = changed ||
			e.args[i].Rename(v, newv)
	}
	oldv, ok := e.Var()
	if ok && oldv == v {
		e.variable = newv
		changed = true
	}
	return changed
}

func (e *Expression) Copy() *Expression {
	if e == nil {
		return nil
	}
	args := make([]Expression, len(e.args))
	for i := range args {
		args[i] = *e.args[i].Copy()
	}
	op := make(LogicOp, len(e.op))
	copy(op, e.op)
	return &Expression{
		variable: e.variable,
		args: args, op: op,
	}
}

func (e *Expression) build(str string) bool {
	token := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			if token == 0 {
				// print("BEGIN\n")
				token++
				continue
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
			if token % 2 == 1 {
				token++
				var newe Expression
				ok := newe.build(str[i : j+1])
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
			return true
		} else if unicode.IsSpace(rune(str[i])) {
			// print("SPACE\n")
			if token > 0 && token % 2 == 0 {
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
		} else if unicode.IsNumber(rune(str[i])) {
			// print("VARIABLE\n")
			variable := uint(0)
			fmt.Sscanf(str[i:], "%d", &variable)
			temp := variable
			digits := 1
			for temp > 9 {
				temp /= 10
				digits++
			}
			i += digits - 1
			if token == 0 {
				e.variable = variable
				return true
			} else if token % 2 == 1 {
				token++
				e.args = append(e.args, Expression{
					variable: variable,
				})
			} else {
				return false
			}
		}
	}
	return false
}

func New(str string) *Expression {
	e := new(Expression)
	ok := e.build(str)
	if ok {
		return e
	}
	return nil
}

func (e *Expression) String() string {
	if e == nil {
		return "()"
	} else if v, ok := e.Var(); ok {
		return fmt.Sprintf("%d", v)
	}
	result := "(" + e.op.String()
	for i := range e.args {
		result += " " + e.args[i].String()
	}
	result += ")"
	return result
}

func (e *Expression) Compiled() *Expression {
	v := e.Variables()
	results := e.results(v)
	e = &Expression{
		op: LogicOp(results),
		args: make([]Expression, v.Len()),
	}
	for i := range e.args {
		e.args[i] = Expression{variable: uint(i)}
	}
	return e
}

func (e *Expression) Simplify() {
	v := e.Variables()
	e.simplify(v)
}

func pattern(results []bool) (uint, bool) {
	fd := uint(0)
	for i := 1; i < len(results); i++ {
		if results[i-1] != results[i] {
			fd = uint(i)
			break
		}
	}
	if fd == 0 {
		return 0, true
	} else if 2*fd > uint(len(results)) {
		return 0, false
	}
	equ := true
	not := true
	b := results[0]
	for i := 0; i < len(results); i++ {
		if i > 0 && uint(i) % fd == 0 {
			b = !b
		}
		if results[i] != b {
			equ = false
		} else {
			not = false
		}
	}
	if equ || not {
		return fd, true
	}
	return 0, false
}

func (e *Expression) simplify(v *Variables) {
	if _, ok := e.Var(); e == nil || ok {
		return
	}
	for i := range e.args {
		e.args[i].simplify(v)
	}
	v.SetAll(false)
	results := e.results(v) /* ! */
	fd, ok := pattern(results)
	if !ok {
		return
	}
	if fd == 0 {
		if results[0] {
			e.op = LogicOpString(True)
		} else {
			e.op = LogicOpString(False)
		}
		e.variable = 0
		e.args = nil
		return
	}
	hash := uint(0)
	for fd > 1 {
		fd /= 2
		hash++
	}
	if !results[0] {
		e.variable = v.vhash(hash)
		e.args = nil
		e.op = nil
	} else {
		e.op = LogicOpString(Not)
		e.variable = 0
		e.args = make([]Expression, 1)
		e.args[0].variable = v.vhash(hash)
	}
}
