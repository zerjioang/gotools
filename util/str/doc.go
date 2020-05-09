package str

/*

Initial benchmark analysis:

BenchmarkStringUtils/to-lower-std-4         	 5000000	       219 ns/op	   4.56 MB/s	      64 B/op	       2 allocs/op
BenchmarkStringUtils/ToLowerAscii-4         	20000000	       121 ns/op	   8.25 MB/s	      32 B/op	       1 allocs/op
BenchmarkStringUtils/len-std-4              	2000000000	         0.54 ns/op	1837.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStringUtils/len-custom-4           	2000000000	         0.46 ns/op	2161.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetJsonBytes/get-bytes-nil-4         	200000000	      5.51 ns/op	 181.60 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetJsonBytes/std-marshal-example-4   	  2000000	      1213 ns/op	   0.82 MB/s	     112 B/op	       3 allocs/op
BenchmarkGetJsonBytes/std-json-go-4           	  1000000	      1100 ns/op	   0.91 MB/s	      96 B/op	       2 allocs/op
BenchmarkBytesToStrings-4                     	 10000000	       122 ns/op	      32 B/op	       1 allocs/op
BenchmarkBytesToStringsUnsafe-4                2000000000	      0.75 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesCompareSame-4                   	100000000	      10.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesCompareDifferent-4              	100000000	      12.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringsCompareSame-4                 	300000000	      4.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringsCompareDifferent-4            	 50000000	      23.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesContains-4                         50000000	      24.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringsContains-4                    	100000000	      17.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesIndex-4                         	100000000	      13.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringIndex-4                        	100000000	      12.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesReplace-4                       	 10000000	       183 ns/op	      32 B/op	       1 allocs/op
BenchmarkStringsReplace-4                         5000000	       262 ns/op	      64 B/op	       2 allocs/op
BenchmarkBytesConcat2-4                       	 20000000	      77.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkStringsConcat2-4                     	 10000000	       110 ns/op	      32 B/op	       1 allocs/op
BenchmarkStringsJoin2-4                       	 10000000	       154 ns/op	      32 B/op	       1 allocs/op
BenchmarkMapHints-4                           	100000000	      14.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapsHints_Dont-4                     	 50000000	      27.1 ns/op	       0 B/op	       0 allocs/op
*/
