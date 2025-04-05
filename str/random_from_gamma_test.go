package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestRandomFromGamma(t *testing.T) {
	testCases := []struct {
		name          string
		length        int
		gamma         string
		expectedError bool
		expectedLen   int
	}{
		{
			name:          "Zero Length",
			length:        0,
			gamma:         "abc",
			expectedError: false,
			expectedLen:   0,
		},
		{
			name:          "Positive Length",
			length:        5,
			gamma:         "abc",
			expectedError: false,
			expectedLen:   5,
		},
		{
			name:          "Empty Gamma",
			length:        5,
			gamma:         "",
			expectedError: false,
			expectedLen:   0,
		},
		{
			name:          "Large Length",
			length:        100,
			gamma:         "abcdefghijklmnopqrstuvwxyz",
			expectedError: false,
			expectedLen:   100,
		},
		{
			name:          "Single Char Gamma",
			length:        10,
			gamma:         "a",
			expectedError: false,
			expectedLen:   10,
		},
		{
			name:          "Unicode Gamma",
			length:        10,
			gamma:         "привет",
			expectedError: false,
			expectedLen:   10,
		},
		{
			name:          "Special Chars Gamma",
			length:        10,
			gamma:         "!@#$%^&*()",
			expectedError: false,
			expectedLen:   10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.RandomFromGamma(tc.length, tc.gamma)
			if len([]rune(result)) != tc.expectedLen {
				t.Errorf("RandomFromGamma(%d, %q) = %d, want %d", tc.length, tc.gamma, len([]rune(result)), tc.expectedLen)
			}

			if tc.expectedLen > 0 && tc.gamma != "" {
				if !str.ContainsOnly(result, tc.gamma) {
					t.Errorf("RandomFromGamma(%d, %q) = %q, want string containing only chars from gamma", tc.length, tc.gamma, result)
				}
			}
		})
	}
}

func TestRandomFromGamma_Uniqueness(t *testing.T) {
	// Test for uniqueness by generating multiple strings and checking for duplicates
	numStrings := 100
	length := 10
	gamma := "abcdefghijklmnopqrstuvwxyz"
	generatedStrings := make(map[string]bool)

	for i := 0; i < numStrings; i++ {
		randomStr := str.RandomFromGamma(length, gamma)
		if _, exists := generatedStrings[randomStr]; exists {
			t.Errorf("Duplicate string generated: %s", randomStr)
		}
		generatedStrings[randomStr] = true
	}
}
