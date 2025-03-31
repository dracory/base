package main

import (
	"testing"
)

func TestConditionalLogic(t *testing.T) {
	// Create test cases
	testCases := []struct {
		name        string
		orderType   string
		totalAmount float64
		expectedSteps []string
	}{
		{"Digital Order", "digital", 100.0, []string{"ProcessOrder", "ApplyDiscount"}},
		{"Physical Order", "physical", 100.0, []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"}},
		{"Subscription Order", "subscription", 100.0, []string{"ProcessOrder", "ApplyDiscount"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := NewConditionalDag(tc.orderType, tc.totalAmount)
			ctx := NewOrderContext(tc.orderType, tc.totalAmount)
			ctx, err := dag.Run(ctx)
			if err != nil {
				t.Errorf("Error running DAG: %v", err)
				return
			}

			stepsExecuted := ctx.Get("stepsExecuted").([]string)
			if len(stepsExecuted) != len(tc.expectedSteps) {
				t.Errorf("Expected %d steps, got %d", len(tc.expectedSteps), len(stepsExecuted))
			}

			for i, step := range stepsExecuted {
				if step != tc.expectedSteps[i] {
					t.Errorf("Expected step %d to be %s, got %s", i, tc.expectedSteps[i], step)
				}
			}
		})
	}
}

func TestNewConditionalDag(t *testing.T) {
	tests := []struct {
		name        string
		orderType   string
		totalAmount float64
		expected    []string
		finalAmount float64
	}{
		{
			name:        "standard order",
			orderType:   "standard",
			totalAmount: 100.0,
			expected:    []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"},
			finalAmount: 126.0, // 100 * 0.9 * 1.2 + 5
		},
		{
			name:        "digital order",
			orderType:   "digital",
			totalAmount: 100.0,
			expected:    []string{"ProcessOrder", "ApplyDiscount", "CalculateTax"},
			finalAmount: 108.0, // 100 * 0.9 * 1.2
		},
		{
			name:        "physical order",
			orderType:   "physical",
			totalAmount: 100.0,
			expected:    []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"},
			finalAmount: 126.0, // 100 * 0.9 * 1.2 + 5
		},
		{
			name:        "subscription order",
			orderType:   "subscription",
			totalAmount: 100.0,
			expected:    []string{"ProcessOrder", "ApplyDiscount"},
			finalAmount: 90.0, // 100 * 0.9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context
			ctx := NewOrderContext(tt.orderType, tt.totalAmount)

			// Create and run DAG
			dag := NewConditionalDag(tt.orderType, tt.totalAmount)
			ctx, err := dag.Run(ctx)
			if err != nil {
				t.Fatal(err)
			}

			// Verify total amount
			actualAmount := ctx.Get("totalAmount").(float64)
			if actualAmount != tt.finalAmount {
				t.Fatalf("expected total amount %.2f, got: %.2f", tt.finalAmount, actualAmount)
			}

			// Verify steps executed in correct order
			actualSteps := ctx.Get("stepsExecuted").([]string)
			if !equalSlices(actualSteps, tt.expected) {
				t.Fatalf("steps executed in wrong order: %v, expected: %v", actualSteps, tt.expected)
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
