package arr

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Min search the minimum value of a collection.
func Min[T constraints.Ordered](collection []T) T {
	return lo.Min(collection)
}
