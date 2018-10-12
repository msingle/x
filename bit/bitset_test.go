package bit

import (
	"fmt"
	"testing"
)

func TestBitArray(t *testing.T) {
	var (
		ba       *BitArray = NewBitArray(32)
		expected []uint8   = []uint8{0x83, 0x0, 0x83, 0x80}
		setter   []int     = []int{1, 2, 8, 17, 18, 24, 32}
	)
	if ba == nil {
		t.Error("unable to initialize.")
	}
	if ba.size != 32 {
		t.Errorf("expected ba.size==32, got %d.", ba.size)
	}
	// set bits
	for _, s := range setter {
		if err := ba.Set(s); err != nil {
			t.Errorf("expected Set(%d)==true; got %v.", s, err)
		}
	}
	// validity check
	for i := 0; i < len(expected); i++ {
		if expected[i] != ba.bits[i] {
			t.Errorf("assertion failed; iteration(%d), expected %d, got %d.", i, expected[i], ba.bits[i])
		}
	}
	if err := ba.Set(33); err == nil {
		t.Errorf("expected Set(33)!=nil; got %v.", err)
	}
	if _, err := ba.Get(33); err == nil {
		t.Errorf("expected Get(33)!=nil; got %v.", err)
	}
	if value, err := ba.Get(8); err != nil {
		t.Errorf("expected Get(8)!=nil; value %d, got %v.", value, err)
	}
	if err := ba.Flip(15); err != nil {
		t.Errorf("expected Flip(15)==nil; got %v.", err)
	}
	if ok, err := ba.IsSet(15); ok && err != nil {
		t.Errorf("expected IsSet(15)==true, nil; got %t, %v.", ok, err)
	}
}

func TestBitArray2(t *testing.T) {
	var (
		ba *BitArray = NewBitArray(16)
	)
	if ba == nil {
		t.Error("unable to initialize.")
	}
	if ba.size != 16 {
		t.Errorf("expected ba.size==32, got %d.", ba.size)
	}
	fmt.Println(ba.Set(1))
	fmt.Println(ba.Set(2))
	fmt.Println(ba.Set(16))
	fmt.Println(ba)
}
