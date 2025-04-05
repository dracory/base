package arr

import (
	"math/rand/v2"
)

// Random returns a random element from the provided slice.
// It uses Go generics, so it works with slices of any comparable type T.
// If the slice is empty, it returns the zero value of type T and false.
func Random[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	// Use rand.N from math/rand/v2 for a random index
	randomIndex := rand.N(len(slice))
	return slice[randomIndex]
}
