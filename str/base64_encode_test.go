package str_test

import (
	"bytes"
	"encoding/base64"
	"testing"
	"github.com/dracory/base/str"
)

func TestBase64Encode_SimpleText(t *testing.T) {
	input := []byte("hello")
	want := []byte("aGVsbG8=")
	got, err := str.Base64Encode(input)
	if err != nil {
		t.Errorf("Base64Encode(%q) failed: %v", input, err)
		return
	}
	if !bytes.Equal(got, want) {
		t.Errorf("Base64Encode(%q) = %q, want %q", input, got, want)
	}
}

func TestBase64Encode_EmptyInput(t *testing.T) {
	input := []byte("")
	want := []byte("")
	got, err := str.Base64Encode(input)
	if err != nil {
		t.Errorf("Base64Encode(%q) failed: %v", input, err)
		return
	}
	if !bytes.Equal(got, want) {
		t.Errorf("Base64Encode(%q) = %q, want %q", input, got, want)
	}
}

func TestBase64Encode_BinaryData(t *testing.T) {
	input := []byte{0x00, 0x01, 0x02, 0x03, 0x04}
	want := []byte("AAECAwQ=")
	got, err := str.Base64Encode(input)
	if err != nil {
		t.Errorf("Base64Encode(%v) failed: %v", input, err)
		return
	}
	if !bytes.Equal(got, want) {
		t.Errorf("Base64Encode(%v) = %q, want %q", input, got, want)
	}
}

func TestBase64Encode_LargeInput(t *testing.T) {
	// Create a large input to test buffer handling
	input := make([]byte, 1024)
	for i := range input {
		input[i] = byte(i % 256)
	}
	want := base64.StdEncoding.EncodeToString(input)
	got, err := str.Base64Encode(input)
	if err != nil {
		t.Errorf("Base64Encode(large input) failed: %v", err)
		return
	}
	if string(got) != want {
		t.Errorf("Base64Encode(large input) produced incorrect encoding")
	}
}
