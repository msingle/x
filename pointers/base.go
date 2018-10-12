// Packge pointers provide facilities for working with
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
	archADDRSIZE = 32 << uintptr(^uintptr(0)>>63)
	archWORDSIZE = archADDRSIZE >> 3
	archMAXTAG   = archWORDSIZE - 1
	archPTRMASK  = ^uintptr((archADDRSIZE >> 5) + 1)
)

var (
	EPTRINVAL  error = errors.New("pointer: invalid.")
	EPTRINVALT error = errors.New("pointer: invalid tag.")
)

var (
	_PTR_         unsafe.Pointer
	_INTERFACE_   interface{}
	archPTRSIZE   uintptr = unsafe.Sizeof(_PTR_)
	sizeINTERFACE uintptr = unsafe.Sizeof(_INTERFACE_)
)
