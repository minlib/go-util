package imagex

import (
	"bytes"
	"github.com/fogleman/gg"
	"github.com/minlib/go-util/filex"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
)

// GetResourceReader get local or remote resource file
func GetResourceReader(pathOrUrl string) (*bytes.Reader, error) {
	if strings.HasPrefix(pathOrUrl, "https://") || strings.HasPrefix(pathOrUrl, "http://") {
		resp, err := http.Get(pathOrUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fileBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	} else {
		fileBytes, err := os.ReadFile(pathOrUrl)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	}
}

// NewDrawable creates a drawable image from an existing image.
// This is useful when you need to perform drawing operations on an existing image.
// Parameters:
//   - srcImage: The source image to create a drawable copy from
//
// Returns:
//   - draw.Image: A new drawable image containing the contents of the source image
func NewDrawable(srcImage image.Image) draw.Image {
	// Get the bounds of the source image
	bounds := srcImage.Bounds()

	// Create a new RGBA image with the same dimensions as the source
	img := image.NewRGBA(bounds)

	// Copy the source image to the new image using draw.Draw for efficiency
	draw.Draw(img, bounds, srcImage, bounds.Min, draw.Src)

	return img
}

// NewImage creates a new image filled with a solid color.
// This function returns a draw.Image which can be used for further drawing operations.
// Parameters:
//   - width: The width of the new image in pixels
//   - height: The height of the new image in pixels
//   - fillColor: The color to fill the image with
//
// Returns:
//   - draw.Image: A new image filled with the specified color
func NewImage(width, height int, fillColor color.Color) draw.Image {
	// Create a new RGBA image with the specified dimensions
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the entire image with the background color using draw.Draw for efficiency
	draw.Draw(img, img.Bounds(), &image.Uniform{C: fillColor}, image.Point{}, draw.Src)

	return img
}

// NewImageFromFile creates a drawable image from an image file.
// This is useful when you need to perform drawing operations on an image loaded from a file.
// Parameters:
//   - srcImagePath: The path to the source image file
//
// Returns:
//   - draw.Image: A new drawable image containing the contents of the source image file
//   - error: Any error encountered while reading or processing the image
func NewImageFromFile(srcImagePath string) (draw.Image, error) {
	// Read the source image from file
	srcImage, err := ReadImage(srcImagePath)
	if err != nil {
		// If there's an error reading the image, return the error
		return nil, err
	}
	img := NewDrawable(srcImage)
	return img, nil
}

// ReadImage reads an image from a file or URL.
// Parameters:
//   - pathOrUrl: The path or URL to the image file
//
// Returns:
//   - image.Image: The decoded image
//   - error: Any error encountered while reading or decoding the image
func ReadImage(pathOrUrl string) (image.Image, error) {
	reader, err := GetResourceReader(pathOrUrl)
	if err != nil {
		return nil, err
	}
	// When loading JPG format images, you must import the "image/jpeg" dependency,
	// otherwise you will get an "unknown format" error
	srcImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return srcImage, err
}

// GetSize gets the size of the image.
// Parameters:
//   - path: The path to the image file
//
// Returns:
//   - int: The width of the image
//   - int: The height of the image
//   - error: Any error encountered while reading the image
func GetSize(path string) (int, int, error) {
	if strings.HasSuffix(path, ".webp") {
		return GetWebpSize(path)
	}
	srcImage, err := ReadImage(path)
	if err != nil {
		return 0, 0, err
	}
	// Get the width and height of the image
	width := srcImage.Bounds().Dx()
	height := srcImage.Bounds().Dy()
	return width, height, nil
}

// SavePNG saves an image as a PNG file.
// Parameters:
//   - src: The source image to save
//   - filename: The path where the PNG file will be saved
//
// Returns:
//   - error: Any error encountered while saving the image
func SavePNG(src image.Image, filename string) error {
	if err := filex.MkdirAll(filename); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, src)
}

// SaveJPG saves an image as a JPG file.
// Parameters:
//   - src: The source image to save
//   - filename: The path where the JPG file will be saved
//   - quality: The quality of the JPG image (1-100)
//
// Returns:
//   - error: Any error encountered while saving the image
func SaveJPG(src image.Image, filename string, quality int) error {
	if err := filex.MkdirAll(filename); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, src, &jpeg.Options{
		Quality: quality,
	})
}

