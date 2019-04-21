package bit

// Slice is a slice of bits
type Slice struct {
	data []uint8
	len  uint
}

// BoundError is panicked when
// the accessed value is bigger than len
type BoundError struct{}

func (e BoundError) Error() string {
	return "out of bounds"
}

// Make makes a bit slice of given bit capacity
func Make(len uint, cap uint) *Slice {
	if len > cap {
		panic(BoundError{})
	}
	s := new(Slice)
	s.len = len
	cap += cap % 8
	s.data = make([]uint8, cap/8)
	return s
}

// Len is the bit length
func (s *Slice) Len() uint {
	return s.len
}

// Cap is the bit capacity
func (s *Slice) Cap() uint {
	return uint(len(s.data)) * 8
}

// Fill sets all bits as the given value
func (s *Slice) Fill(val bool) {
	mask := uint8(0)
	if val {
		mask = ^mask
	}
	for i := s.len - 1; i >= 0; i-- {
		s.data[i] = mask
	}
}

// Set individual bits
func (s *Slice) Set(i uint, val bool) {
	if i >= s.len {
		panic(BoundError{})
	}
	mask := uint8(1 << (i % 8))
	if val {
		s.data[i/8] |= mask
	} else {
		s.data[i/8] &^= mask
	}
}

// Get individual bits
func (s *Slice) Get(i uint) bool {
	if i >= s.len {
		panic(BoundError{})
	}
	mask := uint8(1 << (i % 8))
	return s.data[i/8]&mask != 0
}

// Append works like the builtin append
func (s *Slice) Append(vals ...bool) {
	n := uint(len(vals))
	if s.len > s.Cap() {
		len := s.len + n
		len = (len + len%8) / 8
		temp := make([]uint8, len)
		copy(temp, s.data)
		s.data = temp
	}
	for i, b := range vals {
		if b {
			s.Set(s.len+uint(i), b)
		}
	}
	s.len += n
}

// SetLength will panic if argument is more than s.Len()
func (s *Slice) SetLength(len uint) {
	if len <= s.Len() {
		s.len = len
	} else {
		panic(BoundError{})
	}
}
