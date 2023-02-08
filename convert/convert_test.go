package convert

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	b := IntToBytes(-111)
	fmt.Printf("%T, %v\n", b, b) // []uint8, [255 255 255 255 255 255 255 145]
}

func TestBytesToInt(t *testing.T) {
	b := IntToBytes(-111)
	i := BytesToInt[int](b)
	want := -111
	i64 := BytesToInt[int64](b)
	if i != int(want) {
		t.Errorf("BytesToInt() = %v, want %v", i, want)
	}
	fmt.Printf("%T, %v\n", i, i)     // int, -111
	fmt.Printf("%T, %v\n", i64, i64) // int64, -111
}

func TestIntToString(t *testing.T) {
	print(IntToString(int(100)))
	print(IntToString(int8(8)))
	print(IntToString(int32(16)))
	print(IntToString(int32(32)))
	print(IntToString(int64(64)))
}

func TestStringToInt(t *testing.T) {
	print2(StringToInt[int]("100"))
	print2(StringToInt[int8]("100"))
	print2(StringToInt[int16]("100"))
	print2(StringToInt[int32]("100"))
	print2(StringToInt[int64]("100"))
}

func TestStringToFloat(t *testing.T) {
	print2(StringToFloat[float32]("100.1111"))
	print2(StringToFloat[float32]("100.9999"))
	print2(StringToFloat[float64]("100.1111"))
	print2(StringToFloat[float64]("100.9999"))
}

func TestFloatToString(t *testing.T) {
	print(FloatToString(float32(100.1111), -1))
	print(FloatToString(float32(100.1111), 2))
	print(FloatToString(float32(100.9999), 2))
	print(FloatToString(float64(100.1111), -1))
	print(FloatToString(float64(100.1111), 2))
	print(FloatToString(float64(100.9999), 2))
}

func TestStringToIntSlice(t *testing.T) {
	print2(StringToIntSlice[int16]("1111,222,333", ","))
	print2(StringToIntSlice[int32]("1111,222,333", ","))
	print2(StringToIntSlice[int64]("1111,222,333", ","))
}

func print(i any) {
	fmt.Printf("%T, %v\n", i, i)
}

func print2(s any, err error) {
	fmt.Printf("%T, %v, err: %v\n", s, s, err)
}
