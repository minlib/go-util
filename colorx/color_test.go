package colorx

import (
	"image/color"
	"testing"
)

// TestHex2RGB tests the Hex2RGB function with various inputs.
func TestHex2RGB(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [3]uint8
	}{
		{
			name:     "Valid hex with # prefix",
			input:    "#0b21f3",
			expected: [3]uint8{11, 33, 243},
		},
		{
			name:     "Valid hex without # prefix",
			input:    "0b21f3",
			expected: [3]uint8{11, 33, 243},
		},
		{
			name:     "Black color",
			input:    "#000000",
			expected: [3]uint8{0, 0, 0},
		},
		{
			name:     "White color",
			input:    "#FFFFFF",
			expected: [3]uint8{255, 255, 255},
		},
		{
			name:     "Invalid hex - too short",
			input:    "#fff",
			expected: [3]uint8{255, 255, 255}, // Should return default black
		},
		{
			name:     "Invalid hex - too long",
			input:    "#ffffffff",
			expected: [3]uint8{255, 255, 255}, // Only first 6 chars are used
		},
		{
			name:     "Invalid hex - non-hex characters",
			input:    "#xyz123",
			expected: [3]uint8{0, 0, 0}, // Should return default black
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := Hex2RGB(tt.input)
			if r != tt.expected[0] || g != tt.expected[1] || b != tt.expected[2] {
				t.Errorf("Hex2RGB(%s) = (%d, %d, %d), expected (%d, %d, %d)",
					tt.input, r, g, b, tt.expected[0], tt.expected[1], tt.expected[2])
			}
		})
	}
}

// TestRGB2Hex tests the RGB2Hex function with various inputs.
func TestRGB2Hex(t *testing.T) {
	tests := []struct {
		name     string
		r, g, b  uint8
		expected string
	}{
		{
			name:     "Standard RGB values",
			r:        11,
			g:        33,
			b:        243,
			expected: "#0b21f3",
		},
		{
			name:     "Black color",
			r:        0,
			g:        0,
			b:        0,
			expected: "#000000",
		},
		{
			name:     "White color",
			r:        255,
			g:        255,
			b:        255,
			expected: "#ffffff",
		},
		{
			name:     "Single digit hex values",
			r:        5,
			g:        10,
			b:        15,
			expected: "#050a0f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGB2Hex(tt.r, tt.g, tt.b)
			if result != tt.expected {
				t.Errorf("RGB2Hex(%d, %d, %d) = %s, expected %s",
					tt.r, tt.g, tt.b, result, tt.expected)
			}
		})
	}
}

// TestRGBA2Hex tests the RGBA2Hex function with various inputs.
func TestRGBA2Hex(t *testing.T) {
	tests := []struct {
		name     string
		rgba     color.RGBA
		expected string
	}{
		{
			name:     "White with full opacity",
			rgba:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
			expected: "#ffffffff",
		},
		{
			name:     "Specific color with partial opacity",
			rgba:     color.RGBA{R: 11, G: 33, B: 243, A: 128},
			expected: "#0b21f380",
		},
		{
			name:     "Black with zero opacity",
			rgba:     color.RGBA{R: 0, G: 0, B: 0, A: 0},
			expected: "#00000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBA2Hex(tt.rgba)
			if result != tt.expected {
				t.Errorf("RGBA2Hex(%+v) = %s, expected %s",
					tt.rgba, result, tt.expected)
			}
		})
	}
}

// TestHex2RGBA tests the Hex2RGBA function with various inputs.
func TestHex2RGBA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected color.RGBA
	}{
		{
			name:     "8-digit hex with # prefix",
			input:    "#ffffffff",
			expected: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
		{
			name:     "8-digit hex without # prefix",
			input:    "0b21f380",
			expected: color.RGBA{R: 11, G: 33, B: 243, A: 128},
		},
		{
			name:     "6-digit hex (defaults to full opacity)",
			input:    "#0b21f3",
			expected: color.RGBA{R: 11, G: 33, B: 243, A: 255},
		},
		{
			name:     "6-digit hex uppercase",
			input:    "#0B21F3",
			expected: color.RGBA{R: 11, G: 33, B: 243, A: 255},
		},
		{
			name:     "4-digit hex (RGBA short format)",
			input:    "#F00F", // Red with full opacity
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:     "3-digit hex (RGB short format)",
			input:    "#F00", // Red with full opacity
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:     "Invalid hex (defaults to black with full opacity)",
			input:    "#XYZ",
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Hex2RGBA(tt.input)
			if result != tt.expected {
				t.Errorf("Hex2RGBA(%s) = %+v, expected %+v",
					tt.input, result, tt.expected)
			}
		})
	}
}
