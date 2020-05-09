package cgo

import (
	"testing"
)

func BenchmarkCgoVersion(b *testing.B) {
	b.Run("cgo-version", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = CgoVersion()
		}
	})
}
