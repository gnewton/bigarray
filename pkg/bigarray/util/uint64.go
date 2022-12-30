package util

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func encodeUint64ToBytes(x uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, x)
	return buf
}

func writeUint64AsBytes(w io.Writer, v uint64) error {
	b := encodeUint64ToBytes(v)
	n, err := w.Write(b)
	if n != 8 {
		log.Fatal("Not all bytes written")
	}
	return err
}

func readBytesAsUint64(r io.Reader) (uint64, error) {
	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if n != 8 {
		log.Fatal("Not all bytes read")
	}
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf), nil
}

func decodeUint64ToBytes(buf []byte) (uint64, error) {
	if len(buf) != 8 {
		return 0, fmt.Errorf("Buff too short: have %d; need %d", len(buf), 8)
	}
	return binary.LittleEndian.Uint64(buf), nil
}
