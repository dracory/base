package files_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dracory/base/files"
)

// mockFile implements io.Reader, io.Seeker, io.Closer for testing purposes.
// It remains internal to the test package. It is not a multipart.File
type mockFile struct {
	reader  io.Reader
	seeker  io.Seeker
	name    string
	closed  bool
	readErr error // Optional error to simulate read failures
}

func (mf *mockFile) Read(p []byte) (n int, err error) {
	if mf.readErr != nil {
		return 0, mf.readErr // Simulate a read error
	}
	if mf.closed {
		return 0, errors.New("file already closed")
	}
	if mf.reader == nil {
		return 0, io.EOF // Simulate end of file if no reader
	}
	return mf.reader.Read(p)
}

func (mf *mockFile) Seek(offset int64, whence int) (int64, error) {
	if mf.closed {
		return 0, errors.New("file already closed")
	}
	if mf.seeker == nil {
		// multipart.File doesn't strictly require Seek,
		// but os.File (which CreateTemp returns) does.
		// Depending on how the tested function uses the file,
		// this might or might not be needed.
		// For io.Copy, Seek is not used on the source.
		return 0, errors.New("seek not supported by this mock")
	}
	return mf.seeker.Seek(offset, whence)
}

func (mf *mockFile) Close() error {
	if mf.closed {
		return errors.New("file already closed")
	}
	mf.closed = true
	// Simulate closing the underlying reader if it's a Closer
	if closer, ok := mf.reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// mockMultipartFile implements multipart.File for testing purposes.
// It wraps a mockFile to satisfy the multipart.File interface.
type mockMultipartFile struct {
	*mockFile
}

func (mmf *mockMultipartFile) Name() string {
	return mmf.name
}

func (mmf *mockMultipartFile) ReadAt(p []byte, off int64) (n int, err error) {
	// ReadAt is not used by io.Copy, so we can leave it unimplemented
	// or implement it if needed for other tests.
	return 0, fmt.Errorf("ReadAt not implemented")
}

// --- Test Cases ---

func TestSaveToTempDir_Success(t *testing.T) {
	fileName := "test_image.jpg"
	fileContent := "this is dummy image content"
	contentReader := strings.NewReader(fileContent)

	// Create a mock multipart.File using mockFile
	mock := &mockMultipartFile{
		mockFile: &mockFile{
			reader: contentReader,
			seeker: contentReader, // strings.Reader implements Seeker
		},
	}
	// Call the function under test using the package name
	tempFilePath, err := files.SaveToTempDir(fileName, mock)

	// Assertions
	if err != nil {
		t.Fatalf("SaveToTempDir failed unexpectedly: %v", err)
	}
	if tempFilePath == "" {
		t.Fatal("SaveToTempDir returned an empty path")
	}

	// Ensure the temp file is cleaned up after the test
	t.Cleanup(func() {
		err := os.Remove(tempFilePath)
		if err != nil && !os.IsNotExist(err) { // Don't fail if already removed
			t.Errorf("Failed to remove temp file %s: %v", tempFilePath, err)
		}
	})

	// Verify the file exists
	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		t.Fatalf("Expected temp file to exist, but it doesn't: %s", tempFilePath)
	}

	// Verify the file extension
	expectedExt := ".jpg"
	actualExt := filepath.Ext(tempFilePath)
	if actualExt != expectedExt {
		t.Errorf("Expected file extension '%s', but got '%s'", expectedExt, actualExt)
	}

	// Verify the file content
	savedContentBytes, readErr := os.ReadFile(tempFilePath)
	if readErr != nil {
		t.Fatalf("Failed to read saved temp file %s: %v", tempFilePath, readErr)
	}
	savedContent := string(savedContentBytes)
	if savedContent != fileContent {
		t.Errorf("Expected file content '%s', but got '%s'", fileContent, savedContent)
	}

	// Verify the mock file was closed (optional but good practice)
	// Note: The function under test doesn't close the input file,
	// it closes the *output* file it creates. So this check is not relevant here.
	// if !mock.closed {
	// 	t.Error("Expected the input file to be closed, but it wasn't")
	// }
}

