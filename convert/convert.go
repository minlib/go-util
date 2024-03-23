package convert

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/minlib/go-util/slicex"
	"golang.org/x/exp/constraints"
)

// IntToBytes convert int to bytes
func IntToBytes[E constraints.Integer](i E) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, int64(i))
	return buf.Bytes()
}

// BytesToInt convert bytes to int
func BytesToInt[E constraints.Integer](b []byte) E {
	buf := bytes.NewBuffer(b)
	var res int64
	binary.Read(buf, binary.BigEndian, &res)
	return E(res)
}

// StringToInt convert float to int
func StringToInt[E constraints.Integer](s string) (E, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return E(i), err
}

// IntToString convert int to float
func IntToString[E constraints.Integer](i E) string {
	return strconv.FormatInt(int64(i), 10)
}

// StringToFloat convert float to string
func StringToFloat[E constraints.Float](s string) (E, error) {
	f, err := strconv.ParseFloat(s, 64)
	return E(f), err
}

// FloatToString convert string to float
func FloatToString[E constraints.Float](f E, digits int) string {
	return strconv.FormatFloat(float64(f), 'f', digits, 64)
}

// StringSliceToIntSlice convert string slice to int slice
func StringSliceToIntSlice[E constraints.Integer](s []string) ([]E, error) {
	return slicex.StringToInt[E](s)
}

// IntSliceToStringSlice convert int slice to string slice
func IntSliceToStringSlice[S ~[]E, E constraints.Integer](s S) []string {
	return slicex.IntToString(s)
}

// StringToIntSlice convert string to int slice
func StringToIntSlice[E constraints.Integer](s, sep string) ([]E, error) {
	slice := strings.Split(s, sep)
	return StringSliceToIntSlice[E](slice)
}
