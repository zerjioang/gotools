// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package mem

import (
	"testing"

	"github.com/zerjioang/gotools/lib/metrics/model"
)

func BenchmarkMemStatus(b *testing.B) {
	b.Run("instantiate-struct", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = MemStatusMonitor()
		}
	})
	b.Run("instantiate-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = MemStatusMonitorPtr()
		}
	})
	b.Run("instantiate-internal", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = memStatusMonitor()
		}
	})
	b.Run("struct-start", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		m := MemStatusMonitor()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			m.Start()
		}
	})
	b.Run("ptr-start", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		m := MemStatusMonitorPtr()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			m.Start()
		}
	})

	b.Run("read-memory", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitor()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				m.ReadMemory()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitorPtr()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				m.ReadMemory()
			}
		})
	})

	b.Run("read", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitor()
			wrapper := model.ServerStatusResponse{}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				wrapper = m.Read(wrapper)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitorPtr()
			wrapper := model.ServerStatusResponse{}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				wrapper = m.Read(wrapper)
			}
		})
	})
	b.Run("readptr", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitor()
			wrapper := model.ServerStatusResponse{}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				m.ReadPtr(&wrapper)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			m := MemStatusMonitorPtr()
			wrapper := model.ServerStatusResponse{}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				m.ReadPtr(&wrapper)
			}
		})
	})
}