// ResizeImage resizes an image proportionally to the specified width.
// Parameters:
//   - src: The source image to resize
//   - width: The target width for the resized image
//
// Returns:
//   - image.Image: The resized image
//   - error: Any error encountered during resizing
func ResizeImage(src image.Image, width int) (image.Image, error) {
	bounds := src.Bounds()
	srcWidth, srcHeight := bounds.Dx(), bounds.Dy()
	newHeight := int(float64(srcHeight) * (float64(width) / float64(srcWidth)))
	newImage := image.NewRGBA(image.Rect(0, 0, width, newHeight))
	for x := 0; x < width; x++ {
		for y := 0; y < newHeight; y++ {
			srcX := int(float64(x) * (float64(srcWidth) / float64(width)))
			srcY := int(float64(y) * (float64(srcHeight) / float64(newHeight)))
			if srcX < 0 {
				srcX = 0
			}
			if srcY < 0 {
				srcY = 0
			}
			if srcX >= srcWidth {
				srcX = srcWidth - 1
			}
			if srcY >= srcHeight {
				srcY = srcHeight - 1
			}
			newImage.Set(x, y, src.At(srcX, srcY))
		}
	}
	return newImage, nil
}

// DrawCircleBorder adds a circular border to an image.
// This function creates a circular border around the input image by:
// 1. Calculating the diameter based on the smaller dimension of the image
// 2. Creating a new canvas with dimensions to accommodate the border
// 3. Drawing a circular border on the canvas
// 4. Clipping the source image to fit within the circular area
//
// Parameters:
//   - src: The source image to add a circular border to
//   - borderWidth: The width of the border in pixels (must be positive)
//   - borderColor: The color of the border (can be any color.Color implementation)
//
// Returns:
//   - image.Image: The image with the added circular border
//   - If borderWidth <= 0, returns the original image unchanged
//   - If src has invalid dimensions (width or height <= 0), returns the original image unchanged
func DrawCircleBorder(src image.Image, borderWidth int, borderColor color.Color) image.Image {
	// Handle edge cases
	if borderWidth <= 0 {
		return src // No border to add
	}

	// Get source image dimensions
	srcBounds := src.Bounds()
	srcWidth, srcHeight := srcBounds.Dx(), srcBounds.Dy()

	// Handle invalid image dimensions
	if srcWidth <= 0 || srcHeight <= 0 {
		return src
	}

	// For circular images, draw a circular border
	// Get the diameter (smaller of width and height)
	diameter := srcWidth
	if srcHeight < srcWidth {
		diameter = srcHeight
	}

	// Calculate total size including border
	totalSize := diameter + borderWidth*2

	// Calculate center coordinates for the border circle
	centerX, centerY := float64(totalSize)/2, float64(totalSize)/2

	// Calculate the radius of the border circle
	borderRadius := float64(diameter)/2 + float64(borderWidth)

	// Calculate the offset to center the source image
	offsetX, offsetY := (srcWidth-diameter)/2, (srcHeight-diameter)/2

	// Create a new context with the total size
	dc := gg.NewContext(totalSize, totalSize)

	// Draw border
	dc.DrawCircle(centerX, centerY, borderRadius)
	dc.SetColor(borderColor)
	dc.Fill()

	// Draw a circular clipping path for the inner image
	imageRadius := float64(diameter)/2 - 0.5
	dc.DrawCircle(centerX, centerY, imageRadius)
	dc.Clip()

	// Draw the source image in the center
	dc.DrawImage(src, borderWidth-offsetX, borderWidth-offsetY)

	return dc.Image()
}

