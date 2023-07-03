package poster

import (
	"github.com/skip2/go-qrcode"
	"image"
)

type QRCodeDraw struct {
	X       int
	Y       int
	Size    int
	Content string
	NextStep
}

func (d *QRCodeDraw) Do(c *Context) error {
	qrImage, err := DrawQRCode(d.Content, qrcode.Medium, d.Size)
	if err != nil {
		return err
	}
	qrPoint := image.Point{
		X: d.X,
		Y: d.Y,
	}
	DrawImage(c.Canvas, qrImage, qrImage.Bounds().Min.Sub(qrPoint))
	return nil
}
