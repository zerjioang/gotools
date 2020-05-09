// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package concurrent

/*
go test -bench=. -benchmem -cpu=1,2,4
go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

Performance results

BenchmarkMutexSet               50000000                32.9 ns/op        30.35 MB/s           0 B/op          0 allocs/op
BenchmarkMutexSet-2             20000000                54.0 ns/op        18.51 MB/s           0 B/op          0 allocs/op
BenchmarkMutexSet-4             20000000               101 ns/op           9.86 MB/s           0 B/op          0 allocs/op
BenchmarkMutexGet               100000000               16.2 ns/op        61.71 MB/s           0 B/op          0 allocs/op
BenchmarkMutexGet-2             100000000               47.2 ns/op        21.20 MB/s           0 B/op          0 allocs/op
BenchmarkMutexGet-4             30000000                50.7 ns/op        19.72 MB/s           0 B/op          0 allocs/op
BenchmarkAtomicSet              30000000                47.1 ns/op        21.21 MB/s          16 B/op          1 allocs/op
BenchmarkAtomicSet-2            30000000                46.8 ns/op        21.36 MB/s          16 B/op          1 allocs/op
BenchmarkAtomicSet-4            30000000                38.4 ns/op        26.07 MB/s          16 B/op          1 allocs/op
BenchmarkAtomicGet              2000000000               1.87 ns/op      533.94 MB/s           0 B/op          0 allocs/op
BenchmarkAtomicGet-2            2000000000               1.02 ns/op      984.08 MB/s           0 B/op          0 allocs/op
BenchmarkAtomicGet-4            2000000000               1.02 ns/op      977.47 MB/s           0 B/op          0 allocs/op

*/
