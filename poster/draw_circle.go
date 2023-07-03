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
