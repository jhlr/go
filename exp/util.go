package exp

func floorLog(base, number uint) uint {
	result := uint(0)
	for number >= base {
		number /= base
		result++
	}
	return result
}

func useless(results []bool) []uint {
	l := uint(len(results))
	v :=  make([]uint, 0, floorLog(2, l))
	for i := l/2; i > 0; i /= 2 {
		broken := false
		for j := i; j < l; j++ {
			if results[j-i] != results[j] {
				broken = true
				break
			}
		}
		if !broken {
			v = append(v, floorLog(2, i))
		}
	}
	return v
}

// func pattern(results []bool) (uint, bool) {
// 	fd := uint(0)
// 	for i := 1; i < len(results); i++ {
// 		if results[i-1] != results[i] {
// 			fd = uint(i)
// 			break
// 		}
// 	}
// 	if fd == 0 {
// 		return 0, true
// 	} else if 2*fd > uint(len(results)) {
// 		return fd, false
// 	}
// 	equ := true
// 	not := true
// 	b := results[0]
// 	for i := 0; i < len(results); i++ {
// 		if i > 0 && uint(i)%fd == 0 {
// 			b = !b
// 		}
// 		if results[i] != b {
// 			equ = false
// 		} else {
// 			not = false
// 		}
// 	}
// 	if equ || not {
// 		return fd, true
// 	}
// 	return 0, false
// }
