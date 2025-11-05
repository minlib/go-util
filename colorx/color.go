package colorx

import (
	"image/color"
	"strconv"
	"strings"
)

func toHex(t uint8) string {
	result := strconv.FormatUint(uint64(t), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func RGB2Hex(r, g, b uint8) string {
	red := toHex(r)
	green := toHex(g)
	blue := toHex(b)
	return "#" + red + green + blue
}

func Hex2RGB(hex string) (uint8, uint8, uint8) {
	if strings.HasPrefix(hex, "#") {
		hex = strings.TrimLeft(hex, "#")
	}
	r, _ := strconv.ParseUint(hex[:2], 16, 10)
	g, _ := strconv.ParseUint(hex[2:4], 16, 10)
	b, _ := strconv.ParseUint(hex[4:6], 16, 10)
	return uint8(r), uint8(g), uint8(b)
}

// RGBA2Hex 将 color.RGBA 转换为十六进制字符串
func RGBA2Hex(c color.RGBA) string {
	return "#" + toHex(c.R) + toHex(c.G) + toHex(c.B) + toHex(c.A)
}

// Hex2RGBA 将十六进制颜色字符串转换为 color.RGBA
func Hex2RGBA(hex string) color.RGBA {
	// 移除 # 前缀
	hex = strings.TrimPrefix(hex, "#")

	// 处理不同的十六进制格式
	switch len(hex) {
	case 3: // RGB 格式
		r, _ := strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
		g, _ := strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
		b, _ := strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
		return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	case 4: // RGBA 格式
		r, _ := strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
		g, _ := strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
		b, _ := strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
		a, _ := strconv.ParseUint(string(hex[3])+string(hex[3]), 16, 8)
		return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	case 6: // RRGGBB 格式
		r, _ := strconv.ParseUint(hex[0:2], 16, 8)
		g, _ := strconv.ParseUint(hex[2:4], 16, 8)
		b, _ := strconv.ParseUint(hex[4:6], 16, 8)
		return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	case 8: // RRGGBBAA 格式
		r, _ := strconv.ParseUint(hex[0:2], 16, 8)
		g, _ := strconv.ParseUint(hex[2:4], 16, 8)
		b, _ := strconv.ParseUint(hex[4:6], 16, 8)
		a, _ := strconv.ParseUint(hex[6:8], 16, 8)
		return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	default:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // 默认黑色
	}
}
