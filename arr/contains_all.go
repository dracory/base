package arr

// ContainsAll checks if slice 'a' contains all elements of slice 'b', regardless of order.
func ContainsAll[T comparable](a, b []T) bool {
	if len(b) == 0 {
		return true // An empty slice is considered to be contained in any slice.
	}

	if len(a) < len(b) {
		return false // If 'a' is shorter than 'b', it can't contain all elements of 'b'.
	}

	// Count element occurrences in both slices.
	countA := make(map[T]int)
	countB := make(map[T]int)

	for _, item := range a {
		countA[item]++
	}
	for _, item := range b {
		countB[item]++
	}

	// Check if all elements in 'b' are present in 'a' with sufficient counts.
	for item, count := range countB {
		if countA[item] < count {
			return false
		}
	}

	return true
}
