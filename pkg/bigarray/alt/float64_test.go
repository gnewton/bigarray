package alt

import (
	"fmt"
	"testing"
)

// Positive
func Test_Float64To_ZeroBytes(t *testing.T) {
	b := make([]byte, 8)

	i, err := bytesToFloat64(b)
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Fatal(fmt.Errorf("Should be zero: %s", haveNeed(int64(i), 0)))
	}
}

func Test_Float64From_ZeroBytes(t *testing.T) {
	b := float64ToBytes(float64(0))
	if len(b) != 8 {
		t.Fatal(fmt.Errorf("[]byte array wrong length: %s", haveNeed(int64(len(b)), 8)))
	}
	for i := 0; i < 8; i++ {
		if b[i] != 0 {
			t.Fatal(fmt.Errorf("[]byte array value incorrect at %d: %s", i, haveNeed(int64(b[i]), 0)))
		}
	}
}

func Test_Float64From_FullCycle(t *testing.T) {

	var i, start, end uint64
	start = 0
	end = 100000

	for i = start; i < end; i++ {
		// toBytes
		b := float64ToBytes(float64(i) + .5)
		if len(b) != 8 {
			t.Fatal(fmt.Errorf("[]byte wrong length: %s", haveNeed(int64(len(b)), 8)))
		}

		// Back again
		v, err := bytesToFloat64(b)
		if err != nil {
			t.Fatal(err)
		}
		if v != float64(i)+.5 {
			t.Fatal(fmt.Errorf("Conversion wrong %s", haveNeed(int64(v), int64(i))))
		}
	}
}

// Negative
func Test_Float64_BytesArgBadLength(t *testing.T) {
	b := make([]byte, 7)

	_, err := bytesToFloat64(b)
	if err == nil {
		t.Fatal(shouldBeError())
	}
}

func Test_Float64_BytesArgNil(t *testing.T) {
	var b []byte

	_, err := bytesToFloat64(b)
	if err == nil {
		t.Fatal(shouldBeError())
	}
}
