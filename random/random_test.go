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

func TestNumeric(t *testing.T) {
	fmt.Println(Numeric(50))
}

func TestAlphanumeric(t *testing.T) {
	fmt.Println(Alphanumeric(50))
}

func TestAlphanumericOrSymbol(t *testing.T) {
	fmt.Println(AlphanumericOrSymbol(50))
}

func TestClarityCaptcha(t *testing.T) {
	fmt.Println(ClarityCaptcha(50))
}

func TestNewUUID(t *testing.T) {
	fmt.Println(NewUUID())
	fmt.Println(uuid.NewString())
}

func TestIntRangeZeroFill(t *testing.T) {
	fmt.Println(IntRangeZeroFill(1, 1, 2))       // 01
	fmt.Println(IntRangeZeroFill(12, 12, 2))     // 12
	fmt.Println(IntRangeZeroFill(123, 123, 2))   // 123
	fmt.Println(IntRangeZeroFill(1234, 1234, 2)) // 1234
	fmt.Println(IntRangeZeroFill(123, 123, 4))   // 0123
	fmt.Println(IntRangeZeroFill(1234, 1234, 4)) // 1234
	fmt.Println(IntRangeZeroFill(0, 1234, 4))
}
