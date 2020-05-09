// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package eth

import (
	"testing"
)

func BenchmarkAddress(b *testing.B) {
	b.Run("convert-address", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ConvertAddress(address0)
		}
	})
	b.Run("convert-address-and-hex", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ConvertAddress(address0).Hex()
		}
	})
	b.Run("is-zero-address", func(b *testing.B) {
		b.Run("invalid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsZeroAddress("0x-invalid-address")
			}
		})
		b.Run("valid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsZeroAddress(address0)
			}
		})
	})
	b.Run("is-valid-address", func(b *testing.B) {
		b.Run("invalid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsValidAddress("0x-invalid-address")
			}
		})
		b.Run("valid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsValidAddress(address0)
			}
		})
		b.Run("valid-length-address-ow", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsValidAddressLow(address0)
			}
		})
	})
}

func BenchmarkIsValidBlock(b *testing.B) {
	b.Run("empty-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsValidBlockNumber("")
		}
	})
	b.Run("latest", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsValidBlockNumber("latest")
		}
	})
	b.Run("earliest", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsValidBlockNumber("latest")
		}
	})
	b.Run("pending", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsValidBlockNumber("latest")
		}
	})
	b.Run("hex-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsValidBlockNumber("0xff")
		}
	})
}
