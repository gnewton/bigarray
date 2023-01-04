package drivers

import (
	"github.com/gnewton/bigarray/pkg/bigarray/alt"
	"github.com/gnewton/bigarray/pkg/bigarray/alt/serialize"
	"log"
	"os"
	"testing"
)

// Shouldn't fail tests
func Test_InstantiateReadAllOKValues(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := alt.TmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w alt.WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	w.Done()
}

func Test_WriteValuesInOrder(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := alt.TmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w alt.WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	var s alt.Serializer[uint64]
	s = new(serialize.Uint64Serializer)

	var v uint64
	for v = 0; v < 10000; v++ {
		err := alt.Put(w, s, int64(v), &v)
		if err != nil {
			log.Println(v)
			t.Fatal(err)
		}
	}

	err = w.Done()
	if err != nil {
		t.Fatal(err)
	}
}

// Should fail tests
func Test_WriteValues_OutOfOrder(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := alt.TmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w alt.WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	var s alt.Serializer[uint64]
	s = new(serialize.Uint64Serializer)
	var v uint64 = 123
	err = alt.Put(w, s, 1, &v)
	// Shouldn't fail
	if err == nil {
		t.Fatal(alt.ShouldBeError())
	}
}

func Test_Instantiate_BadSizeOf(t *testing.T) {
	f, err := alt.TmpFile(".")
	defer os.Remove(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewSequentialWriter(f, 0)
	if err == nil {
		t.Fatal(alt.ShouldBeError())
	}
}

func Test_Instantiate_NilFile(t *testing.T) {
	_, err := NewSequentialWriter(nil, 8)
	if err == nil {
		t.Fatal(alt.ShouldBeError())
	}
}
