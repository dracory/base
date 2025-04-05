package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestCountBy(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		predicate  func(item int) bool
		expected   int
	}{
		{
			"empty slice",
			[]int{},
			func(item int) bool { return item > 0 },
			0,
		},
		{
			"all positive",
			[]int{1, 2, 3, 4, 5},
			func(item int) bool { return item > 0 },
			5,
		},
		{
			"all negative",
			[]int{-1, -2, -3, -4, -5},
			func(item int) bool { return item > 0 },
			0,
		},
		{
			"mixed",
			[]int{-1, 0, 1, 2, 3},
			func(item int) bool { return item > 0 },
			3,
		},
		{
			"even numbers",
			[]int{1, 2, 3, 4, 5, 6},
			func(item int) bool { return item%2 == 0 },
			3,
		},
		{
			"odd numbers",
			[]int{1, 2, 3, 4, 5, 6},
			func(item int) bool { return item%2 != 0 },
			3,
		},
		{
			"greater than 3",
			[]int{1, 2, 3, 4, 5, 6},
			func(item int) bool { return item > 3 },
			3,
		},
		{
			"empty slice with different predicate",
			[]int{},
			func(item int) bool { return item%2 == 0 },
			0,
		},
		{
			"all zeros",
			[]int{0, 0, 0, 0},
			func(item int) bool { return item == 0 },
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.CountBy(tt.collection, tt.predicate)
			if actual != tt.expected {
				t.Errorf("arr.CountBy(%v, predicate) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
