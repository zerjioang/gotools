package random_test

import (
	"testing"

	"github.com/zerjioang/gotools/thirdparty/gommon/random"
)

func BenchmarkRandom(b *testing.B) {
	b.Run("default-implementation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		b.Log(random.String(32))
		for n := 0; n < b.N; n++ {
			_ = random.String(32)
		}
	})
	b.Run("string-32-mask", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = random.RandStringBytesMaskImpr(32)
		}
	})
	b.Run("string-32-unsafe", func(b *testing.B) {
		//b.Log(random.RandStringBytesMaskImprSrcUnsafe(32))
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = random.RandStringBytesMaskImprSrcUnsafe(32)
		}
	})
	b.Run("RandomUUID32", func(b *testing.B) {
		//fmt.Println(random.RandStringBytesMaskImprSrcUnsafe32())
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = random.RandomUUID32()
		}
	})
	b.Run("RandomUUID32-local", func(b *testing.B) {
		//fmt.Println(random.RandStringBytesMaskImprSrcUnsafe32())
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = random.RandomUUID32Shared()
		}
	})
}
