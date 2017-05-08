package ring

type Ring struct {
	Value interface{}
	next  *Ring
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = new(Ring)
		p = p.next
	}
	p.next = r
	return r
}

func (r *Ring) Len() int {
	i := 0
	temp := r
	for temp != nil {
		i++
		temp = temp.next
		if temp == r {
			break
		}
	}
	return i
}

func (r *Ring) Move(i int) *Ring {
	if r == nil {
		return nil
	} else if i < 0 {
		len := r.Len()
		i = i % len
		if i < 0 {
			i += len
		}
	}
	for j := 0; j < i; j++ {
		r = r.next
	}
	return r
}

func (r *Ring) Link(nd *Ring) *Ring {
	if nd == nil {
		return r
	} else if r == nil {
		return nd
	}
	last := nd.Move(-1)
	temp := r.next
	r.next = nd
	last.next = temp
	return temp
}

func (r *Ring) Unlink(n int) *Ring {
	if n == 0 {
		return nil
	} else if r == nil || r.next == r {
		return r
	} else if n < 0 {
		r = r.Move(n - 1)
		n = -n
	}
	last := r.Move(n)
	first := r.next
	r.next = last.next
	last.next = first
	return first
}

func (r *Ring) Do(f func(interface{})) {
	for r != nil {
		f(r.Value)
		r = r.next
	}
}
