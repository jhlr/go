package ring

type List struct {
	head, last *Ring
	len, cap   int
}

func NewList(n int) *List {
	if n < 0 {
		n = 0
	}
	f := new(List)
	f.last = New(n + 1)
	f.head = f.last
	f.len = 0
	f.cap = n
	return f
}

// Len returns the number of stored elements
func (l *List) Len() int {
	return l.len
}

// Cap returns how many elements can be
// stored without allocating
func (l *List) Cap() int {
	return l.cap
}

func (l *List) Add(e interface{}) bool {
	if l.len == l.cap {
		l.head.Link(New(1))
		l.cap++
	}
	l.head.Value = e
	l.head = l.head.Move(1)
	l.len++
	return l.len == l.cap
}

func (l *List) Rem() (interface{}, bool) {
	e, ok := l.Get()
	if ok {
		l.last.Value = nil
		l.last = l.last.Move(1)
		l.len--
	}
	return e, ok
}

func (l *List) Get() (interface{}, bool) {
	if l.len == 0 {
		return nil, false
	}
	return l.last.Value, true
}
