package hex

import "testing"

// benchmark to test the overhead of calling ASM code from go
/*
BenchmarkAdd/add-asm-12         	1000000000	         2.01 ns/op	 497.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkAdd/add-go-12          	2000000000	         0.25 ns/op	3998.48 MB/s	       0 B/op	       0 allocs/op
*/
func BenchmarkAdd(b *testing.B) {
	b.Run("empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
		}
	})
	b.Run("add-asm", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Add(5, 6)
		}
	})
	b.Run("add-go", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = 5 + 6
		}
	})
}
