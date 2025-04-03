package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestWords_Basic tests basic word truncation functionality
func TestWords_Basic(t *testing.T) {
	testCases := []struct {
		input    string
		limit    int
		expected string
	}{
		{"Hello World", 1, "Hello..."},
		{"Hello World", 2, "Hello World"},
		{"Hello World", 3, "Hello World"},
		{"A quick brown fox", 2, "A quick..."},
		{"A quick brown fox", 4, "A quick brown fox"},
	}

	for _, tc := range testCases {
		result := str.Words(tc.input, tc.limit)
		if result != tc.expected {
			t.Errorf("Words(%q, %d) = %q, want %q", tc.input, tc.limit, result, tc.expected)
		}
	}
}

// TestWords_CustomEnd tests word truncation with custom end strings
func TestWords_CustomEnd(t *testing.T) {
	testCases := []struct {
		input    string
		limit    int
		end      string
		expected string
	}{
		{"Hello World", 1, "...", "Hello..."},
		{"Hello World", 1, "[...]", "Hello[...]"},
		{"Hello World", 1, "---", "Hello---"},
		{"Hello World", 1, "", "Hello"},
	}

	for _, tc := range testCases {
		result := str.Words(tc.input, tc.limit, tc.end)
		if result != tc.expected {
			t.Errorf("Words(%q, %d, %q) = %q, want %q", tc.input, tc.limit, tc.end, result, tc.expected)
		}
	}
}

// TestWords_SpecialCases tests edge cases and special scenarios
func TestWords_SpecialCases(t *testing.T) {
	testCases := []struct {
		input    string
		limit    int
		expected string
	}{
		{"", 1, ""},
		{"Hello", 0, ""},
		{"Hello World", 0, ""},
		{"Hello World", -1, ""},
		{"  Hello  World  ", 1, "Hello..."},
		{"Hello   World", 1, "Hello..."},
		{"HelloWorld", 1, "HelloWorld"},
		{"HelloWorld", 0, ""},
	}

	for _, tc := range testCases {
		result := str.Words(tc.input, tc.limit)
		if result != tc.expected {
			t.Errorf("Words(%q, %d) = %q, want %q", tc.input, tc.limit, result, tc.expected)
		}
	}
}

// TestWords_Cyrillic tests word truncation with Cyrillic characters
func TestWords_Cyrillic(t *testing.T) {
	testCases := []struct {
		input    string
		limit    int
		expected string
	}{
		{"Привет Мир", 1, "Привет..."},
		{"Привет Мир", 2, "Привет Мир"},
		{"Привет Мир", 3, "Привет Мир"},
	}

	for _, tc := range testCases {
		result := str.Words(tc.input, tc.limit)
		if result != tc.expected {
			t.Errorf("Words(%q, %d) = %q, want %q", tc.input, tc.limit, result, tc.expected)
		}
	}
}

// TestWords_Numbers tests word truncation with numbers
func TestWords_Numbers(t *testing.T) {
	testCases := []struct {
		input    string
		limit    int
		expected string
	}{
		{"123 456 789", 1, "123..."},
		{"123 456 789", 2, "123 456..."},
		{"123 456 789", 3, "123 456 789"},
	}

	for _, tc := range testCases {
		result := str.Words(tc.input, tc.limit)
		if result != tc.expected {
			t.Errorf("Words(%q, %d) = %q, want %q", tc.input, tc.limit, result, tc.expected)
		}
	}
}
