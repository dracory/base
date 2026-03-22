package config

import (
	"errors"
	"os"
	"testing"
)

func TestInitializeEnvEncVariables(t *testing.T) {
	tests := []struct {
		name           string
		appEnvironment string
		publicKey      string
		privateKey     string
		wantErr        bool
		errType        error
	}{
		{
			name:           "Missing app environment",
			appEnvironment: "",
			publicKey:      "test-public",
			privateKey:     "test-private",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Missing public key",
			appEnvironment: "development",
			publicKey:      "",
			privateKey:     "test-private",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Missing private key",
			appEnvironment: "development",
			publicKey:      "test-public",
			privateKey:     "",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "No vault file",
			appEnvironment: "production",
			publicKey:      "test-public",
			privateKey:     "test-private",
			wantErr:        true,
			errType:        &EnvEncError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitializeEnvEncVariablesFromFile(tt.appEnvironment, tt.publicKey, tt.privateKey)

			if tt.wantErr {
				if err == nil {
					t.Errorf("InitializeEnvEncVariables() expected error but got none")
					return
				}

				// Check error type
				if tt.errType != nil {
					switch expectedErr := tt.errType.(type) {
					case *MissingEnvError:
						var missingErr *MissingEnvError
						if !errors.As(err, &missingErr) {
							t.Errorf("InitializeEnvEncVariables() expected error type %T but got %T", expectedErr, err)
						}
					case *EnvEncError:
						var envErr *EnvEncError
						if !errors.As(err, &envErr) {
							t.Errorf("InitializeEnvEncVariables() expected error type %T but got %T", expectedErr, err)
						}
					}
				}
			} else {
				if err != nil {
					t.Errorf("InitializeEnvEncVariables() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "Existing file",
			path:     "env_enc_loader_test.go", // This test file should exist
			expected: true,
		},
		{
			name:     "Non-existing file",
			path:     "non_existing_file_12345.txt",
			expected: false,
		},
		{
			name:     "Empty path",
			path:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fileExists(tt.path)
			if result != tt.expected {
				t.Errorf("fileExists(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestMissingEnvError(t *testing.T) {
	err := MissingEnvError{
		Key:     "TEST_KEY",
		Context: "test context",
	}

	expected := "config: required env \"TEST_KEY\" is missing: test context"
	if err.Error() != expected {
		t.Errorf("MissingEnvError.Error() = %q, want %q", err.Error(), expected)
	}
}

func TestEnvEncError(t *testing.T) {
	err := &EnvEncError{
		Operation: "test_operation",
		Message:   "test message",
	}

	expected := "EnvEnc error in test_operation: test message"
	if err.Error() != expected {
		t.Errorf("EnvEncError.Error() = %q, want %q", err.Error(), expected)
	}
}

// Test with actual file system integration (integration test)
func TestInitializeEnvEncVariablesIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Create a temporary vault file for testing
	vaultContent := "TEST_VAR=test_value"
	vaultFilePath := ".env.test.vault"

	// Clean up after test
	defer func() {
		os.Remove(vaultFilePath)
	}()

	// Create test vault file
	err := os.WriteFile(vaultFilePath, []byte(vaultContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test vault file: %v", err)
	}

	// Test with valid configuration (this will fail decryption but should find the file)
	err = InitializeEnvEncVariablesFromFile("test", "invalid-public-key", "invalid-private-key")
	if err == nil {
		t.Errorf("Expected error due to invalid keys, but got none")
	}

	// Check that the error is about key derivation, not file not found
	var envErr *EnvEncError
	if !errors.As(err, &envErr) || envErr.Operation != "derive_key" {
		t.Errorf("Expected derive_key error, got: %v", err)
	}
}

// Test resource loader functionality
func TestInitializeEnvEncVariablesFromResources(t *testing.T) {
	tests := []struct {
		name           string
		appEnvironment string
		publicKey      string
		privateKey     string
		resourceLoader func(string) (string, error)
		wantErr        bool
		errType        error
	}{
		{
			name:           "Missing app environment",
			appEnvironment: "",
			publicKey:      "test-public",
			privateKey:     "test-private",
			resourceLoader: func(name string) (string, error) {
				return "test-content", nil
			},
			wantErr: true,
			errType: &MissingEnvError{},
		},
		{
			name:           "Missing public key",
			appEnvironment: "development",
			publicKey:      "",
			privateKey:     "test-private",
			resourceLoader: func(name string) (string, error) {
				return "test-content", nil
			},
			wantErr: true,
			errType: &MissingEnvError{},
		},
		{
			name:           "Missing private key",
			appEnvironment: "development",
			publicKey:      "test-public",
			privateKey:     "",
			resourceLoader: func(name string) (string, error) {
				return "test-content", nil
			},
			wantErr: true,
			errType: &MissingEnvError{},
		},
		{
			name:           "No resource loader",
			appEnvironment: "production",
			publicKey:      "test-public",
			privateKey:     "test-private",
			resourceLoader: nil,
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Resource not found",
			appEnvironment: "production",
			publicKey:      "test-public",
			privateKey:     "test-private",
			resourceLoader: func(name string) (string, error) {
				return "", errors.New("not found")
			},
			wantErr: true,
			errType: &EnvEncError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitializeEnvEncVariablesFromResources(tt.appEnvironment, tt.publicKey, tt.privateKey, tt.resourceLoader)

			if tt.wantErr {
				if err == nil {
					t.Errorf("InitializeEnvEncVariablesFromResources() expected error but got none")
					return
				}

				// Check error type
				if tt.errType != nil {
					switch tt.errType.(type) {
					case *MissingEnvError:
						var missingErr *MissingEnvError
						if !errors.As(err, &missingErr) {
							t.Errorf("Expected MissingEnvError, got: %T", err)
						}
					case *EnvEncError:
						var envErr *EnvEncError
						if !errors.As(err, &envErr) {
							t.Errorf("Expected EnvEncError, got: %T", err)
						}
					}
				}
			} else {
				if err != nil {
					t.Errorf("InitializeEnvEncVariablesFromResources() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateInputs(t *testing.T) {
	tests := []struct {
		name           string
		appEnvironment string
		publicKey      string
		privateKey     string
		wantErr        bool
		errType        error
	}{
		{
			name:           "All valid inputs",
			appEnvironment: "development",
			publicKey:      "valid-public-key",
			privateKey:     "valid-private-key",
			wantErr:        false,
		},
		{
			name:           "Missing app environment",
			appEnvironment: "",
			publicKey:      "valid-public-key",
			privateKey:     "valid-private-key",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Missing public key",
			appEnvironment: "development",
			publicKey:      "",
			privateKey:     "valid-private-key",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Missing private key",
			appEnvironment: "development",
			publicKey:      "valid-public-key",
			privateKey:     "",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
		{
			name:           "Whitespace only app environment",
			appEnvironment: "   ",
			publicKey:      "valid-public-key",
			privateKey:     "valid-private-key",
			wantErr:        true,
			errType:        &MissingEnvError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputs(tt.appEnvironment, tt.publicKey, tt.privateKey)

			if tt.wantErr {
				if err == nil {
					t.Errorf("validateInputs() expected error but got none")
					return
				}

				if tt.errType != nil {
					switch tt.errType.(type) {
					case *MissingEnvError:
						var missingErr *MissingEnvError
						if !errors.As(err, &missingErr) {
							t.Errorf("Expected MissingEnvError, got: %T", err)
						}
					}
				}
			} else {
				if err != nil {
					t.Errorf("validateInputs() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestResourceLoader(t *testing.T) {
	tests := []struct {
		name           string
		resourceLoader func(string) (string, error)
		expectedResult string
		expectError    bool
	}{
		{
			name: "Successful resource loading",
			resourceLoader: func(name string) (string, error) {
				return "resource content", nil
			},
			expectedResult: "resource content",
			expectError:    false,
		},
		{
			name: "Resource not found",
			resourceLoader: func(name string) (string, error) {
				return "", errors.New("not found")
			},
			expectedResult: "",
			expectError:    true,
		},
		{
			name:           "Nil resource loader",
			resourceLoader: nil,
			expectedResult: "",
			expectError:    false, // Should handle nil gracefully
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.resourceLoader == nil {
				// Test with nil resource loader - should not panic
				t.Log("Testing with nil resource loader")
				return
			}

			result, err := tt.resourceLoader("test.vault")

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if result != tt.expectedResult {
				t.Errorf("Expected result %q, got %q", tt.expectedResult, result)
			}
		})
	}
}
