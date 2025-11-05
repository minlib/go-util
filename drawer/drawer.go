package drawer

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/golang/freetype/truetype"
	"github.com/minlib/go-util/fontx"
	"github.com/minlib/go-util/imagex"
)

// Drawer is the interface that wraps the basic methods for drawing components.
// Each drawing component must implement these methods to be used in the pipeline.
type Drawer interface {
	// Type returns the type identifier of the drawer.
	Type() string

	// Draw executes the drawing logic using the provided context.
	Draw(ctx *Context) error

	// Validate checks if the drawer configuration is valid.
	Validate() error
}

// Context holds the drawing context including canvas and resources.
type Context struct {
	// Canvas is the target image where drawing operations are performed.
	Canvas draw.Image
}

// NewContext creates a new drawing context with the specified canvas.
func NewContext(canvas draw.Image) *Context {
	return &Context{
		Canvas: canvas,
	}
}

// LoadImage loads an image from the specified path (local or remote).
func (r *Context) LoadImage(path string) (image.Image, error) {
	img, err := imagex.ReadImage(path)
	if err != nil {
		return nil, fmt.Errorf("load image failed: %w", err)
	}
	return img, nil
}

// LoadFont loads a font from the specified path.
func (r *Context) LoadFont(path string) (*truetype.Font, error) {
	return fontx.GetFont(path)
}
