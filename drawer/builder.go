package drawer

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
)

// Builder is a helper for constructing posters with a fluent API.
type Builder struct {
	Canvas   draw.Image
	pipeline *Pipeline
}

// NewBuilder creates a new Builder instance with the specified dimensions.
func NewBuilder(canvas draw.Image) *Builder {
	return &Builder{
		Canvas:   canvas,
		pipeline: NewPipeline(),
	}
}

// AddDrawer adds one or more drawing components to the builder.
func (b *Builder) AddDrawer(drawers ...Drawer) *Builder {
	b.pipeline.AddDrawer(drawers...)
	return b
}

// FromJSONConfig loads drawing components from a JSON configuration string.
func (b *Builder) FromJSONConfig(jsonStr string) (*Builder, error) {
	var configs []struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &configs); err != nil {
		return nil, fmt.Errorf("unmarshal json config failed: %w", err)
	}

	for _, cfg := range configs {
		var drawer Drawer
		switch cfg.Type {
		case "image":
			var d ImageDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, fmt.Errorf("unmarshal image drawer failed: %w", err)
			}
			drawer = &d
		case "text":
			var d TextDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, fmt.Errorf("unmarshal text drawer failed: %w", err)
			}
			drawer = &d
		case "qrcode":
			var d QRCodeDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, fmt.Errorf("unmarshal qrcode drawer failed: %w", err)
			}
			drawer = &d
		default:
			return nil, fmt.Errorf("unsupported drawer type: %s", cfg.Type)
		}
		b.AddDrawer(drawer)
	}
	return b, nil
}

// Build executes the drawing pipeline and returns the resulting image.
func (b *Builder) Build() (image.Image, error) {
	ctx := &Context{
		Canvas: b.Canvas,
	}
	if err := b.pipeline.Execute(ctx); err != nil {
		return nil, fmt.Errorf("build failed: %w", err)
	}
	return ctx.Canvas, nil
}
