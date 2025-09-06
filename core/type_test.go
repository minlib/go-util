// Copyright 2024 Minzhan.com Inc. All rights reserved.

package core

import (
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

// 测试指针转换函数（如Time、String、Bool等）
func TestPointerFunctions(t *testing.T) {
	// 测试Time函数
	now := time.Now()
	timePtr := Time(now)
	if timePtr == nil {
		t.Error("Time() 返回了nil指针")
	}
	if *timePtr != now {
		t.Errorf("Time() 结果不符，预期 %v，实际 %v", now, *timePtr)
	}

	// 测试String函数
	str := "test string"
	strPtr := String(str)
	if strPtr == nil {
		t.Error("String() 返回了nil指针")
	}
	if *strPtr != str {
		t.Errorf("String() 结果不符，预期 %s，实际 %s", str, *strPtr)
	}

	// 测试Bool函数
	boolPtr := Bool(true)
	if boolPtr == nil {
		t.Error("Bool() 返回了nil指针")
	}
	if *boolPtr != true {
		t.Errorf("Bool() 结果不符，预期 %v，实际 %v", true, *boolPtr)
	}

	// 测试Float64函数
	f64 := 3.1415926
	f64Ptr := Float64(f64)
	if f64Ptr == nil {
		t.Error("Float64() 返回了nil指针")
	}
	if *f64Ptr != f64 {
		t.Errorf("Float64() 结果不符，预期 %v，实际 %v", f64, *f64Ptr)
	}

	// 测试Float32函数
	f32 := float32(2.71828)
	f32Ptr := Float32(f32)
	if f32Ptr == nil {
		t.Error("Float32() 返回了nil指针")
	}
	if *f32Ptr != f32 {
		t.Errorf("Float32() 结果不符，预期 %v，实际 %v", f32, *f32Ptr)
	}

	// 测试Int64函数
	i64 := int64(9876543210)
	i64Ptr := Int64(i64)
	if i64Ptr == nil {
		t.Error("Int64() 返回了nil指针")
	}
	if *i64Ptr != i64 {
		t.Errorf("Int64() 结果不符，预期 %d，实际 %d", i64, *i64Ptr)
	}

	// 测试Int32函数
	i32 := int32(12345)
	i32Ptr := Int32(i32)
	if i32Ptr == nil {
		t.Error("Int32() 返回了nil指针")
	}
	if *i32Ptr != i32 {
		t.Errorf("Int32() 结果不符，预期 %d，实际 %d", i32, *i32Ptr)
	}

	// 测试Int16函数
	i16 := int16(32767)
	i16Ptr := Int16(i16)
	if i16Ptr == nil {
		t.Error("Int16() 返回了nil指针")
	}
	if *i16Ptr != i16 {
		t.Errorf("Int16() 结果不符，预期 %d，实际 %d", i16, *i16Ptr)
	}

	// 测试Int8函数
	i8 := int8(127)
	i8Ptr := Int8(i8)
	if i8Ptr == nil {
		t.Error("Int8() 返回了nil指针")
	}
	if *i8Ptr != i8 {
		t.Errorf("Int8() 结果不符，预期 %d，实际 %d", i8, *i8Ptr)
	}

	// 测试Int函数
	i := 10086
	iPtr := Int(i)
	if iPtr == nil {
		t.Error("Int() 返回了nil指针")
	}
	if *iPtr != i {
		t.Errorf("Int() 结果不符，预期 %d，实际 %d", i, *iPtr)
	}

	// 测试Decimal函数
	d := decimal.NewFromInt(100)
	dPtr := Decimal(d)
	if dPtr == nil {
		t.Error("Decimal() 返回了nil指针")
	}
	if *dPtr != d {
		t.Errorf("Decimal() 结果不符，预期 %v，实际 %v", d, *dPtr)
	}
}

// 测试字符串取值函数
func TestStringValueFunctions(t *testing.T) {
	// 测试StringValue
	t.Run("StringValue", func(t *testing.T) {
		// 非nil情况
		s := "hello"
		if got := StringValue(&s); got != s {
			t.Errorf("StringValue(&s) 预期 %s，实际 %s", s, got)
		}
		// nil情况
		if got := StringValue(nil); got != "" {
			t.Errorf("StringValue(nil) 预期空字符串，实际 %s", got)
		}
	})

	// 测试StringValueOrDefault
	t.Run("StringValueOrDefault", func(t *testing.T) {
		// 非nil情况
		s := "world"
		def := "default"
		if got := StringValueOrDefault(&s, def); got != s {
			t.Errorf("StringValueOrDefault(&s, def) 预期 %s，实际 %s", s, got)
		}
		// nil情况
		if got := StringValueOrDefault(nil, def); got != def {
			t.Errorf("StringValueOrDefault(nil, def) 预期 %s，实际 %s", def, got)
		}
	})
}

// 测试布尔取值函数
func TestBoolValueFunctions(t *testing.T) {
	// 测试BoolValue
	t.Run("BoolValue", func(t *testing.T) {
		// 非nil情况（true）
		b := true
		if got := BoolValue(&b); !got {
			t.Error("BoolValue(&true) 预期true，实际false")
		}
		// 非nil情况（false）
		b = false
		if got := BoolValue(&b); got {
			t.Error("BoolValue(&false) 预期false，实际true")
		}
		// nil情况
		if got := BoolValue(nil); got {
			t.Error("BoolValue(nil) 预期false，实际true")
		}
	})

	// 测试BoolValueOrDefault
	t.Run("BoolValueOrDefault", func(t *testing.T) {
		// 非nil情况
		b := false
		if got := BoolValueOrDefault(&b, true); got {
			t.Errorf("BoolValueOrDefault(&false, true) 预期false，实际true")
		}
		// nil情况
		if got := BoolValueOrDefault(nil, true); !got {
			t.Errorf("BoolValueOrDefault(nil, true) 预期true，实际false")
		}
	})
}

// 测试浮点型取值函数
func TestFloatValueFunctions(t *testing.T) {
	// 测试Float64Value
	t.Run("Float64Value", func(t *testing.T) {
		f := 123.456
		if got := Float64Value(&f); got != f {
			t.Errorf("Float64Value(&f) 预期 %v，实际 %v", f, got)
		}
		if got := Float64Value(nil); got != 0 {
			t.Errorf("Float64Value(nil) 预期0，实际 %v", got)
		}
	})

	// 测试Float64ValueOrDefault
	t.Run("Float64ValueOrDefault", func(t *testing.T) {
		f := 789.012
		def := 345.678
		if got := Float64ValueOrDefault(&f, def); got != f {
			t.Errorf("Float64ValueOrDefault(&f, def) 预期 %v，实际 %v", f, got)
		}
		if got := Float64ValueOrDefault(nil, def); got != def {
			t.Errorf("Float64ValueOrDefault(nil, def) 预期 %v，实际 %v", def, got)
		}
	})

	// 测试Float32Value
	t.Run("Float32Value", func(t *testing.T) {
		f := float32(11.22)
		if got := Float32Value(&f); got != f {
			t.Errorf("Float32Value(&f) 预期 %v，实际 %v", f, got)
		}
		if got := Float32Value(nil); got != 0 {
			t.Errorf("Float32Value(nil) 预期0，实际 %v", got)
		}
	})

	// 测试Float32ValueOrDefault
	t.Run("Float32ValueOrDefault", func(t *testing.T) {
		f := float32(33.44)
		def := float32(55.66)
		if got := Float32ValueOrDefault(&f, def); got != f {
			t.Errorf("Float32ValueOrDefault(&f, def) 预期 %v，实际 %v", f, got)
		}
		if got := Float32ValueOrDefault(nil, def); got != def {
			t.Errorf("Float32ValueOrDefault(nil, def) 预期 %v，实际 %v", def, got)
		}
	})
}

// 测试整数型取值函数
func TestIntValueFunctions(t *testing.T) {
	// 测试Int64Value
	t.Run("Int64Value", func(t *testing.T) {
		i := int64(1000)
		if got := Int64Value(&i); got != i {
			t.Errorf("Int64Value(&i) 预期 %d，实际 %d", i, got)
		}
		if got := Int64Value(nil); got != 0 {
			t.Errorf("Int64Value(nil) 预期0，实际 %d", got)
		}
	})

	// 测试Int64ValueOrDefault
	t.Run("Int64ValueOrDefault", func(t *testing.T) {
		i := int64(2000)
		def := int64(3000)
		if got := Int64ValueOrDefault(&i, def); got != i {
			t.Errorf("Int64ValueOrDefault(&i, def) 预期 %d，实际 %d", i, got)
		}
		if got := Int64ValueOrDefault(nil, def); got != def {
			t.Errorf("Int64ValueOrDefault(nil, def) 预期 %d，实际 %d", def, got)
		}
	})

	// 测试Int32Value
	t.Run("Int32Value", func(t *testing.T) {
		i := int32(123)
		if got := Int32Value(&i); got != i {
			t.Errorf("Int32Value(&i) 预期 %d，实际 %d", i, got)
		}
		if got := Int32Value(nil); got != 0 {
			t.Errorf("Int32Value(nil) 预期0，实际 %d", got)
		}
	})

	// 测试Int32ValueOrDefault
	t.Run("Int32ValueOrDefault", func(t *testing.T) {
		i := int32(456)
		def := int32(789)
		if got := Int32ValueOrDefault(&i, def); got != i {
			t.Errorf("Int32ValueOrDefault(&i, def) 预期 %d，实际 %d", i, got)
		}
		if got := Int32ValueOrDefault(nil, def); got != def {
			t.Errorf("Int32ValueOrDefault(nil, def) 预期 %d，实际 %d", def, got)
		}
	})

	// 测试Int16Value
	t.Run("Int16Value", func(t *testing.T) {
		i := int16(12)
		if got := Int16Value(&i); got != i {
			t.Errorf("Int16Value(&i) 预期 %d，实际 %d", i, got)
		}
		if got := Int16Value(nil); got != 0 {
			t.Errorf("Int16Value(nil) 预期0，实际 %d", got)
		}
	})

	// 测试Int16ValueOrDefault
	t.Run("Int16ValueOrDefault", func(t *testing.T) {
		i := int16(34)
		def := int16(56)
		if got := Int16ValueOrDefault(&i, def); got != i {
			t.Errorf("Int16ValueOrDefault(&i, def) 预期 %d，实际 %d", i, got)
		}
		if got := Int16ValueOrDefault(nil, def); got != def {
			t.Errorf("Int16ValueOrDefault(nil, def) 预期 %d，实际 %d", def, got)
		}
	})

	// 测试Int8Value
	t.Run("Int8Value", func(t *testing.T) {
		i := int8(7)
		if got := Int8Value(&i); got != i {
			t.Errorf("Int8Value(&i) 预期 %d，实际 %d", i, got)
		}
		if got := Int8Value(nil); got != 0 {
			t.Errorf("Int8Value(nil) 预期0，实际 %d", got)
		}
	})

	// 测试Int8ValueOrDefault
	t.Run("Int8ValueOrDefault", func(t *testing.T) {
		i := int8(8)
		def := int8(9)
		if got := Int8ValueOrDefault(&i, def); got != i {
			t.Errorf("Int8ValueOrDefault(&i, def) 预期 %d，实际 %d", i, got)
		}
		if got := Int8ValueOrDefault(nil, def); got != def {
			t.Errorf("Int8ValueOrDefault(nil, def) 预期 %d，实际 %d", def, got)
		}
	})

	// 测试IntValue
	t.Run("IntValue", func(t *testing.T) {
		i := 100
		if got := IntValue(&i); got != i {
			t.Errorf("IntValue(&i) 预期 %d，实际 %d", i, got)
		}
		if got := IntValue(nil); got != 0 {
			t.Errorf("IntValue(nil) 预期0，实际 %d", got)
		}
	})

	// 测试IntValueOrDefault（注意修复原注释错误）
	t.Run("IntValueOrDefault", func(t *testing.T) {
		i := 200
		def := 300
		if got := IntValueOrDefault(&i, def); got != i {
			t.Errorf("IntValueOrDefault(&i, def) 预期 %d，实际 %d", i, got)
		}
		if got := IntValueOrDefault(nil, def); got != def {
			t.Errorf("IntValueOrDefault(nil, def) 预期 %d，实际 %d", def, got)
		}
	})
}

// 测试Decimal取值函数
func TestDecimalValueFunctions(t *testing.T) {
	// 测试DecimalValue
	t.Run("DecimalValue", func(t *testing.T) {
		d := decimal.NewFromInt(100)
		if got := DecimalValue(&d); got != d {
			t.Errorf("DecimalValue(&d) 预期 %v，实际 %v", d, got)
		}
		// 测试nil时返回零值
		zero := decimal.Decimal{}
		if got := DecimalValue(nil); got != zero {
			t.Errorf("DecimalValue(nil) 预期零值 %v，实际 %v", zero, got)
		}
	})

	// 测试DecimalValueOrDefault
	t.Run("DecimalValueOrDefault", func(t *testing.T) {
		d := decimal.NewFromInt(200)
		def := decimal.NewFromInt(300)
		if got := DecimalValueOrDefault(&d, def); got != d {
			t.Errorf("DecimalValueOrDefault(&d, def) 预期 %v，实际 %v", d, got)
		}
		if got := DecimalValueOrDefault(nil, def); got != def {
			t.Errorf("DecimalValueOrDefault(nil, def) 预期 %v，实际 %v", def, got)
		}
	})
}
