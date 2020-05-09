// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package base64

import (
	"encoding/base64"
	"testing"
)

func BenchmarkBase64(b *testing.B) {
	b.Run("go", func(b *testing.B) {
		b.Run("encode-decode", func(b *testing.B) {
			msg := "Hello, world fromb64 benchmark"
			raw := []byte(msg)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				encoded := base64.StdEncoding.EncodeToString(raw)
				_, _ = base64.StdEncoding.DecodeString(encoded)
			}
		})
		b.Run("encode", func(b *testing.B) {
			raw := []byte(plainText)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = base64.StdEncoding.EncodeToString(raw)
			}
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.Run("encode-string", func(b *testing.B) {
			raw := []byte(plainText)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = EncodeToString(raw)
			}
		})
		b.Run("encode-stream", func(b *testing.B) {
			raw := []byte(plainText)
			w := new(testWriter).Write
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				EncodeToStream(raw, w)
			}
		})
	})
}
