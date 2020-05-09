// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hashset

import (
	"strconv"
	"testing"
)

func BenchmarkHashSetAtomic(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = NewAtomicHashSet()
		}
	})
	b.Run("instantiate-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			_ = NewAtomicHashSetPtr()
		}
	})
	b.Run("add", func(b *testing.B) {
		b.Run("simple", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			b.SetBytes(1)
			set := NewAtomicHashSet()
			for i := 0; i < b.N; i++ {
				set.Add("test")
			}
		})
		b.Run("10000-items", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			set := NewAtomicHashSet()
			for i := 0; i < b.N; i++ {
				for i := 0; i < 10000; i++ {
					set.Add(strconv.Itoa(i))
				}
			}
		})
	})
	b.Run("contains", func(b *testing.B) {
		b.Run("simple", func(b *testing.B) {

			set := NewAtomicHashSet()
			set.Add("test")

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				set.ContainsString("test")
			}
		})
		b.Run("10000-items", func(b *testing.B) {

			//add 10000 items first
			set := NewAtomicHashSet()
			for i := 0; i < 10000; i++ {
				set.Add(strconv.Itoa(i))
			}

			b.Run("contains-first", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					set.ContainsString("0")
				}
			})
			b.Run("contains-middle", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					set.ContainsString("5000")
				}
			})
			b.Run("contains-last", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					set.ContainsString("9999")
				}
			})
		})
	})
	b.Run("count-0", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		set := NewAtomicHashSet()
		for i := 0; i < b.N; i++ {
			_ = set.Size()
		}
	})
	b.Run("count-10000", func(b *testing.B) {
		//add 10000 items first
		set := NewAtomicHashSet()
		for i := 0; i < 10000; i++ {
			set.Add(strconv.Itoa(i))
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			set.Size()
		}
	})
	b.Run("size", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		set := NewAtomicHashSet()
		for i := 0; i < b.N; i++ {
			_ = set.Size()
		}
	})
	b.Run("size-10000", func(b *testing.B) {
		//add 10000 items first
		set := NewAtomicHashSet()
		for i := 0; i < 10000; i++ {
			set.Add(strconv.Itoa(i))
		}

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			set.Size()
		}
	})

	b.Run("clear", func(b *testing.B) {
		b.Run("clear-standard", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			set := NewAtomicHashSet()
			for i := 0; i < b.N; i++ {
				set.Clear()
			}
		})

		b.Run("clear-fast", func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			set := NewAtomicHashSet()
			for i := 0; i < b.N; i++ {
				set.Clear()
			}
		})
	})
}
