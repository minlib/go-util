package random

import (
	"math/rand"
	"time"
)

const (
	UPPER_CASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER_CASE = "abcdefghijklmnopqrstuvwxyz"
	NUMBER     = "0123456789"
)

// RangeInt  returns, as an int, a non-negative pseudo-random number in the range interval [min,max]
func RangeInt(min, max int) int {
	if min > max {
		panic("the min is greater than max!")
	} else if min == max {
		return min
	}
	return min + rand.Intn(max+1-min)
}

// String 生成随机大写字符串
func String(s string, count int) string {
	length := len(s)
	result := make([]byte, count)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		result[i] = s[r.Intn(length)]
	}
	return string(result)
}

// LowerString 生成随机大写字符串
func UpperString(count int) string {
	return String(UPPER_CASE, count)
}

// LowerString 生成随机小写字符串
func LowerString(count int) string {
	return String(LOWER_CASE, count)
}

// NumberString 生成随机数字字符串
func NumberString(count int) string {
	return String(NUMBER, count)
}
