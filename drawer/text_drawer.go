package drawer

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

// TextDraw represents a component for drawing text.
type TextDraw struct {
	// X is the x-coordinate for drawing the text.
	X float64 `json:"x"`

	// Y is the y-coordinate for drawing the text.
	Y float64 `json:"y"`

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

	// FontPath is the path to the font file.
	FontPath string `json:"fontPath"`

	// Align is the text alignment (AlignLeft, AlignCenter, AlignRight).
	Align TextAlign `json:"align"`

	// LineSpacing is the line spacing multiplier.
	LineSpacing float64 `json:"lineSpacing"`

	// CorrectionX is the X-axis correction value.
	CorrectionX float64 `json:"correctionX"`

	// CorrectionY is the Y-axis correction value.
	CorrectionY float64 `json:"correctionY"`

	// Font is the loaded font object (internal use).
	Font *truetype.Font `json:"-"`
}

// Type returns the type identifier for the TextDraw component.
func (d *TextDraw) Type() string {
	return "text"
}

// Draw executes the text drawing logic.
func (d *TextDraw) Draw(ctx *Context) error {
	// 1. Set default values
	if d.Size == 0 {
		d.Size = 24
	}
	if d.LineSpacing == 0 {
		d.LineSpacing = 1.2
	}

	// 2. Load the font
	var font *truetype.Font
	var err error
	if d.Font != nil {
		font = d.Font
	} else if d.FontPath != "" {
		font, err = ctx.LoadFont(d.FontPath)
		if err != nil {
			return fmt.Errorf("load font failed: %w", err)
		}
	} else {
		return errors.New("fontPath is required")
	}

	// 3. Configure the gg drawing context
	bounds := ctx.Canvas.Bounds()
	ctxGG := gg.NewContext(bounds.Dx(), bounds.Dy())
	ctxGG.DrawImage(ctx.Canvas, 0, 0) // Copy existing canvas content
	ctxGG.SetHexColor(d.Color)
	ctxGG.SetFontFace(truetype.NewFace(font, &truetype.Options{Size: d.Size}))

	// 4. Draw the text
	adjustedX := d.X + d.CorrectionX
	adjustedY := d.Y + d.CorrectionY
	if strings.Contains(d.Content, "\n") {
		w, _ := ctxGG.MeasureMultilineString(d.Content, d.LineSpacing)
		ctxGG.DrawStringWrapped(d.Content, adjustedX, adjustedY, d.AX, d.AY, w, d.LineSpacing, gg.Align(d.Align))
	} else {
		ctxGG.DrawStringAnchored(d.Content, adjustedX, adjustedY, d.AX, d.AY)
	}

	// 5. Merge back to the original canvas
	draw.Draw(ctx.Canvas, bounds, ctxGG.Image(), image.Point{}, draw.Over)
	return nil
}

// Validate checks if the TextDraw configuration is valid.
func (d *TextDraw) Validate() error {
	if d.Content == "" {
		return errors.New("content is required")
	}
	if d.FontPath == "" && d.Font == nil {
		return errors.New("fontPath or Font is required")
	}
	if d.AX < 0 || d.AX > 1 || d.AY < 0 || d.AY > 1 {
		return errors.New("ax and ay must be between 0 and 1")
	}
	return nil
}