func TestSaveToTempDir_NoExtension(t *testing.T) {
	fileName := "filewithoutextension"
	fileContent := "some data"
	contentReader := strings.NewReader(fileContent)
	mock := &mockMultipartFile{
		mockFile: &mockFile{
			reader: contentReader,
			seeker: contentReader},
	}

	tempFilePath, err := files.SaveToTempDir(fileName, mock)
	if err != nil {
		t.Fatalf("SaveToTempDir failed unexpectedly: %v", err)
	}
	t.Cleanup(func() {
		err := os.Remove(tempFilePath)
		if err != nil && !os.IsNotExist(err) {
			t.Errorf("Failed to remove temp file %s: %v", tempFilePath, err)
		}
	})

	// Verify the file extension is empty
	actualExt := filepath.Ext(tempFilePath)
	if actualExt != "" {
		t.Errorf("Expected empty file extension for temp file, but got '%s'", actualExt)
	}

	// Verify content just in case
	savedContentBytes, readErr := os.ReadFile(tempFilePath)
	if readErr != nil {
		t.Fatalf("Failed to read saved temp file %s: %v", tempFilePath, readErr)
	}
	if string(savedContentBytes) != fileContent {
		t.Errorf("File content mismatch. Expected '%s', got '%s'", fileContent, string(savedContentBytes))
	}
}

func TestSaveToTempDir_CopyError(t *testing.T) {
	fileName := "error_file.txt"
	simulatedError := errors.New("simulated copy error")

	// Create a mock multipart.File that will return an error on Read
	mock := &mockMultipartFile{
		mockFile: &mockFile{
			reader:  bytes.NewReader([]byte("some initial data")), // Need some data to trigger Read
			readErr: simulatedError,
		},
	}

	// Call the function under test
	tempFilePath, err := files.SaveToTempDir(fileName, mock)

	// Assertions
	if err == nil {
		// If a file was created despite the error, clean it up
		if tempFilePath != "" {
			os.Remove(tempFilePath) // Attempt cleanup
			t.Errorf("Temp file %s was created despite copy error, attempting cleanup.", tempFilePath)
		}
		t.Fatal("SaveToTempDir should have failed due to copy error, but it didn't")
	}

	// Use errors.Is for checking wrapped errors potentially
	if !errors.Is(err, simulatedError) {
		t.Errorf("Expected error wrapping '%v', but got '%v'", simulatedError, err)
	}

	// The function returns "" on copy error, so tempFilePath should be empty.
	// However, os.CreateTemp might have succeeded before io.Copy failed.
	// The function *should* clean up the temp file in case of io.Copy error,
	// but the current implementation doesn't explicitly do that.
	// We primarily check that the function signals failure correctly.
	if tempFilePath != "" {
		t.Errorf("Expected empty path on copy error, but got '%s'", tempFilePath)
		// Attempt cleanup just in case the function didn't
		os.Remove(tempFilePath)
	}
}

// Note: Testing os.CreateTemp failure directly is hard without manipulating
// the OS environment (e.g., permissions on TempDir), which makes tests brittle.
// We rely on the standard library's testing for os.CreateTemp itself.

// --- Potential Improvements & Considerations ---
// 1. Error Handling in SaveToTempDir: If io.Copy fails, the temporary file created
//    by os.CreateTemp might be left behind. Consider adding cleanup logic within
//    SaveToTempDir for the io.Copy error case.
//    Example:
//    _, errCopy := io.Copy(out, file)
//    if errCopy != nil {
//        out.Close()        // Close the file first
//        os.Remove(out.Name()) // Attempt to remove the temp file
//        return "", errCopy // Return the copy error
//    }
//
// 2. Input File Closing: The SaveToTempDir function does not close the input
//    `multipart.File`. This is generally the responsibility of the caller
//    who opened/received the multipart file. The tests correctly reflect this.
