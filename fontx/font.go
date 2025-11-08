package fontx

import (
	"fmt"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// GetFont loads and parses a TrueType font from the specified file path.
func GetFont(path string) (*truetype.Font, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read font file failed: %w", err)
	}
	f, err := freetype.ParseFont(fileBytes)
	if err != nil {
		return nil, fmt.Errorf("parse font failed: %w", err)
	}
	return f, nil
}

// GetFonts loads and parses multiple TrueType fonts from the specified file paths.
// The keys of the map are the font family names, and the values are the file paths.
func GetFonts(familyPaths map[string]string) (map[string]*truetype.Font, error) {
	fonts := make(map[string]*truetype.Font)
	for family, path := range familyPaths {
		f, err := GetFont(path)
		if err != nil {
			return nil, fmt.Errorf("load font family '%s' failed: %w", family, err)
		}
		fonts[family] = f
	}
	return fonts, nil
}

// GetFontFace creates a new font face with the specified size from a TrueType font.
func GetFontFace(f *truetype.Font, size float64) font.Face {
	return truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
}

// GetFontAndFace combines font loading and face creation into a single function.
func GetFontAndFace(path string, size float64) (font.Face, error) {
	f, err := GetFont(path)
	if err != nil {
		return nil, fmt.Errorf("get font failed: %w", err)
	}
	fontFace := GetFontFace(f, size)
	return fontFace, nil
}
