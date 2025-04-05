package arr_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/dracory/base/arr"
)

func TestGroupBy(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		iteratee   func(item int) int
		expected   map[int][]int
	}{
		{
			name:       "empty slice",
			collection: []int{},
			iteratee: func(item int) int {
				return item % 2
			},
			expected: map[int][]int{},
		},
		{
			name:       "even and odd",
			collection: []int{1, 2, 3, 4, 5, 6},
			iteratee: func(item int) int {
				return item % 2
			},
			expected: map[int][]int{
				0: {2, 4, 6},
				1: {1, 3, 5},
			},
		},
		{
			name:       "all even",
			collection: []int{2, 4, 6, 8},
			iteratee: func(item int) int {
				return item % 2
			},
			expected: map[int][]int{
				0: {2, 4, 6, 8},
			},
		},
		{
			name:       "all odd",
			collection: []int{1, 3, 5, 7},
			iteratee: func(item int) int {
				return item % 2
			},
			expected: map[int][]int{
				1: {1, 3, 5, 7},
			},
		},
		{
			name:       "group by tens",
			collection: []int{1, 10, 11, 20, 21, 30},
			iteratee: func(item int) int {
				return item / 10
			},
			expected: map[int][]int{
				0: {1},
				1: {10, 11},
				2: {20, 21},
				3: {30},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.GroupBy(tt.collection, tt.iteratee)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.GroupBy(%v, iteratee) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestGroupByString(t *testing.T) {
	tests := []struct {
		name       string
		collection []string
		iteratee   func(item string) string
		expected   map[string][]string
	}{
		{
			name:       "empty slice",
			collection: []string{},
			iteratee: func(item string) string {
				return item
			},
			expected: map[string][]string{},
		},
		{
			name:       "group by first letter",
			collection: []string{"apple", "banana", "apricot", "blueberry", "cantaloupe"},
			iteratee: func(item string) string {
				return item[:1]
			},
			expected: map[string][]string{
				"a": {"apple", "apricot"},
				"b": {"banana", "blueberry"},
				"c": {"cantaloupe"},
			},
		},
		{
			name:       "group by length",
			collection: []string{"a", "bb", "ccc", "dd", "eee"},
			iteratee: func(item string) string {
				return strconv.Itoa(len(item)) // Corrected line
			},
			expected: map[string][]string{
				"1": {"a"},
				"2": {"bb", "dd"},
				"3": {"ccc", "eee"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.GroupBy(tt.collection, tt.iteratee)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("arr.GroupBy(%v, iteratee) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
