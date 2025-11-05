package drawer

import (
	"fmt"
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
)

// Pipeline manages the execution of drawing components in sequence.
type Pipeline struct {
	// drawers is the ordered list of drawing components.
	drawers []Drawer

	// output is the path where the final image will be saved.
	output string
}

// NewPipeline creates a new Pipeline instance with optional initial drawers.
func NewPipeline(drawers ...Drawer) *Pipeline {
	return &Pipeline{drawers: drawers}
}

// AddDrawer adds one or more drawing components to the pipeline.
func (p *Pipeline) AddDrawer(drawers ...Drawer) *Pipeline {
	p.drawers = append(p.drawers, drawers...)
	return p
}

// SetOutput sets the output path for saving the final image.
func (p *Pipeline) SetOutput(output string) *Pipeline {
	p.output = output
	return p
}

// Execute runs the drawing pipeline, executing all components in order.
func (p *Pipeline) Execute(ctx *Context) error {
	// 1. Validate all components
	if err := p.validate(); err != nil {
		return fmt.Errorf("pipeline validation failed: %w", err)
	}

	// 2. Execute drawing components in order
	for i, drawer := range p.drawers {
		if err := drawer.Draw(ctx); err != nil {
			return fmt.Errorf("drawer[%s] failed at index %d: %w", drawer.Type(), i, err)
		}
	}

	// 3. Save the image if output path is specified
	if p.output != "" {
		if err := filex.MkdirAll(p.output); err != nil {
			return fmt.Errorf("create directory for output failed: %w", err)
		}
		if err := imagex.SavePNG(ctx.Canvas, p.output); err != nil {
			return fmt.Errorf("save image to %s failed: %w", p.output, err)
		}
	}

	return nil
}

// validate checks if all drawing components in the pipeline are valid.
func (p *Pipeline) validate() error {
	for _, drawer := range p.drawers {
		if err := drawer.Validate(); err != nil {
			return fmt.Errorf("invalid drawer %s: %w", drawer.Type(), err)
		}
	}
	return nil
}
