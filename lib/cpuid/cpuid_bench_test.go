package cpuid

import "testing"

func BenchmarkCpuFeatures(b *testing.B) {
	b.Run("determine-features", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			determineCPUFeatures()
		}
	})
	b.Run("read-features", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetCpuFeatures()
		}
	})
	b.Run("read-features-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetCpuFeaturesPtr()
		}
	})
}
