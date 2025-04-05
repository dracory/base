package arr

import (
	"github.com/samber/lo"
)

// Unique returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
func Unique[T comparable](collection []T) []T {
	return lo.Uniq(collection)
}
