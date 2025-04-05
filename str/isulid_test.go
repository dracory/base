package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIsUlid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid ULID",
			input:    "01H7Y605917404784241000000",
			expected: true,
		},
		{
			name:     "Valid ULID with lowercase",
			input:    "01h7y605917404784241000000",
			expected: false,
		},
		{
			name:     "Invalid ULID - Too Short",
			input:    "01H7Y60591740478424100000",
			expected: false,
		},
		{
			name:     "Invalid ULID - Too Long",
			input:    "01H7Y6059174047842410000000",
			expected: false,
		},
		{
			name:     "Invalid ULID - Invalid Characters",
			input:    "01H7Y60591740478424100000!",
			expected: false,
		},
		{
			name:     "Invalid ULID - Empty String",
			input:    "",
			expected: false,
		},
		{
			name:     "Invalid ULID - Special Characters",
			input:    "01H7Y6059174047842410000@@",
			expected: false,
		},
		{
			name:     "Invalid ULID - Lowercase",
			input:    "01h7y605917404784241000000",
			expected: false,
		},
		{
			name:     "Invalid ULID - Mixed Case",
			input:    "01H7y605917404784241000000",
			expected: false,
		},
		{
			name:     "Invalid ULID - Numbers and Special Characters",
			input:    "0123456789!@#$%^&*()_+=-",
			expected: false,
		},
		{
			name:     "Invalid ULID - Only Numbers",
			input:    "01234567890123456789012345",
			expected: true,
		},
		{
			name:     "Invalid ULID - Only Letters",
			input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expected: true,
		},
		{
			name:     "Invalid ULID - Numbers and Letters",
			input:    "0123456789ABCDEFGHIJKLMNOP",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.IsUlid(tc.input)
			if result != tc.expected {
				t.Errorf("IsUlid(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}
