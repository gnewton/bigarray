package mem

import (
	"github.com/gnewton/bigarray/pkg/bigarray"
	"github.com/gnewton/bigarray/pkg/bigarray/util"
	"testing"
)

func Test_Write1Read1(t *testing.T) {
	var ba bigarray.BigArray[int64]
	size := 20
	index := int64(1)
	value := int64(100)

	ba = newInt64BigArrayMem(size)
	serializer := new(util.Int64Serializer)
	ba.SetSerializer(serializer)

	err := ba.Put(index, value)
	if err != nil {
		t.Fatal(err)
	}
	v, err := ba.Get(index)
	if err != nil {
		t.Fatal(err)
	}
	if v != value {
		t.Fatalf("Fail %d", 23)
	}
}

func Test_IndexTooLarge(t *testing.T) {
	var ba bigarray.BigArray[int64]
	//var foo [8589934592]int

	size := 20
	index := int64(101)
	value := int64(42)

	ba = newInt64BigArrayMem(size)
	serializer := new(util.Int64Serializer)
	ba.SetSerializer(serializer)

	err := ba.Put(index, value)
	if err == nil {
		t.Log(err)
		t.Fatalf("Should fail: index too large")
	}
}
