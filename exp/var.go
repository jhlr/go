package exp

type Variables struct{
	variables map[uint]uint
	values []bool
}

func (e *Expression) Variables() *Variables {
	v := &Variables{
		variables: make(map[uint]uint),
		values: nil,
	}
	count := e.unique(v, 0)
	v.values = make([]bool, count)
	return v
}

func (e *Expression) unique(v *Variables, count uint) uint {
	if vname, ok := e.Var(); ok {
		_, ok = v.variables[vname]
		if !ok {
			v.variables[vname] = count
			count++
		}
		return count
	}
	for i := range e.args {
		count = e.args[i].unique(v, count)
	}
	return count
}

func (v *Variables) vhash(vcode uint) uint {
	for n, c := range v.variables {
		if c == vcode {
			return n
		}
	}
	panic("exp: not found")
}

func (v *Variables) Set(vname uint, b bool) bool {
	vcode, ok := v.variables[vname]
	if !ok {
		return false
	}
	v.values[vcode] = b
	return true
}

func (v *Variables) ForEach(iter func(uint,bool)bool) {
	for vname, vcode := range v.variables {
		v.values[vcode] = iter(vname, v.values[vcode])
	}
}

func (v *Variables) Len() int {
	return len(v.values)
}

func (v *Variables) Get(vname uint) (bool, bool) {
	vcode, ok := v.variables[vname]
	if !ok {
		return false, false
	}
	return v.values[vcode], true
}

func (v *Variables) Increment() bool {
	for i := range v.values {
		if !v.values[i] {
			v.values[i] = true
			return false
		} else {
			v.values[i] = false
		}
	}
	return true
}

func (v *Variables) SetAll(b bool) {
	for i := range v.values {
		v.values[i] = b
	}
}

func (e *Expression) Apply(v *Variables) bool {
	if vname, ok := e.Var(); ok {
		b, ok := v.Get(vname)
		if !ok {
			return false
		}
		return b
	}
	args := make([]bool, len(e.args))
	for i := range args {
		args[i] = e.args[i].Apply(v)
	}
	return e.op.Apply(args...)
}

func (e *Expression) Results() []bool {
	if e == nil {
		return nil
	}
	return e.results(e.Variables())
}

func (e *Expression) results(v *Variables) []bool {
	results := make([]bool, 0, v.Len()*v.Len())
	for {
		results = append(results, e.Apply(v))
		if v.Increment() {
			break
		}
	}
	return results
}
