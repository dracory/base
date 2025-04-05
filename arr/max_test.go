package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{"empty slice", []int{}, 0}, // Should return zero value for empty slice
		{"single element", []int{1}, 1},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 5},
		{"negative numbers", []int{-1, -2, -3}, -1},
		{"mixed numbers", []int{-1, 0, 1}, 1},
		{"duplicate max", []int{1, 5, 2, 5, 3}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Max(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Max(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestMaxFloat(t *testing.T) {
	tests := []struct {
		name       string
		collection []float64
		expected   float64
	}{
		{"empty slice", []float64{}, 0}, // Should return zero value for empty slice
		{"single element", []float64{1.1}, 1.1},
		{"multiple elements", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 5.5},
		{"negative numbers", []float64{-1.1, -2.2, -3.3}, -1.1},
		{"mixed numbers", []float64{-1.1, 0, 1.1}, 1.1},
		{"duplicate max", []float64{1.1, 5.5, 2.2, 5.5, 3.3}, 5.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Max(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Max(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
