// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package detect

/*

Benchmark Results

NOTE: Do always your own benchmarking! This belong to my own customer laptop with 8gb ram and Intel(R) Core(TM) i5-5200U CPU @ 2.20GHz

goos: linux
goarch: amd64
pkg: github.com/zerjioang/gotools/ipv4/detect
BenchmarkIsIpv4Methods/is-valid-ipv4-net-4         	20000000	        80.2 ns/op	  12.47 MB/s	      16 B/op	       1 allocs/op
BenchmarkIsIpv4Methods/is-valid-ipv4-regex-4       	 2000000	       766 ns/op	   1.30 MB/s	      32 B/op	       1 allocs/op
BenchmarkIsIpv4Methods/is-valid-ipv4-simple-4      	10000000	       238 ns/op	   4.20 MB/s	      64 B/op	       1 allocs/op
BenchmarkIsIpv4Methods/is-valid-ipv4-optimized-4   	50000000	        29.6 ns/op	  33.84 MB/s	       0 B/op	       0 allocs/op
PASS
*/
