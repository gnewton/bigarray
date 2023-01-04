package serialize

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray/alt"
	"math"
	//"log"
)

type Float64Serializer struct {
}

func (s *Float64Serializer) Serialize(i uint64) ([]byte, error) {
	return uint64ToBytes(i), nil
}

func (s *Float64Serializer) Deserialize(b []byte) (float64, error) {
	return bytesToFloat64(b)
}

func (s *Float64Serializer) SizeOf() int {
	return 8
}

// See https://stackoverflow.com/questions/43693360/convert-float64-to-byte-array

func float64ToBytes(f float64) []byte {
	var b [8]byte
	// b := make([]byte, 8)

	// n := math.Float64bits(f)
	// b[0] = byte(n >> 56)
	// b[1] = byte(n >> 48)
	// b[2] = byte(n >> 40)
	// b[3] = byte(n >> 32)
	// b[4] = byte(n >> 24)
	// b[5] = byte(n >> 16)
	// b[6] = byte(n >> 8)
	// b[7] = byte(n)

	// return b

	binary.LittleEndian.PutUint64(b[:], math.Float64bits(f))
	return b[:]

}

func bytesToFloat64(b []byte) (float64, error) {
	if b == nil {
		return 0, errors.New("[]byte is nil")
	}
	if len(b) != 8 {
		return 0, fmt.Errorf("[]byte wrong length: %s", alt.HaveNeed(int64(len(b)), 8))
	}
	bits := binary.LittleEndian.Uint64(b)
	return math.Float64frombits(bits), nil
}
