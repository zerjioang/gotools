package atomicvar

import "testing"

func BenchmarkAtomicBoolUin32(b *testing.B) {
	b.Run("set-true", func(b *testing.B) {
		var v AtomicBoolUint32
		b.ResetTimer()
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v.SetTrue()
		}
	})
	b.Run("set-false", func(b *testing.B) {
		var v AtomicBoolUint32
		b.ResetTimer()
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v.SetFalse()
		}
	})
	b.Run("is-true", func(b *testing.B) {
		var v AtomicBoolUint32
		b.ResetTimer()
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v.IsTrue()
		}
	})
	b.Run("is-false", func(b *testing.B) {
		var v AtomicBoolUint32
		b.ResetTimer()
		b.SetBytes(1)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v.IsTrue()
		}
	})
}
