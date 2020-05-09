// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package concurrentbuffer

/*
a simple concurrent buffer implementation example

## Initial performance benchmark results:

BenchmarkConcurrentBuffer/instantiate/struct-4  	2000000000	         0.41 ns/op	2442.17 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/instantiate/ptr-4     	2000000000	         0.34 ns/op	2907.85 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read/struct-4         	20000000	        63.7 ns/op	  15.69 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read/ptr-4            	20000000	        59.8 ns/op	  16.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/struct-4        	20000000	        73.1 ns/op	  13.69 MB/s	      14 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/ptr-4           	20000000	        75.9 ns/op	  13.17 MB/s	      14 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string/struct-4       	20000000	        92.4 ns/op	  10.82 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string/ptr-4          	20000000	        88.6 ns/op	  11.28 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string#01/struct-4    	30000000	        62.2 ns/op	  16.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string#01/ptr-4       	30000000	        59.3 ns/op	  16.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/struct-4          	30000000	        57.5 ns/op	  17.39 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/ptr-4             	30000000	        54.7 ns/op	  18.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/struct-4         	30000000	        60.1 ns/op	  16.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/ptr-4            	20000000	        61.5 ns/op	  16.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/struct-4          	30000000	        53.6 ns/op	  18.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/ptr-4             	30000000	        53.7 ns/op	  18.63 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/struct-4         	30000000	        60.4 ns/op	  16.56 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/ptr-4            	20000000	        58.8 ns/op	  16.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/struct-4    	30000000	        58.1 ns/op	  17.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/ptr-4       	30000000	        56.9 ns/op	  17.57 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/struct-4   	20000000	        68.1 ns/op	  14.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/ptr-4      	20000000	        71.3 ns/op	  14.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/struct-4    	20000000	        74.4 ns/op	  13.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/ptr-4       	20000000	        74.6 ns/op	  13.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/struct-4    	20000000	        62.1 ns/op	  16.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/ptr-4       	20000000	        61.5 ns/op	  16.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/struct-4  	20000000	        73.6 ns/op	  13.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/ptr-4     	20000000	        71.0 ns/op	  14.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/struct-4        	30000000	        54.0 ns/op	  18.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/ptr-4           	30000000	        54.7 ns/op	  18.28 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/struct-4     	30000000	        54.9 ns/op	  18.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/ptr-4        	30000000	        55.6 ns/op	  18.00 MB/s	       0 B/op	       0 allocs/op

## Optimization 1: remove all 'defer' declarations:

BenchmarkConcurrentBuffer/instantiate/struct-4  	2000000000	         0.35 ns/op	2887.77 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/instantiate/ptr-4     	2000000000	         0.39 ns/op	2572.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read/struct-4         	50000000	        23.7 ns/op	  42.26 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read/ptr-4            	50000000	        23.5 ns/op	  42.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/struct-4        	30000000	        34.8 ns/op	  28.75 MB/s	      19 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/ptr-4           	50000000	        30.4 ns/op	  32.89 MB/s	      11 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string/struct-4       	50000000	        36.6 ns/op	  27.35 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string/ptr-4          	50000000	        40.2 ns/op	  24.90 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string#01/struct-4    	100000000	        21.7 ns/op	  46.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string#01/ptr-4       	100000000	        22.5 ns/op	  44.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/struct-4          	100000000	        18.4 ns/op	  54.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/ptr-4             	100000000	        18.4 ns/op	  54.35 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/struct-4         	100000000	        20.7 ns/op	  48.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/ptr-4            	100000000	        20.5 ns/op	  48.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/struct-4          	100000000	        18.4 ns/op	  54.46 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/ptr-4             	100000000	        18.0 ns/op	  55.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/struct-4         	100000000	        21.2 ns/op	  47.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/ptr-4            	100000000	        21.7 ns/op	  46.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/struct-4    	100000000	        22.3 ns/op	  44.76 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/ptr-4       	100000000	        21.9 ns/op	  45.66 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/struct-4   	50000000	        30.2 ns/op	  33.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/ptr-4      	50000000	        29.4 ns/op	  34.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/struct-4    	50000000	        33.9 ns/op	  29.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/ptr-4       	50000000	        33.1 ns/op	  30.20 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/struct-4    	50000000	        25.0 ns/op	  39.99 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/ptr-4       	50000000	        24.4 ns/op	  41.05 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/struct-4  	50000000	        32.5 ns/op	  30.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/ptr-4     	50000000	        33.4 ns/op	  29.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/struct-4        	100000000	        21.5 ns/op	  46.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/ptr-4           	100000000	        21.4 ns/op	  46.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/struct-4     	100000000	        21.6 ns/op	  46.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/ptr-4        	100000000	        21.2 ns/op	  47.19 MB/s	       0 B/op	       0 allocs/op

## Optimization 2: use sync.Mutex as pointer

BenchmarkConcurrentBuffer/instantiate/struct-4  	2000000000	         0.39 ns/op	2540.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/instantiate/ptr-4     	100000000	        17.8 ns/op	  56.19 MB/s	       8 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/read/struct-4         	50000000	        22.9 ns/op	  43.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read/ptr-4            	50000000	        23.1 ns/op	  43.35 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/struct-4        	30000000	        34.5 ns/op	  28.98 MB/s	      19 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/write/ptr-4           	50000000	        31.1 ns/op	  32.19 MB/s	      11 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string/struct-4       	50000000	        35.4 ns/op	  28.25 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string/ptr-4          	50000000	        34.6 ns/op	  28.92 MB/s	       5 B/op	       1 allocs/op
BenchmarkConcurrentBuffer/string#01/struct-4    	100000000	        20.8 ns/op	  48.17 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/string#01/ptr-4       	100000000	        21.6 ns/op	  46.34 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/struct-4          	100000000	        17.9 ns/op	  55.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/cap/ptr-4             	100000000	        18.4 ns/op	  54.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/struct-4         	100000000	        20.9 ns/op	  47.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/grow/ptr-4            	100000000	        20.5 ns/op	  48.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/struct-4          	100000000	        18.4 ns/op	  54.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/len/ptr-4             	100000000	        18.1 ns/op	  55.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/struct-4         	50000000	        20.0 ns/op	  49.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/next/ptr-4            	100000000	        20.4 ns/op	  48.91 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/struct-4    	100000000	        20.8 ns/op	  48.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-byte/ptr-4       	100000000	        20.2 ns/op	  49.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/struct-4   	50000000	        30.6 ns/op	  32.63 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-bytes/ptr-4      	50000000	        30.5 ns/op	  32.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/struct-4    	30000000	        34.2 ns/op	  29.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-from/ptr-4       	50000000	        35.8 ns/op	  27.92 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/struct-4    	50000000	        24.7 ns/op	  40.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-rune/ptr-4       	50000000	        23.4 ns/op	  42.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/struct-4  	50000000	        30.9 ns/op	  32.39 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/read-string/ptr-4     	50000000	        30.4 ns/op	  32.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/struct-4        	100000000	        18.1 ns/op	  55.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/reset/ptr-4           	100000000	        18.4 ns/op	  54.24 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/struct-4     	100000000	        18.4 ns/op	  54.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkConcurrentBuffer/truncate/ptr-4        	100000000	        18.5 ns/op	  53.94 MB/s	       0 B/op	       0 allocs/op

*/
