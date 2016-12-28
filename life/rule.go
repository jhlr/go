package life

// Rule bool is the next state given the arguments
// [2] is the state of the cell
// [9] is the number of neighbors
type Rule [2][9]bool

const (
	t = true
	f = false
)

// ConwayRule B3/S23
func ConwayRule() Rule {
	return [2][9]bool{
		{f, f, f, t, f, f, f, f, f},
		{f, f, t, t, f, f, f, f, f},
	}
}

// SeedsRule B3/S
func SeedsRule() Rule {
	return [2][9]bool{
		{f, f, t, f, f, f, f, f, f},
		{f, f, f, f, f, f, f, f, f},
	}
}

// MorleyRule B368/S245
func MorleyRule() Rule {
	return [2][9]bool{
		{f, f, f, t, f, f, t, f, t},
		{f, f, t, f, t, t, f, f, f},
	}
}

// DiamoebaRule B35678/S678
func DiamoebaRule() Rule {
	return [2][9]bool{
		{f, f, f, t, f, t, t, t, t},
		{f, f, f, f, f, t, t, t, t},
	}
}

// AnnealRule B4678/S35678
func AnnealRule() Rule {
	return [2][9]bool{
		{f, f, f, f, t, f, t, t, t},
		{f, f, f, t, f, t, t, t, t},
	}
}
