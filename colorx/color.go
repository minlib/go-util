// Package colorx provides utilities for color conversion between different color formats.
// It supports conversion between RGB/RGBA and hexadecimal color representations.
package colorx

import (
	"image/color"
	"strconv"
	"strings"
)

// toHex converts a uint8 value to its two-character hexadecimal representation.
// The result is always two characters, with a leading zero if needed.
func toHex(c uint8) string {
	result := strconv.FormatUint(uint64(c), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

// RGB2Hex converts RGB color values to a hexadecimal color string.
// The input values r, g, b should be in the range 0-255.
// Returns a string in the format "#rrggbb".
func RGB2Hex(r, g, b uint8) string {
	return "#" + toHex(r) + toHex(g) + toHex(b)
}

// RGBA2Hex converts a color.RGBA struct to a hexadecimal color string.
// Returns a string in the format "#rrggbbaa".
func RGBA2Hex(c color.RGBA) string {
	return "#" + toHex(c.R) + toHex(c.G) + toHex(c.B) + toHex(c.A)
}

// Hex2RGB converts a hexadecimal color string to RGB values.
// The input hex string can optionally start with "#".
// Returns the red, green, and blue values in the range 0-255.
// Example: Hex2RGB("#ff0000") returns (255, 0, 0).
func Hex2RGB(hex string) (uint8, uint8, uint8) {
	rgba, _ := Hex2RGBA(hex)
	return rgba.R, rgba.G, rgba.B
}

// Hex2RGBA converts a hexadecimal color string to a color.RGBA struct.
// The input hex string can optionally start with "#".
// Supports both short (3 or 4 character) and long (6 or 8 character) hex formats:
//   - 3 characters (e.g., "#f00") expands to 6 characters ("#ff0000")
//   - 4 characters (e.g., "#f00f") expands to 8 characters ("#ff0000ff")
//   - 6 characters (e.g., "#ff0000") adds full opacity (255)
//   - 8 characters (e.g., "#ff0000ff") uses all components as provided
//
// Returns a color.RGBA struct with R, G, B, A values in the range 0-255.
func Hex2RGBA(hex string) (color.RGBA, error) {
	// Remove # prefix if present
	hex = strings.TrimPrefix(hex, "#")
	black := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// Handle different hex formats
	switch len(hex) {
	case 3: // RGB format (e.g., #f00)
		r, err := strconv.ParseUint(string(hex[0]), 16, 4)
		if err != nil {
			return black, err
		}
		g, err := strconv.ParseUint(string(hex[1]), 16, 4)
		if err != nil {
			return black, err
		}
		b, err := strconv.ParseUint(string(hex[2]), 16, 4)
		if err != nil {
			return black, err
		}
		// Expand short format (e.g., f -> ff)
		return color.RGBA{
			R: uint8(r<<4 | r),
			G: uint8(g<<4 | g),
			B: uint8(b<<4 | b),
			A: 255,
		}, nil
	case 4: // RGBA format (e.g., #f00f)
		r, err := strconv.ParseUint(string(hex[0]), 16, 4)
		if err != nil {
			return black, err
		}
		g, err := strconv.ParseUint(string(hex[1]), 16, 4)
		if err != nil {
			return black, err
		}
		b, err := strconv.ParseUint(string(hex[2]), 16, 4)
		if err != nil {
			return black, err
		}
		a, err := strconv.ParseUint(string(hex[3]), 16, 4)
		if err != nil {
			return black, err
		}
		// Expand short format (e.g., f -> ff)
		return color.RGBA{
			R: uint8(r<<4 | r),
			G: uint8(g<<4 | g),
			B: uint8(b<<4 | b),
			A: uint8(a<<4 | a),
		}, nil
	case 6: // RRGGBB format (e.g., #ff0000)
		r, err := strconv.ParseUint(hex[0:2], 16, 8)
		if err != nil {
			return black, err
		}
		g, err := strconv.ParseUint(hex[2:4], 16, 8)
		if err != nil {
			return black, err
		}
		b, err := strconv.ParseUint(hex[4:6], 16, 8)
		if err != nil {
			return black, err
		}
		return color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 255,
		}, nil
	case 8: // RRGGBBAA format (e.g., #ff0000ff)
		r, err := strconv.ParseUint(hex[0:2], 16, 8)
		if err != nil {
			return black, err
		}
		g, err := strconv.ParseUint(hex[2:4], 16, 8)
		if err != nil {
			return black, err
		}
		b, err := strconv.ParseUint(hex[4:6], 16, 8)
		if err != nil {
			return black, err
		}
		a, err := strconv.ParseUint(hex[6:8], 16, 8)
		if err != nil {
			return black, err
		}
		return color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}, nil
	default:
		// Invalid hex format, return default black with full opacity
		return black, nil
	}
}
