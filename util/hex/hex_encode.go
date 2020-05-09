// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex

import (
	"bytes"
	"sync"
	"unsafe"
)

const (
	hextable = "0123456789abcdef"
)

var (
	hextableData = [16]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 102}
	hexpool      *sync.Pool
)

func init() {
	hexpool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// Encode encodes src into EncodedLen(len(src))
// bytes of dst. As a convenience, it returns the number
// of bytes written to dst, but this value is always EncodedLen(len(src)).
// Encode implements hexadecimal encoding.
func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return len(src) * 2
}

// UnsafeEncodeToString returns the hexadecimal encoding of src.
func UnsafeEncodeToString(src []byte) string {
	size := len(src)
	dst := make([]byte, size*2)
	//no bound check
	sourceStart := uintptr(unsafe.Pointer(&src[0]))
	dstStart := uintptr(unsafe.Pointer(&dst[0]))
	tableStart := uintptr(unsafe.Pointer(&hextableData[0]))
	step := unsafe.Sizeof(src[0])

	for i := 0; i < size; i++ {
		indexp := (unsafe.Pointer)(sourceStart + step*uintptr(i))
		v := *(*byte)(indexp)
		//hex encoded values and store
		*(*byte)((unsafe.Pointer)(dstStart + step*uintptr(i*2))) = *(*byte)((unsafe.Pointer)(tableStart + step*uintptr(v>>4)))
		*(*byte)((unsafe.Pointer)(dstStart + step*uintptr(i*2+1))) = *(*byte)((unsafe.Pointer)(tableStart + step*uintptr(v&0x0f)))
	}
	return *(*string)(unsafe.Pointer(&dst))
}

// UnsafeEncodeToString returns the hexadecimal encoding of src.
func UnsafeEncodeToStringPooled(src []byte) string {
	buf := hexpool.Get().(*bytes.Buffer)
	for i := 0; i < len(src); i++ {
		v := src[i]
		buf.WriteByte(hextableData[v>>4])
		buf.WriteByte(hextableData[v&0x0f])
	}
	dst := buf.Bytes()
	buf.Reset()
	hexpool.Put(buf)
	return *(*string)(unsafe.Pointer(&dst))
}

func EncodeString(src []byte) string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(src); i++ {
		v := src[i]
		buf.WriteByte(hextableData[v>>4])
		buf.WriteByte(hextableData[v&0x0f])
	}
	return buf.String()
}
