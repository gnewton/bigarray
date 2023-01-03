package alt

import (
// "log"
)

type Int32Serializer struct {
}

func (s *Int32Serializer) Serialize(i int32) ([]byte, error) {
	return uint32ToBytes(uint32(i)), nil
}

func (s *Int32Serializer) Deserialize(b []byte) (int32, error) {
	ui32, err := bytesToUint32(b)
	return int32(ui32), err
}

func (s *Int32Serializer) SizeOf() int {
	return 4
}
