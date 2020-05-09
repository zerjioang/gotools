// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package str

import (
	"encoding/json"
	"strings"
	"testing"
)

var CompilerHackDoNotOptimize string
var CompilerHackDoNotOptimizeInt int

func BenchmarkStringUtils(b *testing.B) {

	b.Run("to-lower-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var s string
		for n := 0; n < b.N; n++ {
			s = strings.ToLower(val)
		}
		CompilerHackDoNotOptimize = s
	})
	b.Run("ToLowerAscii", func(b *testing.B) {
		b.Run("empty", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			val := ""
			b.ResetTimer()
			var s string
			for n := 0; n < b.N; n++ {
				s = ToLowerAscii(val)
			}
			CompilerHackDoNotOptimize = s
		})
		b.Run("with-content", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			val := "Hello World, This is AWESOME"
			b.ResetTimer()
			var s string
			for n := 0; n < b.N; n++ {
				s = ToLowerAscii(val)
			}
			CompilerHackDoNotOptimize = s
		})
	})
	b.Run("ToLowerAscii-bytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var s string
		for n := 0; n < b.N; n++ {
			s = ToLowerAscii(val)
		}
		CompilerHackDoNotOptimize = s
	})
	b.Run("tolower-bytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var s string
		for n := 0; n < b.N; n++ {
			s = toLower(val)
		}
		CompilerHackDoNotOptimize = s
	})
	b.Run("ToLowerAscii-std-bytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var s string
		for n := 0; n < b.N; n++ {
			s = strings.ToLower(val)
		}
		CompilerHackDoNotOptimize = s
	})

	b.Run("len-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var i int
		for n := 0; n < b.N; n++ {
			i = len(val)
		}
		CompilerHackDoNotOptimizeInt = i
	})
	b.Run("len-custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		var i int
		for n := 0; n < b.N; n++ {
			i = strLen(val)
		}
		CompilerHackDoNotOptimizeInt = i
	})
}

func BenchmarkGetJsonBytes(b *testing.B) {
	b.Run("get-bytes-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		for i := 0; i < b.N; i++ {
			GetJsonBytes(nil)
		}
	})
	b.Run("std-marshal-example", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id      int
			Message string
		}
		test := testStruct{Id: 23554675, Message: "this is a test struct"}
		for i := 0; i < b.N; i++ {
			_, _ = StdMarshal(test)
		}
	})
	b.Run("std-json-go", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id      int
			Message string
		}
		test := testStruct{Id: 23554675, Message: "this is a test struct"}
		for i := 0; i < b.N; i++ {
			_, _ = json.Marshal(test)
		}
	})
}
