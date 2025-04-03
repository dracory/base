package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestIsMatchBasic tests basic matching scenarios
func TestIsMatchBasic(t *testing.T) {
	testCases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"Hello World", []string{"Hello"}, true},
		{"Hello World", []string{"World"}, true},
		{"Hello World", []string{"Hello", "World"}, true},
		{"Hello World", []string{"Hello.*World"}, true},
		{"Hello World", []string{"^Hello"}, true},
		{"Hello World", []string{"World$"}, true},
		{"Hello World", []string{"^World"}, false},
		{"Hello World", []string{"Hello$"}, false},
		{"Hello World", []string{"\\d+"}, false},
		{"Hello World", []string{".*"}, true},
		{"Hello World", []string{"Hello", "Hi"}, true},
		{"Hello World", []string{"Hi", "Bye"}, false},
	}

	for _, tc := range testCases {
		result := str.IsMatch(tc.input, tc.patterns...)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
		}
	}
}

// TestIsMatchEmpty tests empty string and pattern scenarios
func TestIsMatchEmpty(t *testing.T) {
	testCases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"Hello World", []string{""}, false},
		{"Hello World", []string{}, false},
		{"", []string{".*"}, true},
		{"", []string{"Hello"}, false},
		{"", []string{""}, false},
		{"", []string{}, false},
	}

	for _, tc := range testCases {
		result := str.IsMatch(tc.input, tc.patterns...)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
		}
	}
}

// TestIsMatchNumbers tests number pattern matching
func TestIsMatchNumbers(t *testing.T) {
	testCases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"12345", []string{"\\d+"}, true},
		{"12345", []string{"\\d{5}"}, true},
		{"12345", []string{"\\d{4}"}, true},
		{"12345", []string{"\\d{6}"}, false},
		{"Hello123", []string{"\\d+"}, true},
		{"Hello123", []string{"\\d{3}"}, true},
		{"Hello123", []string{"\\d{4}"}, false},
	}

	for _, tc := range testCases {
		result := str.IsMatch(tc.input, tc.patterns...)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
		}
	}
}

// TestIsMatchWordBoundary tests word boundary patterns
func TestIsMatchWordBoundary(t *testing.T) {
	testCases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"Hello123", []string{"\\w+"}, true},
		{"Hello123", []string{"\\D+"}, true},
		{"Hello123", []string{"\\D{5}"}, true},
		{"Hello123", []string{"\\D{6}"}, false},
		{"Hello-World", []string{"\\w+"}, true},
		{"Hello-World", []string{"\\w-\\w"}, true},
		{"Hello-World", []string{"\\w-\\w+"}, true},
		{"Hello-World", []string{"\\w+\\-\\w+"}, true},
		{"Hello-World", []string{"\\w+\\-\\w+$"}, true},
		{"Hello-World", []string{"^\\w+\\-\\w+$"}, true},
	}

	for _, tc := range testCases {
		result := str.IsMatch(tc.input, tc.patterns...)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
		}
	}
}

// TestIsMatchMultiplePatterns tests multiple pattern scenarios
func TestIsMatchMultiplePatterns(t *testing.T) {
	testCases := []struct {
		input    string
		patterns []string
		expected bool
	}{
		{"Hello World", []string{"^Hello.*World$", "^Hi.*Bye$"}, true},
		{"Hello World", []string{"^Hi.*Bye$"}, false},
		{"Hello-World", []string{"^\\w+\\-\\w+$", "^\\d+\\-\\d+$"}, true},
		{"Hello-World", []string{"^\\d+\\-\\d+$"}, false},
		{"Hello-World", []string{"^\\w+\\-\\w+$", "^\\w+\\-\\w+$"}, true},
		{"Hello-World", []string{"^\\w+\\-\\w+$", "^\\w+\\-\\w+$", "^\\w+\\-\\w+$"}, true},
		{"Hello-World", []string{"^\\w+\\-\\w+$", "^\\w+\\-\\w+$", "^\\w+\\-\\w+$", "^\\w+\\-\\w+$"}, true},
	}

	for _, tc := range testCases {
		result := str.IsMatch(tc.input, tc.patterns...)
		if result != tc.expected {
			t.Errorf("IsMatch(%q, %v) = %v, want %v", tc.input, tc.patterns, result, tc.expected)
		}
	}
}
