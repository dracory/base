package main

import (
	"testing"
)

func TestNewConditionalDag(t *testing.T) {
	tests := []struct {
		name        string
		orderType   string
		expected    []string
		totalAmount float64
	}{
		{
			name:        "standard order",
			orderType:   "standard",
			expected:    []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"},
			totalAmount: 126.0, // 100 * 0.9 * 1.2 + 5
		},
		{
			name:        "digital order",
			orderType:   "digital",
			expected:    []string{"ProcessOrder", "ApplyDiscount", "CalculateTax"},
			totalAmount: 108.0, // 100 * 0.9 * 1.2
		},
		{
			name:        "physical order",
			orderType:   "physical",
			expected:    []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"},
			totalAmount: 126.0, // 100 * 0.9 * 1.2 + 5
		},
		{
			name:        "subscription order",
			orderType:   "subscription",
			expected:    []string{"ProcessOrder", "ApplyDiscount"},
			totalAmount: 90.0, // 100 * 0.9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context
			ctx := NewOrderContext(tt.orderType)

			// Create and run DAG
			dag := NewConditionalDag(tt.orderType)
			if err := dag.Run(ctx); err != nil {
				t.Fatal(err)
			}

			// Verify total amount
			if ctx.Get("totalAmount").(float64) != tt.totalAmount {
				t.Fatalf("expected total amount %.2f, got: %.2f", tt.totalAmount, ctx.Get("totalAmount").(float64))
			}

			// Verify steps executed in correct order
			if !equalSlices(ctx.Get("stepsExecuted").([]string), tt.expected) {
				t.Fatalf("steps executed in wrong order: %v, expected: %v", ctx.Get("stepsExecuted").([]string), tt.expected)
			}
		})
	}
}

// equalSlices checks if two slices are equal
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
