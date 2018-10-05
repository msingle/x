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

package structs

import (
	"bytes"
	"fmt"
	"reflect"
)

// CompileStructInfo is a function that takes a
// struct `s`, generates general info table and
// returns it as `string`.
func CompileStructInfo(s interface{}) string {
	const (
		_fmtheader = "[struct: %-16s%-14sfields=%-6.1dsize=%-6.2d  ]\n"
		_fmtbody   = "[  %-22s offset=%-6.1v size=%-6.1d align=%-6d ]\n"
	)
	var (
		stype reflect.Type = reflect.TypeOf(s)
		nf    int          = stype.NumField()
		field reflect.StructField
		buff  bytes.Buffer
	)
	buff.WriteString(fmt.Sprintf(_fmtheader, stype.Name(), " ", nf, stype.Size()))
	for i := 0; i < nf; i++ {
		field = stype.Field(i)
		buff.WriteString(
			fmt.Sprintf(_fmtbody,
				field.Name,
				field.Offset,
				field.Type.Size(),
				field.Type.Align(),
			))
	}
	return buff.String()
}
