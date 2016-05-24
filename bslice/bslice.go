package bslice

type BitSlice []uint8

func New(len uint) BitSlice {
	s := make([]uint8, len/8+1)
	return BitSlice(s)
}

func (s BitSlice) Fill(val bool) {
	mask := uint8(0)
	if val {
		mask = ^mask
	}
	for i := range s {
		s[i] = mask
	}
}

func (s BitSlice) Set(i uint, val bool) {
	mask := uint8(1 << (i % 8))
	if val {
		s[i/8] |= mask
	} else {
		s[i/8] &^= mask
	}
}

func (s BitSlice) Get(i uint) bool {
	mask := uint8(1 << (i % 8))
	return s[i/8]&mask != 0
}

func (s BitSlice) Copy() BitSlice {
	result := make([]uint8, len(s))
	copy(result, s)
	return result
}

func (s BitSlice) HashCode() uint32 {
	const (
		init uint32 = 5381
		mnum uint32 = 33
	)
	result := init
	for _, c := range s {
		result = result*mnum + uint32(c)
	}
	return result
}
