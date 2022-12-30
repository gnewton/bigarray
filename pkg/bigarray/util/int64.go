package util

import (
	//"encoding/binary"
	"fmt"
	"io"
)

///////////////////////////
// Int64
type Int64Serializer struct {
	w io.Writer
	r io.Reader
}

func (m *Int64Serializer) Serialize(i int64) []byte {
	return encodeInt64ToBytes(i)
}

func (m *Int64Serializer) Deserialize(buf []byte) (int64, error) {
	l := len(buf)
	if l == 0 {
		return 0, fmt.Errorf("Buff cannot be zero")
	}
	if l != m.SizeOf() {
		return 0, fmt.Errorf("Buff wrong size: want %d; have %d", l, m.SizeOf())
	}
	return decodeInt64ToBytes(buf)
}

func (m *Int64Serializer) SizeOf() int {
	return 8
}

func encodeInt64ToBytes(x int64) []byte {
	return encodeUint64ToBytes(uint64(x))
}

func decodeInt64ToBytes(buf []byte) (int64, error) {
	ui64, err := decodeUint64ToBytes(buf)
	return int64(ui64), err
}
