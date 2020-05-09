package xor

import (
	"testing"
)

func BenchmarkEncryptDecrypt(b *testing.B) {
	b.Run("encrypt-string", func(b *testing.B) {
		//bench data
		key := "JOKER"
		data := "hello world from xor"

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = EncryptDecrypt(data, key)
		}
	})
	b.Run("encrypt-byte", func(b *testing.B) {
		//bench data
		key := "JOKER"
		data := "hello world from xor"
		keyRaw := []byte(key)
		dataRaw := []byte(data)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = EncryptDecryptBytes(dataRaw, keyRaw)
		}
	})
}
