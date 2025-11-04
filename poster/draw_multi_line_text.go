package poster

import (
	"errors"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
	"image/draw"
	"strings"
)

// TextAlign 文本对齐方式
type TextAlign int

const (
	AlignLeft   TextAlign = iota // 左对齐
	AlignCenter                  // 居中对齐
	AlignRight                   // 右对齐
)

const (
	FlexStart  = 0.0
	FlexCenter = 0.5
	FlexEnd    = 1.0
)

// MultiLineTextDraw 多行文本绘制组件
type MultiLineTextDraw struct {
	X           float64        `json:"x"`           // 横坐标
	Y           float64        `json:"y"`           // 纵坐标
	AX          float64        `json:"ax"`          // 水平锚点比例（0左，0.5中，1右）
	AY          float64        `json:"ay"`          // 垂直锚点比例（0上，0.5中，1下）
	Size        float64        `json:"size"`        // 字体大小
	Color       string         `json:"color"`       // 颜色（十六进制）
	Content     string         `json:"content"`     // 文本内容（支持\n换行）
	FontPath    string         `json:"fontPath"`    // 字体路径
	Align       TextAlign      `json:"align"`       // 对齐方式
	LineSpacing float64        `json:"lineSpacing"` // 行间距倍数
	CorrectionX float64        `json:"correctionX"` // X轴修正值
	CorrectionY float64        `json:"correctionY"` // Y轴修正值
	Font        *truetype.Font `json:"-"`           // 字体对象（内部使用）
}

// Type 实现Drawer接口
func (d *MultiLineTextDraw) Type() string {
	return "multi_line_text"
}

// Draw 实现Drawer接口
func (d *MultiLineTextDraw) Draw(ctx *Context) error {
	// 1. 设置默认值
	if d.Size == 0 {
		d.Size = 24
	}
	if d.LineSpacing == 0 {
		d.LineSpacing = 1.2
	}

	// 2. 加载字体
	var font *truetype.Font
	var err error
	if d.Font != nil {
		font = d.Font
	} else if d.FontPath != "" {
		font, err = ctx.Resources.LoadFont(d.FontPath)
		if err != nil {
			return errors.New("load font failed: " + err.Error())
		}
	} else {
		return errors.New("fontPath is required")
	}

	// 3. 配置gg绘制上下文
	bounds := ctx.Canvas.Bounds()
	ctxGG := gg.NewContext(bounds.Dx(), bounds.Dy())
	ctxGG.DrawImage(ctx.Canvas, 0, 0) // 复制现有画布内容
	ctxGG.SetHexColor(d.Color)
	ctxGG.SetFontFace(truetype.NewFace(font, &truetype.Options{Size: d.Size}))

	// 4. 绘制文本
	adjustedX := d.X + d.CorrectionX
	adjustedY := d.Y + d.CorrectionY
	if strings.Contains(d.Content, "\n") {
		w, _ := ctxGG.MeasureMultilineString(d.Content, d.LineSpacing)
		ctxGG.DrawStringWrapped(d.Content, adjustedX, adjustedY, d.AX, d.AY, w, d.LineSpacing, gg.Align(d.Align))
	} else {
		ctxGG.DrawStringAnchored(d.Content, adjustedX, adjustedY, d.AX, d.AY)
	}

	// 5. 合并回原始画布
	draw.Draw(ctx.Canvas, bounds, ctxGG.Image(), image.Point{}, draw.Over)
	return nil
}

// Validate 实现Drawer接口
func (d *MultiLineTextDraw) Validate() error {
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
