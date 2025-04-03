package str_test

import (
	"testing"
	"github.com/dracory/base/str"
)

// TestMD5_EmptyString tests MD5 with an empty string
func TestMD5_EmptyString(t *testing.T) {
	input := ""
	expected := "d41d8cd98f00b204e9800998ecf8427e"
	
	result := str.MD5(input)
	if result != expected {
		t.Errorf("MD5(%q) = %q, want %q", input, result, expected)
	}
}

// TestMD5_SimpleString tests MD5 with a simple string
func TestMD5_SimpleString(t *testing.T) {
	input := "hello"
	expected := "5d41402abc4b2a76b9719d911017c592"
	
	result := str.MD5(input)
	if result != expected {
		t.Errorf("MD5(%q) = %q, want %q", input, result, expected)
	}
}

// TestMD5_SpecialChars tests MD5 with special characters
func TestMD5_SpecialChars(t *testing.T) {
	input := "hello!@#$%"
	expected := "2d5d9d005fb9c81fab64c601fd8dbff3"
	
	result := str.MD5(input)
	if result != expected {
		t.Errorf("MD5(%q) = %q, want %q", input, result, expected)
	}
}

// TestMD5_LongString tests MD5 with a long string
func TestMD5_LongString(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog"
	expected := "9e107d9d372bb6826bd81d3542a419d6"
	
	result := str.MD5(input)
	if result != expected {
		t.Errorf("MD5(%q) = %q, want %q", input, result, expected)
	}
}

// TestMD5_Unicode tests MD5 with Unicode characters
func TestMD5_Unicode(t *testing.T) {
	input := "你好，世界！"
	expected := "5082079d92a8ef985f59e001d445ff20"
	
	result := str.MD5(input)
	if result != expected {
		t.Errorf("MD5(%q) = %q, want %q", input, result, expected)
	}
}

// TestMD5_SameInput tests MD5 with the same input multiple times
func TestMD5_SameInput(t *testing.T) {
	input := "test"
	expected := "098f6bcd4621d373cade4e832627b4f6"
	
	// Test multiple times to ensure consistency
	for i := 0; i < 10; i++ {
		result := str.MD5(input)
		if result != expected {
			t.Errorf("MD5(%q) = %q, want %q on iteration %d", input, result, expected, i)
		}
	}
}
