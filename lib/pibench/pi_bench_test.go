// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package pibench

import (
	"testing"

	"github.com/zerjioang/gotools/lib/logger"
)

var (
	scoreVar int64
)

func BenchmarkPi(b *testing.B) {
	b.Run("calculate-score", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			calculateScore()
		}
	})

	b.Run("get-score", func(b *testing.B) {
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetScore()
		}
	})
	// we use a local variable to avoid compiler optimizations and to compare benchmark results too
	b.Run("get-score-with-local-variable", func(b *testing.B) {
		var scr int64
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			scr = GetScore()
		}
		if scr != 0 {

		}
	})
	// we use a global variable to avoid compiler optimizations and to compare benchmark results too
	b.Run("get-score-with-global-variable", func(b *testing.B) {
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			scoreVar = GetScore()
		}
		if scoreVar != 0 {

		}
	})
	b.Run("get-pibench-time", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetBenchTime()
		}
	})
}
