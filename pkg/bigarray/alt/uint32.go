package alt

import (
	"encoding/binary"
	"fmt"
	//"log"
)

type Uint32Serializer struct {
}

func (s *Uint32Serializer) Serialize(i uint32) ([]byte, error) {
	return uint32ToBytes(i), nil
}

func (s *Uint32Serializer) Deserialize(b []byte) (uint32, error) {
	return bytesToUint32(b)
}

func (s *Uint32Serializer) SizeOf() int {
	return 4
}

func uint32ToBytes(i uint32) []byte {
	b := make([]byte, 32)
	binary.LittleEndian.PutUint32(b, i)
	return b
}

func bytesToUint32(b []byte) (uint32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("[]byte wrong length: %s", haveNeed(int64(len(b)), int64(4)))
	}
	return binary.LittleEndian.Uint32(b), nil
}
