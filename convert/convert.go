package convert

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/exp/constraints"
	"strconv"
)

// IntToBytes convert int to bytes
func IntToBytes[E constraints.Integer](i E) []byte {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.BigEndian, int64(i))
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

// BytesToInt convert bytes to int
func BytesToInt[E constraints.Integer](b []byte) E {
	buf := bytes.NewBuffer(b)
	var res int64
	err := binary.Read(buf, binary.BigEndian, &res)
	if err != nil {
		return 0
	}
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
