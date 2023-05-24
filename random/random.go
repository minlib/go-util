package random

import (
	"github.com/minlib/go-util/stringx"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	NUMERAL   = "0123456789"
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOL    = "~!@#$%^&*()_+=<>/,./;'[]{}|"
	CAPTCHA   = "ABCDEFGHJKLMNPQRSTWXYabcdefghjkmnprstwxy3456789"
)

// IntRange returns, as an int, a non-negative pseudo-random number in the range interval [min,max]
func IntRange(min, max int) int {
	if min > max {
		panic("the min value cannot be greater than the max value")
	} else if min == max {
		return min
	}
	return min + rand.Intn(max+1-min)
}

// IntRangeZeroFill Returns a string of random numbers,if less than the specified length, preceded by zeros.
func IntRangeZeroFill(min, max, length int) string {
	randNum := IntRange(min, max)
	return stringx.ZeroFill(randNum, length)
}

// Random 随机生成字符串
func Random(s string, count int) string {
	runes := []rune(s)
	length := len(runes)
	result := make([]rune, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result[i] = runes[r.Intn(length)]
	}
	return string(result)
}

// LowerCase 生成随机小写字符串
func LowerCase(count int) string {
	return Random(LOWERCASE, count)
}

// UpperCase 生成随机大写字符串
func UpperCase(count int) string {
	return Random(UPPERCASE, count)
}

// Numeric 生成随机数字字符串
func Numeric(count int) string {
	return Random(NUMERAL, count)
}

// Alphanumeric 生成随机字母数字
func Alphanumeric(count int) string {
	return Random(NUMERAL+LOWERCASE+UPPERCASE, count)
}

// AlphanumericOrSymbol 生成随机字母数字或符号
func AlphanumericOrSymbol(count int) string {
	return Random(NUMERAL+LOWERCASE+UPPERCASE+SYMBOL, count)
}

// ClarityCaptcha 生成明确的验证码（排除容易混淆的字符串）
func ClarityCaptcha(count int) string {
	return Random(CAPTCHA, count)
}

// NewUUID create uuid
func NewUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
