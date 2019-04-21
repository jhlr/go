package queue

// LessFunc returns true if (a < b)
type LessFunc func(a, b interface{}) bool

// PIFO priority in, first out
type PIFO struct {
	data []interface{}
	less LessFunc
}

// NewPIFO will allocate a slice with the given cap
// and use the given func as a parameter to sort the
// elements
func NewPIFO(cap int, less LessFunc) *PIFO {
	if less == nil {
		return nil
	}
	return &PIFO{
		data: make([]interface{}, 0, cap),
		less: less,
	}
}

// place is the binary search func
func (s *PIFO) place(e interface{}) int {
	low := 0
	high := s.Len() - 1
	for i := 0; ; i++ {
		if high-low <= 1 {
			if s.less(s.data[high], e) {
				return high + 1
			} else if s.less(s.data[low], e) {
				return low + 1
			}
			return low
		}
		mid := (low + high) / 2
		if s.less(s.data[mid], e) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

// Add will use a binary search to insert the given
// elem. Complexity of O(log(N)+N)
func (s *PIFO) Add(e interface{}) bool {
	if s.Len() == s.Cap() {
		return false
	} else if s.Len() == 0 {
		s.data = append(s.data, e)
		return true
	}
	p := s.place(e)
	s.data = append(s.data, nil)
	for i := s.Len() - 1; i > p; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[p] = e
	return true
}

// Rem will remove the highest priority elem.
// Complexity of O(1)
func (s *PIFO) Rem() (interface{}, bool) {
	if s.Len() == 0 {
		return nil, false
	}
	l := s.Len() - 1
	elem := s.data[l]
	s.data = s.data[:l]
	return elem, true
}

// Len is the ammount of stored elements
func (s *PIFO) Len() int {
	return len(s.data)
}

// Cap is the maximum ammount of elements
func (s *PIFO) Cap() int {
	return cap(s.data)
}
