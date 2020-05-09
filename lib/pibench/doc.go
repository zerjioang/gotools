// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3
package pibench

/*
go test -tags "dev oss" -pibench=. -benchmem -cpu=1,2,4
go test -tags "dev oss" -pibench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

package functions performance:

BenchmarkPi/calculate-score-12                  26	      48355870 ns/op	   0.00 MB/s	   67922 B/op	      28 allocs/op
BenchmarkPi/get-score-12               	1000000000	         0.251 ns/op	3985.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-pibench-time-12          	1000000000	         0.255 ns/op	3924.03 MB/s	       0 B/op	       0 allocs/op

after caching calculate-score results, since we only execute once at bootime, we get

BenchmarkPi/calculate-score-4         					 241417473	          4.75 ns/op	 210.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-score-4               					1000000000	         0.397 ns/op	2521.17 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-score-with-local-variable-4         	1000000000	         0.450 ns/op	2223.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-score-with-global-variable-4        	1000000000	         0.435 ns/op	2297.75 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-pibench-time-4                        	1000000000	         0.391 ns/op	2555.08 MB/s	       0 B/op	       0 allocs/op

*/
