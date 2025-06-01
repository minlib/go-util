package decimalx

import (
	"github.com/shopspring/decimal"
	"strings"
)

// GetPlaces 获取小数位
func GetPlaces(value decimal.Decimal) int32 {
	valueString := value.String()
	dotIndex := strings.Index(valueString, ".")
	if dotIndex == -1 {
		return 0
	}
	// 获取小数部分
	decimalPart := valueString[dotIndex+1:]
	// 移除末尾的零
	for len(decimalPart) > 0 && decimalPart[len(decimalPart)-1] == '0' {
		decimalPart = decimalPart[:len(decimalPart)-1]
	}
	return int32(len(decimalPart))
}

// IsPlaces 判断小数点的位数
func IsPlaces(value decimal.Decimal, places int32) bool {
	exp := GetPlaces(value)
	return exp == places
}
