package qrcodex

import (
	"github.com/minlib/go-util/imagex"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"math"
)

// Qrcode 生成无默认边框的基础二维码图像
func Qrcode(content string, size int, level qrcode.RecoveryLevel) (image.Image, error) {
	qrCode, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}
	// 禁用库默认添加的边框（减少冗余空白）
	qrCode.DisableBorder = true
	qrcodeImage := qrCode.Image(size)
	return qrcodeImage, nil
}

// QrcodeWithBorder 生成带自定义白色边框的二维码图像
func QrcodeWithBorder(content string, size int, level qrcode.RecoveryLevel) (image.Image, error) {
	baseImage, err := Qrcode(content, size, level)
	if err != nil {
		return nil, err
	}
	// 计算边框宽度：尺寸的5%，且不小于2像素（符合二维码安静区最小要求）
	borderWidth := int(math.Max(float64(size)*0.05, 2))
	// 为二维码添加白色边框（边框颜色为纯白色，确保扫描识别率）
	qrcodeImage, err := imagex.AddBorder(baseImage, borderWidth, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return nil, err
	}
	return qrcodeImage, nil
}
