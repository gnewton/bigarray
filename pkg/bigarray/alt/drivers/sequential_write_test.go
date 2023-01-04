package drivers

import (
	"log"
	"os"
	"testing"
)

// Shouldn't fail tests
func Test_InstantiateReadAllOKValues(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := tmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	w.Done()
}

func Test_WriteValuesInOrder(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := tmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	var s Serializer[uint64]
	s = new(Uint64Serializer)

	var v uint64
	for v = 0; v < 10000; v++ {
		err := Put(w, s, int64(v), &v)
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
	f, err := tmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var w WriterAt
	w, err = NewSequentialWriter(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	var s Serializer[uint64]
	s = new(Uint64Serializer)
	var v uint64 = 123
	err = Put(w, s, 1, &v)
	// Shouldn't fail
	if err == nil {
		t.Fatal(shouldBeError())
	}
}

func Test_Instantiate_BadSizeOf(t *testing.T) {
	f, err := tmpFile(".")
	defer os.Remove(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewSequentialWriter(f, 0)
	if err == nil {
		t.Fatal(shouldBeError())
	}
}

func Test_Instantiate_NilFile(t *testing.T) {
	_, err := NewSequentialWriter(nil, 8)
	if err == nil {
		t.Fatal(shouldBeError())
	}
}
