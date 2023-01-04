package alt

import (
	"fmt"
	//	"log"
)

type ReadAccess int8
type WriteAccess int8

const (
	RandomReadAccess ReadAccess = iota
	SeqentialReadAccess
)

const (
	RandomWriteAccess WriteAccess = iota
	SeqentialWriteAccess
)

type Cardinality int8

const (
	Unknown Cardinality = iota
	Known
)

// Local WriterAt, ReaderAt as here we cannot support all of io.WriterAt/ReaderAt guarantees
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
	Done() error
}

type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
	Done() error
}

/*
        Write            / Read             NItems**        : Reader       / Writer
        -------------------------------------------------------------------------------
	SequentialAccess / SequentialAccess Unknown   : linear       / mmap|*linear
       	SequentialAccess / SequentialAccess Known     : mmap|*linear / mmap|*linear
        SequentialAccess / RandomAccess     Unknown   : linear       / mmap
       	SequentialAccess / RandomAccess     Known     : mmap|*linear / mmap

       	RandomAccess     / SequentialAccess Known     : mmap         / mmap|*linear
       	RandomAccess     / SequentialAccess Unknown   : kv|MC           / kv
       	RandomAccess     / RandomAccess     Known     : mmap         / mmap
       	RandomAccess     / RandomAccess     Unknown   : kv|MC           / kv|MC

    * - Preferred implementation
   ** - Number of total items to be written, at the time of start of write

   Need: Writer: linear, mmap, kv, MC
         Reader: linear, mmap, kv, MC

   Idea: mmap chunker (MC); array of mmaps, each a certain size (10,000 x SizeOf)
         Get([]buf, i int64): which mmap: a[i/(10000*sizeOf)]; where in this mmap? i%(10000*sizeOf)
         (for unknown array size)
*/

type Serializer[T any] interface {
	Serialize(T) ([]byte, error)
	Deserialize([]byte) (T, error)
	SizeOf() int
}

func Put[T any](w WriterAt, s Serializer[T], index int64, value *T) error {
	buf, err := s.Serialize(*value)
	if err != nil {
		return err
	}

	if len(buf) != s.SizeOf() {
		return fmt.Errorf("Serialize byte array wrong length; have: %d; need: %d", len(buf), s.SizeOf())
	}

	n, err := w.WriteAt(buf, index)
	if err != nil {
		return err
	}
	if n != s.SizeOf() {
		return fmt.Errorf("Wrong # bytes written: have %d; want %d", n, s.SizeOf())
	}
	return err
}

func Get[T any](r ReaderAt, s Serializer[T], index int64) (*T, error) {
	sizeOf := s.SizeOf()
	buf := make([]byte, sizeOf)
	n, err := r.ReadAt(buf, index)
	if err != nil {
		return nil, err
	}
	if n != sizeOf {
		return nil, fmt.Errorf("Wrong number of bytes read: %s", HaveNeed(int64(n), int64(sizeOf)))
	}

	v, err := s.Deserialize(buf)

	return &v, err

}

/*
type ThisWriter struct {
}

func (w *ThisWriter) SetSerializer(Serializer[int64]) error {
	return nil
}
func (w *ThisWriter) GetSerializer() Serializer[int64] {
	return nil
}
func (w *ThisWriter) Write(int64, []byte) error {
	return nil
}
*/
