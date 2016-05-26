package exp

type Expression struct {
	variable uint
	op       LogicOp
	args     []Expression
}

func (e *Expression) Var() (uint, bool) {
	if e == nil {
		return 0, false
	}
	return e.variable, len(e.args) == 0 && len(e.op) != 1
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
	return &Expression{op: op,
		variable: e.variable,
		args:     args,
	}
}

func New(str string) *Expression {
	e := new(Expression)
	ok := e.parse(str)
	if ok {
		return e
	}
	return nil
}

func (e *Expression) VarIs(v uint, b bool) {
	for i := range e.args {
		e.args[i].VarIs(v, b)
	}
	thisv, ok := e.Var()
	if ok && thisv == v {
		e.variable = 0
		if e.op != nil {
			e.args = make([]Expression, 1)
			e = &e.args[0]
		}
		if b {
			e.op = LogicOpString(True)
		} else {
			e.op = LogicOpString(False)
		}
	}
}

// func (e *Expression) Premise(p *Expression) {
// 	// in case of contradiction, it does nothing
// 	if len(p.op) == 1 {
// 		// nonsense: F=T or T=T
// 		return
// 	} else if v, ok := p.Var(); ok {
// 		if len(p.op) == 0 {
// 			e.VarIs(v, true)
// 		} else if p.op[0] != p.op[1] {
// 			e.VarIs(v, p.op[1])
// 		}
// 		e.Simplify()
// 		return
// 	}
// }
