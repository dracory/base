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
		{"1A2B3C", []string{"1A2B3C"}},
		{"A1B2C3D", []string{"A1", "B2", "C3", "D"}},
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
		{"HelloWorld123", []string{"HelloWorld123"}},
		{"Hello123World", []string{"Hello123World"}},
		{"HelloWorldABC", []string{"HelloWorldABC"}},
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
