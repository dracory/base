package img

import (
	"encoding/base64"
	"net/http"
)

// ToBase64Url converts a byte slice to a Base64 encoded URL string.
// It detects the MIME type of the image and prepends the appropriate
// data URI scheme header before encoding the byte slice to a Base64 string.
//
// Parameters:
// - imgBytes: The byte slice representing the image data.
//
// Returns:
// - string: The Base64 encoded URL string.
func ToBase64Url(imgBytes []byte) string {
	// Detect the MIME type of the image from the byte slice
	mimeType := http.DetectContentType(imgBytes)

	// Initialize the Base64 encoding string with the appropriate URI scheme header
	base64Encoding := ""

	// Determine the MIME type and prepend the corresponding URI scheme header
	switch mimeType {
	case "image/bmp":
		base64Encoding += "data:image/bmp;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}

	// Encode the byte slice to Base64 and append to the encoding string
	base64Encoding += base64Encode(imgBytes)

	return base64Encoding
}

func base64Encode(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}
