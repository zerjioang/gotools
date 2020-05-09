package entropy

import "testing"

func BenchmarkAvailableEntropy(b *testing.B) {
	b.Run("initial-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = InitialEntropy()
		}
	})
	b.Run("available-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = AvailableEntropy()
		}
	})
}
