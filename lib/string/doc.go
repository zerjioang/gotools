package string

/*
Initial package performance:

go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4

BenchmarkString/create-empty            10000000000              0.25 ns/op     4018.72 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-empty-2          10000000000              0.25 ns/op     3962.90 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-empty-4          10000000000              0.25 ns/op     3967.99 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data        10000000000              0.50 ns/op     2012.06 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data-2      10000000000              0.50 ns/op     2001.43 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data-4      10000000000              0.50 ns/op     1987.04 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard     1000000000               7.03 ns/op      142.17 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard-2   1000000000               7.05 ns/op      141.75 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard-4   1000000000               7.02 ns/op      142.50 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom       1000000000               8.35 ns/op      119.81 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom-2     1000000000               8.38 ns/op      119.37 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom-4     1000000000               8.37 ns/op      119.54 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard       2000000000               5.12 ns/op      195.28 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard-2     2000000000               5.04 ns/op      198.37 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard-4     2000000000               5.06 ns/op      197.53 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom         10000000000              0.38 ns/op     2665.96 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom-2       10000000000              0.37 ns/op     2693.63 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom-4       10000000000              0.37 ns/op     2698.49 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard       10000000000              1.97 ns/op      507.25 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard-2     10000000000              2.01 ns/op      498.62 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard-4     10000000000              2.02 ns/op      495.16 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom         10000000000              0.25 ns/op     4001.91 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom-2       10000000000              0.25 ns/op     3997.96 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom-4       10000000000              0.25 ns/op     4010.45 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard         10000000000              0.25 ns/op     4065.23 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard-2       10000000000              0.25 ns/op     3995.07 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard-4       10000000000              0.25 ns/op     4014.55 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom           10000000000              0.49 ns/op     2039.64 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom-2         10000000000              0.49 ns/op     2023.83 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom-4         10000000000              0.50 ns/op     1986.92 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard       10000000000              0.49 ns/op     2025.69 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard-2     10000000000              0.50 ns/op     2001.74 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard-4     10000000000              0.50 ns/op     1995.20 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom         10000000000              0.25 ns/op     3921.85 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom-2       10000000000              0.26 ns/op     3913.96 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom-4       10000000000              0.25 ns/op     4043.27 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard                   500000000               20.2 ns/op        49.60 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard-2                 500000000               19.8 ns/op        50.46 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard-4                 500000000               20.0 ns/op        50.01 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom                     1000000000              10.2 ns/op        98.26 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom-2                   1000000000               9.98 ns/op      100.21 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom-4                   1000000000               9.99 ns/op      100.14 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-uppercase/standard                   100000000               71.2 ns/op        14.05 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-uppercase/standard-2                 100000000               60.8 ns/op        16.45 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-uppercase/standard-4                 100000000               61.7 ns/op        16.20 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-uppercase/custom                     1000000000               8.51 ns/op      117.57 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-uppercase/custom-2                   1000000000               8.42 ns/op      118.72 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-uppercase/custom-4                   1000000000               8.38 ns/op      119.38 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-capitalize/standard                  100000000              101 ns/op           9.90 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-capitalize/standard-2                100000000               90.3 ns/op        11.08 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-capitalize/standard-4                100000000               89.9 ns/op        11.12 MB/s          32 B/op          1 allocs/op
BenchmarkString/to-capitalize/custom                    10000000000              0.50 ns/op     2019.87 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-capitalize/custom-2                  10000000000              0.49 ns/op     2022.39 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-capitalize/custom-4                  10000000000              0.49 ns/op     2030.80 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom                          500000000               13.5 ns/op        74.35 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom-2                        500000000               13.6 ns/op        73.38 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom-4                        500000000               13.7 ns/op        72.91 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/standard                     100000000              104 ns/op           9.61 MB/s          32 B/op          1 allocs/op
BenchmarkString/title-case/standard-2                   100000000               92.6 ns/op        10.80 MB/s          32 B/op          1 allocs/op
BenchmarkString/title-case/standard-4                   100000000               90.0 ns/op        11.11 MB/s          32 B/op          1 allocs/op
BenchmarkString/title-case/custom                       300000000               25.5 ns/op        39.21 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/custom-2                     300000000               25.7 ns/op        38.89 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/custom-4                     300000000               25.6 ns/op        39.02 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom                 500000000               12.9 ns/op        77.27 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom-2               500000000               13.5 ns/op        73.97 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom-4               500000000               13.1 ns/op        76.55 MB/s           0 B/op          0 allocs/op
BenchmarkString/contains/custom                         500000000               19.6 ns/op        51.02 MB/s           0 B/op          0 allocs/op
BenchmarkString/contains/custom-2                       500000000               19.3 ns/op        51.73 MB/s           0 B/op          0 allocs/op
BenchmarkString/contains/custom-4                       500000000               19.0 ns/op        52.50 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/standard                     3000000000               2.51 ns/op      399.19 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/standard-2                   3000000000               2.53 ns/op      395.87 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/standard-4                   3000000000               2.58 ns/op      387.66 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/custom                       1000000000               7.72 ns/op      129.47 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/custom-2                     1000000000               7.77 ns/op      128.70 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-suffix/custom-4                     1000000000               7.88 ns/op      126.92 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/standard             3000000000               2.52 ns/op      396.10 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/standard-2           3000000000               2.54 ns/op      393.19 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/standard-4           3000000000               2.53 ns/op      395.33 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/custom               2000000000               5.02 ns/op      199.27 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/custom-2             2000000000               5.03 ns/op      198.90 MB/s           0 B/op          0 allocs/op
BenchmarkString/has-prefix/custom-4             2000000000               4.98 ns/op      200.93 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/standard             1000000000               7.81 ns/op      128.12 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/standard-2           1000000000               7.83 ns/op      127.77 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/standard-4           1000000000               7.81 ns/op      128.09 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/custom               1000000000               7.67 ns/op      130.32 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/custom-2             1000000000               7.89 ns/op      126.71 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-numeric/custom-4             1000000000               7.67 ns/op      130.45 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-hexadecimal/custom           200000000               37.4 ns/op        26.74 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-hexadecimal/custom-2         200000000               36.4 ns/op        27.46 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-hexadecimal/custom-4         200000000               36.4 ns/op        27.46 MB/s           0 B/op          0 allocs/op
BenchmarkString/generate-uintptr                10000000000              0.25 ns/op     3984.54 MB/s           0 B/op          0 allocs/op
BenchmarkString/generate-uintptr-2              10000000000              0.25 ns/op     3984.82 MB/s           0 B/op          0 allocs/op
BenchmarkString/generate-uintptr-4              10000000000              0.26 ns/op     3897.27 MB/s           0 B/op          0 allocs/op

BenchmarkAssembly/is-digit-go-12         	2000000000	         0.26 ns/op	3803.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkAssembly/is-digit-asm-12        	300000000	         5.06 ns/op	 197.52 MB/s	       0 B/op	       0 allocs/op

BenchmarkAssembly/is-digit-array-go-12   	200000000	         7.85 ns/op	 127.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkAssembly/is-digit-array-asm-12  	200000000	         7.04 ns/op	 142.00 MB/s	       0 B/op	       0 allocs/op
BenchmarkAssembly/lowercase-go-12        	20000000	        68.4 ns/op	  14.62 MB/s	      32 B/op	       1 allocs/op
BenchmarkAssembly/lowercase-asm-12       	100000000	        12.5 ns/op	  80.03 MB/s	       0 B/op	       0 allocs/op

*/
