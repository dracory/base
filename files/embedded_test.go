package files

import (
	"embed"
	"testing"
)

//go:embed testdata/*.txt
var testFS embed.FS

func TestEmbeddedFileToBytes(t *testing.T) {
	// Create a test file in testdata
	testContent := []byte("test content for bytes\n")

	// Test reading existing file
	bytes, err := EmbeddedFileToBytes(testFS, "testdata/test.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(bytes) == 0 {
		t.Error("Expected non-empty bytes")
	}

	if string(bytes) != string(testContent) {
		t.Errorf("Expected %s, got %s", string(testContent), string(bytes))
	}
}

func TestEmbeddedFileToString(t *testing.T) {
	// Test reading existing file
	str, err := EmbeddedFileToString(testFS, "testdata/test.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if str == "" {
		t.Error("Expected non-empty string")
	}
}

func TestEmbeddedFileToBytes_NotFound(t *testing.T) {
	// Test reading non-existent file
	_, err := EmbeddedFileToBytes(testFS, "testdata/nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestEmbeddedFileToString_NotFound(t *testing.T) {
	// Test reading non-existent file
	_, err := EmbeddedFileToString(testFS, "testdata/nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}
