package convert

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	b := IntToBytes(-111)
	fmt.Printf("%T, %v\n", b, b) // []uint8, [255 255 255 255 255 255 255 145]
}

func TestBytesToInt2(t *testing.T) {
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
