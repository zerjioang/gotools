package ip

/*

Initial benchmark analysis

BenchmarkIpToUint32/ip-to-int-default-4         	10000000	       181 ns/op	   5.50 MB/s	      16 B/op	       1 allocs/op
BenchmarkIpToUint32/ip-to-int-low-4             	30000000	        45.7 ns/op	  21.90 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpToUint32/decode-int-to-string-4      	 5000000	       276 ns/op	   3.62 MB/s	      16 B/op	       2 allocs/op

BenchmarkIpUtils/ip-to-int-default-12         	20000000	        67.4 ns/op	  14.84 MB/s	      16 B/op	       1 allocs/op
BenchmarkIpUtils/ip-to-int-low-12             	50000000	        22.7 ns/op	  44.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpUtils/decode-int-to-string-12      	20000000	        89.0 ns/op	  11.24 MB/s	      16 B/op	       2 allocs/op
BenchmarkIpUtils/ip-to-int-assembly-amd64-12  	100000000	        13.9 ns/op	  72.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpUtils/is-valid-ipv4-net-pkg-12     	30000000	        47.9 ns/op	  20.86 MB/s	      16 B/op	       1 allocs/op
BenchmarkIpUtils/is-valid-ipv4-custom-12      	100000000	        17.9 ns/op	  55.76 MB/s	       0 B/op	       0 allocs/op
BenchmarkIpUtils/is-valid-ipv4-regex-12       	 3000000	       395 ns/op	   2.53 MB/s	      32 B/op	       1 allocs/op
PASS
*/
