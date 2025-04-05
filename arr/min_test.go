package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{"empty slice", []int{}, 0}, // Should return zero value for empty slice
		{"single element", []int{1}, 1},
		{"multiple elements", []int{5, 4, 3, 2, 1}, 1},
		{"negative numbers", []int{-1, -2, -3}, -3},
		{"mixed numbers", []int{-1, 0, 1}, -1},
		{"duplicate min", []int{5, 1, 2, 1, 3}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Min(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Min(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestMinFloat(t *testing.T) {
	tests := []struct {
		name       string
		collection []float64
		expected   float64
	}{
		{"empty slice", []float64{}, 0}, // Should return zero value for empty slice
		{"single element", []float64{1.1}, 1.1},
		{"multiple elements", []float64{5.5, 4.4, 3.3, 2.2, 1.1}, 1.1},
		{"negative numbers", []float64{-1.1, -2.2, -3.3}, -3.3},
		{"mixed numbers", []float64{-1.1, 0, 1.1}, -1.1},
		{"duplicate min", []float64{5.5, 1.1, 2.2, 1.1, 3.3}, 1.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Min(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Min(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
