// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package counter32_test

import (
	"testing"

	"github.com/zerjioang/gotools/lib/counter32"
)

func BenchmarkCounterPtr(b *testing.B) {

	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = counter32.NewCounter32()
		}
	})
	b.Run("add", func(b *testing.B) {
		c := counter32.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Increment()
		}
	})
	b.Run("get", func(b *testing.B) {
		c := counter32.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.Get()
		}
	})
	b.Run("set-n", func(b *testing.B) {
		c := counter32.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Set(uint32(n))
		}
	})
	b.Run("set-fix", func(b *testing.B) {
		c := counter32.NewCounter32()
		x := uint32(55)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Set(x)
		}
	})
	b.Run("unsafe-bytes", func(b *testing.B) {
		c := counter32.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.UnsafeBytes()
		}
	})
	b.Run("unsafe-bytes-fixed", func(b *testing.B) {
		c := counter32.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.UnsafeBytesFixed()
		}
	})
	b.Run("safe-bytes", func(b *testing.B) {
		c := counter32.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.SafeBytes()
		}
	})
}
