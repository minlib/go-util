package poster

import (
	"bytes"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/minlib/go-util/colorx"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

// Text 文字
type Text struct {
	Canvas   draw.Image
	FontPath string  //字体路径
	Content  string  //内容
	Color    string  //颜色
	Size     float64 //大小
	X        int     //横坐标
	Y        int     //纵坐标
}

// NewRGBA returns a new RGBA image with the given bounds.
func NewRGBA(x0, y0, x1, y1 int) *image.RGBA {
	return image.NewRGBA(image.Rect(x0, y0, x1, y1))
}

// GetFont get font
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

// GetResourceReader get local or remote resource file
func GetResourceReader(pathOrUrl string) (*bytes.Reader, error) {
	if strings.HasPrefix(pathOrUrl, "http") {
		resp, err := http.Get(pathOrUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fileBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	} else {
		fileBytes, err := os.ReadFile(pathOrUrl)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	}
}

// ReadImage read a image
func ReadImage(pathOrUrl string) (image.Image, error) {
	reader, err := GetResourceReader(pathOrUrl)
	if err != nil {
		return nil, err
	}
	// 加载JPG格式图片时，必须引入依赖 "image/jpeg"，否则会出错 unknown format
	srcImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return srcImage, err
}

// SavePNG save as png image
func SavePNG(src image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, src)
}

// SaveJPG save as jpg image
func SaveJPG(src image.Image, outputPath string, quality int) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, src, &jpeg.Options{
		Quality: quality,
	})
}

// DrawImage draw image
func DrawImage(dst draw.Image, src image.Image, sp image.Point) {
	draw.Draw(dst, dst.Bounds(), src, sp, draw.Over)
}

// DrawCircle draw circle image
func DrawCircle(srcImage image.Image) image.Image {
	// 获取图片的宽和高
	srcBounds := srcImage.Bounds()
	width := srcBounds.Max.X - srcBounds.Min.X
	height := srcBounds.Max.Y - srcBounds.Min.Y

	// 解决白底问题
	dstImage := NewRGBA(0, 0, width, height)
	DrawImage(dstImage, srcImage, srcBounds.Min)

	// 获取圆的直径
	var diameter int
	if width > height {
		diameter = height
	} else {
		diameter = width
	}
	return NewCircleDraw(dstImage, image.Point{X: 0, Y: 0}, diameter)
}

// DrawText draw text
func DrawText(text *Text) error {
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
	ctx.SetClip(text.Canvas.Bounds())
	//设置目标图像
	ctx.SetDst(text.Canvas)
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

// DrawQRCode draw qrcode image
func DrawQRCode(content string, level qrcode.RecoveryLevel, size int) (image.Image, error) {
	qr, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}
	qrImage := qr.Image(size)
	return qrImage, nil
}
