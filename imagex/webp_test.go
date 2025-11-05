package imagex

import (
	"fmt"
	"image/color"
	"os"
	"testing"
	"time"

	"github.com/minlib/go-util/filex"
)

var libPath = "D:/dev/libwebp-1.3.2-windows-x64/bin/"
var original = "../outputs/original.jpg"
var webpUtil = &WebpUtil{
	LibPath: libPath,
}

// setup is a common method to be executed before each test
func setup() {
	// Add any initialization logic here
	if !filex.Exist(original) {
		img := NewImage(1920, 1080, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		_ = SavePNG(img, original)
	}
}

// TestMain is the entry point for running tests
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestRunCommand(t *testing.T) {
	outputPath := getOutputPath()
	option := &CommandOption{
		LibPath:  libPath,
		Quality:  80,
		Crop:     &Crop{X: 0, Y: 0, W: 1000, H: 1000},
		Resize:   &Resize{Width: 750, Height: 750},
		Original: original,
		Output:   outputPath,
	}
	err := CWebpCommand(option)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCWebpByScaleToFill(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.CWebp(original, outputPath, ScaleToFill, 0, 540, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestCWebpByAspectFill(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.CWebp(original, outputPath, AspectFill, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestCWebpByWidthFix(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.CWebp(original, outputPath, WidthFix, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestCWebpByHeightFix(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.CWebp(original, outputPath, HeightFix, 500, 0, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestScaleToFill(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.ScaleToFill(original, outputPath, 600, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestScaleToFill2(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.ScaleToFill(original, outputPath, 1920, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestAspectFill(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.AspectFill(original, outputPath, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
}

func TestAspectFill2(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.AspectFill(original, outputPath, 300, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	outputWidth, outputHeight, err := GetSize(outputPath)
	if 300 != outputWidth || 300 != outputHeight {
		t.Errorf("width expected: %d, actual: %d, height expected: %d, actual: %d", 300, outputWidth, 300, outputHeight)
	}
}

func TestWidthFix(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.WidthFix(original, outputPath, 400, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	outputWidth, _, err := GetSize(outputPath)
	if 400 != outputWidth {
		t.Errorf("width expected: %d, actual: %d", 400, outputWidth)
	}
}

func TestWidthFix2(t *testing.T) {
	outputPath := getOutputPath()
	originalWidth, _, err := GetSize(original)
	width := originalWidth + 500
	err = webpUtil.WidthFix(original, outputPath, width, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	outputWidth, _, err := GetSize(outputPath)
	if width != outputWidth {
		t.Errorf("width expected: %d, actual: %d", width, outputWidth)
	}
}

func TestWidthLessen(t *testing.T) {
	outputPath := getOutputPath()
	originalWidth, _, err := GetSize(original)
	err = webpUtil.WidthLessen(original, outputPath, originalWidth+500, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	outputWidth, _, err := GetSize(outputPath)
	if originalWidth != outputWidth {
		t.Errorf("width expected: %d, actual: %d", originalWidth, outputWidth)
	}
}

func TestHeightFix(t *testing.T) {
	outputPath := getOutputPath()
	err := webpUtil.HeightFix(original, outputPath, 800, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	_, outputHeight, err := GetSize(outputPath)
	if 800 != outputHeight {
		t.Errorf("height expected: %d, actual: %d", 800, outputHeight)
	}
}

func TestHeightFix2(t *testing.T) {
	outputPath := getOutputPath()
	_, originalHeight, _ := GetSize(original)
	height := originalHeight + 500
	err := webpUtil.HeightFix(original, outputPath, height, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	_, outputHeight, err := GetSize(outputPath)
	if height != outputHeight {
		t.Errorf("height expected: %d, actual: %d", height, outputHeight)
	}
}

func TestHeightLessen(t *testing.T) {
	outputPath := getOutputPath()
	_, originalHeight, _ := GetSize(original)
	err := webpUtil.HeightLessen(original, outputPath, originalHeight+500, 85)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputPath)
	_, outputHeight, err := GetSize(outputPath)
	if originalHeight != outputHeight {
		t.Errorf("height expected: %d, actual: %d", originalHeight, outputHeight)
	}
}

func getOutputPath() string {
	return fmt.Sprintf("%s%d.webp", "../outputs/", time.Now().UnixNano())
}
