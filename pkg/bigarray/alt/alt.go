package alt

type Mode int8

const (
	RandomAccess Mode = iota
	SeqentialAccess
)

type NItems int8

const (
	Known Sized = iota
	Unknown
)

/*
        Write            / Read             NItems**        : Reader       / Writer
        -------------------------------------------------------------------------------
	SequentialAccess / SequentialAccess Unknown   : linear       / mmap|*linear
       	SequentialAccess / SequentialAccess Known     : mmap|*linear / mmap|*linear
        SequentialAccess / RandomAccess     Unknown   : linear       / mmap
       	SequentialAccess / RandomAccess     Known     : mmap|*linear / mmap

       	RandomAccess     / SequentialAccess Known     : mmap         / mmap|*linear
       	RandomAccess     / SequentialAccess Unknown   : kv           / kv
       	RandomAccess     / RandomAccess     Known     : mmap         / mmap
       	RandomAccess     / RandomAccess     Unknown   : kv           / kv

    * - Preferred implementation
   ** - Number of total items to be written, at the time of start of write

   Need: Writer: linear, mmap, kv
         Reader: linear, mmap, kv

   Idea: mmap chunker; array of mmaps, each a certain size (10,000 x SizeOf)
         Get(i int64): which mmap: a[i/10000]; where in this mmap? i%10000
         (for unknown array size)
*/

type Serializer[T any] interface {
	Serialize(T) ([]byte, error)
	SizeOf() int
}

type Deserializer[T any] interface {
	Deserialize([]byte) (T, error)
	SizeOf() int
}

type Writer interface {
	Write(int64, []byte) error
}

type Reader interface {
	Read(int64) ([]byte, error)
}

func Put[T any](w Writer, s Serializer[T], index int64, value *T) error {
	buf, err := s.Serialize(*value)
	if err != nil {
		return err
	}
	return w.Write(index, buf)
}

func Get[T any](w Reader, d Deserializer[T], index int64) (*T, error) {
	return nil, nil
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
