package mime

import "testing"

func BenchmarkToMimetype(b *testing.B) {
	b.Run("to-mime", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ToMimetype(MimeApplicationJSON)
		}
	})
}
