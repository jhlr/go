package queue

type LIFO struct {
	data []interface{}
}

// NewLIFO allocates a slice of the given size
func NewLIFO(cap int) *LIFO {
	return &LIFO{
		data: make([]interface{}, 0, cap),
	}
}

// Add pushes the element onto the stack.
// Complexity of O(1)
func (q *LIFO) Add(e interface{}) (ok bool) {
	if q.Len() == q.Cap() {
		return false
	}
	q.data = append(q.data, e)
	return true
}

// Rem will pop an element. Complexity of O(1)
func (q *LIFO) Rem() (interface{}, bool) {
	if q.Len() == 0 {
		return nil, false
	}
	l := q.Len() - 1
	elem := q.data[l]
	q.data = q.data[:l]
	return elem, true
}

// Len is the ammount of elements stored
func (q *LIFO) Len() int {
	return len(q.data)
}

// Cap is the maximum capacity
func (q *LIFO) Cap() int {
	return cap(q.data)
}
