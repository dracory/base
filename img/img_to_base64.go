package img

import (
	"log"
	"os"
)

// ImgToBase64Url converts an image file to a Base64 encoded URL string.
//
// It reads the specified image file from disk, converts its content into a Base64
// encoded string with a data URI scheme, and returns this string.
//
// Parameters:
// - filePath: The path to the image file to be converted.
//
// Returns:
// - string: The Base64 encoded URL string.
func ImgToBase64Url(filePath string) string {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		// Log an error message if the file cannot be read
		log.Println(err)
		return ""
	}

	// Convert the byte slice to a Base64 encoded URL string
	return ToBase64Url(bytes)
}
