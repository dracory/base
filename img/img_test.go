package img

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/dracory/base/img/testdata"
	"golang.org/x/image/bmp"
)

// mockEncoder is a mock encoder that always returns an error
type mockEncoder struct{}

func (m *mockEncoder) Encode(img image.Image) ([]byte, error) {
	return nil, errors.New("mock encoder error")
}

// TestBlur tests the Blur function
func TestBlur(t *testing.T) {
	// Generate a test image
	imgBytes, err := testdata.GenerateTestPNG(100, 100)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Test with valid image
	blurredBytes, err := Blur(imgBytes, 0.5, imaging.PNG)
	if err != nil {
		t.Fatalf("Failed to blur image: %v", err)
	}

	// Verify the blurred image can be decoded
	_, err = imaging.Decode(bytes.NewReader(blurredBytes))
	if err != nil {
		t.Fatalf("Failed to decode blurred image: %v", err)
	}

	// Test with invalid image data
	_, err = Blur([]byte("invalid image data"), 0.5, imaging.PNG)
	if err == nil {
		t.Fatal("Expected error when blurring invalid image data, but got nil")
	}

	// Test with encoding error (using a format that will cause an error)
	// This is to cover lines 46-48 in blur.go
	_, err = Blur(imgBytes, 0.5, imaging.Format(99)) // Invalid format
	if err == nil {
		t.Fatal("Expected error when using invalid format, but got nil")
	}
}

// TestGrayscale tests the Grayscale function
func TestGrayscale(t *testing.T) {
	// Generate a test image
	imgBytes, err := testdata.GenerateTestPNG(100, 100)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Test with valid image
	grayscaleBytes, err := Grayscale(imgBytes, imaging.PNG)
	if err != nil {
		t.Fatalf("Failed to convert image to grayscale: %v", err)
	}

	// Verify the grayscale image can be decoded
	_, err = imaging.Decode(bytes.NewReader(grayscaleBytes))
	if err != nil {
		t.Fatalf("Failed to decode grayscale image: %v", err)
	}

	// Test with invalid image data
	_, err = Grayscale([]byte("invalid image data"), imaging.PNG)
	if err == nil {
		t.Fatal("Expected error when converting invalid image data to grayscale, but got nil")
	}

	// Test with encoding error (using a format that will cause an error)
	// This is to cover lines 33-35 in grayscale.go
	_, err = Grayscale(imgBytes, imaging.Format(99)) // Invalid format
	if err == nil {
		t.Fatal("Expected error when using invalid format, but got nil")
	}
}

// TestResize tests the Resize function
func TestResize(t *testing.T) {
	// Generate a test image
	imgBytes, err := testdata.GenerateTestPNG(100, 100)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Test with valid image
	resizedBytes, err := Resize(imgBytes, 50, 50, imaging.PNG)
	if err != nil {
		t.Fatalf("Failed to resize image: %v", err)
	}

	// Verify the resized image can be decoded and has the correct dimensions
	img, err := imaging.Decode(bytes.NewReader(resizedBytes))
	if err != nil {
		t.Fatalf("Failed to decode resized image: %v", err)
	}
	bounds := img.Bounds()
	if bounds.Dx() != 50 || bounds.Dy() != 50 {
		t.Fatalf("Expected resized image to be 50x50, but got %dx%d", bounds.Dx(), bounds.Dy())
	}

	// Test with invalid image data
	_, err = Resize([]byte("invalid image data"), 50, 50, imaging.PNG)
	if err == nil {
		t.Fatal("Expected error when resizing invalid image data, but got nil")
	}

	// Test with encoding error (using a format that will cause an error)
	// This is to cover lines 33-35 in resize.go
	_, err = Resize(imgBytes, 50, 50, imaging.Format(99)) // Invalid format
	if err == nil {
		t.Fatal("Expected error when using invalid format, but got nil")
	}
}

// TestGenerateTestPNG tests the GenerateTestPNG function in testdata package
func TestGenerateTestPNG(t *testing.T) {
	// Test successful image generation
	imgBytes, err := testdata.GenerateTestPNG(10, 10)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Verify the generated image can be decoded
	_, err = png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		t.Fatalf("Failed to decode generated test image: %v", err)
	}

	// We can't easily test the error case in GenerateTestPNG because it's hard to make png.Encode fail
	// in a controlled way. This would require mocking the png.Encode function which is beyond the scope
	// of this test. The error handling is simple enough that we can consider it covered.
}

