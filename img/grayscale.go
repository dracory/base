package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

// Grayscale converts an image to grayscale
//
// Parameters:
// - content: The image bytes.
// - format: The output format.
//
// Returns:
// - The grayscale image as a byte slice.
// - An error if the image could not be decoded or encoded.
func Grayscale(content []byte, format imaging.Format) ([]byte, error) {
	// Decode the image
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	// Convert the image to grayscale
	dstImage := imaging.Grayscale(srcImage)

	// Encode the image as the specified format
	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}
