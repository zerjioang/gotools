package codes

import "testing"

func BenchmarkStatusText(b *testing.B) {
	b.Run("status-text", func(b *testing.B) {
		b.Run("continue-optimized", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = StatusTextOptimized(StatusContinue)
			}
		})
		b.Run("continue", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = StatusText(StatusContinue)
			}
		})
	})
	/*
		b.Run("status-text-switch", func(b *testing.B) {
			b.Run("first", func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(1)
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					_ = StatusTextOptimized(StatusContinue)
				}
			})
			b.Run("middle", func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(1)
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					_ = StatusTextOptimized(StatusPermanentRedirect)
				}
			})
			b.Run("last", func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(1)
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					_ = StatusTextOptimized(StatusNetworkAuthenticationRequired)
				}
			})
		})
	*/
	b.Run("is-informational", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsInformational(StatusProcessing)
		}
	})
}
