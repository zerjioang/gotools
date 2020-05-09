// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package disk

/*
initial package performance:

BenchmarkDiskUsage/instantiate-4         	 2000000	       958 ns/op	   1.04 MB/s	       2 B/op	       1 allocs/op
BenchmarkDiskUsage/read-all-4            	 2000000	       951 ns/op	   1.05 MB/s	       2 B/op	       1 allocs/op
BenchmarkDiskUsage/read-used-4           	 2000000	       961 ns/op	   1.04 MB/s	       2 B/op	       1 allocs/op
BenchmarkDiskUsage/read-free-4           	 2000000	       972 ns/op	   1.03 MB/s	       2 B/op	       1 allocs/op

After creating a constructor like function and eval method

BenchmarkDiskUsage/instantiate-4         	2000000000	         0.42 ns/op	2403.88 MB/s	       0 B/op	       0 allocs/op
BenchmarkDiskUsage/read-all-4            	 2000000	       939 ns/op	   1.06 MB/s	       2 B/op	       1 allocs/op
BenchmarkDiskUsage/read-used-4           	 2000000	       939 ns/op	   1.06 MB/s	       2 B/op	       1 allocs/op
BenchmarkDiskUsage/read-free-4           	 2000000	       968 ns/op	   1.03 MB/s	       2 B/op	       1 allocs/op

After creating a ticker based monitor

BenchmarkDiskUsage/instantiate-4         	2000000000	         0.42 ns/op	2391.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkDiskUsage/read-all-4            	200000000	         6.09 ns/op	 164.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkDiskUsage/read-used-4           	300000000	         5.93 ns/op	 168.68 MB/s	       0 B/op	       0 allocs/op
BenchmarkDiskUsage/read-free-4           	200000000	         5.96 ns/op	 167.77 MB/s	       0 B/op	       0 allocs/op
PASS

But now he ave to make sure about concurrent access

*/
