package colorx

import (
	"fmt"
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
