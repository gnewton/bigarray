package alt

import (
	"encoding/binary"
	"fmt"
	//"log"
)

type Uint64Serializer struct {
}

func (s *Uint64Serializer) Serialize(i uint64) ([]byte, error) {
	return uint64ToBytes(i), nil
}

func (s *Uint64Serializer) Deserialize(b []byte) (uint64, error) {
	return bytesToUint64(b)
}

func (s *Uint64Serializer) SizeOf() int {
	return 8
}

func uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func bytesToUint64(b []byte) (uint64, error) {
	if len(b) != 8 {
		return 0, fmt.Errorf("[]byte wrong length: %s", haveNeed(len(b), 8))
	}
	return binary.LittleEndian.Uint64(b), nil
}
