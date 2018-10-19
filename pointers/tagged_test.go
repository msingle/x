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
