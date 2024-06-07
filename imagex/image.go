package imagex

import (
	"bytes"
	"github.com/minlib/go-util/stringx"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

// GetResourceReader get local or remote resource file
func GetResourceReader(pathOrUrl string) (*bytes.Reader, error) {
	if stringx.HasAnyPrefix(pathOrUrl, "https://", "http://") {
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
	srcImage, err := ReadImage(path)
	if err != nil {
		return 0, 0, err
	}
	// 获取图片的宽和高
	srcBounds := srcImage.Bounds()
	width := srcBounds.Max.X - srcBounds.Min.X
	height := srcBounds.Max.Y - srcBounds.Min.Y
	return width, height, nil
}

// GetWebpSize Get the size of the webp image
func GetWebpSize(path string) (int, int, error) {
	file, err := os.Open(path)
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

// SavePNG save as png image
func SavePNG(src image.Image, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, src)
}

// SaveJPG save as jpg image
func SaveJPG(src image.Image, outputPath string, quality int) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return jpeg.Encode(file, src, &jpeg.Options{
		Quality: quality,
	})
}
