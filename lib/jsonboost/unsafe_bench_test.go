package jsonboost

import (
	"testing"
)

func BenchmarkBytesToString(b *testing.B) {
	b.Run("unsafe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = StringToBytes(benchmarkString)
		}
	})
	b.Run("safe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = []byte(benchmarkString)
		}
	})
}
