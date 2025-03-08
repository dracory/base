package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

// Resize resizes the given image to the given width and height, and returns the
// resized image in the given format.
//
// Parameters:
// - content: The image bytes to be resized.
// - width: The width to resize the image to.
// - height: The height to resize the image to.
// - format: The output format of the resized image.
//
// Returns:
// - []byte: The resized image as a byte slice.
// - error: An error if the image could not be decoded or encoded.
func Resize(content []byte, width, height int, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}
