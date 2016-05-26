package exp

func (e *Expression) Flattened() *Expression {
	v := e.Variables()
	results := e.results(v)
	if len(results) == 2 {
		if results[0] == results[1] {
			return &Expression{op: results[:1]}
		} else if !results[0] && results[1] {
			return &Expression{variable: v.vhash(0)}
		}
		return &Expression{op: results,
			variable: v.vhash(0)}
	}
	e = &Expression{op: results,
		args: make([]Expression, v.Len())}
	for vname, vcode := range v.variables {
		e.args[vcode] = Expression{variable: vname}
	}
	return e
}

func (e *Expression) Simplify() uint {
	if e == nil {
		return 0
	} else if vname, ok := e.Var(); ok {
		if len(e.op) >= 2 && e.op[0] == e.op[1] {
			e.VarIs(vname, e.op[0])
			return 0
		}
		return 1
	}
	v := e.Variables()
	results := e.results(v)
	ul := useless(results)
	for i := range ul {
		vname := v.vhash(ul[i])
		e.VarIs(vname, false)
	}
	vars := uint(0)
	for i := range e.args {
		vars += e.args[i].Simplify()
	}
	if vars <= uint(len(e.args)) {
		*e = *e.Flattened()
	}
	return vars
}
