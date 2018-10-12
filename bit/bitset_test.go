package bit

import (
	"testing"
)

func TestBitArraySet(t *testing.T) {
	var (
		ba       *BitArray = NewBitArray(32)
		expected []uint8   = []uint8{0x83, 0x0, 0x83, 0x80}
	)
	if ba.size != 32 {
		t.Errorf("expected ba.size==32, got %d.", ba.size)
	}
	if ok := ba.Set(17); !ok {
		t.Errorf("expected Set(17)==true; got %t.", ok)
	}
	if ok := ba.Set(33); ok {
		t.Errorf("expected Set(33)==false; got %t.", ok)
	}
	// set bits in first bin
	_ = ba.Set(1)
	_ = ba.Set(2)
	_ = ba.Set(8)
	// set bits in second bin
	_ = ba.Set(24)
	_ = ba.Set(18)
	// set bits in fourth bin
	_ = ba.Set(32)
	for i := 0; i < len(expected); i++ {
		if expected[i] != ba.bits[i] {
			t.Errorf("assertion failed; iteration(%d), expected %d, got %d.", i, expected[i], ba.bits[i])
		}
	}
	// t.Log(ba.bits, ba.bits[2])
	// fmt.Printf("%#x\n", ba.bits[2])
	// fmt.Printf("%#x\n", ba.bits[0])
	// fmt.Println(ba)
}
