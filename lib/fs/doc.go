package fs

/*
bufio library based FS module for faster and better i/o operations

Initial package performance:

BenchmarkFs/pagesize-4          	2000000000	         0.25 ns/op	4025.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkFs/exists-4                   2000000	       800 ns/op	   1.25 MB/s	     213 B/op	       2 allocs/op
BenchmarkFs/read-entropy-4         	    100000	     17803 ns/op	   0.06 MB/s	    4112 B/op	       2 allocs/op
BenchmarkFs/read-entropy-16-4         	2000000	       722 ns/op	   1.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkBinaryWrite-4          	   5000000	       314 ns/op	   3.18 MB/s	      76 B/op	      19 allocs/op
BenchmarkBinaryPut-4            	 200000000	         8.51 ns/op	 117.46 MB/s	       0 B/op	       0 allocs/op
*/
