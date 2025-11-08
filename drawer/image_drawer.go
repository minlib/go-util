package drawer

import (
	"errors"
	"fmt"
	"image"
	"image/draw"

	"github.com/minlib/go-util/imagex"
)

// ImageDraw represents a component for drawing images.
type ImageDraw struct {
	// X is the x-coordinate for drawing the image.
	X int `json:"x"`

	// Y is the y-coordinate for drawing the image.
	Y int `json:"y"`

	// Path is the path to the image file (local or remote).
	Path string `json:"path"`

	// Round indicates whether to apply circular cropping to the image.
	Round bool `json:"round"`

	// Width is the target width for resizing the image (0 means no resizing).
	Width int `json:"width"`
}

// Type returns the type identifier for the ImageDraw component.
func (d *ImageDraw) Type() string {
	return "image"
}

// Draw executes the image drawing logic.
func (d *ImageDraw) Draw(ctx *Context) error {
	// 1. Load the image resource
	srcImage, err := imagex.ReadImage(d.Path)
	if err != nil {
		return fmt.Errorf("load image failed: %w", err)
	}

	// 2. Resize the image if needed
	if d.Width > 0 {
		srcImage, err = imagex.ResizeImage(srcImage, d.Width)
		if err != nil {
			return fmt.Errorf("resize image failed: %w", err)
		}
	}

	// 3. Apply circular cropping if needed
	var drawImage image.Image
	if d.Round {
		drawImage = DrawCircle(srcImage)
	} else {
		drawImage = srcImage
	}

	// 4. Draw the image onto the canvas
	draw.Draw(ctx.Canvas, ctx.Canvas.Bounds(), drawImage, image.Point{X: -d.X, Y: -d.Y}, draw.Over)
	return nil
}

// Validate checks if the ImageDraw configuration is valid.
func (d *ImageDraw) Validate(*Context) error {
	if d.Path == "" {
		return errors.New("path is required")
	}
	if d.X < 0 || d.Y < 0 {
		return errors.New("x and y must be non-negative")
	}
	return nil
}
