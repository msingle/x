package bit

import (
	"errors"
	"fmt"
)

// Constants
const (
	// word size
	cWSIZE int = 8
)

// Error messages
var (
	IndexError error = errors.New("BitArray: index out of range")
)

type BitArray struct {
	bits []uint8
	size uint64
}

func NewBitArray(n uint64) (ba *BitArray) {
	ba = &BitArray{
		size: RoundNextP2(n),
	}
	ba.bits = make([]uint8, ba.size/uint64(cWSIZE))
	return ba
}

func (ba *BitArray) Set(bit int) error {
	var (
		mask uint8 // bit mask
		slot int   // index in slice
		item uint8 // bit from slice
	)
	if uint64(bit) > ba.size {
		goto ERROR
	}
	bit -= 1
	mask, slot, item = ba.access(bit)
	item |= mask
	ba.bits[slot] = item
	return nil
ERROR:
	return IndexError
}

func (ba *BitArray) Flip(bit int) error {
	var (
		mask uint8 // bit mask
		slot int   // index in slice
		item uint8 // bit from slice
	)
	if uint64(bit) > ba.size {
		goto ERROR
	}
	bit -= 1
	mask, slot, item = ba.access(bit)
	ba.bits[slot] = ^((^item & 0xFF) ^ mask)
	return nil
ERROR:
	return IndexError
}

func (ba *BitArray) Get(bit int) (uint8, error) {
	var (
		mask uint8 // bit mask
		item uint8 // bit from slice
	)
	if uint64(bit) > ba.size {
		goto ERROR
	}
	bit -= 1
	mask, _, item = ba.access(bit)
	return (item & mask), nil
ERROR:
	return 0x0, IndexError
}

func (ba *BitArray) IsSet(bit int) (ok bool, err error) {
	var (
		value uint8 // bit value
	)
	value, err = ba.Get(bit)
	if err != nil {
		return false, err
	}
	return value != 0, nil
}

func (ba *BitArray) access(bit int) (uint8, int, uint8) {
	var (
		mask uint8 // bit mask
		slot int   // index in slice
		item uint8 // bit from slice
	)
	mask = 0x1 << uint8(bit%cWSIZE)
	slot = ba.findslot(bit)
	fmt.Println("access", bit, mask, slot, item)
	item = ba.bits[slot]
	return mask, slot, item
}

func (ba *BitArray) findslot(n int) int {
	return ((n + cWSIZE - (n % cWSIZE)) / cWSIZE) - 1
}
