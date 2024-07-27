package poster

import (
	"github.com/minlib/go-util/imagex"
	"image"
)

type ImageDraw struct {
	X     int
	Y     int
	Path  string
	Round bool
	Width int
	NextStep
}

func (d *ImageDraw) Do(c *Context) error {
	srcImage, err := imagex.ReadImage(d.Path)
	if err != nil {
		return err
	}
	if d.Width > 0 {
		srcImage, err = imagex.ResizeImage(srcImage, d.Width)
		if err != nil {
			return err
		}
	}
	var newImage image.Image
	if d.Round {
		newImage = DrawCircle(srcImage)
	} else {
		newImage = srcImage
	}
	srcPoint := image.Point{
		X: d.X,
		Y: d.Y,
	}
	DrawImage(c.Canvas, newImage, newImage.Bounds().Min.Sub(srcPoint))
	return nil
}
