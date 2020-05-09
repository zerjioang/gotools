// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package spinlock

import (
	"testing"
)

func BenchmarkSpinLock(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewSpinLock()
		}
	})
	b.Run("lock-once-and-try", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		sl := NewSpinLock()
		sl.Lock()
		for n := 0; n < b.N; n++ {
			sl.TryLock()
		}
	})
	b.Run("lock-get-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		sl := NewSpinLock()
		sl.Lock()
		for n := 0; n < b.N; n++ {
			_ = sl.String()
		}
	})
	b.Run("unlock-get-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		sl := NewSpinLock()
		for n := 0; n < b.N; n++ {
			_ = sl.String()
		}
	})
}
