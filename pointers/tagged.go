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

import "unsafe"

// - MARK: Pointer-Tagging section.

// GetTag returns the tag value from
// low-order bits.
func GetTag(ptr unsafe.Pointer) uint {
	return uint(uintptr(ptr) & uintptr(ArchMAXTAG))
}

// TaggedPointer is a function for tagging pointers.
// It attaches `tag` value to the pointer `ptr` iff
// `tag` <= `ArchMAXTAG` and returns the tagged pointer
// along with error set to `nil`. It panics when
// `tag` > `ArchMAXTAG`, I do too! It's like getting
// headshot by a champagne cork.
func TaggedPointer(ptr unsafe.Pointer, tag uint) (unsafe.Pointer, error) {
	if tag > ArchMAXTAG {
		// flip the table, not this time!
		panic(EPTRINVALT)
	}
	return unsafe.Pointer(uintptr(ptr) | uintptr(tag)), nil
}

// Untag is a function for untagging pointers. It
// returns a `unsafe.Pointer` with low-order bits
// set to 0.
func Untag(ptr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) & ArchPTRMASK)
}

// HasTag returns whether the given pointer `ptr`
// is tagged.
func HasTag(ptr unsafe.Pointer) bool {
	return GetTag(ptr)&ArchMAXTAG > 0
}
