package arr

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](collection []T) T {
	return lo.Sum(collection)
}
