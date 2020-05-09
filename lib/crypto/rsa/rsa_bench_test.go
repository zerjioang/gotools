package rsa

import (
	"testing"

	"github.com/zerjioang/gotools/lib/crypto/serializer/pem"
)

func BenchmarkGenerateRSA(b *testing.B) {
	b.Run("generate-1024", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = GenerateRSA(1024)
		}
	})
	b.Run("encode-private", func(b *testing.B) {
		b.Run("to-pem", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			pk, _ := GenerateRSA(1024)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = pem.RsaPrivateToPEM(pk)
			}
		})
	})
	b.Run("encode-public", func(b *testing.B) {
		b.Run("to-pem", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			pk, _ := GenerateRSA(1024)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = pem.RsaPublicToPEM(&pk.PublicKey)
			}
		})
	})
}
