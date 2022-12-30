package util

import (
	"encoding/binary"
	"fmt"
	"io"
)

func encodeUint64ToBytes(x uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, x)
	return buf
}

func writeUint64AsBytes(w io.Writer, v uint64) error {
	if w == nil {
		return fmt.Errorf("writeUint64AsBytes: Writer is nil")
	}
	b := encodeUint64ToBytes(v)

	n, err := w.Write(b)
	if err != nil {
		return err
	}
	if n != 8 {
		return fmt.Errorf("writeUint64AsBytes: Not enough bytes written. Have %d; need 8", n)
	}
	return nil
}

func readBytesAsUint64(r io.Reader) (uint64, error) {
	if r == nil {
		return 0, fmt.Errorf("readBytesAsUint64: Reader is nil")
	}
	buf := make([]byte, 8)
	n, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, fmt.Errorf("readBytesAsUint64: Not enough bytes read. Have %d; need 8", n)
	}

	return binary.LittleEndian.Uint64(buf), nil
}

func decodeUint64ToBytes(buf []byte) (uint64, error) {
	if len(buf) != 8 {
		return 0, fmt.Errorf("Buff too short: have %d; need %d", len(buf), 8)
	}
	return binary.LittleEndian.Uint64(buf), nil
}
