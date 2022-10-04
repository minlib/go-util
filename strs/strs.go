package strs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FirstToUpper First letter to upper case
func FirstToUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

// FirstToLower First letter to lower case
func FirstToLower(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

// HumpToUnderline 大小驼峰转下划线
func HumpToUnderline(s string) string {
	r := regexp.MustCompile("[A-Z]")
	result := r.ReplaceAllStringFunc(s, func(s string) string {
		return "_" + strings.ToLower(s[0:])
	})
	return strings.TrimLeft(result, "_")
}

// UnderlineToHump 下划线转驼峰
func UnderlineToHump(s string) string {
	r := regexp.MustCompile("_[a-z]")
	str := r.ReplaceAllStringFunc(s, func(s string) string {
		return strings.ToUpper(s[1:])
	})
	return str
}

// UnderlineToUpperHump 下划线转大驼峰
func UnderlineToUpperHump(s string) string {
	str := UnderlineToHump(s)
	return FirstToUpper(str)
}

// UnderlineToLowerHump 下划线转小驼峰
func UnderlineToLowerHump(s string) string {
	str := UnderlineToHump(s)
	return FirstToLower(str)
}

// FixedLengthNumber 固定长度数字的字符串，前面补零
// FixedLengthNumber(1, 5) // 00001
// FixedLengthNumber(100, 6) // 000100
// FixedLengthNumber(123456, 6)  // 123456
// FixedLengthNumber(1234567, 6) // 1234567
func FixedLengthNumber(number, bit int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(bit)+"d", number)
}

// RuneLength 长度
// RuneLength("abc123") 6
// RuneLength("程序员abc123") 9
func RuneLength(s string) int {
	result := []rune(s)
	return len(result)
}

// ReplaceOffset 替换指定位置与数量的字符串
// s 原字符串
// repl 替换的字符
// offset 开始替换字符串的索引
// limit 替换的字符串数
func ReplaceOffset(s string, repl rune, offset, limit int) string {
	result := []rune(s)
	length := len(result)
	if offset < 0 {
		offset = 0
	}
	for i := offset; i < offset+limit && i < length; i++ {
		result[i] = repl
	}
	return string(result)
}

// HideLeft
func HideLeft(s string, offset int) string {
	length := RuneLength(s)
	limit := length - offset + 1
	return HideLeftLimit(s, offset, limit)
}

// HideLeftLimit
func HideLeftLimit(s string, offset, limit int) string {
	return ReplaceOffset(s, '*', offset, limit)
}

// HideRight
func HideRight(s string, offset int) string {
	length := RuneLength(s)
	limit := length - offset
	return HideRightLimit(s, offset, limit)
}

// HideRightLimit
func HideRightLimit(s string, offset, limit int) string {
	length := RuneLength(s)
	offsetNew := length - offset - limit
	if offsetNew < 0 {
		offsetNew = 0
		limit = length - offset
	}
	return ReplaceOffset(s, '*', offsetNew, limit)
}
