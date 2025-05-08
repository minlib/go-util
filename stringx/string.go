package stringx

import (
	"bytes"
	"fmt"
	"github.com/minlib/go-util/slicex"
	"golang.org/x/exp/constraints"
	"regexp"
	"strconv"
	"strings"
)

// HasAnyPrefix tests whether the string s ends with any prefix.
func HasAnyPrefix(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// HasAnySuffix tests whether the string s ends with any suffix.
func HasAnySuffix(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// ContainsAnyString reports whether substr is within any s.
func ContainsAnyString(s string, substrings ...string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// EqualAnyString reports whether the string s equal any t.
func EqualAnyString(s string, ts ...string) bool {
	for _, t := range ts {
		if s == t {
			return true
		}
	}
	return false
}

// EqualAnyFold tests whether the string s equal under simple Unicode case-folding
func EqualAnyFold(s string, ts ...string) bool {
	for _, t := range ts {
		if strings.EqualFold(s, t) {
			return true
		}
	}
	return false
}

// ToString 函数用于将一个接口值转换为字符串类型。
func ToString(i interface{}) string {
	value, _ := i.(string)
	return value
}

// DefaultIfEmpty 函数用于判断字符串是否为空或仅包含空白字符，如果是则返回默认字符串，否则返回原字符串。
func DefaultIfEmpty(str, defaultStr string) string {
	if str != "" && strings.TrimSpace(str) != "" {
		return str
	} else {
		return defaultStr
	}
}

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

// ZeroFill 固定长度数字的字符串，前面补零
// ZeroFill(1, 5) // 00001
// ZeroFill(100, 6) // 000100
// ZeroFill(123456, 6)  // 123456
// ZeroFill(1234567, 6) // 1234567
func ZeroFill(number, bit int) string {
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

// HideLeft 隐藏字符串，从左边指定位置的所有字符串数
func HideLeft(s string, offset int) string {
	length := RuneLength(s)
	limit := length - offset + 1
	return HideLeftLimit(s, offset, limit)
}

// HideLeftLimit  隐藏字符串，从左边开始的字符串数
func HideLeftLimit(s string, offset, limit int) string {
	return ReplaceOffset(s, '*', offset, limit)
}

// HideRight 隐藏字符串，从右边指定位置的所有字符串数
func HideRight(s string, offset int) string {
	length := RuneLength(s)
	limit := length - offset
	return HideRightLimit(s, offset, limit)
}

// HideRightLimit 隐藏字符串，从右边开始限定的字符串数
func HideRightLimit(s string, offset, limit int) string {
	length := RuneLength(s)
	offsetNew := length - offset - limit
	if offsetNew < 0 {
		offsetNew = 0
		limit = length - offset
	}
	return ReplaceOffset(s, '*', offsetNew, limit)
}

// IsEmpty 判断空
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank 判断空白
func IsBlank(s string) bool {
	return s == "" || len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty 判断非空
func IsNotEmpty(s string) bool {
	return s != ""
}

// IsNotBlank 判断非空白
func IsNotBlank(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

// IsAnyEmpty 判断是否存在空（任意一个为空，返回true）
func IsAnyEmpty(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if s == "" {
			return true
		}
	}
	return false
}

// IsAnyBlank 判断是否存在空白（任意一个为空白，返回true）
func IsAnyBlank(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsAnyNotEmpty 判断是否存在非空字符串（任意一个不为空，返回true）
func IsAnyNotEmpty(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if s != "" {
			return true
		}
	}
	return false
}

// IsAnyNotBlank 判断是否存在非空白字符串（任意一个不为空白，返回true）
func IsAnyNotBlank(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if !IsBlank(s) {
			return true
		}
	}
	return false
}

// IsNoneEmpty 判断不存在空（全部不为空，返回true）
func IsNoneEmpty(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if s == "" {
			return false
		}
	}
	return true
}

// IsNoneBlank 判断不存在空白（全部不为空白，返回true）
func IsNoneBlank(strings ...string) bool {
	if len(strings) == 0 {
		return false
	}
	for _, s := range strings {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitToIntegers slices s into all sub integers separated by sep and returns a slice of
// // the sub integers between those separators.
func SplitToIntegers[E constraints.Integer](s, sep string) ([]E, error) {
	slice := strings.Split(s, sep)
	return slicex.StringToInt[E](slice)
}

// SplitToIntegersIgnoreError slices s into all sub integers separated by sep and returns a slice of
// // the sub integers between those separators.
func SplitToIntegersIgnoreError[E constraints.Integer](s, sep string) []E {
	slice := strings.Split(s, sep)
	return slicex.StringToIntIgnoreError[E](slice)
}

// SplitToInt64SliceIgnoreError slices s into all sub integers separated by sep and returns a slice of
// // the sub integers between those separators.
func SplitToInt64SliceIgnoreError(s, sep string) []int64 {
	slice := strings.Split(s, sep)
	return slicex.StringToInt64IgnoreError(slice)
}

// Join 拼接切片
func Join[E constraints.Integer](elems []E, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.FormatInt(int64(elems[0]), 10)
	}
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(int64(elems[0]), 10))
	for _, v := range elems[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(strconv.FormatInt(int64(v), 10))
	}
	return buffer.String()
}
