package strhash

import (
	_ "runtime"
	"unsafe" // required to use //go:linkname
)

//go:noescape
//go:linkname strhash runtime.strhash
func strhash(a unsafe.Pointer, h uintptr) uintptr

func StrHash(str string) uint64 {
	return uint64(strhash(unsafe.Pointer(&str), 0))
}

func ByteHash(data []byte) uint64 {
	return uint64(strhash(unsafe.Pointer(&data), 0))
}
