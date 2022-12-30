package bigarray

import (
	"github.com/gnewton/pkg/bigarray"
	"io"
	"log"
)

type Int64BigArraySimple struct {
	serializer Serializer[int64]
	store      [][]byte
	len        int
}

func (ba *Int64BigArraySimple) SetSerializer(s Serializer[int64]) error {
	ba.serializer = s
	return nil
}
func (ba *Int64BigArraySimple) GetSerializer() Serializer[int64] {
	return ba.serializer
}

func (ba *Int64BigArraySimple) Put(index int64, val int64) error {
	if int(index) > ba.len {
		log.Fatal("Index too large:", index, ba.len)
	}

	ba.store[index] = ba.serializer.Serialize(val)

	return nil
}

func (ba *Int64BigArraySimple) Get(index int64) (int64, error) {
	if index > int64(ba.len) {
		log.Fatal("Index too large:", index, ba.len)
	}
	buf := ba.store[index]

	return ba.serializer.Deserialize(buf)
}

func newInt64BigArraySimple(size int) BigArray[int64] {
	ba := new(Int64BigArraySimple)
	ba.store = make([][]byte, size)
	ba.len = size

	return ba
}

func writeInt64AsBytes(w io.Writer, v int64) error {
	return writeUint64AsBytes(w, uint64(v))
}

func readBytesAsInt64(r io.Reader) (int64, error) {
	ui64, err := readBytesAsUint64(r)

	return int64(ui64), err
}
