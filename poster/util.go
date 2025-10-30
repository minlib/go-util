package poster

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/draw"
	"os"
)

// NewRGBA returns a new RGBA image with the given bounds.
func NewRGBA(x0, y0, x1, y1 int) *image.RGBA {
	return image.NewRGBA(image.Rect(x0, y0, x1, y1))
}

// DrawImage draw image
func DrawImage(dst draw.Image, src image.Image, sp image.Point) {
	draw.Draw(dst, dst.Bounds(), src, sp, draw.Over)
}

// GetFont Load and parse a truetype font from the specified file path
func GetFont(path string) (*truetype.Font, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fileBytes)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// GetFontFace Create a font.Face from a truetype.Font with the specified point size
func GetFontFace(f *truetype.Font, size float64) font.Face {
	return truetype.NewFace(f, &truetype.Options{
		Size: size,
	})
}

// GetFontAndFace Load and parse a truetype font from the specified file path
func GetFontAndFace(path string, size float64) (font.Face, error) {
	f, err := GetFont(path)
	if err != nil {
		return nil, err
	}
	fontFace := GetFontFace(f, size)
	return fontFace, nil
}
