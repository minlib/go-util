package random

import (
	"fmt"
	"testing"
)

func TestRangeIntn(t *testing.T) {
	fmt.Println(RangeInt(0, 5))  // [0,5]
	fmt.Println(RangeInt(5, 10)) // [5,10]
	fmt.Println(RangeInt(1, 2))  // [1,2]
	fmt.Println(RangeInt(1, 1))  // 1
}

func TestString(t *testing.T) {
	fmt.Println(String("ABCDEFGHJKLMNPQRSTWXYZabcdefghjkmnpqrstwxyz23456789", 50))
}

func TestUpperString(t *testing.T) {
	fmt.Println(UpperString(50))
}

func TestLowerString(t *testing.T) {
	fmt.Println(LowerString(50))
}

func TestNumberString(t *testing.T) {
	fmt.Println(NumberString(50))
}
