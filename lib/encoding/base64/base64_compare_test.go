package base64

import (
	"encoding/base64"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Compare(t *testing.T) {
	t.Run("decode-fast", func(t *testing.T) {
		enc := EncodeFast([]byte("hello"))
		assert.Equal(t, string(enc), "aGVsbG8=")
	})

	t.Run("ssse3", func(t *testing.T) {
		h3 := haveSSSE3()
		t.Log(h3)
	})
}

func BenchmarkBase64Compare(b *testing.B) {
	b.Run("encode-standard", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello from golang byte array")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = base64.StdEncoding.EncodeToString(msg)
		}
	})
	b.Run("encode-fast", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello from golang byte array")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = EncodeFast(msg)
		}
	})
	b.Run("encode-ssse3", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello from golang byte array")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = StdEncoding.EncodeToString(msg)
		}
	})
	b.Run("has-ssse3", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = haveSSSE3()
		}
	})
	b.Run("compare", func(b *testing.B) {
		sizes := []int{12, 23, 64, 128, 256, 512, 1024, 2068, 4096, 8192}
		for _, s := range sizes {
			b.Run("standard-"+strconv.Itoa(s), func(b *testing.B) {
				data := make([]byte, s)
				b.SetBytes(int64(len(data)))
				for i := 0; i < b.N; i++ {
					base64.StdEncoding.EncodeToString(data)
				}
			})
			b.Run("ssse3-"+strconv.Itoa(s), func(b *testing.B) {
				data := make([]byte, s)
				b.SetBytes(int64(len(data)))
				for i := 0; i < b.N; i++ {
					StdEncoding.EncodeToString(data)
				}
			})
		}
	})
}
