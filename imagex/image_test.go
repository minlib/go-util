package imagex

import (
	"fmt"
	"image/color"
	"testing"
	"time"
)

func getOutputPath() string {
	return fmt.Sprintf("%s%d.png", "../outputs/", time.Now().UnixNano())
}

func TestNewDrawable(t *testing.T) {
	output := getOutputPath()
	// Create a simple test image
	srcImg := NewImage(50, 50, color.RGBA{R: 0, G: 255, B: 0, A: 255})

	// Convert it to a drawable image
	drawableImg := NewDrawable(srcImg)

	// Just verify it's not nil (basic test)
	if drawableImg == nil {
		t.Error("Expected a valid drawable image, got nil")
	}
	_ = SavePNG(drawableImg, output)
}

func TestNewImageWithTransparentColor(t *testing.T) {
	output := getOutputPath()
	img := NewImage(100, 100, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	_ = SavePNG(img, output)
}

func TestNewImageWithRedColor(t *testing.T) {
	output := getOutputPath()
	img := NewImage(100, 100, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	_ = SavePNG(img, output)
}

func TestNewImageFromFile(t *testing.T) {
	// First create a test image
	output := getOutputPath()
	srcImg := NewImage(50, 50, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	err := SavePNG(srcImg, output)
	if err != nil {
		t.Errorf("Failed to save test image: %v", err)
		return
	}

	// Now test loading it back
	drawableImg, err := NewImageFromFile(output)
	if err != nil {
		t.Errorf("Failed to load image from file: %v", err)
		return
	}

	// Verify it's not nil
	if drawableImg == nil {
		t.Error("Expected a valid drawable image, got nil")
	}
}

func TestNewImageFromFileWithInvalidPath(t *testing.T) {
	// Test with an invalid file path
	_, err := NewImageFromFile("./nonexistent/image.png")
	if err == nil {
		t.Error("Expected an error when loading from nonexistent file, but got nil")
	}
}
