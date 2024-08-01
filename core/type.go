// Copyright 2024 Minzhan.com Inc. All rights reserved.

package core

import (
	"github.com/minlib/go-util/jsonx"
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

// Int32 复制 int64 对象，并返回复制体的指针
func Int32(i int32) *int32 {
	return &i
}

// Decimal 复制 decimal.Decimal 对象，并返回复制体的指针
func Decimal(d decimal.Decimal) *decimal.Decimal {
	return &d
}

// Long 复制 int64 对象，并返回复制体的指针
func Long(i int64) *jsonx.Long {
	return &jsonx.Long{Int64: &i}
}

// DateTime 复制 time.Time 对象，并返回复制体的指针
func DateTime(t time.Time) *jsonx.DateTime {
	return &jsonx.DateTime{Time: t}
}
