package bit

type blktable []int

var (
	mDeBruijnBitPosition [32]int = [32]int{
		0, 9, 1, 10, 13, 21, 2, 29, 11, 14, 16, 18, 22, 25, 3, 30,
		8, 12, 20, 28, 15, 17, 24, 7, 19, 27, 23, 6, 26, 5, 4, 31,
	}

	blocks blktable = blktable{
		8, 8, 8,
		16, 32, 64,
		128, 256, 512,
		1024, 2048, 4096, 8192,
		16384, 32768, 65536, 131072,
		262144, 524288, 1048576, 2097152,
		4194304, 8388608, 16777216,
		33554432, 67108864, 134217728,
		268435456, 536870912, 1073741824,
		2147483648, 4294967296,
	}
)

func (b blktable) bin(num int) int {
	return int(b.lgb2(uint64(num)))
}

func (b blktable) size(num int) int {
	return b[int(b.bin(num))]
}

func (b blktable) lgb2(num uint64) int {
	return mDeBruijnBitPosition[int(uint64(uint32(RoundPrevP2(num)-1)*(uint32)(0x07C4ACDD))>>27)]
}

func (b blktable) nbalgn(num uint64) int {
	if num == 0 {
		return blocks[0]
	}
	return blocks[b.lgb2(num)]
}

func RoundNextP2(num uint64) uint64 {
	num--
	num |= num >> 1
	num |= num >> 2
	num |= num >> 4
	num |= num >> 8
	num |= num >> 16
	num |= num >> 32
	num++
	return num
}

func RoundPrevP2(num uint64) uint64 {
	num--
	num |= num >> 1
	num |= num >> 2
	num |= num >> 4
	num |= num >> 8
	num |= num >> 16
	num |= num >> 32
	num++
	return num - (num >> 1)
}

func NAlignBits(num uint64) int {
	return blocks.nbalgn(num)
}

func Lgb2(num uint64) int {
	return blocks.lgb2(num)
}

// DEV SECTION

func findslot(n int, wsize int, capacity int) int {
	return ((n * wsize) / capacity) / 2
}
