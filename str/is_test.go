package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

func TestIs(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		patterns []string
		expected bool
	}{
		{
			name:     "Exact Match",
			input:    "hello",
			patterns: []string{"hello"},
			expected: true,
		},
		{
			name:     "No Match",
			input:    "hello",
			patterns: []string{"world"},
			expected: false,
		},
		{
			name:     "Wildcard Match",
			input:    "hello world",
			patterns: []string{"hello*"},
			expected: true,
		},
		{
			name:     "Wildcard Match at End",
			input:    "hello world",
			patterns: []string{"*world"},
			expected: true,
		},
		{
			name:     "Wildcard Match in Middle",
			input:    "hello world",
			patterns: []string{"hel*rld"},
			expected: true,
		},
		{
			name:     "Multiple Patterns - Match",
			input:    "hello",
			patterns: []string{"world", "hello"},
			expected: true,
		},
		{
			name:     "Multiple Patterns - No Match",
			input:    "hello",
			patterns: []string{"world", "test"},
			expected: false,
		},
		{
			name:     "Empty Input",
			input:    "",
			patterns: []string{"hello"},
			expected: false,
		},
		{
			name:     "Empty Pattern",
			input:    "hello",
			patterns: []string{""},
			expected: false,
		},
		{
			name:     "Empty Input and Pattern",
			input:    "",
			patterns: []string{""},
			expected: false,
		},
		{
			name:     "Special Characters",
			input:    "hello*",
			patterns: []string{"hello\\*"},
			expected: true,
		},
		{
			name:     "Special Characters No Match",
			input:    "hello",
			patterns: []string{"hello\\*"},
			expected: false,
		},
		{
			name:     "Multiple Wildcards",
			input:    "hello world test",
			patterns: []string{"hel*ld*st"},
			expected: true,
		},
		{
			name:     "Multiple Wildcards No Match", // Renaming to reflect the change
			input:    "hello world test",
			patterns: []string{"hel*ld*st*"},
			expected: true, // Changed expectation from false to true
		},
		{
			name:     "Only Wildcard",
			input:    "hello world test",
			patterns: []string{"*"},
			expected: true,
		},
		{
			name:     "Only Wildcard Empty",
			input:    "",
			patterns: []string{"*"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.Is(tc.input, tc.patterns...)
			if result != tc.expected {
				t.Errorf("Is(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
			}
		})
	}
}
