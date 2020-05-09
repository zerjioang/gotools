//
// Copyright Helix Distributed Ledger. All Rights Reserved.
// SPDX-License-Identifier: GNU GPL v3
//

package monotonic

import "testing"

/*
* Benchmark functions start with Benchmark not Test.

* Benchmark functions are run several times by the testing package.
  The value of b.N will increase each time until the benchmark runner
  is satisfied with the stability of the benchmark. This has some important
  ramifications which we’ll investigate later in this article.

* Each benchmark is run for a minimum of 1 second by defaulb.
  If the second has not elapsed when the Benchmark function returns,
  the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, …
  and the function run again.

* the for loop is crucial to the operation of the benchmark driver
  it must be: for n := 0; n < b.N; n++

* Add b.ReportAllocs() at the beginning of each benchmark to know allocations
* Add b.SetBytes(1) to measure byte transfers

  Optimization info: https://stackimpacb.com/blog/practical-golang-benchmarks/
*/

func BenchmarkMonotonicTime(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		_ = Now()
	}
}

func BenchmarkMonotonicSince(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		_ = Since(0)
	}
}
