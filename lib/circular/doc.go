// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package circular

/*
A circular buffer style 2 elements structure

Initial package performance:

BenchmarkCircular/instantiate-4         	 3000000	       419 ns/op	   2.38 MB/s	    1024 B/op	       1 allocs/op
BenchmarkCircular/size-4                	30000000	        58.3 ns/op	  17.14 MB/s	      64 B/op	       1 allocs/op
BenchmarkCircular/full-4                	20000000	        69.1 ns/op	  14.47 MB/s	      64 B/op	       1 allocs/op
BenchmarkCircular/empty-4               	20000000	        63.3 ns/op	  15.81 MB/s	      64 B/op	       1 allocs/op
BenchmarkCircular/push-4                	50000000	        32.1 ns/op	  31.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkCircular/pop-4                 	100000000	        14.0 ns/op	  71.21 MB/s	       0 B/op	       0 allocs/op

*/
