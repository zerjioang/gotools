// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package counter32

import (
	"encoding/binary"
	"strconv"
	"sync/atomic"
	"unsafe"
)

type Count32 uint32

func (c *Count32) Increment() uint32 {
	return atomic.AddUint32((*uint32)(c), 1)
}

func (c *Count32) Set(v uint32) {
	atomic.StoreUint32((*uint32)(c), v)
}

func (c *Count32) Get() uint32 {
	return atomic.LoadUint32((*uint32)(c))
}

// The code in the question interprets the uint32 as a slice header.
// The resulting slice is not valid and copy faults.
func (c *Count32) UnsafeBytes() []byte {
	v := atomic.LoadUint32((*uint32)(c))
	return (*[4]byte)(unsafe.Pointer(&v))[:]
}

func (c *Count32) JsonBytes() []byte {
	// todo optimize this concatenation and conversion
	var s = strconv.FormatUint(uint64(c.Get()), 10)
	raw := []byte(`{"count": ` + s + `}`)
	return raw
}

// The code in the question interprets the uint32 as a slice header.
// The resulting slice is not valid and copy faults.
func (c *Count32) UnsafeBytesFixed() [4]byte {
	v := atomic.LoadUint32((*uint32)(c))
	return *(*[4]byte)(unsafe.Pointer(&v))
}

// The resulting slice is safe for any kind of slice operations
func (c *Count32) SafeBytes() []byte {
	v := atomic.LoadUint32((*uint32)(c))
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, v)
	return a
}

// constructor like function
func NewCounter32() Count32 {
	var c Count32
	return c
}
