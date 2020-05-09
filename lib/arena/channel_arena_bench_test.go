package arena

import (
	"testing"
)

func BenchmarkArena(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.SetBytes(1)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewChannelArena(1000, 1024)
		}
	})
	b.Run("push-pop", func(b *testing.B) {
		a := NewChannelArena(1000, 1024)

		b.SetBytes(1)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			a.Push(a.Pop())
		}
	})
	b.Run("insert", func(b *testing.B) {
		a := NewChannelArena(1000, 1024)

		b.SetBytes(1)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			l := a.Pop()
			l[0] = 'a'
			l[1] = 'b'
			l[2] = 'c'
			a.Push(l)
		}
	})
}
