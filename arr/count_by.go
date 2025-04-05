package arr

import "github.com/samber/lo"

// CountBy counts the number of elements in the collection for which predicate is true.
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	return lo.CountBy(collection, predicate)
}
