package convert

import (
	"bytes"
	"encoding/binary"

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
