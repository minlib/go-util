package poster

import (
	"github.com/golang/freetype"
	"github.com/minlib/go-util/colorx"
	"image"
	"image/color"
	"image/draw"
)

// Text 文字
type Text struct {
	X        int     // 横坐标
	Y        int     // 纵坐标
	Size     float64 // 大小
	Color    string  // 颜色
	Content  string  // 内容
	FontPath string  // 字体路径
}

type TextDraw struct {
	X        int     // 横坐标
	Y        int     // 纵坐标
	Size     float64 // 大小
	Color    string  // 颜色
	Content  string  // 内容
	FontPath string  // 字体路径
	NextStep
}

func (d *TextDraw) Do(c *Context) error {
	if d.Size == 0 {
		d.Size = 24
	}
	text := &Text{
		X:        d.X,
		Y:        d.Y,
		Size:     d.Size,
		Color:    d.Color,
		Content:  d.Content,
		FontPath: d.FontPath,
	}
	err := DrawText(c.Canvas, text)
	if err != nil {
		return err
	}
	return nil
}

// DrawText draw text
func DrawText(canvas draw.Image, text *Text) error {
	fontType, err := GetFont(text.FontPath)
	if err != nil {
		return err
	}
	ctx := freetype.NewContext()
	//设置屏幕每英寸的分辨率
	ctx.SetDPI(72)
	//设置用于绘制文本的字体
	ctx.SetFont(fontType)
	//以磅为单位设置字体大小
	ctx.SetFontSize(text.Size)
	//设置剪裁矩形以进行绘制
	ctx.SetClip(canvas.Bounds())
	//设置目标图像
	ctx.SetDst(canvas)
	r, g, b := colorx.Hex2RGB(text.Color)
	//设置绘制操作的源图像，通常为 image.Uniform
	ctx.SetSrc(image.NewUniform(color.RGBA{R: r, G: g, B: b, A: 255}))
	pt := freetype.Pt(text.X, text.Y)
	_, err = ctx.DrawString(text.Content, pt)
	if err != nil {
		return err
	}
	return nil
}
