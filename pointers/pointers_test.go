package pointers

import (
	"fmt"
	"testing"
	"unsafe"
)

type Sample struct {
	value int
}

type SliceIFTest struct {
	nodes []interface{}
	sid   string
}

type SliceTest struct {
	nodes []unsafe.Pointer
	sid   string
	rdi   uint64
}

type ArrayTest struct {
	nodes [16]unsafe.Pointer
	sid   string
}

type tststore struct {
	counter uint64
	nodes   *SliceTest
}

type ststaligned struct {
	nodes                  []unsafe.Pointer // storage with capacity `size`, pow2
	wri, rdi, maxrdi, size uint64           // write, read, max-read and size (mask) indexes
	count                  uint64           // occupancy counter
	addrdata               uintptr
}

// - MARK: Test-structs section.

type tstsample struct {
	value int
}

type tstslice struct {
	nodes []unsafe.Pointer
	sid   string
}

func init() {
	const _archfmt = "========================\n%-6s%8d\n%-6s%12d\n%-6s%11d\n========================\n"
	fmt.Printf(_archfmt,
		"Address size", archADDRSIZE,
		"Tag size", archMAXTAG,
		"Word size", archWORDSIZE)
}

func TestAtomicArray(t *testing.T) {
	var (
		s    *Sample    = &Sample{value: 8}
		a    *ArrayTest = &ArrayTest{}
		sptr unsafe.Pointer
		sval *Sample
		ok   bool
	)
	if !SetArraySlot(unsafe.Pointer(&a.nodes), 6, archPTRSIZE, unsafe.Pointer(s)) {
		t.Fatal("inconsistent state, cannot set slot.")
	}
	fmt.Println(a.nodes)
	sptr, ok = PopArraySlot(unsafe.Pointer(&a.nodes), 6, archPTRSIZE)
	if !ok {
		t.Fatal("inconsistent state, cannot pop slot.")
	}
	sval = (*Sample)(sptr)
	if sval == nil {
		t.Fatal("inconsistent state.")
	} else if sval.value != 8 {
		t.Fatal("assertion failed.")
	}
	fmt.Println("Pointer(S) has Value(", sval.value, ")")
	if a.nodes[6] != nil {
		t.Fatal("inconsistent state.")
	}
}

func TestAtomicSliceWithInterface(t *testing.T) {
	var (
		s    *Sample      = &Sample{value: 8}
		a    *SliceIFTest = &SliceIFTest{nodes: make([]interface{}, 16)}
		sptr unsafe.Pointer
		sval *Sample
		ok   bool
	)
	if !SetSliceSlotI(unsafe.Pointer(&a.nodes), 0, sizeINTERFACE, unsafe.Pointer(s)) {
		t.Fatal("inconsistent state, cannot set slot.")
	}
	if a.nodes[0] == nil {
		t.Fatal("inconsistent state, value is not inserted.")
	} else {
		fmt.Println(a.nodes[2], " is the head value", a.nodes[1])
	}
	fmt.Println("nodes(iface):", a.nodes)
	sptr, ok = PopSliceSlot(unsafe.Pointer(&a.nodes), 0, sizeINTERFACE)
	if !ok {
		t.Fatal("inconsistent state, cannot pop slot.")
	}
	sval = *(**Sample)(sptr)
	if sval == nil {
		t.Fatal("inconsistent state.")
	} else if sval.value != 8 {
		t.Fatal("assertion failed.")
	}
	fmt.Println("Pointer(S) has Value(", sval.value, ")")
}

func TestAtomicSlice(t *testing.T) {
	var (
		s    *Sample    = &Sample{value: 8}
		a    *SliceTest = &SliceTest{nodes: make([]unsafe.Pointer, 16)}
		sptr unsafe.Pointer
		sval *Sample
		ok   bool
	)
	if !SetSliceSlot(unsafe.Pointer(&a.nodes), 0, archPTRSIZE, unsafe.Pointer(s)) {
		t.Fatal("inconsistent state, cannot set slot.")
	}
	if a.nodes[0] == nil {
		t.Fatal("inconsistent state, value is not inserted.")
	}
	fmt.Printf("nodes: %#x\n", a.nodes)
	sptr, ok = PopSliceSlot(unsafe.Pointer(&a.nodes), 0, archPTRSIZE)
	if !ok {
		t.Fatal("inconsistent state, cannot pop slot.")
	}
	fmt.Printf("nodes: %#x\n", a.nodes)
	sval = (*Sample)(sptr)
	if sval == nil {
		t.Fatal("inconsistent state.")
	} else if sval.value != 8 {
		t.Fatal("assertion failed.")
	}
	fmt.Println("Pointer(S) has Value(", sval.value, ")")
}

func TestCASTag(t *testing.T) {
	var (
		s    *Sample = &Sample{value: 8}
		sptr unsafe.Pointer
		ok   bool
		err  error
	)
	sptr, err = TaggedPointer(unsafe.Pointer(s), 1)
	if err != nil {
		t.Fatal("inconsistent state, expected err==nil.")
	}
	fmt.Println("sptr: ", sptr)
	if sptr, ok = CompareAndSwapPointerTag(sptr, 1, 2); sptr == nil || !ok {
		t.Fatal("inconsistent state, expected !nil, true.", sptr, ok)
	}
	fmt.Println("sptr: swap1->2 ", sptr)
	if tag := GetTag(sptr); tag != 2 {
		t.Fatal("assertion failed, expected 2. got ", tag)
	}
}

