package drawer

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
	"image/draw"
	"strings"
)

// TextDraw represents a component for drawing text.
type TextDraw struct {
	// Left is the x-coordinate for drawing the text.
	Left float64 `json:"left"`

	// Right is the x-coordinate for drawing the text.
	Right float64 `json:"right"`

	// Top is the y-coordinate for drawing the text.
	Top float64 `json:"top"`

	// AX is the horizontal anchor point ratio (0=left, 0.5=center, 1=right).
	AX float64 `json:"ax"`

	// AY is the vertical anchor point ratio (0=top, 0.5=center, 1=bottom).
	AY float64 `json:"ay"`

	// Size is the font size for the text.
	Size float64 `json:"size"`

	// Color is the text color in hexadecimal format.
	Color string `json:"color"`

	// Content is the text content to be drawn (supports \n for line breaks).
	Content string `json:"content"`

	// FontFamily is the font family to be used for drawing the text.
	FontFamily string `json:"fontFamily"`

	// Align is the text alignment (AlignLeft, AlignCenter, AlignRight).
	Align TextAlign `json:"align"`

	// LineSpacing is the line spacing multiplier.
	LineSpacing float64 `json:"lineSpacing"`

	// CorrectionX is the Left-axis correction value.
	CorrectionX float64 `json:"correctionX"`

	// CorrectionY is the Top-axis correction value.
	CorrectionY float64 `json:"correctionY"`
}

// Type returns the type identifier for the TextDraw component.
func (d *TextDraw) Type() string {
	return "text"
}

// Draw executes the text drawing logic.
func (d *TextDraw) Draw(ctx *Context) error {
	if d.Content == "" || strings.TrimSpace(d.Content) == "" {
		return nil
	}
	// Set default values
	if d.Size <= 0 {
		d.Size = 24
	}
	if d.LineSpacing <= 0 {
		d.LineSpacing = 1
	}
	// Configure the gg drawing context
	bounds := ctx.Canvas.Bounds()
	ctxGG := gg.NewContext(bounds.Dx(), bounds.Dy())
	font := ctx.Fonts[d.FontFamily]
	fontFace := truetype.NewFace(font, &truetype.Options{Size: d.Size})
	if fontFace == nil {
		return fmt.Errorf("font family '%s' not found", d.FontFamily)
	}
	ctxGG.DrawImage(ctx.Canvas, 0, 0) // Copy existing canvas content
	ctxGG.SetHexColor(d.Color)
	ctxGG.SetFontFace(fontFace)

	// Draw the text
	adjustedX := d.Left + d.CorrectionX
	adjustedY := d.Top + d.CorrectionY

	if d.Align == AlignRight && d.Right > 0 {
		adjustedX = float64(bounds.Dx()) - d.Right
	}

	w, _ := ctxGG.MeasureMultilineString(d.Content, d.LineSpacing)
	ctxGG.DrawStringWrapped(d.Content, adjustedX, adjustedY, d.AX, d.AY, w, d.LineSpacing, gg.Align(AlignMap[string(d.Align)]))

	// 5. Merge back to the original canvas
	draw.Draw(ctx.Canvas, bounds, ctxGG.Image(), image.Point{}, draw.Over)
	return nil
}

// Validate checks if the TextDraw configuration is valid.
func (d *TextDraw) Validate(ctx *Context) error {
	if d.FontFamily == "" {
		return errors.New("fontFamily is required")
	}
	if ctx.Fonts[d.FontFamily] == nil {
		return fmt.Errorf("font family '%s' not found", d.FontFamily)
	}
	if d.AX < 0 || d.AX > 1 || d.AY < 0 || d.AY > 1 {
		return errors.New("ax and ay must be between 0 and 1")
	}
	return nil
}
