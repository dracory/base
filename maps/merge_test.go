package maps_test

import (
	"reflect"
	"testing"

	"github.com/dracory/base/maps"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		maps     []map[string]any
		expected map[string]any
	}{
		{
			name:     "empty maps",
			maps:     []map[string]any{},
			expected: map[string]any{},
		},
		{
			name:     "single map",
			maps:     []map[string]any{{"a": 1}},
			expected: map[string]any{"a": 1},
		},
		{
			name:     "multiple maps no overlap",
			maps:     []map[string]any{{"a": 1}, {"b": 2}, {"c": 3}},
			expected: map[string]any{"a": 1, "b": 2, "c": 3},
		},
		{
			name:     "multiple maps with overlap",
			maps:     []map[string]any{{"a": 1}, {"a": 2}, {"a": 3}},
			expected: map[string]any{"a": 3},
		},
		{
			name:     "multiple maps with mixed overlap",
			maps:     []map[string]any{{"a": 1, "b": 1}, {"b": 2, "c": 2}, {"c": 3, "d": 3}},
			expected: map[string]any{"a": 1, "b": 2, "c": 3, "d": 3},
		},
		{
			name:     "empty and non-empty maps",
			maps:     []map[string]any{{}, {"a": 1}, {}},
			expected: map[string]any{"a": 1},
		},
		{
			name:     "multiple maps with different keys",
			maps:     []map[string]any{{"a": 1, "b": 2}, {"c": 3, "d": 4}, {"e": 5, "f": 6}},
			expected: map[string]any{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6},
		},
		{
			name:     "mixed value types",
			maps:     []map[string]any{{"a": 1}, {"b": "hello"}, {"c": true}},
			expected: map[string]any{"a": 1, "b": "hello", "c": true},
		},
		{
			name:     "nil map",
			maps:     []map[string]any{{"a": 1}, nil, {"b": 2}},
			expected: map[string]any{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maps.Merge(tt.maps...)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("maps.Merge(%v) = %v, want %v", tt.maps, actual, tt.expected)
			}
		})
	}
}
