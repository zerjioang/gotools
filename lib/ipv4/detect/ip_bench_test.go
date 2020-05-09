// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package detect

import (
	"testing"
)

func BenchmarkIsIpv4Methods(b *testing.B) {
	b.Run("is-valid-ipv4-net", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Net("10.41.132.6")
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
	b.Run("is-valid-ipv4-simple", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4Simple("10.41.132.6")
		}
	})
	b.Run("is-valid-ipv4-optimized", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsIpv4("10.41.132.6")
		}
	})
}
