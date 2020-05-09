package multihash_test

import (
	"testing"

	"github.com/zerjioang/gotools/multihash"
)

// BenchmarkMultihash/encode-12         	  396752	      2872 ns/op	   0.35 MB/s	     128 B/op	       3 allocs/op
// BenchmarkMultihash/encode-12         	  411987	      2855 ns/op	   0.35 MB/s	     112 B/op	       2 allocs/op
// BenchmarkMultihash/encode-12         	  432103	      2855 ns/op	   0.35 MB/s	      96 B/op	       1 allocs/op
func BenchmarkMultihash(b *testing.B) {
	b.Run("encode-md2", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		plain := "hello-world"
		msg := []byte(plain)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = multihash.Encode(multihash.Md2, msg)
		}
	})
}
