// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package counter32

/*

go test -bench=. -benchmem -benchtime=2s -cpu=1,2,4

this package implements an concurrency safe atomic uint32 data structure

Initial performance:

BenchmarkCounterPtr/instantiate                 3000000000               0.31 ns/op     3225.37 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-2               3000000000               0.31 ns/op     3235.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-4               3000000000               0.31 ns/op     3180.21 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add                         500000000                6.03 ns/op      165.73 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-2                       500000000                5.91 ns/op      169.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-4                       500000000                5.94 ns/op      168.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get                         3000000000               0.33 ns/op     2998.91 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-2                       3000000000               0.34 ns/op     2969.39 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-4                       3000000000               0.34 ns/op     2950.34 MB/s           0 B/op          0 allocs/op

As can be seen non pointer based implementation is much slower ,and thus, is use is not recommended

next iteration: we add a method to covnert uint32 to bytes using json encoder with following results

BenchmarkCounterPtr/instantiate                 3000000000               0.31 ns/op     3217.65 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-2               3000000000               0.31 ns/op     3230.79 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-4               3000000000               0.31 ns/op     3209.56 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add                         500000000                5.88 ns/op      170.06 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-2                       500000000                5.89 ns/op      169.77 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-4                       500000000                5.90 ns/op      169.52 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get                         3000000000               0.33 ns/op     3001.97 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-2                       3000000000               0.33 ns/op     3019.26 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-4                       3000000000               0.35 ns/op     2857.33 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n                       500000000                5.98 ns/op      167.24 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n-2                     500000000                5.98 ns/op      167.15 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n-4                     500000000                6.28 ns/op      159.22 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix                     500000000                6.33 ns/op      157.97 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix-2                   500000000                6.30 ns/op      158.67 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix-4                   500000000                6.24 ns/op      160.18 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/bytes                       10000000               229 ns/op           4.35 MB/s          16 B/op          2 allocs/op
BenchmarkCounterPtr/bytes-2                     20000000               210 ns/op           4.74 MB/s          16 B/op          2 allocs/op
BenchmarkCounterPtr/bytes-4                     20000000               212 ns/op           4.71 MB/s          16 B/op          2 allocs/op

next iteration: remove json encoder with unsafe data conversion

BenchmarkCounterPtr/instantiate                 3000000000               0.33 ns/op     3007.12 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-2               3000000000               0.33 ns/op     3060.02 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-4               3000000000               0.32 ns/op     3082.09 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add                         500000000                6.21 ns/op      161.02 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-2                       500000000                5.87 ns/op      170.26 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-4                       500000000                5.88 ns/op      170.17 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get                         3000000000               0.32 ns/op     3167.30 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-2                       3000000000               0.33 ns/op     3073.03 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-4                       3000000000               0.31 ns/op     3202.22 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n                       500000000                5.86 ns/op      170.72 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n-2                     500000000                5.89 ns/op      169.90 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-n-4                     500000000                5.86 ns/op      170.73 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix                     500000000                5.86 ns/op      170.52 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix-2                   500000000                5.87 ns/op      170.46 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/set-fix-4                   500000000                5.87 ns/op      170.28 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes                3000000000               0.34 ns/op     2951.80 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes-2              3000000000               0.31 ns/op     3198.11 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes-4              3000000000               0.31 ns/op     3196.46 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes-fixed          3000000000               0.31 ns/op     3212.32 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes-fixed-2        3000000000               0.31 ns/op     3207.14 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/unsafe-bytes-fixed-4        3000000000               0.31 ns/op     3212.62 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/safe-bytes                  3000000000               0.42 ns/op     2406.24 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/safe-bytes-2                3000000000               0.41 ns/op     2410.94 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/safe-bytes-4                3000000000               0.42 ns/op     2385.62 MB/s           0 B/op          0 allocs/op

*/
