package str_test

import (
	"testing"

	"github.com/dracory/base/str"
)

// TestUcSplitBasic tests basic cases of UcSplit
func TestUcSplitBasic(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"Hello", []string{"Hello"}},
		{"hello", []string{"hello"}},
		{"HelloWorld", []string{"Hello", "World"}},
		{"Hello World", []string{"Hello ", "World"}},
		{"helloWorld", []string{"hello", "World"}},
		{"HelloWORLD", []string{"Hello", "W", "O", "R", "L", "D"}},
	}

	for _, tc := range testCases {
		result := str.UcSplit(tc.input)
		if !equalStringSlices(result, tc.expected) {
			t.Errorf("UcSplit(%q) = got: %v, expected %v", tc.input, result, tc.expected)
		}
	}
}

// TestUcSplitMultipleUppercase tests cases with multiple uppercase characters
func TestUcSplitMultipleUppercase(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"ABC", []string{"A", "B", "C"}},
		{"AaBbCc", []string{"Aa", "Bb", "Cc"}},
		{"HelloABCWorld", []string{"Hello", "A", "B", "C", "World"}},
		{"AAABBBCCC", []string{"A", "A", "A", "B", "B", "B", "C", "C", "C"}},
		{"HelloAAABBBCCCWorld", []string{"Hello", "A", "A", "A", "B", "B", "B", "C", "C", "C", "World"}},
	}

	for _, tc := range testCases {
		result := str.UcSplit(tc.input)
		if !equalStringSlices(result, tc.expected) {
			t.Errorf("UcSplit(%q) = got: %v, expected: %v", tc.input, result, tc.expected)
		}
	}
}

// TestUcSplitSpecialCases tests special cases including numbers and spaces
func TestUcSplitSpecialCases(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"A1B2C3", []string{"A1", "B2", "C3"}},
		{"1A2B3C", []string{"1", "A2", "B3", "C"}},
		{"A1B2C3D", []string{"A1", "B2", "C3", "D"}},
		{"Hello123", []string{"Hello123"}},
		{"123Hello", []string{"123", "Hello"}},
	}

	for _, tc := range testCases {
		result := str.UcSplit(tc.input)
		if !equalStringSlices(result, tc.expected) {
			t.Errorf("UcSplit(%q) = got: %v, expected: %v", tc.input, result, tc.expected)
		}
	}
}

// TestUcSplitCyrillic tests cases with Cyrillic characters
func TestUcSplitCyrillic(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"АаБбВв", []string{"Аа", "Бб", "Вв"}},
		{"ПриветМир", []string{"Привет", "Мир"}},
		{"ПриветМИР", []string{"Привет", "М", "И", "Р"}},
		{"Привет мИр", []string{"Привет м", "Ир"}},
	}

	for _, tc := range testCases {
		result := str.UcSplit(tc.input)
		if !equalStringSlices(result, tc.expected) {
			t.Errorf("UcSplit(%q) = got: %v, expected: %v", tc.input, result, tc.expected)
		}
	}
}

// TestUcSplitMixedCases tests mixed cases with numbers and uppercase characters
func TestUcSplitMixedCases(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"HelloWorld123", []string{"Hello", "World123"}},
		{"Hello123World", []string{"Hello123", "World"}},
		{"HelloWorldABC", []string{"Hello", "World", "A", "B", "C"}},
		{"Hello123ABCWorld", []string{"Hello123", "A", "B", "C", "World"}},
		{"HelloABC123World", []string{"Hello", "A", "B", "C123", "World"}},
	}

	for _, tc := range testCases {
		result := str.UcSplit(tc.input)
		if !equalStringSlices(result, tc.expected) {
			t.Errorf("UcSplit(%q) = got: %v, expected: %v", tc.input, result, tc.expected)
		}
	}
}

// Helper function to compare string slices
func equalStringSlices(a, b []string) bool {
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