func TestRDCSS(t *testing.T) {
	var (
		expected = &tststore{
			counter: 16,
			nodes:   &SliceTest{make([]unsafe.Pointer, 0), "replaced", 0},
		}
		c = &tststore{
			counter: 16,
			nodes:   &SliceTest{make([]unsafe.Pointer, 0), "original", 0},
		}
		crdiptr unsafe.Pointer
	)
	crdiptr = unsafe.Pointer(&c.counter)
	if !RDCSS(
		(*unsafe.Pointer)(crdiptr),
		(unsafe.Pointer)(unsafe.Pointer(uintptr(expected.counter))),
		(*unsafe.Pointer)(unsafe.Pointer(&c.nodes)),
		unsafe.Pointer(c.nodes),
		unsafe.Pointer(expected.nodes),
	) {
		t.Fatal("Expected RDCSS to succeed")
	}
	if c.nodes == nil {
		t.Fatal("inconsistent state.")
	}
	if c.nodes.sid != "replaced" {
		t.Fatal("inconsistent state.", c.nodes)
	}
}

func TestRDCSS2(t *testing.T) {
	var (
		item  = &Sample{16}
		sl    [8]unsafe.Pointer
		slptr *[8]unsafe.Pointer
		n     *Sample
	)
	slptr = &sl
	if !SetArraySlot(unsafe.Pointer(slptr), 4, unsafe.Sizeof(slptr), unsafe.Pointer(item)) {
		t.Fatal("inconsistent state.")
	}
	nsptr, ok := PopArraySlot(unsafe.Pointer(slptr), 4, unsafe.Sizeof(slptr))
	if !ok {
		t.Fatal("inconsistent state.")
	}
	n = (*Sample)(nsptr)
	if n == nil {
		t.Fatal("inconsistent state.")
	} else if n.value != 16 {
		t.Fatal("inconsistent state.")
	}
}

func TestRDCSS2x(t *testing.T) {
	var (
		s       *tstsample     = &tstsample{value: 64}
		n       *tstsample     = &tstsample{value: 128}
		ring    *ststaligned   = &ststaligned{nodes: make([]unsafe.Pointer, 8)}
		slotptr unsafe.Pointer = unsafe.Pointer(OffsetSliceSlot(unsafe.Pointer(&ring.nodes), 1, archPTRSIZE))
		rdiPtr  unsafe.Pointer = unsafe.Pointer(&ring.rdi)
		vptr    **tstsample    // value pointer
	)
	if !SetSliceSlot(unsafe.Pointer(&ring.nodes), 1, archPTRSIZE, unsafe.Pointer(&s)) {
		t.Fatal("inconsistent state, can't write to slice/slot.")
	}
	if ring.nodes[1] == nil {
		t.Fatal("assertion failed, expected non-nil.")
	}
	vptr = (**tstsample)(ring.nodes[1])
	if vptr == nil {
		t.Fatal("assertion failed, expected non-nil.")
	} else if *vptr == nil {
		t.Fatal("assertion failed, expected non-nil.")
	}
	if (*vptr) != s {
		t.Fatal("assertion failed, expected equal as 's' is written to slot 1.")
	}
	if (((*vptr).value != s.value) && s.value == 64) || s.value != 64 {
		t.Fatal("assertion failed, invalid pointers.")
	}
	fmt.Printf("nodes: %#x\n", ring.nodes)
	if !RDCSS(
		(*unsafe.Pointer)(unsafe.Pointer(&rdiPtr)),
		(unsafe.Pointer)(unsafe.Pointer(rdiPtr)),
		(*unsafe.Pointer)(unsafe.Pointer(slotptr)),
		(unsafe.Pointer)(unsafe.Pointer(&s)),
		(unsafe.Pointer)(unsafe.Pointer(n)),
	) {
		t.Fatal("inconsistent state, RDCSS failure.")
	}
	if s == nil {
		t.Fatal("assertion failed, fatal condition from incorrect pointer mutation.")
	} else if s.value != 64 {
		t.Fatal("assertion failed, invalid pointer.")
	}
	fmt.Printf("nodes: %#x\n", ring.nodes)
	if ring.nodes[1] == nil {
		t.Fatal("assertion failed, expected non-nil for occupied slot.")
	}
	if s == nil {
		t.Fatal("assertion failed, fatal condition from incorrect pointer mutation.")
	}
	vptr = (**tstsample)(slotptr)
	if vptr == nil {
		t.Fatal("assertion failed, expected non-nil pointer.")
	} else if *vptr == nil {
		t.Fatal("assertion failed, expected non-nil pointer.")
	}
	if (*vptr) != n {
		t.Fatal("assertion failed, expected equal as 's' is swapped with 'n'.")
	}
	if (((*vptr).value != n.value) && n.value == 128) || n.value != 128 {
		t.Fatal("assertion failed, invalid pointers.")
	}
	fmt.Println("slotptr:", *(**tstsample)(slotptr), "addr:", slotptr, &n, ring.nodes, ring.nodes[1])
}
