package alt

type Mode int8

const (
	RandomAccess Mode = iota
	OrderedAccess
)

type Sized int8

const (
	KnownSize Sized = iota
	UnknownSize
)

/*

	OrderedAccess/OrderedAccess UnknownSize: linear/mmap|*linear
       	OrderedAccess/OrderedAccess KnownSize  : mmap|*linear/mmap|*linear
        OrderedAccess/RandomAccess UnknownSize: linear/mmap
       	OrderedAccess/RandomAccess KnownSize  : mmap|*linear/mmap

       	RandomAccess/OrderedAccess KnownSize   : mmap/mmap|linear
       	RandomAccess/OrderedAccess UnknownSize : kv/linear
       	RandomAccess/RandomAccess KnownSize   : mmap/mmap
       	RandomAccess/RandomAccess UnknownSize : kv/mmap

*/

type Serializer[T any] interface {
	Serialize(T) ([]byte, error)
	SizeOf() int
}

type Deserializer[T any] interface {
	Deserialize([]byte) (T, error)
	SizeOf() int
}

type Writer[T any] interface {
	SetSerializer(Serializer[T]) error
	GetSerializer() Serializer[T]
	Write(int64, []byte) error
}

type Reader[T any] interface {
	SetDeserializer(Deserializer[T]) error
	GetDeserializer() Deserializer[T]
	Read(int64) ([]byte, error)
}

func Put[T any](w Writer[T], index int64, value *T) error {
	buf, err := w.GetSerializer().Serialize(*value)
	if err != nil {
		return err
	}
	return w.Write(index, buf)
}

func Get[T any](w Reader[T], index int64) (*T, error) {
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
