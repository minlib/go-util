package imagex

import (
	"bytes"
	"github.com/minlib/go-util/filex"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetResourceReader get local or remote resource file
func GetResourceReader(pathOrUrl string) (*bytes.Reader, error) {
	if strings.HasPrefix(pathOrUrl, "https://") || strings.HasPrefix(pathOrUrl, "http://") {
		resp, err := http.Get(pathOrUrl)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		fileBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	} else {
		fileBytes, err := os.ReadFile(pathOrUrl)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(fileBytes), nil
	}
}

// NewDrawable creates a drawable image from an existing image.
// This is useful when you need to perform drawing operations on an existing image.
// Parameters:
//   - srcImage: The source image to create a drawable copy from
//
// Returns:
//   - draw.Image: A new drawable image containing the contents of the source image
func NewDrawable(srcImage image.Image) draw.Image {
	// Get the bounds of the source image
	bounds := srcImage.Bounds()

	// Create a new RGBA image with the same dimensions as the source
	img := image.NewRGBA(bounds)

	// Copy the source image to the new image using draw.Draw for efficiency
	draw.Draw(img, bounds, srcImage, bounds.Min, draw.Src)

	return img
}

// NewImage creates a new image filled with a solid color.
// This function returns a draw.Image which can be used for further drawing operations.
// Parameters:
//   - width: The width of the new image in pixels
//   - height: The height of the new image in pixels
//   - fillColor: The color to fill the image with
//
// Returns:
//   - draw.Image: A new image filled with the specified color
func NewImage(width, height int, fillColor color.Color) draw.Image {
	// Create a new RGBA image with the specified dimensions
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill the entire image with the background color using draw.Draw for efficiency
	draw.Draw(img, img.Bounds(), &image.Uniform{C: fillColor}, image.Point{}, draw.Src)

	return img
}

// NewImageFromFile creates a drawable image from an image file.
// This is useful when you need to perform drawing operations on an image loaded from a file.
// Parameters:
//   - srcImagePath: The path to the source image file
//
// Returns:
//   - draw.Image: A new drawable image containing the contents of the source image file
//   - error: Any error encountered while reading or processing the image
func NewImageFromFile(srcImagePath string) (draw.Image, error) {
	// Read the source image from file
	srcImage, err := ReadImage(srcImagePath)
	if err != nil {
		// If there's an error reading the image, return the error
		return nil, err
	}
	img := NewDrawable(srcImage)
	return img, nil
}

// ReadImage read a image
func ReadImage(pathOrUrl string) (image.Image, error) {
	reader, err := GetResourceReader(pathOrUrl)
	if err != nil {
		return nil, err
	}
	// 加载JPG格式图片时，必须引入依赖 "image/jpeg"，否则会出错 unknown format
	srcImage, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return srcImage, err
}

// GetSize get the size of the image
func GetSize(path string) (int, int, error) {
	if strings.HasSuffix(path, ".webp") {
		return GetWebpSize(path)
	}
	srcImage, err := ReadImage(path)
	if err != nil {
		return 0, 0, err
	}
	// 获取图片的宽和高
	width := srcImage.Bounds().Dx()
	height := srcImage.Bounds().Dy()
	return width, height, nil
}

// SavePNG save as png image
func SavePNG(src image.Image, filename string) error {
	if err := filex.MkdirAll(filename); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, src)
}

// SaveJPG save as jpg image
func SaveJPG(src image.Image, filename string, quality int) error {
	if err := filex.MkdirAll(filename); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, src, &jpeg.Options{
		Quality: quality,
	})
}

// ResizeImage 图像按比例缩放到指定宽度
func ResizeImage(src image.Image, width int) (image.Image, error) {
	bounds := src.Bounds()
	srcWidth, srcHeight := bounds.Dx(), bounds.Dy()
	newHeight := int(float64(srcHeight) * (float64(width) / float64(srcWidth)))
	newImage := image.NewRGBA(image.Rect(0, 0, width, newHeight))
	for x := 0; x < width; x++ {
		for y := 0; y < newHeight; y++ {
			srcX := int(float64(x) * (float64(srcWidth) / float64(width)))
			srcY := int(float64(y) * (float64(srcHeight) / float64(newHeight)))
			if srcX < 0 {
				srcX = 0
			}
			if srcY < 0 {
				srcY = 0
			}
			if srcX >= srcWidth {
				srcX = srcWidth - 1
			}
			if srcY >= srcHeight {
				srcY = srcHeight - 1
			}
			newImage.Set(x, y, src.At(srcX, srcY))
		}
	}
	return newImage, nil
}

// AddBorder 图片添加边框
func AddBorder(src image.Image, borderWidth int, backgroundColor color.Color) (image.Image, error) {
	bounds := src.Bounds()
	srcWidth, srcHeight := bounds.Dx(), bounds.Dy()
	// 计算带边框的新尺寸（两侧各增加borderWidth）
	newWidth := srcWidth + 2*borderWidth
	newHeight := srcHeight + 2*borderWidth
	// 创建新图像作为画布（自定义背景颜色，如 白色：color.RGBA{R: 255, G: 255, B: 255, A: 255}）
	newImage := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.Draw(newImage, newImage.Bounds(), &image.Uniform{C: backgroundColor}, image.Point{}, draw.Src)
	// 将原图片绘制到新图像的中间（留出边框空间）
	rectangle := image.Rect(borderWidth, borderWidth, srcWidth+borderWidth, srcHeight+borderWidth)
	draw.Draw(newImage, rectangle, src, bounds.Min, draw.Src)
	return newImage, nil
}
