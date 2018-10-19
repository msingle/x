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

// packge pointers provide facilities for working with
// pointers and atomics.
package pointers

import (
	"errors"
	"unsafe"
)

/**
* TODO
* . extend to cover 64bit spare bits ( 19 bits )
**/

const (
	ArchADDRSIZE = 32 << uintptr(^uintptr(0)>>63)
	ArchWORDSIZE = ArchADDRSIZE >> 3
	ArchMAXTAG   = ArchWORDSIZE - 1
	ArchPTRMASK  = ^uintptr((ArchADDRSIZE >> 5) + 1)
)

var (
	EPTRINVAL  error = errors.New("pointer: invalid.")
	EPTRINVALT error = errors.New("pointer: invalid tag.")
)

var (
	_PTR_         unsafe.Pointer
	_INTERFACE_   interface{}
	ArchPTRSIZE   uintptr = unsafe.Sizeof(_PTR_)
	sizeINTERFACE uintptr = unsafe.Sizeof(_INTERFACE_)
)
