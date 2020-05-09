// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex

/*
initial benchmark

BenchmarkEncode/encode-stdlib-4         	10000000	       125 ns/op	   7.98 MB/s	      64 B/op	       2 allocs/op
BenchmarkEncode/encode-fast-4           	10000000	       127 ns/op	   7.83 MB/s	      64 B/op	       2 allocs/op

After using unsafe for string returning

BenchmarkEncode/encode-stdlib-4         	10000000	       124 ns/op	   8.03 MB/s	      64 B/op	       2 allocs/op
BenchmarkEncode/encode-fast-4           	20000000	        80.4 ns/op	  12.44 MB/s	      32 B/op	       1 allocs/op

we add a pooled version to compare too

BenchmarkEncode/encode-stdlib-4         	10000000	       130 ns/op	   7.66 MB/s	      64 B/op	       2 allocs/op
BenchmarkEncode/encode-fast-4           	20000000	        82.4 ns/op	  12.14 MB/s	      32 B/op	       1 allocs/op
BenchmarkEncode/encode-fast-pooled-4    	10000000	       199 ns/op	   5.02 MB/s	       0 B/op	       0 allocs/op

we add more unsafe usage to avoid bound check in all array operations
*/
