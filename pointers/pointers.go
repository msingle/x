/* MIT License
*
* Copyright (c) 2018 Mike Taghavi <mitghi[at]gmail.com>
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
 */

package pointers

import (
	"sync/atomic"
	"unsafe"
)

// - MARK: Atomics section.

// CASSliceSlot is a function that performs a CAS operation
// on a given slice slot by performing pointer arithmitic
// to find slot address. `addr` is a pointer to slice,
// `data` is a pointer to old value to be compared,
// `target` is a pointer to the new value,  `index` is
// the slot number and `ptrsize` is the slice value size.
// It returns true when succesfull.
func CASSliceSlot(addr unsafe.Pointer, data unsafe.Pointer, target unsafe.Pointer, index int, ptrsize uintptr) bool {
	var (
		tptr *unsafe.Pointer
		cptr unsafe.Pointer
	)
	tptr = (*unsafe.Pointer)(unsafe.Pointer(*(*uintptr)(addr) + (ptrsize * uintptr(index))))
	cptr = unsafe.Pointer(tptr)
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(cptr)),
		(unsafe.Pointer)(unsafe.Pointer(target)),
		(unsafe.Pointer)(unsafe.Pointer(data)),
	)
}

// CASSliceSlotPtr is a function that performs a CAS operation
// on a given slice slot by performing pointer arithmitic
// to find slot pointer address. `addr` is a pointer to slice,
// `data` is a pointer to old value to be compared,
// `target` is a pointer to the new value,  `index` is
// the slot number and `ptrsize` is the slice value size.
// It returns true when succesfull.
func CASSliceSlotPtr(addr unsafe.Pointer, data unsafe.Pointer, target unsafe.Pointer, index int, ptrsize uintptr) bool {
	var (
		tptr *unsafe.Pointer
		cptr unsafe.Pointer
	)
	tptr = (*unsafe.Pointer)(unsafe.Pointer((uintptr)(addr) + (ptrsize * uintptr(index))))
	cptr = unsafe.Pointer(tptr)
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(cptr)),
		(unsafe.Pointer)(unsafe.Pointer(target)),
		(unsafe.Pointer)(unsafe.Pointer(data)),
	)
}

// CASArraySlot is a function that performs a CAS operation
// on a given array slot by performing pointer arithmitic
// to find slot address. `addr` is a pointer to array,
// `data` is a pointer to old value to be compared,
// `target` is a pointer to the new value,  `index` is
// the slot number and `ptrsize` is the slice value size.
// It returns true when succesfull.
func CASArraySlot(addr unsafe.Pointer, data unsafe.Pointer, target unsafe.Pointer, index int, ptrsize uintptr) bool {
	var (
		tptr *unsafe.Pointer
		cptr unsafe.Pointer
	)
	tptr = (*unsafe.Pointer)(unsafe.Pointer((uintptr)(addr) + (ptrsize * uintptr(index))))
	cptr = unsafe.Pointer(tptr)
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(cptr)),
		(unsafe.Pointer)(unsafe.Pointer(target)),
		(unsafe.Pointer)(unsafe.Pointer(data)),
	)
}

// OffsetArraySlot takes a array pointer and returns
// slot address by adding `index` times `ptrsize` bytes
// to slice data pointer.
func OffsetArraySlot(addr unsafe.Pointer, index int, ptrsize uintptr) unsafe.Pointer {
	return unsafe.Pointer((*unsafe.Pointer)(unsafe.Pointer((uintptr)(addr) + (ptrsize * uintptr(index)))))
}

// OffsetSliceSlot takes a slice pointer and returns
// slot address by adding `index` times `ptrsize` bytes
// to slice data pointer.
func OffsetSliceSlot(addr unsafe.Pointer, index int, ptrsize uintptr) unsafe.Pointer {
	return unsafe.Pointer(*(*uintptr)(addr) + (ptrsize * uintptr(index)))
}

// SetSliceSlot is a wrapper function that writes `d`
// to the given slice slot iff its nil and returns
// true when succesfull.
func SetSliceSlot(addr unsafe.Pointer, index int, ptrsize uintptr, d unsafe.Pointer) bool {
	return CASSliceSlot(addr, d, nil, index, ptrsize)
}

