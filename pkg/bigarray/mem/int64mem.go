package mem

import (
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray"
)

// Simple, array-backed store, primarily for testing/validation

type Int64BigArrayMem struct {
	serializer bigarray.Serializer[int64]
	store      [][]byte
	len        int
}

func (ba *Int64BigArrayMem) SetSerializer(s bigarray.Serializer[int64]) error {
	if s == nil {
		return fmt.Errorf("Int64BigArrayMem.SetSerializer: Serializer is nil")
	}
	ba.serializer = s
	return nil
}
func (ba *Int64BigArrayMem) GetSerializer() bigarray.Serializer[int64] {
	return ba.serializer
}

func (ba *Int64BigArrayMem) Put(index int64, val int64) error {
	if index < 0 {
		return fmt.Errorf("Index < 0; =%d", index)
	}
	if int(index) > ba.len {
		return fmt.Errorf("Index too large. Have %d; need %d", index, ba.len)
	}

	ba.store[index] = ba.serializer.Serialize(val)

	return nil
}

func (ba *Int64BigArrayMem) Get(index int64) (int64, error) {
	if index > int64(ba.len) {
		return 0, fmt.Errorf("Index too large. Have %d; need %d", index, ba.len)
	}
	buf := ba.store[index]

	return ba.serializer.Deserialize(buf)
}

func newInt64BigArrayMem(size int) bigarray.BigArray[int64] {
	ba := new(Int64BigArrayMem)
	ba.store = make([][]byte, size)
	ba.len = size

	return ba
}
