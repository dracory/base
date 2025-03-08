package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

// Blur blurs an image by the given amount.
//
// The amount is a float in the range of 0 to 1.0, where 0 is the original
// image and 1.0 is a fully blurred image.
//
// The format parameter is the format of the output image. It should be the
// same as the format of the input image.
//
// Example:
//
//	blurAmount := 0.5
//	blurFormat := imaging.JPEG
//	blurImage, err := Blur(imageBytes, blurAmount, blurFormat)
//
// Parameters:
//   - content: the input image as a byte slice.
//   - blur: the amount to blur the image by.
//   - format: the format of the output image.
//
// Returns:
//   - []byte: the blurred image as a byte slice.
//   - error: an error if the image could not be decoded or encoded.
func Blur(content []byte, blur float64, format imaging.Format) ([]byte, error) {
	// Decode the input image.
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	// Create a blurred version of the image.
	dstImage := imaging.Blur(srcImage, blur)

	// Encode the image to the specified format.
	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	// Return the blurred image as a byte slice.
	return buffer.Bytes(), errImageEncode
}
