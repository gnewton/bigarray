package alt

import (
	"errors"
	"log"
	"os"
	"testing"
)

// Positive tests
func Test_InstantiateAllOKValues(t *testing.T) {
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

// Negative tests
func Test_Instantiate_BadSizeOf(t *testing.T) {
	f, err := tmpFile(".")
	defer os.Remove(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewSequentialWriter(f, 0)
	if err == nil {
		t.Fatal(err)
	}
}

func Test_Instantiate_NilFile(t *testing.T) {
	_, err := NewSequentialWriter(nil, 8)
	if err == nil {
		t.Fatal(errors.New("Should be nil"))
	}
}