// DrawRectangleBorder adds a rectangular border to an image with optional rounded corners.
// This function creates a rectangular border around the input image by:
// 1. Calculating new dimensions to accommodate the border
// 2. Creating a new canvas with the expanded dimensions
// 3. Drawing either a rectangle or rounded rectangle for the border
// 4. Drawing the source image centered within the border
//
// Parameters:
//   - src: The source image to add a border to
//   - borderWidth: The width of the border in pixels (must be positive)
//   - borderColor: The color of the border (can be any color.Color implementation)
//   - radius: The corner radius for rounded rectangles (0 for sharp corners)
//
// Returns:
//   - image.Image: The image with the added rectangular border
//   - If borderWidth <= 0, returns the original image unchanged
//   - If src has invalid dimensions (width or height <= 0), returns the original image unchanged
func DrawRectangleBorder(src image.Image, borderWidth int, borderColor color.Color, radius float64) image.Image {
	// Handle edge cases
	if borderWidth <= 0 {
		return src // No border to add
	}

	// Get source image dimensions
	srcBounds := src.Bounds()
	srcWidth, srcHeight := srcBounds.Dx(), srcBounds.Dy()

	// Handle invalid image dimensions
	if srcWidth <= 0 || srcHeight <= 0 {
		return src
	}

	// Calculate new dimensions with border (both sides)
	newWidth := srcWidth + 2*borderWidth
	newHeight := srcHeight + 2*borderWidth

	// Create a new context with the new dimensions
	dc := gg.NewContext(newWidth, newHeight)

	// Draw border with or without rounded corners
	if radius <= 0 {
		// Sharp corners
		dc.DrawRectangle(0, 0, float64(newWidth), float64(newHeight))
	} else {
		// Rounded corners
		dc.DrawRoundedRectangle(0, 0, float64(newWidth), float64(newHeight), radius)
	}

	dc.SetColor(borderColor)
	dc.Fill()

	// Draw the source image in the center
	dc.DrawImage(src, borderWidth, borderWidth)

	return dc.Image()
}

