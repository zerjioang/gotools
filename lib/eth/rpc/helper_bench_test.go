// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"math/big"
	"testing"
)

func BenchmarkParseInt(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = ParseInt("0x143")
		_, _ = ParseInt("143")
		_, _ = ParseInt("0xaaa")
		_, _ = ParseInt("1*29")
	}
}

func BenchmarkParseBigInt(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _, _ = ParseBigInt("0xabc")
		_, _, _ = ParseBigInt("$%1")
	}
}

func BenchmarkIntToHex(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		IntToHex(1000000000000000000)
		IntToHex(111)
	}
}

func BenchmarkBigToHex(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = big.NewInt(0).SetString("1000000000000000000", 10)
		_, _ = big.NewInt(0).SetString("100000000000000000000", 10)
		_, _ = big.NewInt(0).SetString("0", 10)
	}
}
