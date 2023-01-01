package alt

import (
	"encoding/binary"
	"fmt"
)

// uint64
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

// int64
func int64ToBytes(i int64) []byte {
	return uint64ToBytes(uint64(i))
}

func bytesToInt64(b []byte) (uint64, error) {
	return bytesToUint64(b)

}
