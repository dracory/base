package str_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/dracory/base/str"
)

func TestToBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []byte
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: []byte{},
		},
		{
			name:     "Simple String",
			input:    "hello",
			expected: []byte("hello"),
		},
		{
			name:     "String with Spaces",
			input:    "hello world",
			expected: []byte("hello world"),
		},
		{
			name:     "String with Special Characters",
			input:    "hello!@#$",
			expected: []byte("hello!@#$"),
		},
		{
			name:     "Unicode String",
			input:    "привет",
			expected: []byte("привет"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := str.ToBytes(tc.input)

			// Check if the result has the same length as the expected byte slice
			if len(result) != len(tc.expected) {
				t.Errorf("ToBytes(%q) = %v, want %v", tc.input, result, tc.expected)
			}

			// Check if the result has the same underlying data as the expected byte slice
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ToBytes(%q) = %v, want %v", tc.input, result, tc.expected)
			}

			// Check if the result and the input string share the same underlying data
			// by comparing their unsafe pointers.
			if len(tc.input) > 0 {
				stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&tc.input))
				sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&result))
				if stringHeader.Data != sliceHeader.Data {
					t.Errorf("ToBytes(%q) does not share the same underlying data with the input string", tc.input)
				}
			}
		})
	}
}
