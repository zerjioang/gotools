package strhash

import "testing"

func BenchmarkStrHash(b *testing.B) {
	x0 := "abcdefghijklmnopqrstuvwxyz"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		StrHash(x0)
	}
}

func BenchmarkMapHash(b *testing.B) {
	key := "abcdefghijklmnopqrstuvwxyz"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(1)
	t := map[string]int64{
		key: 99,
	}
	for n := 0; n < b.N; n++ {
		_ = t[key]
	}
}
