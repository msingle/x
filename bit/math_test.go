package bit

import "testing"

func TestMath(t *testing.T) {
	const n = 16
	var (
		alignBit   int = blocks.nbalgn(n)
		logN       int
		arrayIndex int
	)
	arrayIndex = blocks.bin(alignBit)
	logN = blocks.lgb2(uint64(arrayIndex))
	// output values
	t.Log(alignBit)
	t.Log(arrayIndex)
	t.Log(logN)
	t.Log(RoundPrevP2(32))
	t.Log(findslot(32, 8, 32))
}
