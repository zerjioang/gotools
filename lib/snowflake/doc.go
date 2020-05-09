package snowflake

/*
// Original package: https://github.com/bwmarrin/snowflake
Initial package performance:

go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4

BenchmarkParseBase32-4           	50000000	        24.0 ns/op	  41.64 MB/s	       0 B/op	       0 allocs/op
BenchmarkBase32-4                	20000000	        65.2 ns/op	  15.33 MB/s	      16 B/op	       1 allocs/op
BenchmarkParseBase58-4           	50000000	        23.5 ns/op	  42.61 MB/s	       0 B/op	       0 allocs/op
BenchmarkBase58-4                	20000000	        66.6 ns/op	  15.02 MB/s	      16 B/op	       1 allocs/op
BenchmarkGenerate-4              	 5000000	       243 ns/op	   4.10 MB/s	       0 B/op	       0 allocs/op
BenchmarkGenerateMaxSequence-4   	20000000	        99.6 ns/op	  10.04 MB/s	       0 B/op	       0 allocs/op
BenchmarkUnmarshal-4             	20000000	        96.4 ns/op	  10.38 MB/s	      32 B/op	       1 allocs/op
BenchmarkMarshal-4               	20000000	        97.4 ns/op	  10.27 MB/s	      32 B/op	       1 allocs/op
*/
