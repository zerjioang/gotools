package bounds

import "testing"

// this is a test file for bounc checking explanation purposes

// go test -bench=BenchmarkSum*
// go test -gcflags="-d=ssa/check_bce/debug=1" ./bounds_bench_test.go

// BenchmarkSumForward-12                  	2000000000	         0.74 ns/op	1352.35 MB/s	       0 B/op	       0 allocs/op
// BenchmarkSumForwardWithBoundsHack-12    	2000000000	         0.25 ns/op	4066.94 MB/s	       0 B/op	       0 allocs/op
// BenchmarkSumBackward-12                 	2000000000	         0.49 ns/op	2037.94 MB/s	       0 B/op	       0 allocs/op

func BenchmarkSumForward(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	// Create a slice of five integers
	var nums []int
	size := 5
	for i := 0; i < size; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		sum := nums[0] + nums[1] + nums[2] + nums[3] + nums[4]
		_ = sum
	}
}

func BenchmarkSumForwardWithBoundsHack(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	// Create a slice of five integers
	var nums []int

	size := 5
	for i := 0; i < size; i++ {
		nums = append(nums, i)
	}

	_ = nums[size-1] // early bounds check to guarantee safety of writes
	for n := 0; n < b.N; n++ {
		sum := nums[0] + nums[1] + nums[2] + nums[3] + nums[4]
		_ = sum
	}
}

func BenchmarkSumBackward(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	// Create a slice of five integers
	var nums []int
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	for n := 0; n < b.N; n++ {
		// Sum over the array backwards
		sum := nums[4] + nums[3] + nums[2] + nums[1] + nums[0]
		_ = sum
	}
}
