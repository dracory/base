package config

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestLoadAccumulator_Add(t *testing.T) {
	acc := &LoadAccumulator{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	acc.Add(err1)
	acc.Add(nil) // Should be ignored
	acc.Add(err2)

	if len(acc.errs) != 2 {
		t.Errorf("LoadAccumulator.Add() collected %d errors, want 2", len(acc.errs))
	}
}

func TestLoadAccumulator_MustString(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		context   string
		envValue  string
		wantValue string
		wantError bool
	}{
		{
			name:      "success case",
			key:       "TEST_MUST_STRING_SUCCESS",
			context:   "test context",
			envValue:  "test-value",
			wantValue: "test-value",
			wantError: false,
		},
		{
			name:      "error case",
			key:       "TEST_MUST_STRING_ERROR",
			context:   "test context",
			envValue:  "",
			wantValue: "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(tt.key, tt.envValue)
			defer os.Unsetenv(tt.key)

			acc := &LoadAccumulator{}
			value := acc.MustString(tt.key, tt.context)

			if value != tt.wantValue {
				t.Errorf("LoadAccumulator.MustString() = %q, want %q", value, tt.wantValue)
			}

			hasError := len(acc.errs) > 0
			if hasError != tt.wantError {
				t.Errorf("LoadAccumulator.MustString() hasError = %v, want %v", hasError, tt.wantError)
			}
		})
	}
}

func TestLoadAccumulator_MustWhen(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		key       string
		context   string
		value     string
		wantError bool
	}{
		{
			name:      "condition false - no error",
			condition: false,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "",
			wantError: false,
		},
		{
			name:      "condition true with value - no error",
			condition: true,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "some-value",
			wantError: false,
		},
		{
			name:      "condition true without value - error",
			condition: true,
			key:       "TEST_KEY",
			context:   "test context",
			value:     "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := &LoadAccumulator{}
			acc.MustWhen(tt.condition, tt.key, tt.context, tt.value)

			hasError := len(acc.errs) > 0
			if hasError != tt.wantError {
				t.Errorf("LoadAccumulator.MustWhen() hasError = %v, want %v", hasError, tt.wantError)
			}
		})
	}
}

func TestLoadAccumulator_Err(t *testing.T) {
	t.Run("returns nil when no errors", func(t *testing.T) {
		acc := &LoadAccumulator{}
		if err := acc.Err(); err != nil {
			t.Errorf("LoadAccumulator.Err() = %v, want nil", err)
		}
	})

	t.Run("returns ValidationError when errors present", func(t *testing.T) {
		acc := &LoadAccumulator{}
		acc.Add(errors.New("error 1"))
		acc.Add(errors.New("error 2"))

		err := acc.Err()
		if err == nil {
			t.Fatal("LoadAccumulator.Err() = nil, want error")
		}

		if _, ok := err.(ValidationError); !ok {
			t.Errorf("LoadAccumulator.Err() type = %T, want ValidationError", err)
		}
	})
}

func TestValidationError_Error(t *testing.T) {
	t.Run("empty errors", func(t *testing.T) {
		err := ValidationError{errs: []error{}}
		expected := "config: validation failed"
		if got := err.Error(); got != expected {
			t.Errorf("ValidationError.Error() = %q, want %q", got, expected)
		}
	})

	t.Run("single error", func(t *testing.T) {
		err := ValidationError{errs: []error{errors.New("test error")}}
		got := err.Error()
		if !strings.Contains(got, "config: validation failed:") {
			t.Errorf("ValidationError.Error() should contain header")
		}
		if !strings.Contains(got, "test error") {
			t.Errorf("ValidationError.Error() should contain error message")
		}
	})

	t.Run("multiple errors", func(t *testing.T) {
		err := ValidationError{errs: []error{
			errors.New("error 1"),
			errors.New("error 2"),
			errors.New("error 3"),
		}}
		got := err.Error()
		if !strings.Contains(got, "error 1") {
			t.Errorf("ValidationError.Error() should contain first error")
		}
		if !strings.Contains(got, "error 2") {
			t.Errorf("ValidationError.Error() should contain second error")
		}
		if !strings.Contains(got, "error 3") {
			t.Errorf("ValidationError.Error() should contain third error")
		}
	})
}

func TestValidationError_Errors(t *testing.T) {
	originalErrs := []error{
		errors.New("error 1"),
		errors.New("error 2"),
	}
	err := ValidationError{errs: originalErrs}

	gotErrs := err.Errors()

	if len(gotErrs) != len(originalErrs) {
		t.Errorf("ValidationError.Errors() length = %d, want %d", len(gotErrs), len(originalErrs))
	}

	// Verify it's a defensive copy by modifying the returned slice
	gotErrs[0] = errors.New("modified")
	if err.errs[0].Error() == "modified" {
		t.Error("ValidationError.Errors() should return defensive copy")
	}
}

func TestLoadAccumulator_Integration(t *testing.T) {
	// Set up environment
	os.Setenv("TEST_VALID_KEY", "valid-value")
	defer os.Unsetenv("TEST_VALID_KEY")

	acc := &LoadAccumulator{}

	// Mix of valid and invalid operations
	value1 := acc.MustString("TEST_VALID_KEY", "should succeed")
	value2 := acc.MustString("TEST_MISSING_KEY", "should fail")
	acc.MustWhen(true, "TEST_REQUIRED_KEY", "required when condition is true", "")
	acc.MustWhen(false, "TEST_OPTIONAL_KEY", "not required when condition is false", "")

	if value1 != "valid-value" {
		t.Errorf("First MustString() = %q, want %q", value1, "valid-value")
	}

	if value2 != "" {
		t.Errorf("Second MustString() = %q, want empty string", value2)
	}

	err := acc.Err()
	if err == nil {
		t.Fatal("LoadAccumulator.Err() = nil, want error")
	}

	validationErr, ok := err.(ValidationError)
	if !ok {
		t.Fatalf("LoadAccumulator.Err() type = %T, want ValidationError", err)
	}

	if len(validationErr.Errors()) != 2 {
		t.Errorf("ValidationError has %d errors, want 2", len(validationErr.Errors()))
	}
}