// DrawShadow adds a shadow effect to an image using the gg library.
// Parameters:
//   - src: The source image to add a shadow to
//   - offsetX: The horizontal offset of the shadow
//   - offsetY: The vertical offset of the shadow
//   - blurRadius: The blur radius of the shadow (0 for no blur)
//   - shadowColor: The color of the shadow
//   - circular: If true, draws a circular shadow (for circular images), otherwise draws a rectangular shadow
//
// Returns:
//   - image.Image: The image with the added shadow
func DrawShadow(src image.Image, offsetX, offsetY, blurRadius int, shadowColor color.Color, circular bool) image.Image {
	srcBounds := src.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// Calculate shadow padding based on blur radius
	shadowPadding := int(math.Ceil(float64(blurRadius) * 3))

	if circular {
		// For circular images, draw a circular shadow
		// Get the diameter (smaller of width and height)
		diameter := srcWidth
		if srcHeight < srcWidth {
			diameter = srcHeight
		}

		// Calculate canvas size
		canvasWidth := diameter + shadowPadding*2
		canvasHeight := diameter + shadowPadding*2

		// Create a new context for drawing
		dc := gg.NewContext(canvasWidth, canvasHeight)

		// Draw circular shadow using multiple samples to simulate blur
		centerX := float64(canvasWidth)/2 + float64(offsetX)
		centerY := float64(canvasHeight)/2 + float64(offsetY)
		radius := float64(diameter)/2 + 1

		// Create shadow with simulated blur effect
		if blurRadius > 0 {
			// Use multiple passes with decreasing alpha to simulate blur
			alpha := 0.5
			steps := blurRadius * 2
			if steps < 1 {
				steps = 1
			}

			for i := 0; i < steps; i++ {
				// Calculate radius for this step
				stepRadius := radius + float64(i)*0.5

				// Set color with reduced alpha for blur effect
				shadowAlpha := uint8(alpha * 255 / float64(steps-i))
				var colorWithAlpha color.Color

				if shadowColor != nil {
					// Extract original color and apply our alpha
					r, g, b, _ := shadowColor.RGBA()
					colorWithAlpha = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), shadowAlpha}
				} else {
					colorWithAlpha = color.RGBA{0, 0, 0, shadowAlpha} // Default black shadow
				}

				dc.DrawCircle(centerX, centerY, stepRadius)
				dc.SetColor(colorWithAlpha)
				dc.Fill()

				// Reduce alpha for next iteration
				alpha *= 0.8
			}
		} else {
			// No blur, just draw simple circular shadow
			dc.DrawCircle(centerX, centerY, radius)
			if shadowColor != nil {
				dc.SetColor(shadowColor)
			} else {
				dc.SetColor(color.RGBA{0, 0, 0, 128}) // Default semi-transparent black
			}
			dc.Fill()
		}

		// Draw a circular clipping path for the main image
		dc.Identity()
		centerX, centerY = float64(canvasWidth)/2, float64(canvasHeight)/2

		// Draw the source image in the center
		imageRadius := float64(diameter)/2 - 0.5
		dc.DrawCircle(centerX, centerY, imageRadius)
		dc.Clip()

		// Calculate the offset to center the source image
		offsetXCenter := (srcWidth - diameter) / 2
		offsetYCenter := (srcHeight - diameter) / 2

		dc.DrawImage(src, shadowPadding-offsetXCenter, shadowPadding-offsetYCenter)

		return dc.Image()
	} else {
		// Original rectangular shadow implementation
		// Calculate canvas size
		canvasWidth := srcWidth + shadowPadding*2
		canvasHeight := srcHeight + shadowPadding*2

		// Create a new context for drawing
		dc := gg.NewContext(canvasWidth, canvasHeight)

		// Draw shadow using multiple samples to simulate blur
		centerX := float64(canvasWidth)/2 + float64(offsetX)
		centerY := float64(canvasHeight)/2 + float64(offsetY)

		// Create shadow with simulated blur effect
		if blurRadius > 0 {
			// Use multiple passes with decreasing alpha to simulate blur
			alpha := 0.5
			steps := blurRadius * 2
			if steps < 1 {
				steps = 1
			}

			for i := 0; i < steps; i++ {
				// Calculate dimensions for this step
				width := float64(srcWidth) + float64(i)*0.5
				height := float64(srcHeight) + float64(i)*0.5

				// Set color with reduced alpha for blur effect
				shadowAlpha := uint8(alpha * 255 / float64(steps-i))
				var colorWithAlpha color.Color

				if shadowColor != nil {
					// Extract original color and apply our alpha
					r, g, b, _ := shadowColor.RGBA()
					colorWithAlpha = color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: shadowAlpha}
				} else {
					colorWithAlpha = color.RGBA{A: shadowAlpha} // Default black shadow
				}

				dc.DrawRectangle(centerX-float64(width)/2, centerY-float64(height)/2, width, height)
				dc.SetColor(colorWithAlpha)
				dc.Fill()

				// Reduce alpha for next iteration
				alpha *= 0.8
			}
		} else {
			// No blur, just draw simple shadow
			width := float64(srcWidth)
			height := float64(srcHeight)

			dc.DrawRectangle(centerX-width/2, centerY-height/2, width, height)
			if shadowColor != nil {
				dc.SetColor(shadowColor)
			} else {
				dc.SetColor(color.RGBA{A: 128}) // Default semi-transparent black
			}
			dc.Fill()
		}

		// Draw the source image in the center
		centerX = float64(canvasWidth) / 2
		centerY = float64(canvasHeight) / 2
		dc.DrawImage(src, int(centerX)-srcWidth/2, int(centerY)-srcHeight/2)

		return dc.Image()
	}
}

// DrawCircle creates a circular version of the source image.
// This function only handles basic circular cropping without border or shadow.
// Parameters:
//   - src: The source image to convert to a circle
//
// Returns:
//   - image.Image: The circular version of the source image
func DrawCircle(src image.Image) image.Image {
	// Get image dimensions
	srcBounds := src.Bounds()
	width := srcBounds.Dx()
	height := srcBounds.Dy()

	// Get the diameter (smaller of width and height)
	diameter := width
	if height < width {
		diameter = height
	}

	// Create a new context with the diameter as both width and height
	dc := gg.NewContext(diameter, diameter)

	// Draw a circular clipping path
	centerX, centerY := float64(diameter)/2, float64(diameter)/2
	radius := float64(diameter)/2 - 0.5 // Subtract 0.5 for better anti-aliasing

	// Create circular path
	dc.DrawCircle(centerX, centerY, radius)
	dc.Clip()

	// Calculate the offset to center the source image
	offsetX := (width - diameter) / 2
	offsetY := (height - diameter) / 2

	// Draw the source image onto the circular context
	dc.DrawImage(src, -offsetX, -offsetY)

	// Return the resulting image
	return dc.Image()
}
