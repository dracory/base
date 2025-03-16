package testdata

import (
	"bytes"
	"image/png"
	"testing"
)

// TestGenerateTestPNG tests the GenerateTestPNG function
func TestGenerateTestPNG(t *testing.T) {
	// Test successful image generation
	imgBytes, err := GenerateTestPNG(10, 10)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Verify the generated image can be decoded
	_, err = png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		t.Fatalf("Failed to decode generated test image: %v", err)
	}
}
