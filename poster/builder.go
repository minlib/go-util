package poster

import (
	"encoding/json"
	"errors"
	"image"
)

// Builder 海报构建器
type Builder struct {
	width     int
	height    int
	resources *ResourceLoader
	pipeline  *Pipeline
}

// NewPosterBuilder 创建构建器
func NewPosterBuilder(width, height int) *Builder {
	return &Builder{
		width:     width,
		height:    height,
		resources: &ResourceLoader{},
		pipeline:  NewPipeline(),
	}
}

// AddDrawer 添加绘制组件
func (b *Builder) AddDrawer(drawers ...Drawer) *Builder {
	b.pipeline.AddDrawer(drawers...)
	return b
}

// FromJSONConfig 从JSON配置加载组件
func (b *Builder) FromJSONConfig(jsonStr string) (*Builder, error) {
	var configs []struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &configs); err != nil {
		return nil, err
	}

	for _, cfg := range configs {
		var drawer Drawer
		switch cfg.Type {
		case "image":
			var d ImageDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, err
			}
			drawer = &d
		case "text":
			var d TextDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, err
			}
			drawer = &d
		case "multi_line_text":
			var d MultiLineTextDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, err
			}
			drawer = &d
		case "qrcode":
			var d QRCodeDraw
			if err := json.Unmarshal(cfg.Data, &d); err != nil {
				return nil, err
			}
			drawer = &d
		default:
			return nil, errors.New("unsupported drawer type: " + cfg.Type)
		}
		b.AddDrawer(drawer)
	}
	return b, nil
}

// Build 执行绘制并返回结果
func (b *Builder) Build() (image.Image, error) {
	ctx := &Context{
		Canvas:    NewRGBA(0, 0, b.width, b.height),
		Resources: b.resources,
	}

	if err := b.pipeline.Execute(ctx); err != nil {
		return nil, err
	}
	return ctx.Canvas, nil
}
