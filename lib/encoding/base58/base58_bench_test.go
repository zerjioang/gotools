package base58

import (
	"testing"
)

func BenchmarkInternal(b *testing.B) {
	// BenchmarkInternal/create-alphabet-4         	 5439754	       215 ns/op	   4.66 MB/s	     192 B/op	       1 allocs/op
	// BenchmarkInternal/create-alphabet-4         	 7477903	       158 ns/op	   6.32 MB/s	     192 B/op	       1 allocs/op
	// working as struct
	// BenchmarkInternal/create-alphabet-4         	12094263	      94.0 ns/op	  10.63 MB/s	       0 B/op	       0 allocs/op
	b.Run("create-alphabet", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
		}
	})
	// BenchmarkInternal/decode-4         	 1391650	       777 ns/op	   1.29 MB/s	     304 B/op	       3 allocs/op
	b.Run("decode-fast", func(b *testing.B) {
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			_, _ = FastBase58Decoding(src)
		}
	})

	// BenchmarkInternal/decode-trivial-4         	  319533	      4451 ns/op	   0.22 MB/s	     432 B/op	      38 allocs/op
	b.Run("decode-trivial", func(b *testing.B) {
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			_, _ = TrivialBase58Decoding(src)
		}
	})

	// BenchmarkInternal/decode-bitcoin-4         	  258448	      4362 ns/op	   0.23 MB/s	     232 B/op	       8 allocs/op
	b.Run("decode-bitcoin", func(b *testing.B) {
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			_ = BitcoinDecode(src)
		}
	})

	// BenchmarkInternal/encode-fast-4         	 1000000	      1000 ns/op	   1.00 MB/s	      96 B/op	       2 allocs/op
	b.Run("encode-fast", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello world from base58")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = FastBase58Encoding(msg)
		}
	})
	b.Run("encode-fast-2", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello world from base58")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = FastBase58Encoding2(msg)
		}
	})

	// BenchmarkInternal/encode-bitcoin-4         	  231429	      4767 ns/op	   0.21 MB/s	     448 B/op	      36 allocs/op
	b.Run("encode-bitcoin", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello world from base58")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = BitcoinEncode(msg)
		}
	})

	// BenchmarkInternal/encode-trivial-4         	  164035	      6937 ns/op	   0.14 MB/s	    1408 B/op	      67 allocs/op
	b.Run("encode-trivial", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		msg := []byte("hello world from base58")
		b.SetBytes(int64(len(msg)))
		for i := 0; i < b.N; i++ {
			_ = TrivialBase58Encoding(msg)
		}
	})

	// BenchmarkInternal/encode-decode-4         	  599040	      1862 ns/op	   0.54 MB/s	     432 B/op	       5 allocs/op
	b.Run("encode-decode-fast", func(b *testing.B) {
		// num := Base58Decode([]byte(vv))
		// chk := Base58Encode(num)
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			num, err := FastBase58Decoding(src)
			if err == nil {
				chk := FastBase58Encoding(num)
				if src != chk {
				}
			}
		}
	})

	// BenchmarkInternal/encode-decode-bitcoin-4         	  128122	      8453 ns/op	   0.12 MB/s	     648 B/op	      43 allocs/op
	b.Run("encode-decode-bitcoin", func(b *testing.B) {
		// num := Base58Decode([]byte(vv))
		// chk := Base58Encode(num)
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			num := BitcoinDecode(src)
			chk := BitcoinEncode(num)
			if src != chk {
			}
		}
	})

	// BenchmarkInternal/encode-decode-trivial-4   	  108061	     10449 ns/op	   0.10 MB/s	    1904 B/op	     106 allocs/op
	b.Run("encode-decode-trivial", func(b *testing.B) {
		// num := Base58Decode([]byte(vv))
		// chk := Base58Encode(num)
		src := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
		b.ResetTimer()
		b.ReportAllocs()
		b.SetBytes(int64(len(src)))
		for i := 0; i < b.N; i++ {
			num, err := TrivialBase58Decoding(src)
			if err == nil {
				chk := TrivialBase58Encoding(num)
				if src != chk {
				}
			}
		}
	})
}
