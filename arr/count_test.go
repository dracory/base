package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 5},
		{"negative numbers", []int{-1, -2, -3}, 3},
		{"mixed numbers", []int{-1, 0, 1}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Count(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Count(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
