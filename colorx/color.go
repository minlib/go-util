package colorx

import (
	"strconv"
	"strings"
)

func rgb2HexPart(t uint8) string {
	result := strconv.FormatUint(uint64(t), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func RGB2Hex(r, g, b uint8) string {
	red := rgb2HexPart(r)
	green := rgb2HexPart(g)
	blue := rgb2HexPart(b)
	return "#" + red + green + blue
}

func Hex2RGB(hex string) (uint8, uint8, uint8) {
	if strings.HasPrefix(hex, "#") {
		hex = strings.TrimLeft(hex, "#")
	}
	r, _ := strconv.ParseUint(hex[:2], 16, 10)
	g, _ := strconv.ParseUint(hex[2:4], 16, 10)
	b, _ := strconv.ParseUint(hex[4:], 16, 10)
	return uint8(r), uint8(g), uint8(b)
}
