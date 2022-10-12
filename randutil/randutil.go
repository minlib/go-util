package randutil

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	NUMERAL          = "0123456789"
	LOWER_CASE       = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CAPTCHA_EXPLICIT = "ABCDEFGHJKLMNPQRSTWXYabcdefghjkmnprstwxy3456789"
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

// UpperCase 生成随机大写字符串
func UpperCase(count int) string {
	return Random(UPPER_CASE, count)
}

// LowerCase 生成随机小写字符串
func LowerCase(count int) string {
	return Random(LOWER_CASE, count)
}

// Numeral 生成随机数字字符串
func Numeral(count int) string {
	return Random(NUMERAL, count)
}

// NumeralOrCase 生成随机数字或字母字符串
func NumeralOrCase(count int) string {
	return Random(NUMERAL+UPPER_CASE+LOWER_CASE, count)
}

// CaptchaExplicit 生成明确的验证码（排除容易混淆的字符串）
func CaptchaExplicit(count int) string {
	return Random(CAPTCHA_EXPLICIT, count)
}

// NewUUID create uuid
func NewUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
