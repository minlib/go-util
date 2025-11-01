package qrcodex

import (
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
	"github.com/skip2/go-qrcode"
	"log"
	"testing"
	"time"
)

func TestQrcodeWithBorder(t *testing.T) {
	outputPath := "d:/output/" + time.Now().Format("20060102150405") + ".png"
	_ = filex.MkdirAll(outputPath)

	qrcodeImage, err := QrcodeWithBorder("https://www.minzhan.com", 200, qrcode.High)
	if err != nil {
		t.Fatalf("生成二维码失败: %v", err)
	}

	if err = imagex.SavePNG(qrcodeImage, outputPath); err != nil {
		t.Fatalf("保存二维码失败: %v", err)
	}

	log.Println("已生成带边框的二维码")
	log.Println(outputPath)
}
