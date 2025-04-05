package str_test

import (
	"regexp"
	"testing"

	"github.com/dracory/base/str"
)

func TestRandom(t *testing.T) {
	testCases := []struct {
		name          string
		length        int
		expectedError bool
	}{
		{"Zero Length", 0, false},
		{"Positive Length", 10, false},
		{"Large Length", 100, false},
		{"Negative Length", -1, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectedError {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Random(%d) did not panic, but expected to", tc.length)
					}
				}()
				str.Random(tc.length)
			} else {
				result := str.Random(tc.length)
				if len(result) != tc.length {
					t.Errorf("Random(%d) = %d, want %d", tc.length, len(result), tc.length)
				}

				// Check if the result is a valid base64 string
				matched, _ := regexp.MatchString(`^[a-zA-Z0-9+/]*={0,2}$`, result)
				if !matched && tc.length > 0 {
					t.Errorf("Random(%d) = %q, want a valid base64 string", tc.length, result)
				}
			}
		})
	}
}

func TestRandom_Uniqueness(t *testing.T) {
	// Test for uniqueness by generating multiple strings and checking for duplicates
	numStrings := 100
	length := 10
	generatedStrings := make(map[string]bool)

	for i := 0; i < numStrings; i++ {
		randomStr := str.Random(length)
		if _, exists := generatedStrings[randomStr]; exists {
			t.Errorf("Duplicate string generated: %s", randomStr)
		}
		generatedStrings[randomStr] = true
	}
}
