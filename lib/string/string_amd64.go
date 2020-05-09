// +build !noasm
// +build !appengine

package string

import "unsafe"

//go:noescape
func _isNumericArray(buf, len, res unsafe.Pointer)

//go:noescape
func _isDigit(b byte) (result byte)

//go:noescape
func _lowerCase(buf, len, res unsafe.Pointer)

func IsNumericArray(src []byte) bool {
	var result byte
	_isNumericArray(
		unsafe.Pointer(&src[0]),
		unsafe.Pointer(uintptr(len(src))),
		unsafe.Pointer(&result),
	)
	return result == 1
}

func LowerCase(src []byte) {
	_lowerCase(
		unsafe.Pointer(&src[0]),
		unsafe.Pointer(uintptr(len(src))),
		unsafe.Pointer(&src[0]),
	)
}

func IsDigit(b byte) bool {
	r := _isDigit(b)
	return r == 1
}
