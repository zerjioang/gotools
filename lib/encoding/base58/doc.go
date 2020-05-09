package base58

/*
Package base58 provides fast implementation of base58 encoding.

Base58 Usage

To decode a base58 string:

	encoded := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
	buf, _ := base58.Decode(encoded)

To encode the same data:

	encoded := base58.Encode(buf)

With custom alphabet

  customAlphabet := base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
  encoded := base58.EncodeAlphabet(buf, customAlphabet)

Benchmarking

BenchmarkInternal/decode-fast-4             	 1491794	       799 ns/op	   1.25 MB/s	     304 B/op	       3 allocs/op
BenchmarkInternal/decode-trivial-4          	  339266	      3540 ns/op	   0.28 MB/s	     432 B/op	      38 allocs/op
BenchmarkInternal/decode-bitcoin-4          	  291103	      3867 ns/op	   0.26 MB/s	     232 B/op	       8 allocs/op
BenchmarkInternal/encode-fast-4             	 1352503	       889 ns/op	   1.12 MB/s	      96 B/op	       2 allocs/op
BenchmarkInternal/encode-bitcoin-4          	  258210	      4549 ns/op	   0.22 MB/s	     448 B/op	      36 allocs/op
BenchmarkInternal/encode-trivial-4          	  171396	      6656 ns/op	   0.15 MB/s	    1408 B/op	      67 allocs/op
BenchmarkInternal/encode-decode-fast-4      	  596247	      1811 ns/op	   0.55 MB/s	     432 B/op	       5 allocs/op
BenchmarkInternal/encode-decode-bitcoin-4   	  132553	      8578 ns/op	   0.12 MB/s	     648 B/op	      43 allocs/op
BenchmarkInternal/encode-decode-trivial-4   	  108301	     10463 ns/op	   0.10 MB/s	    1904 B/op	     106 allocs/op

*/
