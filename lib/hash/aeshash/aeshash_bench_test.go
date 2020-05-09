// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package aeshash_test

import (
	"testing"

	"github.com/zerjioang/gotools/lib/hash/aeshash"
)

func BenchmarkAesHash(b *testing.B) {
	b.Run("aeshash-cheese", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = aeshash.Hash("cheese")
		}
	})
	b.Run("aeshash-random-example-test", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = aeshash.Hash("random-example-test")
		}
	})
}
