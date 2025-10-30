package poster

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"image"
	"image/draw"
	"strings"
)

// MultiLineText 多行文字
type MultiLineText struct {
	X           float64        // 横坐标
	Y           float64        // 纵坐标
	AX          float64        // 水平锚点比例（0左，0.5中，1右）
	AY          float64        // 垂直锚点比例（0上，0.5中，1下）
	Size        float64        // 大小
	Color       string         // 颜色
	Content     string         // 内容
	FontPath    string         // 字体路径
	Font        *truetype.Font // 字体对象
	Align       TextAlign      // 对齐方式
	LineSpacing float64        // 行间距
}

type MultiLineTextDraw struct {
	X           float64        // 横坐标
	Y           float64        // 纵坐标
	AX          float64        // 水平锚点比例（0左，0.5中，1右）
	AY          float64        // 垂直锚点比例（0上，0.5中，1下）
	Size        float64        // 大小
	Color       string         // 颜色
	Content     string         // 内容
	FontPath    string         // 字体路径
	Font        *truetype.Font // 字体对象
	Align       TextAlign      // 对齐方式
	LineSpacing float64        // 行间距
	CorrectionX float64        // 横坐标修正值
	CorrectionY float64        // 纵坐标修正值
	NextStep
}

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

func (d *MultiLineTextDraw) Do(c *Context) error {
	if d.Size == 0 {
		d.Size = 24
	}
	multiLineText := &MultiLineText{
		X:           d.X + d.CorrectionX,
		Y:           d.Y + d.CorrectionY,
		AX:          d.AX,
		AY:          d.AY,
		Size:        d.Size,
		Color:       d.Color,
		Content:     d.Content,
		FontPath:    d.FontPath,
		Font:        d.Font,
		Align:       d.Align,
		LineSpacing: d.LineSpacing,
	}
	err := DrawMultiLineText(c.Canvas, multiLineText)
	if err != nil {
		return err
	}
	return nil
}

// DrawMultiLineText draw multi Line text
func DrawMultiLineText(canvas draw.Image, text *MultiLineText) error {
	// 获取图片的宽和高
	srcBounds := canvas.Bounds()
	width := srcBounds.Dx()
	height := srcBounds.Dy()
	ctx := gg.NewContext(width, height)
	// 绘制背景图片
	ctx.DrawImage(canvas, 0, 0)
	ctx.SetHexColor(text.Color)
	// 加载字体
	if text.Font != nil {
		fontFace := GetFontFace(text.Font, text.Size)
		ctx.SetFontFace(fontFace)
	} else if text.FontPath != "" {
		fontFace, err := GetFontAndFace(text.FontPath, text.Size)
		if err != nil {
			return err
		}
		ctx.SetFontFace(fontFace)
		//if err := ctx.LoadFontFace(text.FontPath, text.Size); err != nil {
		//	return err
		//}
	}
	if strings.Contains(text.Content, "\n") {
		w, _ := ctx.MeasureMultilineString(text.Content, text.LineSpacing)
		ctx.DrawStringWrapped(text.Content, text.X, text.Y, text.AX, text.AY, w, text.LineSpacing, gg.Align(text.Align))
	} else {
		ctx.DrawStringAnchored(text.Content, text.X, text.Y, text.AX, text.AY)
	}
	// 将绘制结果合并回原始画布
	draw.Draw(canvas, srcBounds, ctx.Image(), image.Point{}, draw.Over)
	return nil
}
