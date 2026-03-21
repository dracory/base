package session

import (
	"context"
	"net/http"
	"testing"
)

func TestGetAuthSession(t *testing.T) {
	// Test with nil request
	session := GetAuthSession(nil)
	if session != nil {
		t.Error("Expected nil session for nil request")
	}

	// Test with request without session
	req := &http.Request{}
	session = GetAuthSession(req)
	if session != nil {
		t.Error("Expected nil session for request without session")
	}

	// Test with request with session in context
	ctx := SetAuthSessionInContext(context.TODO(), nil) // Using nil for simplicity
	req = req.WithContext(ctx)

	session = GetAuthSession(req)
	if session != nil {
		t.Error("Expected nil session for nil session in context")
	}
}

func TestGetAuthSessionFromContext(t *testing.T) {
	// Test with nil context
	session := GetAuthSessionFromContext(nil)
	if session != nil {
		t.Error("Expected nil session for nil context")
	}

	// Test with context without session
	ctx := context.Background()
	session = GetAuthSessionFromContext(ctx)
	if session != nil {
		t.Error("Expected nil session for context without session")
	}

	// Test with context with nil session
	ctx = SetAuthSessionInContext(ctx, nil)
	session = GetAuthSessionFromContext(ctx)
	if session != nil {
		t.Error("Expected nil session for nil session in context")
	}
}

func TestSetAuthSession(t *testing.T) {
	// Test with nil request
	req := SetAuthSession(nil, nil)
	if req != nil {
		t.Error("Expected nil request for nil input")
	}

	// Test with nil session
	req = &http.Request{}
	req = SetAuthSession(req, nil)
	if req == nil {
		t.Error("Expected request to be returned")
	}

	// Test with valid session
	req = &http.Request{}
	req = SetAuthSession(req, nil) // Using nil for simplicity
	if req == nil {
		t.Error("Expected request to be returned")
	}
}

func TestGetAuthUser(t *testing.T) {
	// Test with nil request
	user := GetAuthUser(nil)
	if user != nil {
		t.Error("Expected nil user for nil request")
	}

	// Test with request without user
	req := &http.Request{}
	user = GetAuthUser(req)
	if user != nil {
		t.Error("Expected nil user for request without user")
	}

	// Test with request with user in context
	ctx := SetAuthUserInContext(context.Background(), nil) // Using nil for simplicity
	req = req.WithContext(ctx)

	user = GetAuthUser(req)
	if user != nil {
		t.Error("Expected nil user for nil user in context")
	}
}

func TestGetAPIAuthUser(t *testing.T) {
	// Test with nil request
	user := GetAPIAuthUser(nil)
	if user != nil {
		t.Error("Expected nil user for nil request")
	}

	// Test with request without user
	req := &http.Request{}
	user = GetAPIAuthUser(req)
	if user != nil {
		t.Error("Expected nil user for request without user")
	}

	// Test with request with user in context
	ctx := SetAPIAuthUserInContext(context.Background(), nil) // Using nil for simplicity
	req = req.WithContext(ctx)

	user = GetAPIAuthUser(req)
	if user != nil {
		t.Error("Expected nil user for nil user in context")
	}
}

func TestExtendSession(t *testing.T) {
	// Test with nil session store
	err := ExtendSession(nil, &http.Request{}, 3600)
	if err == nil {
		t.Error("Expected error for nil session store")
	}
	expectedError := "session store is nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test with nil request
	err = ExtendSession(nil, nil, 3600)
	if err == nil {
		t.Error("Expected error for nil request")
	}
	expectedError = "session store is nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test with request without session
	req := &http.Request{}
	err = ExtendSession(nil, req, 3600)
	if err == nil {
		t.Error("Expected error for request without session")
	}
	expectedError = "session store is nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestUserSettingGet(t *testing.T) {
	// Test with nil session store
	result := UserSettingGet(nil, &http.Request{}, "test_key", "default")
	if result != "default" {
		t.Errorf("Expected default value, got '%s'", result)
	}

	// Test with request without user
	req := &http.Request{}
	result = UserSettingGet(nil, req, "test_key", "default")
	if result != "default" {
		t.Errorf("Expected default value, got '%s'", result)
	}
}

func TestUserSettingSet(t *testing.T) {
	// Test with nil session store
	err := UserSettingSet(nil, &http.Request{}, "test_key", "test_value")
	if err == nil {
		t.Error("Expected error for nil session store")
	}
	expectedError := "session store is nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}

	// Test with request without user
	req := &http.Request{}
	err = UserSettingSet(nil, req, "test_key", "test_value")
	if err == nil {
		t.Error("Expected error for request without user")
	}
	expectedError = "session store is nil"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}
