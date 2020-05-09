// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package stack

import (
	"errors"
	"testing"
)

func BenchmarkError(b *testing.B) {
	b.Run("generate-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Nil()
		}
	})
	b.Run("generate-default", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = New("default")
		}
	})
	b.Run("none", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		stackErr := New("default")
		for n := 0; n < b.N; n++ {
			_ = stackErr.None()
		}
	})
	b.Run("error", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		stackErr := New("default")
		for n := 0; n < b.N; n++ {
			_ = stackErr.Error()
		}
	})
	b.Run("ret", func(b *testing.B) {
		b.Run("nil", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = Ret(nil)
			}
		})
		b.Run("cause", func(b *testing.B) {
			cause := errors.New("default cause as example")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = Ret(cause)
			}
		})
	})
}
