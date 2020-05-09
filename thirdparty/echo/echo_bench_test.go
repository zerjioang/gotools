package echo

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	b.Run("adquire-context", func(b *testing.B) {
		e := New()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = e.AcquireContext()
		}
	})
	b.Run("adquire-release-context", func(b *testing.B) {
		e := New()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c := e.AcquireContext()
			e.ReleaseContext(c)
		}
	})
}
