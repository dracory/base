package arr

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Max searches the maximum value of a collection.
func Max[T constraints.Ordered](collection []T) T {
	return lo.Max(collection)
}
