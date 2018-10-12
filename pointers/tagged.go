package pointers

import "unsafe"

// - MARK: Pointer-Tagging section.

// GetTag returns the tag value from
// low-order bits.
func GetTag(ptr unsafe.Pointer) uint {
	return uint(uintptr(ptr) & uintptr(archMAXTAG))
}

// TaggedPointer is a function for tagging pointers.
// It attaches `tag` value to the pointer `ptr` iff
// `tag` <= `archMAXTAG` and returns the tagged pointer
// along with error set to `nil`. It panics when
// `tag` > `archMAXTAG`, I do too! It's like getting
// headshot by a champagne cork.
func TaggedPointer(ptr unsafe.Pointer, tag uint) (unsafe.Pointer, error) {
	if tag > archMAXTAG {
		// flip the table, not this time!
		panic(EPTRINVALT)
	}
	return unsafe.Pointer(uintptr(ptr) | uintptr(tag)), nil
}

// Untag is a function for untagging pointers. It
// returns a `unsafe.Pointer` with low-order bits
// set to 0.
func Untag(ptr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) & archPTRMASK)
}

// HasTag returns whether the given pointer `ptr`
// is tagged.
func HasTag(ptr unsafe.Pointer) bool {
	return GetTag(ptr)&archMAXTAG > 0
}
