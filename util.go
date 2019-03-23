package algorithms

func slicesEqual(x, y [][]string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if !equal(x[i], y[i]) {
			return false
		}
	}
	return true

}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
