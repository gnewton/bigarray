package bigarray

type BigArray[T any] interface {
	SetSerializer(Serializer[T]) error
	GetSerializer() Serializer[T]
	Put(index int64, val T) error
	Get(index int64) (T, error)
	Done() error
}

type Serializer[T any] interface {
	Deserialize([]byte) (T, error)
	Serialize(T) []byte
	SizeOf() int
}

/*
type BigArrayWriter[T any] interface {
	SetSerializer(Serializer[T])
	GetSerializer() Serializer[T]
	Put(index int64, val T)
}

type BigArrayReader[T any] interface {
	SetDeserializer(Serializer[T])
	GetDeserializer() Serializer[T]
	Get(index int64) T
}
*/

// BigArray implementations
// - Trivial: backed by T[]
// - Simple: backed by byte[]
// - Writes
