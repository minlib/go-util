package random

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestRangeIntn(t *testing.T) {
	fmt.Println(RangeInt(0, 5))  // [0,5]
	fmt.Println(RangeInt(5, 10)) // [5,10]
	fmt.Println(RangeInt(1, 2))  // [1,2]
	fmt.Println(RangeInt(1, 1))  // 1
}

func TestRandom(t *testing.T) {
	fmt.Println(Random("ABCDEFGHJKLMNPQRSTWXYZabcdefghjkmnpqrstwxyz23456789", 50))
	fmt.Println(Random("我不是猪", 50))
}

func TestUpperCase(t *testing.T) {
	fmt.Println(UpperCase(50))
}

func TestLowerCase(t *testing.T) {
	fmt.Println(LowerCase(50))
}

func TestNumeral(t *testing.T) {
	fmt.Println(Numeral(50))
}

func TestNumeralOrCase(t *testing.T) {
	fmt.Println(NumeralOrCase(50))
}

func TestCaptchaExplicit(t *testing.T) {
	fmt.Println(CaptchaExplicit(50))
}

func TestNewUUID(t *testing.T) {
	fmt.Println(NewUUID())
	fmt.Println(uuid.NewString())
}
