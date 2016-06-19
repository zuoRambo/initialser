package initialser

import (
	"image"
	"fmt"
)

// Hex parses a "html" hex color-string, either in the 3 "#f0c" or 6 "#ff1034" digits form.
func hexToRGBA(hex string) (*image.Uniform, error) {
	format := "#%02x%02x%02x"
	factor := 1.0 / 255.0
	if len(hex) == 4 {
		format = "#%1x%1x%1x"
		factor = 1.0 / 15.0
	}
	if len(hex) == 6 {
		hex = "#" + hex
	}

	var r, g, b uint8
	n, err := fmt.Sscanf(hex, format, &r, &g, &b)
	if err != nil {
		return image.NewUniform(Color{}), err
	}
	if n != 3 {
		return image.NewUniform(Color{}), fmt.Errorf("color: %v is not a hex-color", hex)
	}
	return image.NewUniform(Color{float64(r) * factor, float64(g) * factor, float64(b) * factor}), nil

}

// A color is stored internally using sRGB (standard RGB) values in the range 0-1
type Color struct {
	R, G, B float64
}

// Implement the Go color.Color interface.
func (col Color) RGBA() (r, g, b, a uint32) {
	r = uint32(col.R * 65535.0)
	g = uint32(col.G * 65535.0)
	b = uint32(col.B * 65535.0)
	a = 0xFFFF
	return
}


