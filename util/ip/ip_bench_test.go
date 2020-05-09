// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ip

import (
	"testing"
)

func BenchmarkIpUtils(b *testing.B) {

	b.Run("ip-to-int-default", func(b *testing.B) {
		val := "10.41.132.6"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Ip2int(val)
		}
	})
	b.Run("ip-to-int-low", func(b *testing.B) {
		val := "10.41.132.6"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Ip2intLow(val)
		}
	})

	b.Run("decode-int-to-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Int2ip(170492934)
		}
	})

	b.Run("ip-to-int-assembly-amd64", func(b *testing.B) {
		example := []byte("10.41.132.6")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IpToInt2(example)
		}
	})
	b.Run("is-valid-ipv4-net-pkg", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Net("10.41.132.6")
		}
	})
	// BenchmarkIpUtils/is-valid-ipv4-4         	 5000000	       209 ns/op	   4.78 MB/s	      64 B/op	       1 allocs/op
	// BenchmarkIpUtils/is-valid-ipv4-4             10000000	       210 ns/op	   4.75 MB/s	      64 B/op	       1 allocs/op
	b.Run("is-valid-ipv4-custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4("10.41.132.6")
		}
	})
	b.Run("is-valid-ipv4-regex", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Regex("10.41.132.6")
		}
	})
}
