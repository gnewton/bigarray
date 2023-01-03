package alt

import (
// "log"
)

type Int64Serializer struct {
}

func (s *Int64Serializer) Serialize(i int64) ([]byte, error) {
	return uint64ToBytes(uint64(i)), nil
}

func (s *Int64Serializer) Deserialize(b []byte) (int64, error) {
	ui64, err := bytesToUint64(b)
	return int64(ui64), err
}

func (s *Int64Serializer) SizeOf() int {
	return 8
}
