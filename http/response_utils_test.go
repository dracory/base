package http

import (
	"bytes"
	"errors"
	"testing"
)

// mockCloser implements io.Closer for testing
type mockCloser struct {
	shouldError bool
	closed      bool
}

func (m *mockCloser) Close() error {
	m.closed = true
	if m.shouldError {
		return errors.New("close error")
	}
	return nil
}

func TestSafeCloseResponseBody(t *testing.T) {
	// Test with nil body
	SafeCloseResponseBody(nil)
	// Should not panic

	// Test with successful close
	closer := &mockCloser{shouldError: false}
	SafeCloseResponseBody(closer)
	if !closer.closed {
		t.Error("Expected body to be closed")
	}

	// Test with close error (should log error but not panic)
	closerWithError := &mockCloser{shouldError: true}
	// Capture logs to verify error is logged
	// Note: In a real test environment, you might want to capture slog output
	SafeCloseResponseBody(closerWithError)
	if !closerWithError.closed {
		t.Error("Expected body to be closed even with error")
	}
}

// Test with bytes.Buffer which implements io.Closer via nopCloser
func TestSafeCloseResponseBodyWithBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("test data"))
	// bytes.Buffer doesn't implement Close, so we need to use io.NopCloser
	nopCloser := &nopCloser{buf}
	SafeCloseResponseBody(nopCloser)
}

// nopCloser wraps a ReadCloser with a no-op Close method for testing
type nopCloser struct {
	*bytes.Buffer
}

func (n *nopCloser) Close() error {
	return nil
}
