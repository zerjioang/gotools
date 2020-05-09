// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package circular

import (
	"testing"
	"unsafe"
)

var myInt = 1

func BenchmarkCircular(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = NewCircular(128)
		}
	})
	b.Run("size", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		buff := NewCircular(128)
		for i := 0; i < b.N; i++ {
			buff.Size()
		}
	})
	b.Run("full", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		buff := NewCircular(128)
		for i := 0; i < b.N; i++ {
			buff.Full()
		}
	})
	b.Run("empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		buff := NewCircular(128)
		for i := 0; i < b.N; i++ {
			buff.Empty()
		}
	})
	b.Run("push", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		myBuf := NewCircular(128)
		for i := 0; i < b.N; i++ {
			myBuf.Push(unsafe.Pointer(&myInt))
		}
	})
	b.Run("pop", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		myBuf := NewCircular(128)
		for i := 0; i < b.N; i++ {
			myBuf.Push(unsafe.Pointer(&myInt))
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			myBuf.Pop()
		}
	})
}