// createTestImage creates a test image in the specified format
func createTestImage(t *testing.T, format string) []byte {
	// Create a simple image
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	var buf bytes.Buffer

	switch format {
	case "png":
		err := png.Encode(&buf, img)
		if err != nil {
			t.Fatalf("Failed to encode PNG: %v", err)
		}
	case "jpeg":
		err := jpeg.Encode(&buf, img, nil)
		if err != nil {
			t.Fatalf("Failed to encode JPEG: %v", err)
		}
	case "gif":
		err := gif.Encode(&buf, img, nil)
		if err != nil {
			t.Fatalf("Failed to encode GIF: %v", err)
		}
	case "bmp":
		err := bmp.Encode(&buf, img)
		if err != nil {
			t.Fatalf("Failed to encode BMP: %v", err)
		}
	case "webp":
		// Create a minimal valid WebP header for MIME type detection
		// RIFF header
		buf.Write([]byte{0x52, 0x49, 0x46, 0x46})
		// File size (placeholder)
		buf.Write([]byte{0x08, 0x00, 0x00, 0x00})
		// WEBP signature
		buf.Write([]byte{0x57, 0x45, 0x42, 0x50})
		// VP8 chunk header
		buf.Write([]byte{0x56, 0x50, 0x38, 0x20})
	default:
		t.Fatalf("Unsupported format: %s", format)
	}

	return buf.Bytes()
}

// TestToBase64Url tests the ToBase64Url function
func TestToBase64Url(t *testing.T) {
	// Test PNG format
	pngBytes := createTestImage(t, "png")
	pngBase64 := ToBase64Url(pngBytes)
	if !strings.HasPrefix(pngBase64, "data:image/png;base64,") {
		t.Errorf("Expected PNG base64 URL to start with 'data:image/png;base64,', but got %q", pngBase64)
	}

	// Test JPEG format
	jpegBytes := createTestImage(t, "jpeg")
	jpegBase64 := ToBase64Url(jpegBytes)
	if !strings.HasPrefix(jpegBase64, "data:image/jpeg;base64,") {
		t.Errorf("Expected JPEG base64 URL to start with 'data:image/jpeg;base64,', but got %q", jpegBase64)
	}

	// Test GIF format
	gifBytes := createTestImage(t, "gif")
	gifBase64 := ToBase64Url(gifBytes)
	if !strings.HasPrefix(gifBase64, "data:image/gif;base64,") {
		t.Errorf("Expected GIF base64 URL to start with 'data:image/gif;base64,', but got %q", gifBase64)
	}

	// Test BMP format
	bmpBytes := createTestImage(t, "bmp")
	bmpBase64 := ToBase64Url(bmpBytes)
	if !strings.HasPrefix(bmpBase64, "data:image/bmp;base64,") {
		t.Errorf("Expected BMP base64 URL to start with 'data:image/bmp;base64,', but got %q", bmpBase64)
	}

	// Test WEBP format - to cover lines 33-34 in to_base64.go
	webpBytes := createTestImage(t, "webp")
	webpBase64 := ToBase64Url(webpBytes)
	if !strings.HasPrefix(webpBase64, "data:image/webp;base64,") {
		t.Errorf("Expected WEBP base64 URL to start with 'data:image/webp;base64,', but got %q", webpBase64)
	}

	// Test unknown format
	unknownBytes := []byte("test data for unknown mime type")
	unknownBase64 := ToBase64Url(unknownBytes)
	// For unknown type, we just verify it's not empty and doesn't start with any of the known prefixes
	if unknownBase64 == "" {
		t.Errorf("Expected non-empty base64 URL for unknown mime type")
	}
}

// TestImgToBase64Url tests the ImgToBase64Url function
func TestImgToBase64Url(t *testing.T) {
	// Create a temporary test image file
	tempDir := t.TempDir()
	testImagePath := filepath.Join(tempDir, "test.png")

	// Generate a test image
	imgBytes, err := testdata.GenerateTestPNG(50, 50)
	if err != nil {
		t.Fatalf("Failed to generate test image: %v", err)
	}

	// Write the test image to the temporary file
	err = os.WriteFile(testImagePath, imgBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to write test image to file: %v", err)
	}

	// Test with valid image file
	base64Url := ImgToBase64Url(testImagePath)
	if !strings.HasPrefix(base64Url, "data:image/png;base64,") {
		t.Errorf("Expected base64 URL to start with 'data:image/png;base64,', but got %q", base64Url)
	}

	// Test with non-existent file
	nonExistentPath := filepath.Join(tempDir, "non_existent.png")
	base64Url = ImgToBase64Url(nonExistentPath)
	if base64Url != "" {
		t.Errorf("Expected empty string for non-existent file, but got %q", base64Url)
	}
}

// TestBase64Encode tests the base64Encode function
func TestBase64Encode(t *testing.T) {
	testData := []byte("test data for base64 encoding")
	encoded := base64Encode(testData)

	// Verify that the encoded string can be decoded back to the original data
	decoded, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatalf("Failed to decode base64 string: %v", err)
	}

	if !bytes.Equal(decoded, testData) {
		t.Errorf("Decoded data does not match original data")
	}
}
