// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package circular

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// Circular is our circular buffer
type Circular struct {
	read, write uint64
	lastWrite   uint64
	maskVal     uint64
	data        []unsafe.Pointer
}

// constructor like function for new circular
func NewCircular(size uint64) *Circular {
	return &Circular{
		read:    1,
		write:   1,
		data:    make([]unsafe.Pointer, size),
		maskVal: size - 1,
	}
}

// Size is the size of the Circular
func (b Circular) Size() uint64 {
	return atomic.LoadUint64(&b.write) - atomic.LoadUint64(&b.read)
}

// Empty will tell you if the Circular is empty
func (b Circular) Empty() bool {
	return atomic.LoadUint64(&b.write) == atomic.LoadUint64(&b.read)
}

// Full returns true if the Circular is "full"
func (b Circular) Full() bool {
	return b.Size() == (b.maskVal + 1)
}

func (b Circular) mask(val uint64) uint64 {
	return val & b.maskVal
}

// Push places an item onto the ring Circular
func (b *Circular) Push(object unsafe.Pointer) {
	index := atomic.AddUint64(&b.write, 1) - 1
	atomic.StorePointer(&b.data[index&b.maskVal], object)
	for !atomic.CompareAndSwapUint64(&b.lastWrite, index-1, index) {
		// Gosched() is used to force scheduler update in such cases.
		runtime.Gosched()
	}
}

// Pop returns the next item on the ring Circular
func (b *Circular) Pop() unsafe.Pointer {
	for atomic.LoadUint64(&b.write) <= atomic.LoadUint64(&b.read) {
		// Gosched() is used to force scheduler update in such cases.
		runtime.Gosched()
	}

	index := atomic.AddUint64(&b.read, 1) - 1
	for index > atomic.LoadUint64(&b.write) {
		// Gosched() is used to force scheduler update in such cases.
		runtime.Gosched()
	}
	return atomic.LoadPointer(&b.data[index&b.maskVal])
}
func (b *Circular) PopBoolean() bool {
	ptr := b.Pop()
	return *(*bool)(unsafe.Pointer(ptr))
}
func (b *Circular) PushBoolean(val bool) {
	b.Push(unsafe.Pointer(&val))
}
