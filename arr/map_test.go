package arr_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/dracory/base/arr"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		iteratee   func(item int, index int) string
		expected   []string
	}{
		{
			name:       "empty slice",
			collection: []int{},
			iteratee: func(item int, index int) string {
				return ""
			},
			expected: []string{},
		},
		{
			name:       "single element",
			collection: []int{1},
			iteratee: func(item int, index int) string {
				return "number: " + strconv.Itoa(item)
			},
			expected: []string{"number: 1"},
		},
		{
			name:       "multiple elements",
			collection: []int{1, 2, 3, 4, 5},
			iteratee: func(item int, index int) string {
				return "number: " + strconv.Itoa(item)
			},
			expected: []string{"number: 1", "number: 2", "number: 3", "number: 4", "number: 5"},
		},
		{
			name:       "negative numbers",
			collection: []int{-1, -2, -3},
			iteratee: func(item int, index int) string {
				return "number: " + strconv.Itoa(item)
			},
			expected: []string{"number: -1", "number: -2", "number: -3"},
		},
		{
			name:       "mixed numbers",
			collection: []int{-1, 0, 1},
			iteratee: func(item int, index int) string {
				return "number: " + strconv.Itoa(item)
			},
			expected: []string{"number: -1", "number: 0", "number: 1"},
		},
		{
			name:       "index based mapping",
			collection: []int{10, 20, 30},
			iteratee: func(item int, index int) string {
				return "index: " + strconv.Itoa(index) + ", value: " + strconv.Itoa(item/10)
			},
			expected: []string{"index: 0, value: 1", "index: 1, value: 2", "index: 2, value: 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Map(tt.collection, tt.iteratee)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Map(%v, iteratee) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestMapString(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		iteratee   func(item string, index int) int
		expected   []int
	}{
		{
			name:       "empty slice",
			collection: []string{},
			iteratee: func(item string, index int) int {
				return 0
			},
			expected: []int{},
		},
		{
			name:       "single element",
			collection: []string{"a"},
			iteratee: func(item string, index int) int {
				return len(item)
			},
			expected: []int{1},
		},
		{
			name:       "multiple elements",
			collection: []string{"a", "bb", "ccc", "dddd", "eeeee"},
			iteratee: func(item string, index int) int {
				return len(item)
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:       "index based mapping",
			collection: []string{"a", "b", "c"},
			iteratee: func(item string, index int) int {
				return index
			},
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Map(tt.collection, tt.iteratee)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.Map(%v, iteratee) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
