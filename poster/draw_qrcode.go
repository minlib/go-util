package poster

import (
	"errors"
	"github.com/minlib/go-util/qrcodex"
	"github.com/skip2/go-qrcode"
	"image"
)

// QRCodeDraw 二维码绘制组件
type QRCodeDraw struct {
	X       int    `json:"x"`       // 横坐标
	Y       int    `json:"y"`       // 纵坐标
	Size    int    `json:"size"`    // 二维码尺寸
	Content string `json:"content"` // 二维码内容（URL或文本）
}

// Type 实现Drawer接口
func (d *QRCodeDraw) Type() string {
	return "qrcode"
}

// Draw 实现Drawer接口
func (d *QRCodeDraw) Draw(ctx *Context) error {
	// 生成二维码图片
	qrImage, err := qrcodex.QrcodeWithBorder(d.Content, qrcode.Medium, d.Size)
	if err != nil {
		return errors.New("generate qrcode failed: " + err.Error())
	}
	// 绘制到画布
	DrawImage(ctx.Canvas, qrImage, qrImage.Bounds().Min.Sub(image.Point{X: d.X, Y: d.Y}))
	return nil
}

// Validate 实现Drawer接口
func (d *QRCodeDraw) Validate() error {
	if d.Content == "" {
		return errors.New("content is required")
	}
	if d.Size <= 0 {
		return errors.New("size must be positive")
	}
	if d.X < 0 || d.Y < 0 {
		return errors.New("x and y must be non-negative")
	}
	return nil
}

//// DrawQRCode draw qrcode image
//func DrawQRCode(content string, level qrcode.RecoveryLevel, size int) (image.Image, error) {
//	qr, err := qrcode.New(content, level)
//	if err != nil {
//		return nil, err
//	}
//	qrImage := qr.Image(size)
//	return qrImage, nil
//}
