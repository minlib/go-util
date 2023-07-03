package poster

import (
	"fmt"
	"github.com/minlib/go-util/filex"
	"testing"
	"time"
)

func TestCircle(t *testing.T) {
	srcImage, err := ReadImage("./assets/template.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	newImage := DrawCircle(srcImage)
	if err = SavePNG(newImage, getOutputPath()); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func TestPoster(t *testing.T) {
	start := &StartDraw{}
	ctx := &Context{
		Canvas: NewRGBA(0, 0, 750, 1334),
	}
	// 绘制背景图
	backgroundDraw := &ImageDraw{
		X:    0,
		Y:    0,
		Path: "./assets/template.png",
	}
	// 绘制二维码
	qrCodeDraw := &QRCodeDraw{
		X:       30,
		Y:       860,
		Size:    250,
		Content: "http://www.minzhan.com",
	}
	// 绘制本地头像
	avatarDraw2 := &ImageDraw{
		X:     30,
		Y:     250,
		Path:  "./assets/avatar.jpg",
		Round: true,
	}
	// 绘制远程头像
	avatarDraw1 := &ImageDraw{
		X:     30,
		Y:     50,
		Path:  "https://minzhan.net/uploads/image/avatar.png",
		Round: true,
	}
	// 绘制文字
	textDraw1 := &TextDraw{
		X:        180,
		Y:        105,
		Size:     26,
		Color:    "#FFFFFF",
		Content:  "这里是大标题1",
		FontPath: "./assets/msyh.ttf",
	}
	// 绘制文字
	textDraw2 := &TextDraw{
		X:        180,
		Y:        150,
		Size:     20,
		Color:    "#FFFFFF",
		Content:  "这里是小标题2",
		FontPath: "./assets/msyh.ttf",
	}

	// 完成绘制，并导出图片
	finishDraw := &FinishDraw{
		Output: getOutputPath(),
	}

	start.
		SetNext(backgroundDraw).
		SetNext(avatarDraw1).
		SetNext(avatarDraw2).
		SetNext(textDraw1).
		SetNext(textDraw2).
		SetNext(qrCodeDraw).
		SetNext(finishDraw)
	if err := start.Run(ctx); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func getOutputPath() string {
	fmt.Println(time.Now().UnixMicro())
	outputPath := "d:/output/" + time.UnixMilli(time.Now().UnixMilli()).Format("20060102150405.000") + ".png"
	filex.MkdirAll(outputPath)
	return outputPath
}
