package poster

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/minlib/go-util/colorx"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"os"
	"strings"
)

// Text 文字
type Text struct {
	Content  string  // 内容
	FontPath string  // 字体路径
	Color    string  // 颜色
	Size     float64 // 大小
	X        int     // 横坐标
	Y        int     // 纵坐标
}

// MultiLineText 多行文字
type MultiLineText struct {
	Content     string    // 内容
	FontPath    string    // 字体路径
	Color       string    // 颜色
	Size        float64   // 大小
	X           float64   // 横坐标
	Y           float64   // 纵坐标
	AX          float64   // 水平锚点比例（0左，0.5中，1右）
	AY          float64   // 垂直锚点比例（0上，0.5中，1下）
	LineSpacing float64   // 行间距
	Align       TextAlign // 对齐方式
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

// NewRGBA returns a new RGBA image with the given bounds.
func NewRGBA(x0, y0, x1, y1 int) *image.RGBA {
	return image.NewRGBA(image.Rect(x0, y0, x1, y1))
}

// GetFont Load and parse a truetype font from the specified file path
func GetFont(path string) (*truetype.Font, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fileBytes)
	if err != nil {
		return nil, err
	}
	return font, nil
}

// GetFontFace Create a font.Face from a truetype.Font with the specified point size
func GetFontFace(f *truetype.Font, points float64) (font.Face, error) {
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
	})
	return face, nil
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
	if err := ctx.LoadFontFace(text.FontPath, text.Size); err != nil {
		panic(err)
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

// DrawImage draw image
func DrawImage(dst draw.Image, src image.Image, sp image.Point) {
	draw.Draw(dst, dst.Bounds(), src, sp, draw.Over)
}

// DrawCircle draw circle image
func DrawCircle(canvas image.Image) image.Image {
	// 获取图片的宽和高
	srcBounds := canvas.Bounds()
	width := srcBounds.Max.X - srcBounds.Min.X
	height := srcBounds.Max.Y - srcBounds.Min.Y

	// 解决白底问题
	dstImage := NewRGBA(0, 0, width, height)
	DrawImage(dstImage, canvas, srcBounds.Min)

	// 获取圆的直径
	var diameter int
	if width > height {
		diameter = height
	} else {
		diameter = width
	}
	return NewCircleDraw(dstImage, image.Point{X: 0, Y: 0}, diameter)
}

// DrawQRCode draw qrcode image
func DrawQRCode(content string, level qrcode.RecoveryLevel, size int) (image.Image, error) {
	qr, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}
	qrImage := qr.Image(size)
	return qrImage, nil
}
