package arr

import (
	"reflect"
	"testing"
)

func TestIndexRemove(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		index    int
		expected []int
	}{
		{"remove from middle", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"remove first element", []int{1, 2, 3, 4}, 0, []int{2, 3, 4}},
		{"remove last element", []int{1, 2, 3, 4}, 3, []int{1, 2, 3}},
		{"remove from empty slice", []int{}, 0, []int{}},
		{"out of bounds index (negative)", []int{1, 2, 3, 4}, -1, []int{1, 2, 3, 4}},
		{"out of bounds index (greater than slice length)", []int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{"remove from slice with single element", []int{1}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IndexRemove(tt.slice, tt.index)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("IndexRemove(%v, %d) = %v, want %v", tt.slice, tt.index, actual, tt.expected)
			}
		})
	}
}
