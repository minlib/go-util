package webp

import (
	"fmt"
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
	"image/color"
	"strconv"
	"testing"
	"time"
)

var libPath = "D:\\dev\\libwebp-1.3.2-windows-x64\\bin\\"
var original = "D:/temp/original/1234.jpg"
var outputPath = "D:/temp/output/"
var webpUtil = &WebpUtil{
	LibPath: libPath,
}

func TestSetup(t *testing.T) {
	if !filex.Exist(original) {
		_ = filex.MkdirAll(original)
		red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
		imagex.CreateImage(original, 1920, 1080, red)
	}
}

func TestRunCommand(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	option := &CommandOption{
		LibPath:  libPath,
		Quality:  80,
		Crop:     &Crop{X: 0, Y: 0, W: 1000, H: 1000},
		Resize:   &Resize{Width: 750, Height: 750},
		Original: original,
		Output:   output,
	}
	err := CWebp(option)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestScaleToFill(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.ScaleToFill(original, output, 600, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestScaleToFill2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.ScaleToFill(original, output, 1920, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestAspectFill(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.AspectFill(original, output, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestAspectFill2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.AspectFill(original, output, 300, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	webpWidth, webpHeight, err := GetWebpSize(output)
	if 300 != webpWidth || 300 != webpHeight {
		t.Errorf("width expected: %d, actual: %d, height expected: %d, actual: %d", 300, webpWidth, 300, webpHeight)
	}
}

func TestWidthFix(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.WidthFix(original, output, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	webpWidth, _, err := GetWebpSize(output)
	if 400 != webpWidth {
		t.Errorf("width expected: %d, actual: %d", 400, webpWidth)
	}
}

func TestWidthFix2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	imgWidth, _, err := imagex.GetSize(original)
	err = webpUtil.WidthFix(original, output, imgWidth+500, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	webpWidth, _, err := GetWebpSize(output)
	if imgWidth != webpWidth {
		t.Errorf("width expected: %d, actual: %d", imgWidth, webpWidth)
	}
}

func TestHeightFix(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.HeightFix(original, output, 800, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	_, webpHeight, err := GetWebpSize(output)
	if 800 != webpHeight {
		t.Errorf("height expected: %d, actual: %d", 800, webpHeight)
	}
}

func TestHeightFix2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.HeightFix(original, output, 1200, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	_, imgHeight, err := imagex.GetSize(output)
	_, webpHeight, err := GetWebpSize(output)

	if imgHeight != webpHeight {
		t.Errorf("height expected: %d, actual: %d", imgHeight, webpHeight)
	}
}
