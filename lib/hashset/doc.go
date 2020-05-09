// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hashset

/*
Initial Benchmarking

BenchmarkHashSet/instantiate            200000000                8.28 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/instantiate-2          200000000                9.15 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/instantiate-4          200000000                9.60 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/instantiate-ptr        200000000                8.06 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/instantiate-ptr-2      200000000                8.16 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/instantiate-ptr-4      200000000                8.30 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet/add/simple             50000000                28.0 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/add/simple-2           50000000                28.0 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/add/simple-4           50000000                28.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/add/10000-items            2000            917202 ns/op           39365 B/op       9900 allocs/op
BenchmarkHashSet/add/10000-items-2          2000            947855 ns/op           39369 B/op       9900 allocs/op
BenchmarkHashSet/add/10000-items-4          2000            913496 ns/op           39368 B/op       9900 allocs/op
BenchmarkHashSet/contains/simple        50000000                22.9 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/simple-2      100000000               23.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/simple-4      50000000                26.0 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-first            50000000                30.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-first-2          50000000                35.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-first-4          50000000                31.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-middle           50000000                33.2 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-middle-2         50000000                33.5 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-middle-4         50000000                36.9 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-last             50000000                32.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-last-2           50000000                31.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-last-4           50000000                32.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-0                                        100000000               18.1 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-0-2                                      100000000               17.8 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-0-4                                      100000000               17.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-10000                                    100000000               17.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-10000-2                                  100000000               17.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/count-10000-4                                  100000000               17.7 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size                                           100000000               17.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size-2                                         100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size-4                                         100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size-10000                                     100000000               17.4 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size-10000-2                                   100000000               17.3 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/size-10000-4                                   100000000               17.6 ns/op             0 B/op          0 allocs/op
BenchmarkHashSet/clear                                          20000000                64.8 ns/op            48 B/op          1 allocs/op
BenchmarkHashSet/clear-2                                        20000000                65.1 ns/op            48 B/op          1 allocs/op
BenchmarkHashSet/clear-4                                        20000000                69.0 ns/op            48 B/op          1 allocs/op
*/
