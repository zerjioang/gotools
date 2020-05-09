// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package eth

/*

Initial benchmark results:

BenchmarkAddress/convert-address-4         	 10000000	          188 ns/op	   5.31 MB/s	      48 B/op	       1 allocs/op
BenchmarkAddress/is-zero-address-4         	300000000	         4.52 ns/op	 221.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkAddress/is-valid-address-4        	  1000000	         1164 ns/op	   0.86 MB/s	       0 B/op	       0 allocs/op

BenchmarkIsValidBlock/empty-string-4         	300000000	         4.30 ns/op	 232.46 MB/s	       0 B/op	       0 allocs/op
BenchmarkIsValidBlock/latest-4               	300000000	         3.90 ns/op	 256.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkIsValidBlock/earliest-4             	300000000	         4.25 ns/op	 235.46 MB/s	       0 B/op	       0 allocs/op
BenchmarkIsValidBlock/pending-4              	500000000	         3.98 ns/op	 251.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkIsValidBlock/hex-string-4           	100000000	        12.4 ns/op	  80.42 MB/s	       0 B/op	       0 allocs/op

optimization 1: we add a length check in each method

BenchmarkAddress/convert-address-4     						    	  10000000	          187 ns/op	   5.34 MB/s	      48 B/op	       1 allocs/op
BenchmarkAddress/is-zero-address/invalid-length-address-4         	2000000000	         0.41 ns/op	2415.20 MB/s	       0 B/op	       0 allocs/op
BenchmarkAddress/is-zero-address/valid-length-address-4           	 300000000	         4.46 ns/op	 224.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkAddress/is-valid-address/invalid-length-address-4        	 500000000	         3.28 ns/op	 305.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkAddress/is-valid-address/valid-length-address-4          	   1000000	         1126 ns/op	   0.89 MB/s	       0 B/op	       0 allocs/op

optimization 2: eth address validation is done using for loop instead of regular expresion
optimization results: 30x speedup

BenchmarkAddress/is-valid-address/valid-length-address-4          	30000000	        48.6 ns/op	  20.59 MB/s	       0 B/op	       0 allocs/op

*/
