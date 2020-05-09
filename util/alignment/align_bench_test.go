// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package util

import (
	"fmt"
	"testing"
	"unsafe"
)

type myStructOptimized struct {
	myBool  bool    // 1 byte
	myInt   int32   // 4 bytes
	myFloat float64 // 8 bytes
}

type myStructNotOptimized struct {
	myBool  bool    // 1 byte
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
}

func init() {
	fmt.Println("unaligned struct size: ", int64(unsafe.Sizeof(myStructNotOptimized{})))
	fmt.Println("aligned struct size: ", int64(unsafe.Sizeof(myStructOptimized{})))
}

// conclusion:
// aligned struct uses much less amount of RAM while its been used
func BenchmarkStructs(b *testing.B) {
	b.Run("unaligned", func(b *testing.B) {
		b.ReportAllocs()
		//b.SetBytes(int64(unsafe.Sizeof(myStructNotOptimized{})))
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = new(myStructNotOptimized)
		}
	})
	b.Run("aligned", func(b *testing.B) {
		b.ReportAllocs()
		//b.SetBytes(int64(unsafe.Sizeof(myStructOptimized{})))
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = new(myStructOptimized)
		}
	})
}
