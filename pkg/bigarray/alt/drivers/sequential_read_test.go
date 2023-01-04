package drivers

import (
	//"errors"
	"fmt"
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

	var nitems int64 = 1000
	var index int64

	for index = 0; index < nitems; index++ {
		v := uint64(index)
		err := Put(w, s, index, &v)
		if err != nil {
			log.Println(index)
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
		if *v != uint64(index) {
			t.Fatal(fmt.Errorf("Value incorrect: %s", haveNeed(int64(*v), index)))
		}
	}

}
