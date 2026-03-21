package types

import (
	"testing"
	"time"
)

func TestFlashMessage(t *testing.T) {
	// Test creating a flash message
	message := FlashMessage{
		Type:    "success",
		Message: "Operation completed successfully",
		Url:     "/dashboard",
		Time:    time.Now().Format(time.RFC3339),
	}

	// Test field values
	if message.Type != "success" {
		t.Errorf("Expected type 'success', got '%s'", message.Type)
	}

	if message.Message != "Operation completed successfully" {
		t.Errorf("Expected message 'Operation completed successfully', got '%s'", message.Message)
	}

	if message.Url != "/dashboard" {
		t.Errorf("Expected url '/dashboard', got '%s'", message.Url)
	}

	if message.Time == "" {
		t.Error("Expected time to be set, got empty string")
	}
}

func TestFlashMessage_ZeroValue(t *testing.T) {
	// Test zero value
	var message FlashMessage

	if message.Type != "" {
		t.Errorf("Expected empty type, got '%s'", message.Type)
	}

	if message.Message != "" {
		t.Errorf("Expected empty message, got '%s'", message.Message)
	}

	if message.Url != "" {
		t.Errorf("Expected empty url, got '%s'", message.Url)
	}

	if message.Time != "" {
		t.Errorf("Expected empty time, got '%s'", message.Time)
	}
}

func TestFlashMessage_WithOptionalFields(t *testing.T) {
	// Test with minimal required fields
	message := FlashMessage{
		Type:    "error",
		Message: "Something went wrong",
	}

	if message.Type != "error" {
		t.Errorf("Expected type 'error', got '%s'", message.Type)
	}

	if message.Message != "Something went wrong" {
		t.Errorf("Expected message 'Something went wrong', got '%s'", message.Message)
	}

	// Optional fields should be empty
	if message.Url != "" {
		t.Errorf("Expected empty url, got '%s'", message.Url)
	}

	if message.Time != "" {
		t.Errorf("Expected empty time, got '%s'", message.Time)
	}
}

func TestFlashMessage_CommonTypes(t *testing.T) {
	// Test common flash message types
	testCases := []struct {
		flashType string
		expected  string
	}{
		{"success", "success"},
		{"error", "error"},
		{"info", "info"},
		{"warning", "warning"},
	}

	for _, tc := range testCases {
		message := FlashMessage{
			Type:    tc.flashType,
			Message: "Test message",
		}

		if message.Type != tc.expected {
			t.Errorf("Expected type '%s', got '%s'", tc.expected, message.Type)
		}
	}
}
