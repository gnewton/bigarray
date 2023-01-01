package alt

import (
	"errors"
	"testing"
)

// Positive
func Test_Uint64_OK(t *testing.T) {
	b := make([]byte, 8)

	_, err := bytesToUint64(b)
	if err != nil {
		t.Fatal(err)
	}
}

// Negative
func Test_Uint64_BadLength(t *testing.T) {
	b := make([]byte, 7)

	_, err := bytesToUint64(b)
	if err == nil {
		t.Fatal(errors.New("Should be nil"))
	}
}
