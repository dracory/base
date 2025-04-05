package arr_test

import (
	"testing"

	"github.com/dracory/base/arr"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed numbers", []int{-1, 0, 1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Sum(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Sum(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestSumFloat(t *testing.T) {
	tests := []struct {
		name       string
		collection []float64
		expected   float64
	}{
		{"empty slice", []float64{}, 0},
		{"single element", []float64{1.1}, 1.1},
		{"multiple elements", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 16.5},
		{"negative numbers", []float64{-1.1, -2.2, -3.3}, -6.6},
		{"mixed numbers", []float64{-1.1, 0, 1.1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Sum(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Sum(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}

func TestSumComplex(t *testing.T) {
	tests := []struct {
		name       string
		collection []complex128
		expected   complex128
	}{
		{"empty slice", []complex128{}, 0},
		{"single element", []complex128{1 + 1i}, 1 + 1i},
		{"multiple elements", []complex128{1 + 1i, 2 + 2i, 3 + 3i}, 6 + 6i},
		{"negative numbers", []complex128{-1 - 1i, -2 - 2i, -3 - 3i}, -6 - 6i},
		{"mixed numbers", []complex128{-1 - 1i, 0, 1 + 1i}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := arr.Sum(tt.collection)
			if actual != tt.expected {
				t.Errorf("arr.Sum(%v) = %v, want %v", tt.collection, actual, tt.expected)
			}
		})
	}
}
