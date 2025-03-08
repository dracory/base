package arr

import (
	"reflect"
	"testing"
)

func TestIndexMoveUp(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{"move up from middle", []int{1, 2, 3, 4, 5}, 2, []int{1, 3, 2, 4, 5}},
		{"move up from last", []int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 5, 4}},
		{"move up from first", []int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{"out of bounds", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{"empty slice", []int{}, 0, []int{}},
		{"single element", []int{1}, 0, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IndexMoveUp(tt.slice, tt.index)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("IndexMoveUp(%v, %d) = %v, want %v", tt.slice, tt.index, actual, tt.expected)
			}
		})
	}
}
