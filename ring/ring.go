package ring

// Ring is a single linked list
type Ring struct {
	Value interface{}
	next  *Ring
}

// New will make a ring with N elements
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	temp := r
	for i := 1; i < n; i++ {
		temp.next = new(Ring)
		temp = temp.next
	}
	temp.next = r
	return r
}

// Len rotates the ring counting the elements.
// Complexity of O(n)
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

// Move will rotate R to the given side.
// Negative I will call r.Len()
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

// Link will insert other between r and r.next
func (r *Ring) Link(other *Ring) *Ring {
	if other == nil {
		return r
	} else if r == nil {
		return other
	}
	last := other.Move(-1)
	temp := r.next
	r.next = other
	last.next = temp
	return temp
}

// Unlink will select N elements forward or backwards
// remove them from R and make a new ring with them
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
