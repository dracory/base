package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsUuid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid UUID",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3d479",
			expected: true,
		},
		{
			name:     "Valid UUID Uppercase",
			input:    "F47AC10B-58CC-4372-A567-0E02B2C3D479",
			expected: true,
		},
		{
			name:     "Invalid UUID - Too Short",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3d47",
			expected: false,
		},
		{
			name:     "Invalid UUID - Too Long",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3d4790",
			expected: false,
		},
		{
			name:     "Invalid UUID - Invalid Characters",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3d47!",
			expected: false,
		},
		{
			name:     "Valid UUID - Missing Hyphens",
			input:    "f47ac10b58cc4372a5670e02b2c3d479",
			expected: true,
		},
		{
			name:     "Invalid UUID - Extra Hyphens",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3-d479",
			expected: false,
		},
		{
			name:     "Invalid UUID - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid UUID - Special Characters",
			input:    "f47ac10b-58cc-4372-a567-0e02b2c3d4@@",
			expected: false,
		},
		{
			name:     "Invalid UUID - Numbers and Special Characters",
			input:    "01234567-8901-2345-6789-0123456789!@",
			expected: false,
		},
		{
			name:     "Valid UUID - Only Numbers",
			input:    "01234567-8901-2345-6789-012345678901",
			expected: true,
		},
		{
			name:     "Invalid UUID - Only Letters",
			input:    "abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef",
			expected: false,
		},
		{
			name:     "Valid UUID - Numbers and Letters",
			input:    "01234567-89ab-cdef-0123-456789abcdef",
			expected: true,
		},
		{
			name:     "Valid UUID - Numbers and Letters - No Dashes",
			input:    "0123456789abcdef0123456789abcdef",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsUuid(tc.input)
			if result != tc.expected {
				t.Errorf("IsUuid(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
