package qrcodex

import (
	"fmt"
	"github.com/minlib/go-util/imagex"
	"github.com/skip2/go-qrcode"
	"log"
	"testing"
	"time"
)

func TestQrcode(t *testing.T) {
	outputPath := getOutputPath()
	qrcodeImage, err := Qrcode("https://www.minzhan.com", qrcode.High, 128)
	if err != nil {
		t.Fatalf("生成二维码失败: %v", err)
	}
	if err = imagex.SavePNG(qrcodeImage, outputPath); err != nil {
		t.Fatalf("保存二维码失败: %v", err)
	}
	log.Printf("已生成二维码: %s", outputPath)
}

func TestQrcodeWithBorder(t *testing.T) {
	outputPath := getOutputPath()
	qrcodeImage, err := QrcodeWithBorder("https://www.minzhan.com", qrcode.High, 128)
	if err != nil {
		t.Fatalf("生成二维码失败: %v", err)
	}
	if err = imagex.SavePNG(qrcodeImage, outputPath); err != nil {
		t.Fatalf("保存二维码失败: %v", err)
	}
	log.Printf("已生成带边框的二维码: %s", outputPath)
}

func getOutputPath() string {
	return fmt.Sprintf("%s%d.png", "../outputs/", time.Now().UnixNano())
}
