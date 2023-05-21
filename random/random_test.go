package random

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestIntRange(t *testing.T) {
	fmt.Println(IntRange(0, 5))  // [0,5]
	fmt.Println(IntRange(5, 10)) // [5,10]
	fmt.Println(IntRange(1, 2))  // [1,2]
	fmt.Println(IntRange(1, 1))  // 1
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

func TestNextZeroFill(t *testing.T) {
	fmt.Println(NextZeroFill(1, 1, 2))       // 01
	fmt.Println(NextZeroFill(12, 12, 2))     // 12
	fmt.Println(NextZeroFill(123, 123, 2))   // 123
	fmt.Println(NextZeroFill(1234, 1234, 2)) // 1234

	fmt.Println(NextZeroFill(123, 123, 4))   // 0123
	fmt.Println(NextZeroFill(1234, 1234, 4)) // 1234
}
