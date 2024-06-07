package imagex

import (
	"fmt"
	"testing"
)

func TestGetWebpSize(t *testing.T) {
	width, height, err := GetWebpSize("E:\\wwwroot\\uploads\\store12345678910abc\\thumb\\202406\\0222f2db2cddaccd4bc5bd0d6cc94ca1cb48.webp")
	fmt.Println(width, height, err)
}
