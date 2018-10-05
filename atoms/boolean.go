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

package atoms

import (
	"sync/atomic"
)

// Boolean is atomic type
type Boolean uint32

// NewBoolean allocates and initialize a new
// atomic boolean and returns a pointer to it.
func NewBoolean() *Boolean {
	return new(Boolean)
}

// Set sets the state to `v`.
func (b *Boolean) Set(v bool) {
	atomic.StoreUint32((*uint32)(b), btoi(v))
}

// value is a internal receiver function
// that returns the raw value.
func (b *Boolean) value() uint32 {
	return atomic.LoadUint32((*uint32)(b))
}

// Get returns the value.
func (b *Boolean) Get() (value bool) {
	return atomic.LoadUint32((*uint32)(b)) == 1
}

// Is compares whether the internal value
// is equal to `v`.
func (b *Boolean) Is(v bool) (ok bool) {
	return atomic.LoadUint32((*uint32)(b)) == btoi(v)
}

// Flip reverses the internal value
func (b *Boolean) Flip() (bool, bool) {
	var (
		curr uint32 = (*b).value()
		n    uint32 = curr ^ 0x1
	)
	if atomic.CompareAndSwapUint32((*uint32)(b), curr, n) {
		return itob(n), true
	}

	return false, false
}

// CAS performs atomic Compare-and-Swap
// operation and returns the new value.
func (b *Boolean) CAS(o, n bool) (ok bool) {
	return atomic.CompareAndSwapUint32((*uint32)(b), btoi(o), btoi(n))
}

// btoi converts a `bool` to `uint32`.
func btoi(b bool) uint32 {
	if b {
		return 1
	}

	return 0
}

// btoiR converts a `bool` to `uint32`
// and returns the reverse value.
func btoiR(b bool) uint32 {
	if b {
		return 0
	}

	return 1
}

// itob converts a `uint32` to
// `boolean`.
func itob(ui uint32) bool {
	if ui == 1 {
		return true
	}

	return false
}
