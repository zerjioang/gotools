package pearson

import "testing"

func BenchmarkNewPearsonHasher(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = NewPearsonHasher()
		}
	})
	b.Run("complete", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			hasher := NewPearsonHasher()
			_ = hasher.Salted([]byte("foo"), []byte("bar"))
			_ = hasher.Do([]byte("bar"))
		}
	})
	b.Run("hasher-reuse", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		hasher := NewPearsonHasher()
		for i := 0; i < b.N; i++ {
			_ = hasher.Salted([]byte("foo"), []byte("bar"))
			_ = hasher.Do([]byte("bar"))
		}
	})
	b.Run("not-salted", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		hasher := NewPearsonHasher()
		for i := 0; i < b.N; i++ {
			_ = hasher.Do([]byte("bar"))
		}
	})
	b.Run("salted", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		hasher := NewPearsonHasher()
		for i := 0; i < b.N; i++ {
			_ = hasher.Salted([]byte("foo"), []byte("bar"))
			_ = hasher.Do([]byte("bar"))
		}
	})
}
