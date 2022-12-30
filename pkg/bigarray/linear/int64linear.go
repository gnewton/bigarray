package linear

import (
	"bufio"
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray"
	"os"
)

// Simple, array-backed store, primarily for testing/validation

type Int64BigArrayLinear struct {
	serializer bigarray.Serializer[int64]
	index      int64
	file       *os.File
	filename   string
	writer     *bufio.Writer
	frozen     bool
	len        int
}

func (ba *Int64BigArrayLinear) SetSerializer(s bigarray.Serializer[int64]) error {
	if s == nil {
		return fmt.Errorf("Int64BigArrayLinear.SetSerializer: Serializer is nil")
	}
	ba.serializer = s
	return nil
}
func (ba *Int64BigArrayLinear) GetSerializer() bigarray.Serializer[int64] {
	return ba.serializer
}

func (ba *Int64BigArrayLinear) Put(index int64, val int64) error {
	if index < 0 {
		return fmt.Errorf("Index < 0; =%d", index)
	}
	if int(index) > ba.len {
		return fmt.Errorf("Index too large. Have %d; need %d", index, ba.len)
	}

	return nil
}

func (ba *Int64BigArrayLinear) Get(index int64) (int64, error) {
	if index > int64(ba.len) {
		return 0, fmt.Errorf("Index too large. Have %d; need %d", index, ba.len)
	}

	return 0, nil
}

func newInt64BigArrayLinear(size int) bigarray.BigArray[int64] {
	ba := new(Int64BigArrayLinear)
	ba.len = size

	return ba
}
