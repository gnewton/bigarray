package alt

import (
	//"errors"
	"log"
	"os"
	"testing"
)

// Shouldn't fail tests

func Test_InstantiateAllOKValues(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := tmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var r ReaderAt
	r, err = NewSequentialReader(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	r.Done()
}

func Test_Write1000Read1000ValuesInOrder(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// WRITE
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

	var nitems uint64 = 1000
	var v uint64
	for v = 0; v < nitems; v++ {
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

	// READ
	f, err = os.Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	var r ReaderAt
	r, err = NewSequentialReader(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	var index uint64
	for index = 0; index < nitems; index++ {
		v, err := Get(r, s, int64(index))
		if err != nil {
			log.Println(err)
			t.Fatal(err)
		}
		if err != nil {
			log.Println(v)
			t.Fatal(err)
		}
	}

}
