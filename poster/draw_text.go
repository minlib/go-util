package poster

import (
	"errors"
	"github.com/golang/freetype"
	"github.com/minlib/go-util/colorx"
	"image"
	"image/color"
)

// TextDraw 文本绘制组件
type TextDraw struct {
	X        int     `json:"x"`        // 横坐标
	Y        int     `json:"y"`        // 纵坐标
	Size     float64 `json:"size"`     // 字体大小
	Color    string  `json:"color"`    // 颜色（十六进制）
	Content  string  `json:"content"`  // 文本内容
	FontPath string  `json:"fontPath"` // 字体路径
}

// Type 实现Drawer接口
func (d *TextDraw) Type() string {
	return "text"
}

// Draw 实现Drawer接口
func (d *TextDraw) Draw(ctx *Context) error {
	// 1. 加载字体
	font, err := ctx.Resources.LoadFont(d.FontPath)
	if err != nil {
		return errors.New("load font failed: " + err.Error())
	}

	// 2. 配置绘制上下文
	ftCtx := freetype.NewContext()
	ftCtx.SetDPI(72)
	ftCtx.SetFont(font)
	ftCtx.SetFontSize(d.Size)
	ftCtx.SetClip(ctx.Canvas.Bounds())
	ftCtx.SetDst(ctx.Canvas)

	// 3. 设置文本颜色
	r, g, b := colorx.Hex2RGB(d.Color)
	ftCtx.SetSrc(image.NewUniform(color.RGBA{R: r, G: g, B: b, A: 255}))

	// 4. 绘制文本
	_, err = ftCtx.DrawString(d.Content, freetype.Pt(d.X, d.Y))
	if err != nil {
		return errors.New("draw text failed: " + err.Error())
	}
	return nil
}

// Validate 实现Drawer接口
func (d *TextDraw) Validate() error {
	if d.Content == "" {
		return errors.New("content is required")
	}
	if d.FontPath == "" {
		return errors.New("fontPath is required")
	}
	if d.Size <= 0 {
		return errors.New("size must be positive")
	}
	return nil
}
