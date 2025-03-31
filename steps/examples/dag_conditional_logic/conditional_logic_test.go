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
		expectedAmount float64
	}{
		{"Digital Order", "digital", 100.0, []string{"ProcessOrder", "ApplyDiscount", "CalculateTax"}, 108.0},
		{"Physical Order", "physical", 100.0, []string{"ProcessOrder", "ApplyDiscount", "AddShipping", "CalculateTax"}, 114.0},
		{"Subscription Order", "subscription", 100.0, []string{"ProcessOrder", "ApplyDiscount"}, 90.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := RunConditionalExample(tc.orderType, tc.totalAmount)
			if err != nil {
				t.Errorf("Error running DAG: %v", err)
				return
			}

			stepsExecuted := data["stepsExecuted"].([]string)
			if len(stepsExecuted) != len(tc.expectedSteps) {
				t.Errorf("Expected %d steps, got %d", len(tc.expectedSteps), len(stepsExecuted))
				return
			}

			for i, step := range stepsExecuted {
				if step != tc.expectedSteps[i] {
					t.Errorf("Expected step %d to be %s, got %s", i, tc.expectedSteps[i], step)
				}
			}

			totalAmount := data["totalAmount"].(float64)
			if totalAmount != tc.expectedAmount {
				t.Errorf("Expected total amount %.2f, got %.2f", tc.expectedAmount, totalAmount)
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
