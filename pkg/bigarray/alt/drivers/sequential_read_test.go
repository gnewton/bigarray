package drivers

import (
	//"errors"
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray/alt"
	"github.com/gnewton/bigarray/pkg/bigarray/alt/serialize"
	"log"
	"os"
	"testing"
)

// Shouldn't fail tests

func Test_InstantiateAllOKValues(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := alt.TmpFile(".")
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}

	var r alt.ReaderAt
	r, err = NewSequentialReader(f, 8)
	if err != nil {
		t.Fatal(err)
	}
	r.Done()
}

func Test_Write1000Read1000ValuesInOrder(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// WRITE
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

	var nitems int64 = 1000
	var index int64

	for index = 0; index < nitems; index++ {
		v := uint64(index)
		err := alt.Put(w, s, index, &v)
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

	var r alt.ReaderAt
	r, err = NewSequentialReader(f, 8)
	if err != nil {
		t.Fatal(err)
	}

	for index = 0; index < nitems; index++ {
		v, err := alt.Get(r, s, int64(index))
		if err != nil {
			log.Println(err)
			t.Fatal(err)
		}
		if err != nil {
			log.Println(v)
			t.Fatal(err)
		}
		if *v != uint64(index) {
			t.Fatal(fmt.Errorf("Value incorrect: %s", alt.HaveNeed(int64(*v), index)))
		}
	}

}
