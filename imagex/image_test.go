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

// TestNewDrawable tests creating a drawable image from an existing image
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

// TestNewImageWithTransparentColor tests creating an image with transparent color
func TestNewImageWithTransparentColor(t *testing.T) {
	output := getOutputPath()
	img := NewImage(100, 100, color.RGBA{R: 0, G: 0, B: 0, A: 0})
	_ = SavePNG(img, output)
}

// TestNewImageWithRedColor tests creating an image with red color
func TestNewImageWithRedColor(t *testing.T) {
	output := getOutputPath()
	img := NewImage(100, 100, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	_ = SavePNG(img, output)
}

// TestNewImageFromFile tests creating an image from a file
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

// TestNewImageFromFileWithInvalidPath tests creating an image from an invalid file path
func TestNewImageFromFileWithInvalidPath(t *testing.T) {
	// Test with an invalid file path
	_, err := NewImageFromFile("./nonexistent/image.png")
	if err == nil {
		t.Error("Expected an error when loading from nonexistent file, but got nil")
	}
}

// TestReadImageForBmp tests reading image bytes for bmp
func TestReadImageForBmp(t *testing.T) {
	image, err := ReadImage("../assets/images/test.bmp")
	fmt.Println(image, err)
}

// TestReadImage tests reading image bytes for webp
func TestReadImageForWebp(t *testing.T) {
	image, err := ReadImage("../assets/images/test.webp")
	fmt.Println(image, err)
}

// TestGetSizeForBmp tests getting image size for bmp
func TestGetSizeForBmp(t *testing.T) {
	width, height, err := GetSize("../assets/images/test.bmp")
	fmt.Println(width, height, err)
}

// TestGetSizeForWebp tests getting image size for webp
func TestGetSizeForWebp(t *testing.T) {
	width, height, err := GetSize("../assets/images/test.webp")
	fmt.Println(width, height, err)
}

// TestDrawRectangleBorder tests adding a border to an image
func TestDrawRectangleBorder(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(200, 200, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Test adding a border to an image
	borderWidth := 10
	borderColor := color.RGBA{R: 255, G: 255, B: 255, A: 255} // White border
	withBorder := DrawRectangleBorder(srcImg, borderWidth, borderColor, 0)

	_ = SavePNG(withBorder, output)
}

// TestDrawRectangleBorder tests adding a border to an image
func TestDrawRectangleBorder2(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(200, 200, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Test adding a border to an image
	borderWidth := 10
	borderColor := color.RGBA{R: 255, G: 255, B: 255, A: 255} // White border
	withBorder := DrawRectangleBorder(srcImg, borderWidth, borderColor, 10)
	_ = SavePNG(withBorder, output)
}

// TestDrawCircleBorder tests adding a border to an image
func TestDrawCircleBorder(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(200, 200, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Test adding a border to an image
	borderWidth := 10
	borderColor := color.RGBA{R: 255, G: 255, B: 255, A: 255} // White border
	withBorder := DrawCircleBorder(srcImg, borderWidth, borderColor)
	_ = SavePNG(withBorder, output)
}

// TestAddShadow tests adding a shadow to an image
func TestAddShadow(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(200, 200, color.RGBA{R: 0, G: 255, B: 0, A: 255})

	// Test adding a shadow to an image
	offsetX, offsetY := 5, 5
	blurRadius := 3
	shadowColor := color.RGBA{R: 0, G: 0, B: 0, A: 128} // Semi-transparent black shadow
	withShadow := DrawShadow(srcImg, offsetX, offsetY, blurRadius, shadowColor, false)

	_ = SavePNG(withShadow, output)
}

// TestAddShadowWithoutBlur tests adding a shadow without blur to an image
func TestAddShadowWithoutBlur(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(200, 200, color.RGBA{R: 0, G: 255, B: 0, A: 255})

	// Test adding a shadow without blur
	offsetX, offsetY := 3, 3
	blurRadius := 10                                  // No blur
	shadowColor := color.RGBA{R: 0, G: 0, B: 0, A: 0} // Semi-transparent black shadow
	withShadow := DrawShadow(srcImg, offsetX, offsetY, blurRadius, shadowColor, false)

	_ = SavePNG(withShadow, output)
}

// TestAddBorderAndShadow tests adding border and shadow to a circular image
// This test demonstrates using AddBorder and AddShadow with circular=true parameter to handle circular images
func TestAddBorderAndShadow(t *testing.T) {
	output := getOutputPath()
	// Step 1: Create a rectangular base image (300x300 red)
	srcImg := NewImage(300, 400, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Step 2: Add circular border (5px black) - circular=true enables circular mask support
	borderWidth := 5
	borderColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}                            // Black border
	circularWithBorder := DrawRectangleBorder(srcImg, borderWidth, borderColor, 0) // Enable circular mode

	// Step 3: Add circular shadow (3x3 offset, 2px blur, semi-transparent black) - circular=true ensures shadow follows circular shape
	offsetX, offsetY := 3, 3
	blurRadius := 2
	shadowColor := color.RGBA{R: 0, G: 0, B: 0, A: 100}                                          // 40% opaque black shadow
	finalImg := DrawShadow(circularWithBorder, offsetX, offsetY, blurRadius, shadowColor, false) // Enable circular mode

	_ = SavePNG(finalImg, output)
}

// TestDrawCircle tests basic circular cropping functionality
func TestDrawCircle(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(300, 300, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Test basic circular cropping
	drawableImg := DrawCircle(srcImg)

	_ = SavePNG(drawableImg, output)
}

// TestDrawCircleOnRectangularImage tests circular cropping on rectangular image
func TestDrawCircleOnRectangularImage(t *testing.T) {
	output := getOutputPath()
	srcImg := NewImage(300, 400, color.RGBA{R: 0, G: 255, B: 0, A: 255}) // Rectangular image

	// Test circular cropping on rectangular image
	drawableImg := DrawCircle(srcImg)

	_ = SavePNG(drawableImg, output)
}

// TestDrawCircleWithBorderAndShadow tests creating a circular image with border and shadow
// This test demonstrates how to use the enhanced AddBorder and AddShadow functions with circular support
func TestDrawCircleWithBorderAndShadow(t *testing.T) {
	output := getOutputPath()
	// Create a rectangular test image
	srcImg := NewImage(300, 300, color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Step 1: Convert the image to a circle
	circularImg := DrawCircle(srcImg)

	// Step 2: Add a CIRCULAR border to the circular image (using the enhanced AddBorder with circular=true)
	borderWidth := 5
	borderColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}                           // Black border
	circularWithBorder := DrawCircleBorder(circularImg, borderWidth, borderColor) // Note the 'true' parameter

	// Step 3: Add a CIRCULAR shadow to the circular image with border (using the enhanced AddShadow with circular=true)
	offsetX, offsetY := 3, 3
	blurRadius := 2
	shadowColor := color.RGBA{R: 0, G: 0, B: 0, A: 100}                                         // Semi-transparent black shadow
	finalImg := DrawShadow(circularWithBorder, offsetX, offsetY, blurRadius, shadowColor, true) // Note the 'true' parameter

	_ = SavePNG(finalImg, output)
}
