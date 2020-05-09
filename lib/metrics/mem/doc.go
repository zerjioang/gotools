// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package mem

/*
initial package performance:

BenchmarkMemStatus/instantiate-struct-4         	2000000000	         0.34 ns/op	2973.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-ptr-4            	2000000000	         0.34 ns/op	2939.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-internal-4       	 2000000	      1099 ns/op	   0.91 MB/s	    6184 B/op	       3 allocs/op
BenchmarkMemStatus/start-4                      	  1000000	      1563 ns/op	   0.64 MB/s	     524 B/op	       1 allocs/op
PASS

after using atomic value in order to avoid multiple monitor() goroutines, the benchmark is:

BenchmarkMemStatus/struct-start-4               	300000000	         3.94 ns/op	 253.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/ptr-start-4                  	500000000	         3.73 ns/op	 268.01 MB/s	       0 B/op	       0 allocs/op

at this point, we are ready to add all remaining benchmarks:

BenchmarkMemStatus/instantiate-struct-4         	2000000000	         0.34 ns/op	2972.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-ptr-4            	2000000000	         0.34 ns/op	2906.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-internal-4       	 1000000	      1134 ns/op	   0.88 MB/s	    6184 B/op	       3 allocs/op
BenchmarkMemStatus/struct-start-4               	300000000	         4.39 ns/op	 227.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/ptr-start-4                  	300000000	         4.33 ns/op	 231.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/read-memory/struct-4         	  200000	      8369 ns/op	   0.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/read-memory/ptr-4            	  200000	      8630 ns/op	   0.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/read/struct-4                	10000000	       173 ns/op	   5.76 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/read/ptr-4                   	10000000	       191 ns/op	   5.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/readptr/struct-4             	10000000	       163 ns/op	   6.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/readptr/ptr-4                	10000000	       153 ns/op	   6.52 MB/s	       0 B/op	       0 allocs/op
*/
