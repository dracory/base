package testdata

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

// GenerateTestPNG creates a simple test PNG image with the given width and height
func GenerateTestPNG(width, height int) ([]byte, error) {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the image with a simple pattern
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Create a simple gradient pattern
			r := uint8((x * 255) / width)
			g := uint8((y * 255) / height)
			b := uint8(((x + y) * 255) / (width + height))
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	// Encode the image to PNG
	var buf bytes.Buffer
	_ = png.Encode(&buf, img)

	return buf.Bytes(), nil
}
