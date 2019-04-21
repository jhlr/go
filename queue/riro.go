package queue

import (
	"math/rand"
)

// RIRO random in, random out
type RIRO struct {
	data []interface{}
	rand rand.Rand
}

// NewRIRO allocates a slice with the given cap
// and seeds itself
func NewRIRO(cap int, seed int64) *RIRO {
	s := &RIRO{
		data: make([]interface{}, 0, cap),
	}
	s.rand.Seed(seed)
	return s
}

// Add inserts the given elem into a random
// location or, if full, returns false.
// Complexity of O(1)
func (q *RIRO) Add(e interface{}) bool {
	if q.Len() == q.Cap() {
		return false
	} else if q.Len() == 0 {
		q.data = append(q.data, e)
		return true
	}
	r := q.rand.Int() % q.Len()
	q.data = append(q.data, q.data[r])
	q.data[r] = e
	return true
}

// Rem removes an elem from a random location
// or, if empty, returns (nil, false).
// Complexity of O(1)
func (q *RIRO) Rem() (interface{}, bool) {
	if q.Len() == 0 {
		return nil, false
	}
	r := q.rand.Int() % q.Len()
	elem := q.data[r]
	q.data[r] = q.data[q.Len()-1]
	q.data = q.data[:q.Len()-1]
	return elem, true
}

// Len returns the ammount of stored elements
func (q *RIRO) Len() int {
	return len(q.data)
}

// Cap returns the maximum capacity
func (q *RIRO) Cap() int {
	return cap(q.data)
}
