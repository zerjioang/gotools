package md2

import (
	"testing"
)

func BenchmarkMd2(b *testing.B) {
	b.Run("encode-md2-reset", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		plain := "hello-world"
		h := NewMd2()
		_, _ = h.Write([]byte(plain))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			h.Reset()
		}
	})
}
