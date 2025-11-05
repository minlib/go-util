// Package poster provides functionality for creating posters with various drawing components.
package poster

import (
	"errors"
	"fmt"
	"image"
	"image/draw"

	"github.com/minlib/go-util/qrcodex"
	"github.com/skip2/go-qrcode"
)

// QRCodeDraw represents a component for drawing QR codes.
type QRCodeDraw struct {
	// X is the x-coordinate for drawing the QR code.
	X int `json:"x"`

	// Y is the y-coordinate for drawing the QR code.
	Y int `json:"y"`

	// Size is the size of the QR code.
	Size int `json:"size"`

	// Content is the content to be encoded in the QR code (URL or text).
	Content string `json:"content"`
}

// Type returns the type identifier for the QRCodeDraw component.
func (d *QRCodeDraw) Type() string {
	return "qrcode"
}

// Draw executes the QR code drawing logic.
func (d *QRCodeDraw) Draw(ctx *Context) error {
	// Generate the QR code image
	qrImage, err := qrcodex.QrcodeWithBorder(d.Content, qrcode.Medium, d.Size)
	if err != nil {
		return fmt.Errorf("generate qrcode failed: %w", err)
	}
	// Draw the QR code onto the canvas
	// imagex.DrawImage(ctx.Canvas, qrImage, qrImage.Bounds().Min.Sub(image.Point{X: d.X, Y: d.Y}))
	point := qrImage.Bounds().Min.Sub(image.Point{X: d.X, Y: d.Y})
	draw.Draw(ctx.Canvas, ctx.Canvas.Bounds(), qrImage, point, draw.Over)
	return nil
}

// Validate checks if the QRCodeDraw configuration is valid.
func (d *QRCodeDraw) Validate() error {
	if d.Content == "" {
		return errors.New("content is required")
	}
	if d.Size <= 0 {
		return errors.New("size must be positive")
	}
	if d.X < 0 || d.Y < 0 {
		return errors.New("x and y must be non-negative")
	}
	return nil
}
