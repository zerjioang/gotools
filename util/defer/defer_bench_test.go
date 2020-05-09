package _defer

import (
	"testing"
)

func BenchmarkDeferYes(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	t := 0
	for n := 0; n < b.N; n++ {
		doDefer(&t)
	}
}

func BenchmarkDeferNo(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()
	t := 0
	for i := 0; i < b.N; i++ {
		doNoDefer(&t)
	}
}
