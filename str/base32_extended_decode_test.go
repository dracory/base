package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestBase32ExtendedDecode(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid Base32 Extended",
			input:       "JBSWY3DPEB3W64TMMQ======",
			expected:    "Hello world",
			expectError: false,
		},
		{
			name:        "Empty String",
			input:       "",
			expected:    "",
			expectError: false,
		},
		{
			name:        "Invalid Base32 Extended",
			input:       "JBSWY3DP!B3W64TMMQ======",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			decoded, err := str.Base32ExtendedDecode([]byte(tc.input))

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if string(decoded) != tc.expected {
					t.Errorf("Expected decoded string to be %q, but got %q", tc.expected, string(decoded))
				}
			}
		})
	}
}
