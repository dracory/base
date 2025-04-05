package arr

// Equals checks if two slices are equal.
//
// This function assumes that the slices contain comparable types.
//
// It first checks if the slices have the same length. If not, it
// immediately returns false.
//
// If the slices have the same length, it then checks if the elements
// are equal. If it finds an element that is not equal, it immediately
// returns false.
//
// If it checks all elements and finds no unequal elements, it
// returns true.
func Equals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
