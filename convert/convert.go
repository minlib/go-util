package convert

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"golang.org/x/exp/constraints"
)

func IntToBytes[E constraints.Integer](i E) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, int64(i))
	return buf.Bytes()
}

func BytesToInt[E constraints.Integer](b []byte) E {
	buf := bytes.NewBuffer(b)
	var res int64
	binary.Read(buf, binary.BigEndian, &res)
	return E(res)
}

func IntToString[E constraints.Integer](i E) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt[E constraints.Integer](s string) (E, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return E(i), err
}

func FloatToString[E constraints.Float](f E, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}

func StringToFloat[E constraints.Float](s string) (E, error) {
	f, err := strconv.ParseFloat(s, 64)
	return E(f), err
}
