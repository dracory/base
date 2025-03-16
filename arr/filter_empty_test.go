package arr

import (
	"reflect"
	"testing"
)

func TestFilterEmpty(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{"nil slice", nil, nil},
		{"empty slice", []string{}, []string{}},
		{"only empty strings", []string{"", "", ""}, []string{}},
		{"mix of empty and non-empty strings", []string{"", "hello", "", "world"}, []string{"hello", "world"}},
		{"only non-empty strings", []string{"hello", "world", "foo"}, []string{"hello", "world", "foo"}},
		{"single non-empty string", []string{"hello"}, []string{"hello"}},
		{"single empty string", []string{""}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FilterEmpty(tt.slice)
			if tt.slice == nil {
				if actual != nil {
					t.Errorf("FilterEmpty(nil) = %v, want nil", actual)
				}
				return
			}
			if len(tt.expected) == 0 && len(actual) == 0 {
				// special case for empty slices
				return
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("FilterEmpty(%v) = %v, want %v", tt.slice, actual, tt.expected)
			}
		})
	}
}
