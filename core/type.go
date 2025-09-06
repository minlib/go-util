// Copyright 2024 Minzhan.com Inc. All rights reserved.

package core

import (
	"github.com/shopspring/decimal"
	"time"
)

// Time 复制 time.Time 对象，并返回复制体的指针
func Time(t time.Time) *time.Time {
	return &t
}

// String 复制 string 对象，并返回复制体的指针
func String(s string) *string {
	return &s
}

// Bool 复制 bool 对象，并返回复制体的指针
func Bool(b bool) *bool {
	return &b
}

// Float64 复制 float64 对象，并返回复制体的指针
func Float64(f float64) *float64 {
	return &f
}

// Float32 复制 float32 对象，并返回复制体的指针
func Float32(f float32) *float32 {
	return &f
}

// Int64 复制 int64 对象，并返回复制体的指针
func Int64(i int64) *int64 {
	return &i
}

// Int32 复制 int32 对象，并返回复制体的指针
func Int32(i int32) *int32 {
	return &i
}

// Int16 复制 int16 对象，并返回复制体的指针
func Int16(i int16) *int16 {
	return &i
}

// Int8 复制 int8 对象，并返回复制体的指针
func Int8(i int8) *int8 {
	return &i
}

// Int 复制 int 对象，并返回复制体的指针
func Int(i int) *int {
	return &i
}

// Decimal 复制 decimal.Decimal 对象，并返回复制体的指针
func Decimal(d decimal.Decimal) *decimal.Decimal {
	return &d
}

// StringValue 从*string获取值，为nil时返回空字符串
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// StringValueOrDefault 从*string获取值，为nil时返回默认值
func StringValueOrDefault(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return *s
}

// BoolValue 从*bool获取值，为nil时返回false
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// BoolValueOrDefault 从*bool获取值，为nil时返回默认值
func BoolValueOrDefault(b *bool, defaultValue bool) bool {
	if b == nil {
		return defaultValue
	}
	return *b
}

// Float64Value 从*float64获取值，为nil时返回0
func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// Float64ValueOrDefault 从*float64获取值，为nil时返回默认值
func Float64ValueOrDefault(f *float64, defaultValue float64) float64 {
	if f == nil {
		return defaultValue
	}
	return *f
}

// Float32Value 从*float32获取值，为nil时返回0
func Float32Value(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

// Float32ValueOrDefault 从*float32获取值，为nil时返回默认值
func Float32ValueOrDefault(f *float32, defaultValue float32) float32 {
	if f == nil {
		return defaultValue
	}
	return *f
}

// Int64Value 从*int64获取值，为nil时返回0
func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// Int64ValueOrDefault 从*int64获取值，为nil时返回默认值
func Int64ValueOrDefault(i *int64, defaultValue int64) int64 {
	if i == nil {
		return defaultValue
	}
	return *i
}

// Int32Value 从*int32获取值，为nil时返回0
func Int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// Int32ValueOrDefault 从*int32获取值，为nil时返回默认值
func Int32ValueOrDefault(i *int32, defaultValue int32) int32 {
	if i == nil {
		return defaultValue
	}
	return *i
}

// Int16Value 从*int16获取值，为nil时返回0
func Int16Value(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// Int16ValueOrDefault 从*int16获取值，为nil时返回默认值
func Int16ValueOrDefault(i *int16, defaultValue int16) int16 {
	if i == nil {
		return defaultValue
	}
	return *i
}

// Int8Value 从*int8获取值，为nil时返回0
func Int8Value(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// Int8ValueOrDefault 从*int8获取值，为nil时返回默认值
func Int8ValueOrDefault(i *int8, defaultValue int8) int8 {
	if i == nil {
		return defaultValue
	}
	return *i
}

// IntValue 从*int获取值，为nil时返回0
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// IntValueOrDefault 从*int获取值，为nil时返回默认值
func IntValueOrDefault(i *int, defaultValue int) int {
	if i == nil {
		return defaultValue
	}
	return *i
}

// DecimalValue 从*decimal.Decimal获取值，为nil时返回零值
func DecimalValue(d *decimal.Decimal) decimal.Decimal {
	if d == nil {
		return decimal.Decimal{}
	}
	return *d
}

// DecimalValueOrDefault 从*decimal.Decimal获取值，为nil时返回默认值
func DecimalValueOrDefault(d *decimal.Decimal, defaultValue decimal.Decimal) decimal.Decimal {
	if d == nil {
		return defaultValue
	}
	return *d
}
