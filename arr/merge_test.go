package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "empty slices",
			slices:   [][]int{},
			expected: []int{},
		},
		{
			name:     "single slice",
			slices:   [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "multiple slices",
			slices:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty slice in the middle",
			slices:   [][]int{{1, 2}, {}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "empty slice at the beginning",
			slices:   [][]int{{}, {1, 2}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "empty slice at the end",
			slices:   [][]int{{1, 2}, {3, 4}, {}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "nil slice in the middle",
			slices:   [][]int{{1, 2}, nil, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "nil slice at the beginning",
			slices:   [][]int{nil, {1, 2}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "nil slice at the end",
			slices:   [][]int{{1, 2}, {3, 4}, nil},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "mixed nil and empty slices",
			slices:   [][]int{{1, 2}, nil, {}, {3, 4}, nil, {}},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Merge(tt.slices...)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Merge(%v) = %v, want %v", tt.slices, actual, tt.expected)
			}
		})
	}
}

func TestMergeString(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]string
		expected []string
	}{
		{
			name:     "empty slices",
			slices:   [][]string{},
			expected: []string{},
		},
		{
			name:     "single slice",
			slices:   [][]string{{"a", "b", "c"}},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "multiple slices",
			slices:   [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "empty slice in the middle",
			slices:   [][]string{{"a", "b"}, {}, {"c", "d"}},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "empty slice at the beginning",
			slices:   [][]string{{}, {"a", "b"}, {"c", "d"}},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "empty slice at the end",
			slices:   [][]string{{"a", "b"}, {"c", "d"}, {}},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "nil slice in the middle",
			slices:   [][]string{{"a", "b"}, nil, {"c", "d"}},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "nil slice at the beginning",
			slices:   [][]string{nil, {"a", "b"}, {"c", "d"}},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "nil slice at the end",
			slices:   [][]string{{"a", "b"}, {"c", "d"}, nil},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "mixed nil and empty slices",
			slices:   [][]string{{"a", "b"}, nil, {}, {"c", "d"}, nil, {}},
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Merge(tt.slices...)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Merge(%v) = %v, want %v", tt.slices, actual, tt.expected)
			}
		})
	}
}
