// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package convert

/*

Benchmark Results

NOTE: Do always your own benchmarking! This belong to my own customer laptop with 8gb ram and Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz

goos: linux
goarch: amd64
pkg: github.com/zerjioang/gotools/ipv4/convert
BenchmarkIpUtils/ip-to-int-default-4         	20000000	        93.7 ns/op	  10.67 MB/s	      16 B/op	       1 allocs/op
BenchmarkIpUtils/ip-to-int-low-4             	50000000	        32.8 ns/op	  30.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpUtils/decode-int-to-string-4      	10000000	       146 ns/op	   6.82 MB/s	      16 B/op	       2 allocs/op
BenchmarkIpUtils/ip-to-int-assembly-amd64-4  	50000000	        24.2 ns/op	  41.27 MB/s	       0 B/op	       0 allocs/op
PASS
*/
