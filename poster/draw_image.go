package poster

import (
	"errors"
	"github.com/minlib/go-util/imagex"
	"image"
)

// ImageDraw 图片绘制组件
type ImageDraw struct {
	X     int    `json:"x"`     // 横坐标
	Y     int    `json:"y"`     // 纵坐标
	Path  string `json:"path"`  // 图片路径（本地/远程）
	Round bool   `json:"round"` // 是否圆形裁剪
	Width int    `json:"width"` // 宽度（0表示原始尺寸）
}

// Type 实现Drawer接口
func (d *ImageDraw) Type() string {
	return "image"
}

// Draw 实现Drawer接口
func (d *ImageDraw) Draw(ctx *Context) error {
	// 1. 加载图片资源
	srcImage, err := ctx.Resources.LoadImage(d.Path)
	if err != nil {
		return errors.New("load image failed: " + err.Error())
	}

	// 2. 处理图片缩放
	if d.Width > 0 {
		srcImage, err = imagex.ResizeImage(srcImage, d.Width)
		if err != nil {
			return errors.New("resize image failed: " + err.Error())
		}
	}

	// 3. 处理圆形裁剪
	var drawImage image.Image
	if d.Round {
		drawImage = DrawCircle(srcImage) // 复用原有DrawCircle实现
	} else {
		drawImage = srcImage
	}

	// 4. 绘制到画布
	DrawImage(ctx.Canvas, drawImage, image.Point{d.X, d.Y}) // 复用util.go中的DrawImage
	return nil
}

// Validate 实现Drawer接口
func (d *ImageDraw) Validate() error {
	if d.Path == "" {
		return errors.New("path is required")
	}
	if d.X < 0 || d.Y < 0 {
		return errors.New("x and y must be non - negative")
	}
	return nil
}
