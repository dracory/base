package arr

import (
	"github.com/samber/lo"
)

// Split returns an array of elements split into groups the length of size. If array can't be split evenly,
func Split[T any](collection []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	return lo.Chunk(collection, size)
}
