package queue

// CICO custom in, custom out
type CICO struct {
	data    []interface{}
	in, out func() int
}

func NewCICO(cap int, in, out func() int) *CICO {
	if in == nil || out == nil {
		return nil
	}
	return &CICO{
		data: make([]interface{}, 0, cap),
		in:   in, out: out,
	}
}

func (q *CICO) Add(e interface{}) bool {
	if q.Len() == q.Cap() {
		return false
	}
	k := q.in() % (q.Len() + 1)
	if k < 0 {
		k += q.Len() + 1
	}
	q.data = append(q.data, nil)
	for i := q.Len() - 1; i > k; i-- {
		q.data[i] = q.data[i-1]
	}
	q.data[k] = e
	return true
}

func (q *CICO) Rem() (interface{}, bool) {
	if q.Len() == 0 {
		return nil, false
	}
	k := q.out() % q.Len()
	if k < 0 {
		k += q.Len()
	}
	e := q.data[k]
	for i := k; i < q.Len()-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.data = q.data[:q.Len()-1]
	return e, true
}

// Len returns the ammount of stored elements
func (q *CICO) Len() int {
	return len(q.data)
}

// Cap returns the maximum capacity
func (q *CICO) Cap() int {
	return cap(q.data)
}
