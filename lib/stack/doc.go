// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package stack

/*
go test -bench=. -benchtime=2s -benchmem -cpu=1,2,4

BenchmarkError/generate-nil             3000000000               0.45 ns/op     2210.55 MB/s           0 B/op          0 allocs/op
BenchmarkError/generate-nil-2           3000000000               0.40 ns/op     2476.82 MB/s           0 B/op          0 allocs/op
BenchmarkError/generate-nil-4           3000000000               0.45 ns/op     2213.32 MB/s           0 B/op          0 allocs/op
BenchmarkError/generate-default         3000000000               0.43 ns/op     2330.30 MB/s           0 B/op          0 allocs/op
BenchmarkError/generate-default-2       3000000000               0.44 ns/op     2283.35 MB/s           0 B/op          0 allocs/op
BenchmarkError/generate-default-4       3000000000               0.41 ns/op     2421.93 MB/s           0 B/op          0 allocs/op
BenchmarkError/none                     3000000000               0.39 ns/op     2561.25 MB/s           0 B/op          0 allocs/op
BenchmarkError/none-2                   3000000000               0.39 ns/op     2595.88 MB/s           0 B/op          0 allocs/op
BenchmarkError/none-4                   3000000000               0.38 ns/op     2637.06 MB/s           0 B/op          0 allocs/op
BenchmarkError/error                    3000000000               0.40 ns/op     2522.10 MB/s           0 B/op          0 allocs/op
BenchmarkError/error-2                  3000000000               0.39 ns/op     2595.33 MB/s           0 B/op          0 allocs/op
BenchmarkError/error-4                  3000000000               0.40 ns/op     2494.51 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/nil                  1000000000               3.88 ns/op      257.64 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/nil-2                1000000000               7.38 ns/op      135.46 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/nil-4                500000000                7.08 ns/op      141.23 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/cause                300000000               11.7 ns/op        85.13 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/cause-2              300000000               12.0 ns/op        83.43 MB/s           0 B/op          0 allocs/op
BenchmarkError/ret/cause-4              300000000               12.6 ns/op        79.62 MB/s           0 B/op          0 allocs/op
*/
