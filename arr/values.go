package arr

import (
	"github.com/samber/lo"
)

// Values creates an array of the map values.
func Values[K comparable, V any](in map[K]V) []V {
	return lo.Values(in)
}
