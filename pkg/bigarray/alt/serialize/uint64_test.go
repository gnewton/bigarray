package serialize

import (
	"fmt"
	"github.com/gnewton/bigarray/pkg/bigarray/alt"
	"testing"
)

// Positive
func Test_Uint64To_ZeroBytes(t *testing.T) {
	b := make([]byte, 8)

	i, err := bytesToUint64(b)
	if err != nil {
		t.Fatal(err)
	}
	if i != 0 {
		t.Fatal(fmt.Errorf("Should be zero: %s", alt.HaveNeed(int64(i), 0)))
	}
}

func Test_Uint64From_ZeroBytes(t *testing.T) {
	b := uint64ToBytes(uint64(0))
	if len(b) != 8 {
		t.Fatal(fmt.Errorf("[]byte array wrong length: %s", alt.HaveNeed(int64(len(b)), 8)))
	}
	for i := 0; i < 8; i++ {
		if b[i] != 0 {
			t.Fatal(fmt.Errorf("[]byte array value incorrect at %d: %s", i, alt.HaveNeed(int64(b[i]), 0)))
		}
	}
}

func Test_Uint64From_FullCycle(t *testing.T) {

	var i, start, end uint64
	start = 0
	end = 100000

	for i = start; i < end; i++ {
		// toBytes
		b := uint64ToBytes(i)
		if len(b) != 8 {
			t.Fatal(fmt.Errorf("[]byte wrong length: %s", alt.HaveNeed(int64(len(b)), 8)))
		}

		// Back again
		v, err := bytesToUint64(b)
		if err != nil {
			t.Fatal(err)
		}
		if v != i {
			t.Fatal(fmt.Errorf("Conversion wrong %s", alt.HaveNeed(int64(v), int64(i))))
		}
	}
}

// Negative
func Test_Uint64_BytesArgBadLength(t *testing.T) {
	b := make([]byte, 7)

	_, err := bytesToUint64(b)
	if err == nil {
		t.Fatal(alt.ShouldBeError())
	}
}

func Test_Uint64_BytesArgNil(t *testing.T) {
	var b []byte

	_, err := bytesToUint64(b)
	if err == nil {
		t.Fatal(alt.ShouldBeError())
	}
}
