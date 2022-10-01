package core

import (
	"regexp"
	"strings"
)

// FirstToUpper First letter to upper case
func FirstToUpper(str string) string {
	return strings.ToUpper(str[0:1]) + str[1:]
}

// FirstToLower First letter to lower case
func FirstToLower(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
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
