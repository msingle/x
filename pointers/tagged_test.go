package pointers

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestTag(t *testing.T) {
	const sval = 8
	var (
		s    *Sample = &Sample{value: sval}
		sptr unsafe.Pointer
		tag  uint
		err  error
	)
	sptr, err = TaggedPointer(unsafe.Pointer(s), 1)
	if err != nil {
		t.Fatal("assertion failed, expected==nil.")
	}
	s = (*Sample)(sptr)
	tag = GetTag(sptr)
	fmt.Println("Tag: ", tag)
	sptr = Untag(unsafe.Pointer(s))
	s = (*Sample)(sptr)
	if s == nil || s.value != sval {
		t.Fatal("assertion failed, expected s.value==sval.")
	}
	tag = GetTag(sptr)
	if tag != 0 {
		t.Fatal("assertion failed, expected tag==0.")
	}
}
