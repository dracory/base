package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestEach(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   []int // Expected values after the iteratee is applied
	}{
		{
			name:       "empty slice",
			collection: []int{},
			expected:   []int{},
		},
		{
			name:       "single element",
			collection: []int{1},
			expected:   []int{2}, // Assuming iteratee adds 1
		},
		{
			name:       "multiple elements",
			collection: []int{1, 2, 3, 4, 5},
			expected:   []int{2, 3, 4, 5, 6}, // Assuming iteratee adds 1
		},
		{
			name:       "negative numbers",
			collection: []int{-1, -2, -3},
			expected:   []int{0, -1, -2}, // Assuming iteratee adds 1
		},
		{
			name:       "mixed numbers",
			collection: []int{-1, 0, 1},
			expected:   []int{0, 1, 2}, // Assuming iteratee adds 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the collection to avoid modifying the original test data
			collectionCopy := make([]int, len(tt.collection))
			copy(collectionCopy, tt.collection)

			// Create a copy of the collection to avoid modifying the original test data
			expectedCopy := make([]int, len(tt.expected))
			copy(expectedCopy, tt.expected)

			// Apply the iteratee function to the collection
			arr.Each(collectionCopy, func(item int, index int) {
				collectionCopy[index] = item + 1
			})

			// Check if the modified collection matches the expected values
			if len(collectionCopy) != len(expectedCopy) {
				t.Errorf("arr.Each(%v) modified collection length = %v, want %v", tt.collection, len(collectionCopy), len(expectedCopy))
			}

			for i := range collectionCopy {
				if collectionCopy[i] != expectedCopy[i] {
					t.Errorf("arr.Each(%v) modified collection = %v, want %v", tt.collection, collectionCopy, expectedCopy)
					break
				}
			}
		})
	}
}
