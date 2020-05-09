// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package util

/*

# Initial package benchmarking

BenchmarkGetJsonBytes/get-bytes-nil-4         	30000000	        40.4 ns/op	  24.77 MB/s	      16 B/op	       1 allocs/op
BenchmarkGetJsonBytes/fast-marshal-example-4  	 3000000	       641 ns/op	   1.56 MB/s	     120 B/op	       4 allocs/op
BenchmarkGetJsonBytes/std-marshal-example-4   	 2000000	       636 ns/op	   1.57 MB/s	     120 B/op	       4 allocs/op
BenchmarkGetJsonBytes/std-json-go-4           	 3000000	       507 ns/op	   1.97 MB/s	      96 B/op	       2 allocs/op

BenchmarkGenerateUUID/uuid-4         	 		 2000000	       845 ns/op	   1.18 MB/s	      64 B/op	       2 allocs/op

BenchmarkIpToUint32/convert-bytes-4         				500000000	         2.81 ns/op	 355.52 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpToUint32/convert-string-4        			   	200000000	        10.3 ns/op	  96.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpToUint32/convert-string-unsafe-inline-4         	1000000000	         2.86 ns/op	 349.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpToUint32/convert-string-unsafe-4                	500000000	         3.10 ns/op	 322.44 MB/s	       0 B/op	       0 allocs/op

BenchmarkStringUtils/to-lower-std-4         	10000000	       137 ns/op	   7.29 MB/s	      64 B/op	       2 allocs/op
BenchmarkStringUtils/ToLowerAscii-4         	20000000	        63.1 ns/op	  15.85 MB/s	      32 B/op	       1 allocs/op

# Data sizes

int8 	8 bits 	    -128 to 127
int16 	16 bits 	-215 to 215 -1
int32 	32 bits 	-231 to 231 -1
int64 	64 bits 	-263 to 263 -1

uint8 	8 bits 	    0 to 255
uint16 	16 bits 	0 to 216 -1
uint32 	32 bits 	0 to 232 -1
uint64 	64 bits 	0 to 264 -1

float32 32 bits
float 64 64 bits
*/
