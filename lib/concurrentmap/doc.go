// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package concurrentmap

/*
concurrent safe map

Initial Benchmark

BenchmarkItems-4                            	     300	   5061622 ns/op	 1939256 B/op	     385 allocs/op
BenchmarkMarshalJson-4                      	     100	  13079398 ns/op	 3295527 B/op	   20612 allocs/op
BenchmarkStrconv-4                          	50000000	        39.5 ns/op	       7 B/op	       0 allocs/op
BenchmarkSingleInsertAbsent-4               	 3000000	       512 ns/op	     119 B/op	       1 allocs/op
BenchmarkSingleInsertAbsentSyncMap-4        	 1000000	      1359 ns/op	     187 B/op	       5 allocs/op
BenchmarkSingleInsertPresent-4              	30000000	        51.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkSingleInsertPresentSyncMap-4       	20000000	       113 ns/op	      16 B/op	       1 allocs/op
--- FAIL: BenchmarkMultiInsertDifferentSyncMap
BenchmarkMultiInsertDifferent_1_Shard-4     	  300000	      6883 ns/op	     721 B/op	      12 allocs/op
BenchmarkMultiInsertDifferent_16_Shard-4    	 1000000	      1272 ns/op	     330 B/op	      11 allocs/op
BenchmarkMultiInsertDifferent_32_Shard-4    	 1000000	      1255 ns/op	     330 B/op	      11 allocs/op
BenchmarkMultiInsertDifferent_256_Shard-4   	 1000000	      1804 ns/op	     339 B/op	      12 allocs/op
BenchmarkMultiInsertSame-4                  	  300000	      5374 ns/op	     201 B/op	      10 allocs/op
BenchmarkMultiInsertSameSyncMap-4           	  500000	      7499 ns/op	     728 B/op	      31 allocs/op
BenchmarkMultiGetSame-4                     	 2000000	       633 ns/op	      19 B/op	       0 allocs/op
BenchmarkMultiGetSameSyncMap-4              	 3000000	       531 ns/op	       5 B/op	       0 allocs/op
BenchmarkMultiGetSetDifferentSyncMap-4      	  200000	     11782 ns/op	     844 B/op	      35 allocs/op
BenchmarkMultiGetSetDifferent_1_Shard-4     	  300000	      7067 ns/op	     405 B/op	      13 allocs/op
BenchmarkMultiGetSetDifferent_16_Shard-4    	  500000	      4618 ns/op	     347 B/op	      12 allocs/op
BenchmarkMultiGetSetDifferent_32_Shard-4    	 1000000	      3166 ns/op	     340 B/op	      12 allocs/op
BenchmarkMultiGetSetDifferent_256_Shard-4   	 1000000	      1767 ns/op	     339 B/op	      12 allocs/op
BenchmarkMultiGetSetBlockSyncMap-4          	 1000000	      2737 ns/op	     480 B/op	      30 allocs/op
BenchmarkMultiGetSetBlock_1_Shard-4         	  500000	      7254 ns/op	     255 B/op	      10 allocs/op
BenchmarkMultiGetSetBlock_16_Shard-4        	  500000	      4935 ns/op	     181 B/op	      10 allocs/op
BenchmarkMultiGetSetBlock_32_Shard-4        	 1000000	      2932 ns/op	     162 B/op	      10 allocs/op
BenchmarkMultiGetSetBlock_256_Shard-4       	 1000000	      1392 ns/op	     160 B/op	      10 allocs/op
BenchmarkKeys-4   	                               50000	     34673 ns/op	    3696 B/op	       4 allocs/op

*/
