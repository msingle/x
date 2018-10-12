package pointers

import (
	"sync/atomic"
	"unsafe"
)

// - MARK: Multi-Word Compare-and-Swap Operation section.

// RDCSSDescriptor is descriptor for Multi-Word CAS. RDCSS
// is defined as a restricted form of CAS2 operating atomi-
// cally as follow:
//
// word_t RDCSS(word_t *a1,
//              word_t o1,
//              word_t *a2,
//              word_t o2,
//              word_t n) {
//   r = *a2;
//   if ((r  == o2) && (*a1 == o1)) *a2 = n;
//   return r;
// }
type RDCSSDescriptor struct {
	a1 *unsafe.Pointer // control address
	o1 unsafe.Pointer  // expected value
	a2 *unsafe.Pointer // data address
	o2 unsafe.Pointer  // old value
	n  unsafe.Pointer  // new value
}

// RDCSS performs a Double-Compare Single-Swap atomic
// operation. It attempts to change data address pointer
// `a2` to a `rdcssDescriptor` by comparing it against
// old value `o2`. When successfull, the pointer is changed
// to new value `n` or re-instiated to `o2` in case of
// unsuccessfull operation; A descriptor is active when
// referenced from `a2`. Pointer tagging is used to distinct
// `rdcssDescriptor` pointers.
func RDCSS(a1 *unsafe.Pointer, o1 unsafe.Pointer, a2 *unsafe.Pointer, o2 unsafe.Pointer, n unsafe.Pointer) bool {
	// Paper: A Practical Multi-Word Compare-and-Swap Operation
	//        by Timothy L. Harris, Keir Fraser and Ian A. Pratt;
	//        University of Cambridge Computer Laboratory, Cambridge,
	//        UK.
	var (
		desc *RDCSSDescriptor = &RDCSSDescriptor{a1, o1, a2, o2, n}
		dptr unsafe.Pointer
	)
	// add `0x1` tag
	dptr, _ = TaggedPointer(unsafe.Pointer(desc), 1)
	if atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(desc.a2)),
		(unsafe.Pointer)(desc.o2),
		(unsafe.Pointer)(dptr),
	) {
		return RDCSSComplete(dptr)
	}
	return false
}

// RDCSSComplete performs the second stage when descriptor
// is succesfully stored in `a2`. It finishes the operation
// by swapping `a2` with target pointer `n`. The operation
// is successfull, when `a2` is not pointing to RDCSSDescriptor.
// In case of unsucessfull operation, `a2` is swapped with `o2` and
// returns false. Note, `RDCSSDescriptor` pointers have a 0x1
// tag attached to low-order bits.
func RDCSSComplete(d unsafe.Pointer) bool {
	var (
		desc   *RDCSSDescriptor
		tgdptr unsafe.Pointer = d
		dptr   unsafe.Pointer = Untag(d)
	)
	desc = (*RDCSSDescriptor)(dptr)
	if (*desc.a1 == desc.o1) && atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(desc.a2)),
		(unsafe.Pointer)(unsafe.Pointer(tgdptr)),
		(unsafe.Pointer)(desc.n),
	) {
		return true
	}
	if !atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(desc.a2)),
		(unsafe.Pointer)(tgdptr),
		(unsafe.Pointer)(desc.o2),
	) {
		// TODO
		// . restore ( unable to restore case )
	}
	return false
}

// IsRDCSSDescriptor checks whether the given pointer
// `addr` is pointong to `RDCSSDescriptor`or not. According
// to original paper ( Section 6.2 ), `RDCSSDescriptor`
// pointers can be made distinct by non-zero low-order
// bits. A pointer is pointing to `RDCSSDescriptor` iff
// `0x1` is present.
func IsRDCSSDescriptor(addr unsafe.Pointer) bool {
	return HasTag(addr)
}
