package queue

// FIFO works cycling the elements through
// the slice so the elements dont have to shift
type FIFO struct {
	data       []interface{}
	first, len int
}

// NewFIFO allocates a slice of given size
func NewFIFO(cap int) *FIFO {
	return &FIFO{
		data:  make([]interface{}, cap, cap),
		first: 0, len: 0,
	}
}

// Add will place an elem to the end of the queue
func (q *FIFO) Add(e interface{}) bool {
	if q.len == q.Cap() {
		return false
	}
	i := (q.first + q.len) % q.Cap()
	q.data[i] = e
	q.len++
	return true
}

// Rem will pop an elem from the start of the queue
func (q *FIFO) Rem() (interface{}, bool) {
	if q.len == 0 {
		return nil, false
	}
	e := q.data[q.first]
	q.data[q.first] = nil
	q.first = (q.first + 1) % q.Cap()
	q.len--
	return e, true
}

// Len is the ammount of stored elem
func (q *FIFO) Len() int {
	return q.len
}

// Cap is the maximum capacity
func (s *FIFO) Cap() int {
	return len(s.data)
}
