package drawer

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// CircleDraw represents a circular image drawing component.
type CircleDraw struct {
	image    image.Image
	point    image.Point
	diameter int
}

// NewCircleDraw creates a new CircleDraw instance.
func NewCircleDraw(img image.Image, p image.Point, d int) CircleDraw {
	return CircleDraw{img, p, d}
}

// ColorModel returns the color model of the circular image.
func (d CircleDraw) ColorModel() color.Model {
	return d.image.ColorModel()
}

// Bounds returns the bounds of the circular image.
func (d CircleDraw) Bounds() image.Rectangle {
	return image.Rect(0, 0, d.diameter, d.diameter)
}

// At returns the color at the specified coordinates in the circular image.
func (d CircleDraw) At(x, y int) color.Color {
	// Calculate the distance from the pixel to the center
	diameter := d.diameter
	centerX, centerY := float64(diameter)/2, float64(diameter)/2
	distance := math.Sqrt(math.Pow(float64(x)-centerX, 2) + math.Pow(float64(y)-centerY, 2))

	// Return transparent color if outside the circle
	if distance > float64(diameter)/2 {
		// return color.RGBA{R: 255, G: 255, B: 255, A: 0} // Transparent
		return color.RGBA{A: 0} // Transparent
	}

	// Return the color from the original image
	return d.image.At(d.point.X+x, d.point.Y+y)
}

// DrawCircle creates a circular version of the source image.
func DrawCircle(srcImage image.Image) image.Image {
	// Get image dimensions
	srcBounds := srcImage.Bounds()
	width := srcBounds.Dx()
	height := srcBounds.Dy()

	// Create target image
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(dstImage, dstImage.Bounds(), srcImage, srcBounds.Min, draw.Over)

	// Get the diameter (smaller of width and height)
	diameter := width
	if height < width {
		diameter = height
	}

	// Return the circular drawing object
	return NewCircleDraw(dstImage, image.Point{X: 0, Y: 0}, diameter)
}
