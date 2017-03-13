package bit

// Slice is a slice of bits
type Slice []uint8

// New makes a bit slice of given bit capacity
func New(len int) Slice {
	if len < 0 {
		_ = make([]uint8, len)
	}
	s := make([]uint8, len/8+1)
	return Slice(s)
}

// Len is the bit capacity
func (s Slice) Len() int {
	return len(s) * 8
}

// Fill sets all bits as the given value
func (s Slice) Fill(val bool) {
	mask := uint8(0)
	if val {
		mask = ^mask
	}
	for i := range s {
		s[i] = mask
	}
}

// Set sets individual bits
func (s Slice) Set(i uint, val bool) {
	mask := uint8(1 << (i % 8))
	if val {
		s[i/8] |= mask
	} else {
		s[i/8] &^= mask
	}
}

// Get gets individual bits
func (s Slice) Get(i uint) bool {
	mask := uint8(1 << (i % 8))
	return s[i/8]&mask != 0
}

// Copy will allocate an exact copy
func (s Slice) Copy() Slice {
	result := make([]uint8, len(s))
	copy(result, s)
	return result
}
