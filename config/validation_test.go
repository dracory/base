package config

import (
	"os"
	"testing"
)

func TestMissingEnvError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      MissingEnvError
		expected string
	}{
		{
			name:     "with context",
			err:      MissingEnvError{Key: "TEST_KEY", Context: "required for testing"},
			expected: `config: required env "TEST_KEY" is missing: required for testing`,
		},
		{
			name:     "without context",
			err:      MissingEnvError{Key: "TEST_KEY", Context: ""},
			expected: `config: required env "TEST_KEY" is missing`,
		},
		{
			name:     "with whitespace context",
			err:      MissingEnvError{Key: "TEST_KEY", Context: "   "},
			expected: `config: required env "TEST_KEY" is missing`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("MissingEnvError.Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestEnsureRequired(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		key       string
		context   string
		wantError bool
	}{
		{
			name:      "returns nil when value present",
			value:     "some-value",
			key:       "TEST_KEY",
			context:   "test context",
			wantError: false,
		},
		{
			name:      "returns error when empty",
			value:     "",
			key:       "TEST_KEY",
			context:   "test context",
			wantError: true,
		},
		{
			name:      "returns error when whitespace only",
			value:     "   ",
			key:       "TEST_KEY",
			context:   "test context",
			wantError: true,
		},
		{
			name:      "returns nil when value with whitespace",
			value:     "  value  ",
			key:       "TEST_KEY",
			context:   "test context",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EnsureRequired(tt.value, tt.key, tt.context)
			if (err != nil) != tt.wantError {
				t.Errorf("EnsureRequired() error = %v, wantError %v", err, tt.wantError)
			}
			if err != nil {
				if _, ok := err.(MissingEnvError); !ok {
					t.Errorf("EnsureRequired() error type = %T, want MissingEnvError", err)
				}
			}
		})
	}
}

func TestRequireWhen(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		key       string
		context   string
		value     string
		wantError bool
	}{
		{
			name:      "skips when condition false",
			condition: false,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "",
			wantError: false,
		},
		{
			name:      "validates when condition true and value present",
			condition: true,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "some-value",
			wantError: false,
		},
		{
			name:      "returns error when condition true and value empty",
			condition: true,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RequireWhen(tt.condition, tt.key, tt.context, tt.value)
			if (err != nil) != tt.wantError {
				t.Errorf("RequireWhen() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestRequireString(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		context   string
		envValue  string
		wantValue string
		wantError bool
	}{
		{
			name:      "returns value when present",
			key:       "TEST_REQUIRE_STRING_PRESENT",
			context:   "test context",
			envValue:  "test-value",
			wantValue: "test-value",
			wantError: false,
		},
		{
			name:      "trims whitespace",
			key:       "TEST_REQUIRE_STRING_TRIM",
			context:   "test context",
			envValue:  "  test-value  ",
			wantValue: "test-value",
			wantError: false,
		},
		{
			name:      "returns error when empty",
			key:       "TEST_REQUIRE_STRING_EMPTY",
			context:   "test context",
			envValue:  "",
			wantValue: "",
			wantError: true,
		},
		{
			name:      "returns error when whitespace only",
			key:       "TEST_REQUIRE_STRING_WHITESPACE",
			context:   "test context",
			envValue:  "   ",
			wantValue: "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable
			os.Setenv(tt.key, tt.envValue)
			defer os.Unsetenv(tt.key)

			value, err := RequireString(tt.key, tt.context)
			if (err != nil) != tt.wantError {
				t.Errorf("RequireString() error = %v, wantError %v", err, tt.wantError)
			}
			if value != tt.wantValue {
				t.Errorf("RequireString() value = %q, want %q", value, tt.wantValue)
			}
		})
	}
}
