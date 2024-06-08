package webp

import (
	"fmt"
	"github.com/minlib/go-terminal/cmd"
	"github.com/minlib/go-util/filex"
	"github.com/minlib/go-util/imagex"
	"github.com/shopspring/decimal"
	"golang.org/x/image/webp"
	"os"
)

type WebpUtil struct {
	LibPath string
}

type CommandOption struct {
	LibPath  string
	Quality  int
	Crop     *Crop
	Resize   *Resize
	Original string
	Output   string
}

type Resize struct {
	Width  int
	Height int
}

type Crop struct {
	X int
	Y int
	W int
	H int
}

func CWebp(option *CommandOption) error {
	if err := filex.MkdirAll(option.Output); err != nil {
		return err
	}
	var params string
	if option.Quality > 0 {
		params += fmt.Sprintf(" -q %d", option.Quality)
	}
	if option.Crop != nil {
		params += fmt.Sprintf(" -crop %d %d %d %d", option.Crop.X, option.Crop.Y, option.Crop.W, option.Crop.H)
	}
	if option.Resize != nil {
		params += fmt.Sprintf(" -resize %d %d", option.Resize.Width, option.Resize.Height)
	}
	command := fmt.Sprintf("%scwebp%s %s -o %s", option.LibPath, params, option.Original, option.Output)
	c := cmd.Command(command)
	_, err := c.CombinedOutput()
	return err
}

// GetWebpSize Get the size of the webp image
func GetWebpSize(filename string) (int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	cfg, err := webp.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}

// ScaleToFill 缩放模式，不保持纵横比缩放图片，使图片的宽高完全拉伸至填满 image 元素
func (s *WebpUtil) ScaleToFill(original, output string, width, height, quality int) error {
	return CWebp(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   &Resize{Width: width, Height: height},
		Original: original,
		Output:   output,
	})
}

// AspectFill 缩放模式，保持纵横比缩放图片，只保证图片的短边能完全显示出来。也就是说，图片通常只在水平或垂直方向是完整的，另一个方向将会发生截取。
func (s *WebpUtil) AspectFill(original, output string, width, quality int) error {
	imgWidth, imgHeight, err := imagex.GetSize(original)
	if err != nil {
		return err
	}
	var crop *Crop
	var resize *Resize
	if imgWidth > imgHeight {
		x := (imgWidth - imgHeight) / 2
		crop = &Crop{X: x, Y: 0, W: imgHeight, H: imgHeight}
	} else if imgHeight > imgWidth {
		y := (imgHeight - imgWidth) / 2
		crop = &Crop{X: 0, Y: y, W: imgWidth, H: imgWidth}
	}
	if width > 0 {
		resize = &Resize{Width: width, Height: width}
	}
	return CWebp(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     crop,
		Resize:   resize,
		Original: original,
		Output:   output,
	})
}

// WidthFix 缩放模式，宽度不变，高度自动变化，保持原图宽高比不变
func (s *WebpUtil) WidthFix(original, output string, width, quality int) error {
	imgWidth, imgHeight, err := imagex.GetSize(original)
	if err != nil {
		return err
	}
	var resize *Resize
	if width > 0 && width < imgWidth {
		ratio := decimal.NewFromInt(int64(width)).Div(decimal.NewFromInt(int64(imgWidth)))
		height := int(ratio.Mul(decimal.NewFromInt(int64(imgHeight))).IntPart())
		resize = &Resize{Width: width, Height: height}
	}
	return CWebp(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   resize,
		Original: original,
		Output:   output,
	})
}

// HeightFix 缩放模式，高度不变，宽度自动变化，保持原图宽高比不变
func (s *WebpUtil) HeightFix(original, output string, height, quality int) error {
	imgWidth, imgHeight, err := imagex.GetSize(original)
	if err != nil {
		return err
	}
	var resize *Resize
	if height > 0 && height < imgHeight {
		ratio := decimal.NewFromInt(int64(height)).Div(decimal.NewFromInt(int64(imgHeight)))
		width := int(ratio.Mul(decimal.NewFromInt(int64(imgWidth))).IntPart())
		resize = &Resize{Width: width, Height: height}
	}
	return CWebp(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   resize,
		Original: original,
		Output:   output,
	})
}
