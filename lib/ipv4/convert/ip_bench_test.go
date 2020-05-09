// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package convert

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
			_ = IpToIntAssemblyAmd64(example)
		}
	})
}
