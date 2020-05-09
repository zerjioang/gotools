// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package fastime

/*
Initial package performance:

BenchmarkFastTime/fastime-now-4         			30000000	        53.7 ns/op	  18.63 MB/s	       0 B/op	       0 allocs/op
BenchmarkFastTime/fastime-now-unix-4    			30000000	        54.3 ns/op	  18.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardTime/standard-now-4    			30000000	        49.8 ns/op	  20.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardTime/standard-now-unix-4         	30000000	        49.3 ns/op	  20.29 MB/s	       0 B/op	       0 allocs/op

## Package scape analysis results

```bash
/usr/local/go/bin/go \
	test -c -gcflags '-m -m -l' \
	-o /tmp/___fastime_test_go github.com/zerjioang/gotools/lib/eth/fastime
```

```
github.com/zerjioang/gotools/lib/eth/fastime
# github.com/zerjioang/gotools/lib/eth/fastime [github.com/zerjioang/gotools/lib/eth/fastime.test]
core/eth/fastime/fastime_bench_test.go:8:24: leaking param: b
core/eth/fastime/fastime_bench_test.go:8:24:    from b (passed to call[argument escapes]) at core/eth/fastime/fastime_bench_test.go:10:7
core/eth/fastime/fastime_bench_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:10:23:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:10:23:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:10:23
core/eth/fastime/fastime_bench_test.go:10:23:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:17:28: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:17:28:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:17:28: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:17:28:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:17:28
core/eth/fastime/fastime_bench_test.go:17:28:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:10:28: BenchmarkFastTime.func1 b does not escape
core/eth/fastime/fastime_bench_test.go:17:33: BenchmarkFastTime.func2 b does not escape
core/eth/fastime/fastime_bench_test.go:26:28: leaking param: b
core/eth/fastime/fastime_bench_test.go:26:28:   from b (passed to call[argument escapes]) at core/eth/fastime/fastime_bench_test.go:28:7
core/eth/fastime/fastime_bench_test.go:28:24: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:28:24:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:28:24: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:28:24:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:28:24
core/eth/fastime/fastime_bench_test.go:28:24:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:35:29: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:35:29:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:35:29: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:35:29:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:35:29
core/eth/fastime/fastime_bench_test.go:35:29:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:28:29: BenchmarkStandardTime.func1 b does not escape
core/eth/fastime/fastime_bench_test.go:35:34: BenchmarkStandardTime.func2 b does not escape
core/eth/fastime/fastime_test.go:8:19: leaking param: t
core/eth/fastime/fastime_test.go:8:19:  from t (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:10:7
core/eth/fastime/fastime_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_test.go:10:23:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_test.go:10:23:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:10:23
core/eth/fastime/fastime_test.go:10:23:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:16:28: func literal escapes to heap
core/eth/fastime/fastime_test.go:16:28:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:16:28: func literal escapes to heap
core/eth/fastime/fastime_test.go:16:28:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:16:28
core/eth/fastime/fastime_test.go:16:28:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:10:28: TestFastTime.func1 t does not escape
core/eth/fastime/fastime_test.go:16:33: TestFastTime.func2 t does not escape
core/eth/fastime/fastime_test.go:25:23: leaking param: t
core/eth/fastime/fastime_test.go:25:23:         from t (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:27:7
core/eth/fastime/fastime_test.go:27:24: func literal escapes to heap
core/eth/fastime/fastime_test.go:27:24:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:27:24: func literal escapes to heap
core/eth/fastime/fastime_test.go:27:24:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:27:24
core/eth/fastime/fastime_test.go:27:24:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:31:29: func literal escapes to heap
core/eth/fastime/fastime_test.go:31:29:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:31:29: func literal escapes to heap
core/eth/fastime/fastime_test.go:31:29:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:31:29
core/eth/fastime/fastime_test.go:31:29:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:29:4: t.common escapes to heap
core/eth/fastime/fastime_test.go:29:4:  from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:27:29: leaking param: t
core/eth/fastime/fastime_test.go:27:29:         from t.common (dot of pointer) at core/eth/fastime/fastime_test.go:29:4
core/eth/fastime/fastime_test.go:27:29:         from t.common (address-of) at core/eth/fastime/fastime_test.go:29:4
core/eth/fastime/fastime_test.go:27:29:         from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8: tm3 escapes to heap
core/eth/fastime/fastime_test.go:29:8:  from ... argument (arg to ...) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8:  from *(... argument) (indirection) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8:  from ... argument (passed to call[argument content escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:34:4: t.common escapes to heap
core/eth/fastime/fastime_test.go:34:4:  from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:31:34: leaking param: t
core/eth/fastime/fastime_test.go:31:34:         from t.common (dot of pointer) at core/eth/fastime/fastime_test.go:34:4
core/eth/fastime/fastime_test.go:31:34:         from t.common (address-of) at core/eth/fastime/fastime_test.go:34:4
core/eth/fastime/fastime_test.go:31:34:         from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8: u escapes to heap
core/eth/fastime/fastime_test.go:34:8:  from ... argument (arg to ...) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8:  from *(... argument) (indirection) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8:  from ... argument (passed to call[argument content escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:29:8: TestStandardTime.func1 ... argument does not escape
core/eth/fastime/fastime_test.go:34:8: TestStandardTime.func2 ... argument does not escape
<autogenerated>:1: (*Duration).Nanoseconds .this does not escape
<autogenerated>:1: (*FastTime).Add .this does not escape
<autogenerated>:1: (*FastTime).Unix .this does not escape
# github.com/zerjioang/gotools/lib/eth/fastime.test
/tmp/go-build640814093/b001/_testmain.go:48:42: testdeps.TestDeps literal escapes to heap
/tmp/go-build640814093/b001/_testmain.go:48:42:         from testdeps.TestDeps literal (passed to call[argument escapes]) at $WORK/b001/_testmain.go:48:24
```

As can be seen our fastime variables tm1 and tm2 are not scaped to Heap.

go test -bench=. -benchmem -benchtime=2s -cpu=1,2,4

BenchmarkFastTime/fastime-now                   50000000                59.5 ns/op        16.80 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-2                 50000000                59.2 ns/op        16.90 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-4                 50000000                62.6 ns/op        15.97 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-unix              50000000                53.8 ns/op        18.58 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-unix-2            50000000                55.8 ns/op        17.91 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-unix-4            50000000                54.2 ns/op        18.44 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-nanos             50000000                54.6 ns/op        18.30 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-nanos-2           50000000                52.9 ns/op        18.90 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-now-nanos-4           50000000                52.5 ns/op        19.03 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-1           3000000000               0.62 ns/op     1603.94 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-1-2         3000000000               0.66 ns/op     1525.52 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-1-4         3000000000               0.63 ns/op     1590.14 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-2           3000000000               0.34 ns/op     2981.54 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-2-2         3000000000               0.32 ns/op     3110.56 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-2-4         3000000000               0.32 ns/op     3095.06 MB/s           0 B/op          0 allocs/op
BenchmarkFastTime/fastime-from-time-3-4         	20000000	        51.3 ns/op	  19.49 MB/s	       0 B/op	       0 allocs/op

BenchmarkStandardTime/standard-now              50000000                48.6 ns/op        20.56 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-2            50000000                48.8 ns/op        20.49 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-4            50000000                51.2 ns/op        19.52 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-unix                 50000000                51.7 ns/op        19.33 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-unix-2               50000000                52.3 ns/op        19.13 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-unix-4               50000000                51.5 ns/op        19.40 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-nanos                50000000                52.1 ns/op        19.20 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-nanos-2              50000000                52.4 ns/op        19.07 MB/s           0 B/op          0 allocs/op
BenchmarkStandardTime/standard-now-nanos-4              50000000                52.6 ns/op        19.00 MB/s           0 B/op          0 allocs/op

*/
