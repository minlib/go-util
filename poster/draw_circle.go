package poster

import (
	"image"
	"image/color"
	"math"
)

type CircleDraw struct {
	image    image.Image
	point    image.Point
	diameter int
}

func NewCircleDraw(img image.Image, p image.Point, d int) CircleDraw {
	return CircleDraw{img, p, d}
}

func (d CircleDraw) ColorModel() color.Model {
	return d.image.ColorModel()
}

func (d CircleDraw) Bounds() image.Rectangle {
	return image.Rect(0, 0, d.diameter, d.diameter)
}

func (d CircleDraw) At(x, y int) color.Color {
	diameter := d.diameter
	sqrt := math.Sqrt(math.Pow(float64(x-diameter/2), 2) + math.Pow(float64(y-diameter/2), 2))
	if sqrt > float64(diameter)/2 {
		return d.image.ColorModel().Convert(color.RGBA{R: 255, G: 255, B: 255, A: 0})
	} else {
		return d.image.At(d.point.X+x, d.point.Y+y)
	}
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
