package bit

const (
	// word size
	cWSIZE int = 8
)

type BitArray struct {
	bits []uint8
	size uint64
}

// Get
// Flip
// IsSet

func NewBitArray(n uint64) (ba *BitArray) {
	ba = &BitArray{size: RoundNextP2(n)}
	ba.bits = make([]uint8, ba.size/uint64(cWSIZE))
	return ba
}

func (ba *BitArray) Set(bit int) bool {
	var (
		mask uint8 // bit mask
		slot int   // index in slice
		item uint8 // bit from slice
	)
	if uint64(bit) > ba.size {
		goto ERROR
	}
	bit -= 1
	mask = 0x1 << uint8(bit%cWSIZE)
	slot = ba.findslot(bit)
	item = ba.bits[slot]
	item |= mask
	ba.bits[slot] = item
	return true
ERROR:
	return false
}

// func (ba *BitArray) Get(bit int) (bit uint8) {
// 	var (
// 		mask uint8 = 0x1 << (bit % cWSIZE) // bit mask
// 		slot int   = ba.findslot(bit)      // index in slice
// 		item uint8 = ba.bits[slot]         // bit from slice
// 	)
//   return item &
// }

func (ba *BitArray) findslot(n int) int {
	return ((n * cWSIZE) / int(ba.size)) / 2
}

// DEV SECTION

func findslot(n int, wsize int, capacity int) int {
	return ((n * wsize) / capacity) / 2
}