// SetSliceSlotPtr is a wrapper function that writes `d`
// to the given slice slot opinter iff its nil and returns
// true when succesfull.
func SetSliceSlotPtr(addr unsafe.Pointer, index int, ptrsize uintptr, d unsafe.Pointer) bool {
	return CASSliceSlotPtr(addr, d, nil, index, ptrsize)
}

// SetSliceSlotI is a wrapper function that writes `d`
// to the given slice slot iff its nil and return
// true when succesfull. Note, it differs from
// `SetSliceSlot` because `d` is written as a pointer
// to `interface{}`.
func SetSliceSlotI(addr unsafe.Pointer, index int, ptrsize uintptr, d interface{}) bool {
	return CASSliceSlot(addr, unsafe.Pointer(&d), nil, index, ptrsize)
}

// SetArraySlot is a wrapper function that writes `d`
// to the given array slot iff its nil. It returns
// true when succesfull.
func SetArraySlot(addr unsafe.Pointer, index int, ptrsize uintptr, d unsafe.Pointer) bool {
	return CASArraySlot(addr, d, nil, index, ptrsize)
}

// LoadArraySlot takes a array pointer and loads
// slot address by adding `index` times `ptrsize` bytes
// to slice data pointer.
func LoadArraySlot(addr unsafe.Pointer, index int, ptrsize uintptr) unsafe.Pointer {
	var (
		tptr *unsafe.Pointer
	)
	tptr = (*unsafe.Pointer)(unsafe.Pointer((uintptr)(addr) + (ptrsize * uintptr(index))))
	return atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(tptr)))
}

// LoadSliceSlot takes a slice pointer and loads
// slot address by adding `index` times `ptrsize` bytes
// to slice data pointer.
func LoadSliceSlot(addr unsafe.Pointer, index int, ptrsize uintptr) unsafe.Pointer {
	var (
		bin *unsafe.Pointer
	)
	bin = (*unsafe.Pointer)(unsafe.Pointer(*(*uintptr)(addr) + (ptrsize * uintptr(index))))
	return atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(bin)))
}

// PopArraySlot is a wrapper function that pops
// `index` slot of array iff its nil. It returns
// a pointer and true when succesfull.
func PopArraySlot(addr unsafe.Pointer, index int, ptrsize uintptr) (unsafe.Pointer, bool) {
	var (
		slot unsafe.Pointer = LoadArraySlot(addr, index, ptrsize)
	)
	if !CASArraySlot(addr, nil, slot, index, ptrsize) {
		return nil, false
	}

	return slot, true
}

// PopSliceSlot is a wrapper function that pops
// `index` slot of slice iff its nil. It returns
// a pointer and true when succesfull.
func PopSliceSlot(addr unsafe.Pointer, index int, ptrsize uintptr) (unsafe.Pointer, bool) {
	var (
		slot unsafe.Pointer = LoadSliceSlot(addr, index, ptrsize)
	)
	if !CASSliceSlot(addr, nil, slot, index, ptrsize) {
		return nil, false
	}

	return slot, true
}

// CompareAndSwapPointerTag performs CAS operation
// and swaps `source` to `source` with new tag
// when comparision is successfull. It reutrns a
// pointer and boolean to to indicate its success.
func CompareAndSwapPointerTag(source unsafe.Pointer, oldtag uint, newtag uint) (unsafe.Pointer, bool) {
	if oldtag > ArchMAXTAG || newtag > ArchMAXTAG {
		panic(EPTRINVALT)
	}
	var (
		sraw   unsafe.Pointer = Untag(source)
		sptr   unsafe.Pointer
		target unsafe.Pointer
	)
	sptr, _ = TaggedPointer(sraw, oldtag)
	target, _ = TaggedPointer(sraw, newtag)
	if atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&sptr)),
		(unsafe.Pointer)(source),
		(unsafe.Pointer)(target),
	) {
		return target, true
	}

	return nil, false
}
