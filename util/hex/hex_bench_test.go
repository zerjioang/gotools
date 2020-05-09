// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex_test

import (
	"encoding/hex"
	"testing"

	gohex "github.com/tmthrgd/go-hex"
	hex2 "github.com/zerjioang/gotools/util/hex"
)

func BenchmarkEncode(b *testing.B) {
	b.Run("encode-stdlib", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = hex.EncodeToString(data)
		}
	})
	b.Run("encode-fast", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = hex2.UnsafeEncodeToString(data)
		}
	})
	b.Run("encode-gohex", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = gohex.EncodeToString(data)
		}
	})
	b.Run("encode-fast-pooled", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = hex2.UnsafeEncodeToStringPooled(data)
		}
	})
	b.Run("decode-stdlib", func(b *testing.B) {
		data := "746869732d69732d612d74657374"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = hex.DecodeString(data)
		}
	})
	b.Run("decode-gohex", func(b *testing.B) {
		data := "746869732d69732d612d74657374"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = gohex.DecodeString(data)
		}
	})
	/*b.Run("decode-fast", func(b *testing.B) {
		data := "746869732d69732d612d74657374"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_,_ = UnsafeDecodeString(data)
		}
	})*/
}
