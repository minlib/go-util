package imagex

import (
	"errors"
	"fmt"
	"github.com/minlib/go-terminal/cmd"
	"github.com/minlib/go-util/filex"
	"golang.org/x/image/webp"
	"os"
	"strings"
)

var (
	ErrBadDimension = errors.New("图片尺寸错误")
)

const (
	ScaleToFill  = int8(1)
	AspectFill   = int8(2)
	WidthFix     = int8(3)
	HeightFix    = int8(4)
	WidthLessen  = int8(5)
	HeightLessen = int8(6)
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

func CWebpCommand(option *CommandOption) error {
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
	if bytes, err := c.CombinedOutput(); err != nil {
		errorMessage := string(bytes)
		// Bad picture dimension. Maximum width and height allowed is 16383 pixels.
		if strings.Contains(errorMessage, "Bad picture dimension") {
			return ErrBadDimension
		}
		return err
	}
	return nil
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

// CWebp 根据缩放模式生成
func (s *WebpUtil) CWebp(original, output string, mode int8, width, height, quality int) error {
	if mode == ScaleToFill {
		return s.ScaleToFill(original, output, width, height, quality)
	} else if mode == AspectFill {
		return s.AspectFill(original, output, width, quality)
	} else if mode == WidthFix {
		return s.WidthFix(original, output, width, quality)
	} else if mode == HeightFix {
		return s.HeightFix(original, output, height, quality)
	} else if mode == WidthLessen {
		return s.WidthLessen(original, output, width, quality)
	} else if mode == HeightLessen {
		return s.HeightLessen(original, output, height, quality)
	} else {
		return s.WidthFix(original, output, width, quality)
	}
}

// ScaleToFill 缩放模式，不保持纵横比缩放图片，使图片的宽高完全拉伸至填满 image 元素
func (s *WebpUtil) ScaleToFill(original, output string, width, height, quality int) error {
	return CWebpCommand(&CommandOption{
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
	imgWidth, imgHeight, err := GetSize(original)
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
	return CWebpCommand(&CommandOption{
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
	return CWebpCommand(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   &Resize{Width: width, Height: 0},
		Original: original,
		Output:   output,
	})
}

// HeightFix 缩放模式，高度不变，宽度自动变化，保持原图宽高比不变
func (s *WebpUtil) HeightFix(original, output string, height, quality int) error {
	return CWebpCommand(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   &Resize{Width: 0, Height: height},
		Original: original,
		Output:   output,
	})
}

// WidthLessen 缩小模式，宽度不变，高度自动变化，保持原图宽高比不变，宽度不会超过原图
func (s *WebpUtil) WidthLessen(original, output string, width, quality int) error {
	var resize *Resize
	if width > 0 {
		imgWidth, _, err := GetSize(original)
		if err != nil {
			return err
		}
		if width < imgWidth {
			resize = &Resize{Width: width, Height: 0}
		}
	}
	return CWebpCommand(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   resize,
		Original: original,
		Output:   output,
	})
}

// HeightLessen 缩小模式，高度不变，宽度自动变化，保持原图宽高比不变，宽度不会超过原图
func (s *WebpUtil) HeightLessen(original, output string, height, quality int) error {
	var resize *Resize
	if height > 0 {
		_, imgHeight, err := GetSize(original)
		if err != nil {
			return err
		}
		if height < imgHeight {
			resize = &Resize{Width: 0, Height: height}
		}
	}
	return CWebpCommand(&CommandOption{
		LibPath:  s.LibPath,
		Quality:  quality,
		Crop:     nil,
		Resize:   resize,
		Original: original,
		Output:   output,
	})
}
