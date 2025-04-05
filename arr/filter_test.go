package arr_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/arr"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		predicate  func(item int, index int) bool
		expected   []int
	}{
		{
			name:       "empty slice",
			collection: []int{},
			predicate: func(item int, index int) bool {
				return item > 0
			},
			expected: []int{},
		},
		{
			name:       "all positive",
			collection: []int{1, 2, 3, 4, 5},
			predicate: func(item int, index int) bool {
				return item > 0
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "all negative",
			collection: []int{-1, -2, -3, -4, -5},
			predicate: func(item int, index int) bool {
				return item > 0
			},
			expected: []int{},
		},
		{
			name:       "mixed",
			collection: []int{-1, 0, 1, 2, 3},
			predicate: func(item int, index int) bool {
				return item > 0
			},
			expected: []int{1, 2, 3},
		},
		{
			name:       "even numbers",
			collection: []int{1, 2, 3, 4, 5, 6},
			predicate: func(item int, index int) bool {
				return item%2 == 0
			},
			expected: []int{2, 4, 6},
		},
		{
			name:       "odd numbers",
			collection: []int{1, 2, 3, 4, 5, 6},
			predicate: func(item int, index int) bool {
				return item%2 != 0
			},
			expected: []int{1, 3, 5},
		},
		{
			name:       "greater than 3",
			collection: []int{1, 2, 3, 4, 5, 6},
			predicate: func(item int, index int) bool {
				return item > 3
			},
			expected: []int{4, 5, 6},
		},
		{
			name:       "index based filtering",
			collection: []int{10, 20, 30, 40, 50},
			predicate: func(item int, index int) bool {
				return index%2 == 0 // Keep elements at even indices
			},
			expected: []int{10, 30, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Filter(tt.collection, tt.predicate)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Filter(%v, predicate) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		predicate  func(item string, index int) bool
		expected   []string
	}{
		{
			name:       "empty slice",
			collection: []string{},
			predicate: func(item string, index int) bool {
				return len(item) > 0
			},
			expected: []string{},
		},
		{
			name:       "filter by length",
			collection: []string{"a", "bb", "ccc", "dddd", "eeeee"},
			predicate: func(item string, index int) bool {
				return len(item) > 2
			},
			expected: []string{"ccc", "dddd", "eeeee"},
		},
		{
			name:       "filter by content",
			collection: []string{"apple", "banana", "orange", "grape"},
			predicate: func(item string, index int) bool {
				return item[0] == 'a'
			},
			expected: []string{"apple"},
		},
		{
			name:       "index based filtering",
			collection: []string{"a", "b", "c", "d", "e"},
			predicate: func(item string, index int) bool {
				return index%2 == 0 // Keep elements at even indices
			},
			expected: []string{"a", "c", "e"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Filter(tt.collection, tt.predicate)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Filter(%v, predicate) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
