package fastrand

/*
source: https://github.com/valyala/fastrand

Initial package performance:

BenchmarkUint32n-4                     	100000000	        18.2 ns/op	  55.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkRNGUint32n-4                  	300000000	         4.16 ns/op	 240.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkRNGUint32nWithLock-4          	20000000	        95.8 ns/op	  10.44 MB/s	       0 B/op	       0 allocs/op
BenchmarkRNGUint32nArray-4             	30000000	        47.4 ns/op	  21.08 MB/s	       0 B/op	       0 allocs/op
BenchmarkMathRandInt31n-4              	10000000	       135 ns/op	   7.36 MB/s	       0 B/op	       0 allocs/op
BenchmarkMathRandRNGInt31n-4           	100000000	        10.9 ns/op	  92.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkMathRandRNGInt31nWithLock-4   	20000000	       114 ns/op	   8.70 MB/s	       0 B/op	       0 allocs/op
BenchmarkMathRandRNGInt31nArray-4      	20000000	        71.4 ns/op	  14.00 MB/s	       0 B/op	       0 allocs/op

*/
