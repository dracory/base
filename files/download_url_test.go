package files_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dracory/base/files" // Adjust this import path if necessary
)

// mockReaderCloser simulates an io.ReadCloser that can return errors on Read.
type mockReaderCloser struct {
	reader  io.Reader
	readErr error
	closed  bool
}

func (m *mockReaderCloser) Read(p []byte) (n int, err error) {
	if m.closed {
		return 0, errors.New("read on closed reader")
	}
	if m.readErr != nil {
		return 0, m.readErr
	}
	if m.reader == nil {
		return 0, io.EOF
	}
	return m.reader.Read(p)
}

func (m *mockReaderCloser) Close() error {
	m.closed = true
	return nil
}

func TestDownloadURL_Success(t *testing.T) {
	expectedContent := "This is the file content from the server."

	// 1. Setup mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, expectedContent)
	}))
	// Ensure server is closed after the test
	t.Cleanup(server.Close)

	// 2. Create temporary file path
	tempDir := t.TempDir() // Creates a temporary directory cleaned up automatically
	localFilePath := filepath.Join(tempDir, "downloaded_file.txt")

	// 3. Call the function under test
	err := files.DownloadURL(server.URL, localFilePath)

	// 4. Assertions
	if err != nil {
		t.Fatalf("DownloadURL failed unexpectedly: %v", err)
	}

	// Verify file content
	actualContentBytes, readErr := os.ReadFile(localFilePath)
	if readErr != nil {
		t.Fatalf("Failed to read downloaded file %s: %v", localFilePath, readErr)
	}
	actualContent := string(actualContentBytes)

	if actualContent != expectedContent {
		t.Errorf("File content mismatch.\nExpected: %q\nActual:   %q", expectedContent, actualContent)
	}
}

func TestDownloadURL_HttpError(t *testing.T) {
	// 1. Setup mock HTTP server to return 404
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
	}))
	t.Cleanup(server.Close)

	// 2. Create temporary file path
	tempDir := t.TempDir()
	localFilePath := filepath.Join(tempDir, "should_not_matter.txt")

	// 3. Call the function under test
	err := files.DownloadURL(server.URL, localFilePath)

	// 4. Assertions
	// The current implementation of DownloadURL doesn't check the status code.
	// It will proceed and try to save the "Not Found" body.
	// Therefore, we expect success (nil error) according to the current code logic.
	// A potential improvement to DownloadURL would be to check resp.StatusCode.
	if err != nil {
		t.Fatalf("DownloadURL failed unexpectedly even with 404 (current behavior): %v", err)
	}

	// Verify the "error page" content was downloaded
	actualContentBytes, readErr := os.ReadFile(localFilePath)
	if readErr != nil {
		t.Fatalf("Failed to read downloaded file %s even on 404: %v", localFilePath, readErr)
	}
	if string(actualContentBytes) != "Not Found" {
		t.Errorf("Expected 'Not Found' body to be saved on 404, but got: %q", string(actualContentBytes))
	}
}

func TestDownloadURL_NetworkError(t *testing.T) {
	// 1. Use an invalid URL that won't resolve
	invalidURL := "http://invalid-url-that-should-not-exist-12345.xyz"

	// 2. Create temporary file path (shouldn't be created)
	tempDir := t.TempDir()
	localFilePath := filepath.Join(tempDir, "should_not_be_created.txt")

	// 3. Call the function under test
	err := files.DownloadURL(invalidURL, localFilePath)

	// 4. Assertions
	if err == nil {
		// Clean up if file was unexpectedly created
		os.Remove(localFilePath)
		t.Fatal("DownloadURL should have failed for an invalid URL, but it didn't")
	}

	// Check if the file was created (it shouldn't have been)
	if _, statErr := os.Stat(localFilePath); statErr == nil || !os.IsNotExist(statErr) {
		t.Errorf("File %s was created despite network error", localFilePath)
		os.Remove(localFilePath) // Clean up
	}
}

func TestDownloadURL_FileCreateError(t *testing.T) {
	// 1. Setup mock HTTP server (content doesn't matter much here)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "dummy content")
	}))
	t.Cleanup(server.Close)

	// 2. Provide an invalid local file path (e.g., path includes non-existent dir)
	// Note: Creating a truly invalid path reliably across OSes can be tricky.
	// An empty path often works, or a path into a non-existent directory.
	invalidLocalFilePath := filepath.Join("non_existent_dir_for_test", "cannot_create.txt")
	// Ensure the directory doesn't exist (best effort)
	os.RemoveAll(filepath.Dir(invalidLocalFilePath))

	// 3. Call the function under test
	err := files.DownloadURL(server.URL, invalidLocalFilePath)

	// 4. Assertions
	if err == nil {
		t.Fatal("DownloadURL should have failed due to file creation error, but it didn't")
	}

	// Check if the error seems related to file path issues (optional, error type varies)
	// Example check (might need adjustment based on OS):
	if !strings.Contains(err.Error(), "no such file or directory") && !strings.Contains(err.Error(), "cannot find the path") {
		t.Logf("Warning: Received error '%v', which might not be the expected file creation error type.", err)
	}
}

func TestDownloadURL_CopyError(t *testing.T) {
	// 1. Setup mock HTTP server with a body that errors during Read
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We need to hijack the response writer to control the body precisely
		hijacker, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
			t.Log("Hijacking not supported by test server ResponseWriter") // Log for debugging test setup
			return
		}
		conn, _, err := hijacker.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			t.Logf("Hijack failed: %v", err) // Log for debugging test setup
			return
		}
		// Write headers manually
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nContent-Length: 100\r\n\r\n") // Example headers
		// Write some initial data
		fmt.Fprintf(conn, "initial data")
		// Close the connection prematurely to cause an error during io.Copy in the client
		conn.Close()
	}))
	t.Cleanup(server.Close)

	// 2. Create temporary file path
	tempDir := t.TempDir()
	localFilePath := filepath.Join(tempDir, "partially_downloaded.txt")

	// 3. Call the function under test
	err := files.DownloadURL(server.URL, localFilePath)

	// 4. Assertions
	if err == nil {
		t.Fatal("DownloadURL should have failed due to io.Copy error, but it didn't")
	}

	// Check if the error is an IO error (it might be wrapped)
	// Note: The exact error from a premature close can vary (e.g., io.ErrUnexpectedEOF, net specific errors)
	// We check if it's *not* nil, as simulating the exact copy error is complex.
	t.Logf("Received expected error during copy: %v", err) // Log the error for info

	// Optional: Check if the file contains only partial data (or is empty)
	partialContentBytes, readErr := os.ReadFile(localFilePath)
	if readErr != nil && !os.IsNotExist(readErr) { // File might not exist if Create failed first, or might be partially written
		t.Logf("Could not read potentially partial file %s: %v", localFilePath, readErr)
	} else if readErr == nil {
		t.Logf("Partial file content: %q", string(partialContentBytes))
		// Add assertion if specific partial content is expected
		if !strings.HasPrefix(string(partialContentBytes), "initial data") {
			t.Errorf("Partial content mismatch or unexpected content: %q", string(partialContentBytes))
		}
	}
}
