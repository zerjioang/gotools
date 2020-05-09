package browser

import (
	"testing"
)

func BenchmarkHasGraphicInterface(b *testing.B) {
	b.Run("detectUI", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = detectUI()
		}
	})
	b.Run("HasGraphicInterface", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = HasGraphicInterface()
		}
	})
}
