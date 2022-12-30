package linear

import (
	"bufio"
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray"
	"github.com/gnewton/bigarray/pkg/bigarray/util"
	"log"
	"os"
)

// Simple, array-backed store, primarily for testing/validation

type Int64BigArrayLinear struct {
	serializer bigarray.Serializer[int64]
	index      int64
	file       *os.File
	filename   string
	writer     *bufio.Writer
	reader     *bufio.Reader
	frozen     bool
	len        int
}

func newInt64BigArrayLinear(tmpdir string) (bigarray.BigArray[int64], error) {
	if len(tmpdir) == 0 {
		tmpdir = "."
	}
	ba := new(Int64BigArrayLinear)
	ba.frozen = false

	var err error
	ba.file, err = os.CreateTemp(tmpdir, "big_")
	if err != nil {
		return nil, err
	}

	fileInfo, err := ba.file.Stat()
	if err != nil {
		return nil, err
	}
	ba.filename = fileInfo.Name()
	log.Println("Opened: ", ba.filename)

	ba.writer = bufio.NewWriter(ba.file)

	return ba, nil
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
	if ba.frozen {
		return fmt.Errorf("Attempting to put after frozen; index=%d  value=%d", index, val)
	}
	if index < 0 {
		return fmt.Errorf("Index < 0; =%d", index)
	}
	if index != ba.index {
		return fmt.Errorf("Index does not match next in linear: Have %d; need %d", index, ba.index)
	}

	ba.index++
	return util.WriteInt64AsBytes(ba.writer, val)

}

func (ba *Int64BigArrayLinear) Get(index int64) (int64, error) {
	if !ba.frozen {
		ba.frozen = true
		ba.index = 0
		err := ba.writer.Flush()
		if err != nil {
			return -1, err
		}
		// close+flush ba.writer
		// open ba.reader
		ba.reader = bufio.NewReader(ba.file)
	}

	if index != ba.index {
		return -1, fmt.Errorf("Index does not match next in linear: Have %d; need %d", index, ba.index)
	}
	ba.index++
	return util.ReadBytesAsInt64(ba.reader)
}

func (ba *Int64BigArrayLinear) Done() error {
	// close all; remove tmp file

	return nil
}
