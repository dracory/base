package arr

import "github.com/samber/lo"

// Each iterates over elements of collection and invokes iteratee for each element.
func Each[T any](collection []T, iteratee func(item T, index int)) {
	lo.ForEach(collection, iteratee)
}
