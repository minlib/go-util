package drawer

import (
	"image/draw"

	"github.com/golang/freetype/truetype"
)

// Drawer is the interface that wraps the basic methods for drawing components.
// Each drawing component must implement these methods to be used in the pipeline.
type Drawer interface {
	// Type returns the type identifier of the drawer.
	Type() string

	// Draw executes the drawing logic using the provided context.
	Draw(ctx *Context) error

	// Validate checks if the drawer configuration is valid.
	Validate(ctx *Context) error
}

// Context holds the drawing context including canvas and resources.
type Context struct {
	// Canvas is the target image where drawing operations are performed.
	Canvas draw.Image
	// Fonts is a map of font families and their corresponding
	Fonts map[string]*truetype.Font
}

// NewContext creates a new drawing context with the specified canvas.
func NewContext(canvas draw.Image, fonts map[string]*truetype.Font) *Context {
	if fonts == nil {
		fonts = make(map[string]*truetype.Font)
	}
	return &Context{
		Canvas: canvas,
		Fonts:  fonts,
	}
}
