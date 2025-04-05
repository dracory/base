package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		size       int
		expected   [][]int
	}{
		{
			name:       "empty slice",
			collection: []int{},
			size:       2,
			expected:   [][]int{},
		},
		{
			name:       "single element",
			collection: []int{1},
			size:       2,
			expected:   [][]int{{1}},
		},
		{
			name:       "multiple elements, even split",
			collection: []int{1, 2, 3, 4},
			size:       2,
			expected:   [][]int{{1, 2}, {3, 4}},
		},
		{
			name:       "multiple elements, odd split",
			collection: []int{1, 2, 3, 4, 5},
			size:       2,
			expected:   [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:       "size larger than collection",
			collection: []int{1, 2, 3},
			size:       5,
			expected:   [][]int{{1, 2, 3}},
		},
		{
			name:       "size of 1",
			collection: []int{1, 2, 3, 4},
			size:       1,
			expected:   [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:       "size of 0",
			collection: []int{1, 2, 3, 4},
			size:       0,
			expected:   [][]int{},
		},
		{
			name:       "negative size",
			collection: []int{1, 2, 3, 4},
			size:       -1,
			expected:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Split(tt.collection, tt.size)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Split(%v, %d) = %v, want %v", tt.collection, tt.size, actual, tt.expected)
			}
		})
	}
}

func TestSplitString(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		size       int
		expected   [][]string
	}{
		{
			name:       "empty slice",
			collection: []string{},
			size:       2,
			expected:   [][]string{},
		},
		{
			name:       "single element",
			collection: []string{"a"},
			size:       2,
			expected:   [][]string{{"a"}},
		},
		{
			name:       "multiple elements, even split",
			collection: []string{"a", "b", "c", "d"},
			size:       2,
			expected:   [][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			name:       "multiple elements, odd split",
			collection: []string{"a", "b", "c", "d", "e"},
			size:       2,
			expected:   [][]string{{"a", "b"}, {"c", "d"}, {"e"}},
		},
		{
			name:       "size larger than collection",
			collection: []string{"a", "b", "c"},
			size:       5,
			expected:   [][]string{{"a", "b", "c"}},
		},
		{
			name:       "size of 1",
			collection: []string{"a", "b", "c", "d"},
			size:       1,
			expected:   [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
		},
		{
			name:       "size of 0",
			collection: []string{"a", "b", "c", "d"},
			size:       0,
			expected:   [][]string{},
		},
		{
			name:       "negative size",
			collection: []string{"a", "b", "c", "d"},
			size:       -1,
			expected:   [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Split(tt.collection, tt.size)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Split(%v, %d) = %v, want %v", tt.collection, tt.size, actual, tt.expected)
			}
		})
	}
}
