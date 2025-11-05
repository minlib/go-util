package colorx

import (
	"fmt"
	"image/color"
	"testing"
)

func TestHex2RGB(t *testing.T) {
	fmt.Println(Hex2RGB("#0b21f3"))
	fmt.Println(Hex2RGB("#000000"))
	fmt.Println(Hex2RGB("#FFFFFF"))
}

func TestRGB2Hex(t *testing.T) {
	fmt.Println(RGB2Hex(11, 33, 243))
	fmt.Println(RGB2Hex(0, 0, 0))
	fmt.Println(RGB2Hex(255, 255, 255))
}

func TestRGBA2Hex(t *testing.T) {
	rgba := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	hex := RGBA2Hex(rgba)
	fmt.Println(hex) // 应该输出 "#ffffffff"

	rgba2 := color.RGBA{R: 11, G: 33, B: 243, A: 128}
	hex2 := RGBA2Hex(rgba2)
	fmt.Println(hex2) // 应该输出 "#0b21f380"
}

func TestHex2RGBA(t *testing.T) {
	rgba := Hex2RGBA("#ffffffff")
	fmt.Printf("RGBA: %+v\n", rgba) // 应该输出 {R:255 G:255 B:255 A:255}

	rgba2 := Hex2RGBA("#0b21f380")
	fmt.Printf("RGBA: %+v\n", rgba2) // 应该输出 {R:11 G:33 B:243 A:128}
}
