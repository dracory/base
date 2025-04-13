package img

import (
	"net/url"
	"strings"

	"github.com/spf13/cast"
)

type PicsumURLOptions struct {
	ID        int
	Blur      int // 1 to 10
	Grayscale bool
	Seed      string // if you want random image (but staying the same)
}

// PicsumURL generates an image URL for the Lorem Picsum online service
// More info can be found at its website: https://picsum.photos/
func PicsumURL(width int, height int, opt PicsumURLOptions) string {
	// Use strings.Builder for efficient string concatenation
	var sb strings.Builder
	sb.WriteString("https://picsum.photos") // Base URL without trailing slash initially

	if opt.Seed != "" {
		sb.WriteString("/seed/")
		sb.WriteString(opt.Seed)
	}
	if opt.ID != 0 {
		sb.WriteString("/id/")
		sb.WriteString(cast.ToString(opt.ID))
	}

	// Add dimensions
	sb.WriteString("/")
	sb.WriteString(cast.ToString(width))
	sb.WriteString("/")
	sb.WriteString(cast.ToString(height))

	// Handle query parameters more robustly using net/url
	queryParams := url.Values{}
	if opt.Grayscale {
		// Picsum uses the presence of the key for grayscale
		queryParams.Add("grayscale", "")
	}
	if opt.Blur > 0 {
		// Clamp blur value to the 1-10 range as per Picsum docs
		blurVal := opt.Blur
		if blurVal > 10 {
			blurVal = 10
		}
		queryParams.Add("blur", cast.ToString(blurVal))
	}

	// Encode and append query parameters if any exist
	encodedQuery := queryParams.Encode()
	if encodedQuery != "" {
		sb.WriteString("?")
		sb.WriteString(encodedQuery)
	}

	return sb.String()
}
