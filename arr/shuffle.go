package arr

import (
	"github.com/samber/lo/mutable"
)

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
func Shuffle[T any](collection []T) []T {
	mutable.Shuffle(collection)
	return collection
}
