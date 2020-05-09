// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package badips

import "testing"

func BenchmarkBadIPs(b *testing.B) {
	b.Run("get-list", func(b *testing.B) {
		Init("./list_any_3", true)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetBadIPList()
		}
	})

	b.Run("get-list-parallel", func(b *testing.B) {
		Init("./list_any_3", true)
		b.RunParallel(func(pb *testing.PB) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for pb.Next() {
				_ = GetBadIPList()
			}
		})
	})

	b.Run("first-item-access", func(b *testing.B) {
		Init("./list_any_3", true)
		l := GetBadIPList()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = l.Contains("31.6.220.31")
		}
	})

	b.Run("first-item-access-parallel", func(b *testing.B) {
		Init("./list_any_3", true)
		l := GetBadIPList()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains("31.6.220.31")
			}
		})
	})

	b.Run("last-item-access", func(b *testing.B) {
		Init("./list_any_3", true)

		l := GetBadIPList()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = l.Contains("121.31.56.58")
		}
	})

	b.Run("last-item-access-parallel", func(b *testing.B) {
		Init("./list_any_3", true)
		l := GetBadIPList()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains("121.31.56.58")
			}
		})
	})

	b.Run("is-blacklisted", func(b *testing.B) {
		Init("./list_any_3", true)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsBackListedIp("123.169.201.52")
		}
	})

	b.Run("is-blacklisted-parallel", func(b *testing.B) {
		Init("./list_any_3", true)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = IsBackListedIp("123.169.201.52")
			}
		})
	})
}
