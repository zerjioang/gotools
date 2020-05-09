package codec

import "testing"

func BenchmarkEncoding(b *testing.B) {
	b.Run("BytesToUint64", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = BytesToUint64(dataBytes)
		}
	})
	b.Run("Uint64ToBytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Uint64ToBytes(dataUint64)
		}
	})
}
