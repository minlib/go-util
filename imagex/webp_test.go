package imagex

import (
	"fmt"
	"github.com/minlib/go-util/filex"
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
		CreateImage(original, 1920, 1080, red)
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
	err := CWebpCommand(option)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCWebpByScaleToFill(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.CWebp(original, output, ScaleToFill, 0, 540, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestCWebpByAspectFill(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.CWebp(original, output, AspectFill, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestCWebpByWidthFix(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.CWebp(original, output, WidthFix, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func TestCWebpByHeightFix(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	err := webpUtil.CWebp(original, output, HeightFix, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
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
	outputWidth, outputHeight, err := GetSize(output)
	if 300 != outputWidth || 300 != outputHeight {
		t.Errorf("width expected: %d, actual: %d, height expected: %d, actual: %d", 300, outputWidth, 300, outputHeight)
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
	outputWidth, _, err := GetSize(output)
	if 400 != outputWidth {
		t.Errorf("width expected: %d, actual: %d", 400, outputWidth)
	}
}

func TestWidthFix2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	originalWidth, _, err := GetSize(original)
	width := originalWidth + 500
	err = webpUtil.WidthFix(original, output, width, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	outputWidth, _, err := GetSize(output)
	if width != outputWidth {
		t.Errorf("width expected: %d, actual: %d", width, outputWidth)
	}
}

func TestWidthLessen(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	originalWidth, _, err := GetSize(original)
	err = webpUtil.WidthLessen(original, output, originalWidth+500, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	outputWidth, _, err := GetSize(output)
	if originalWidth != outputWidth {
		t.Errorf("width expected: %d, actual: %d", originalWidth, outputWidth)
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
	_, outputHeight, err := GetSize(output)
	if 800 != outputHeight {
		t.Errorf("height expected: %d, actual: %d", 800, outputHeight)
	}
}

func TestHeightFix2(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	_, originalHeight, _ := GetSize(original)
	height := originalHeight + 500
	err := webpUtil.HeightFix(original, output, height, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	_, outputHeight, err := GetSize(output)
	if height != outputHeight {
		t.Errorf("height expected: %d, actual: %d", height, outputHeight)
	}
}

func TestHeightLessen(t *testing.T) {
	output := outputPath + strconv.Itoa(int(time.Now().UnixNano())) + ".webp"
	_, originalHeight, _ := GetSize(original)
	err := webpUtil.HeightLessen(original, output, originalHeight+500, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
	_, outputHeight, err := GetSize(output)
	if originalHeight != outputHeight {
		t.Errorf("height expected: %d, actual: %d", originalHeight, outputHeight)
	}
}
