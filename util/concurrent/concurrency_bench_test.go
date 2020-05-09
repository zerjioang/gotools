// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package concurrent

import (
	"sync"
	"sync/atomic"
	"testing"
)

type ConfigWithMutex struct {
	mut      *sync.RWMutex
	endpoint string
}

type Config struct {
	endpoint string
}

func BenchmarkMutexSet(b *testing.B) {
	config := ConfigWithMutex{}
	config.mut = new(sync.RWMutex)
	b.ReportAllocs()
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.mut.Lock()
			config.endpoint = "api.example.com"
			config.mut.Unlock()
		}
	})
}

func BenchmarkMutexGet(b *testing.B) {
	config := ConfigWithMutex{endpoint: "api.example.com"}
	config.mut = new(sync.RWMutex)
	b.ReportAllocs()
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.mut.RLock()
			_ = config.endpoint
			config.mut.RUnlock()
		}
	})
}

func BenchmarkAtomicSet(b *testing.B) {
	var config = new(atomic.Value)
	c := Config{endpoint: "api.example.com"}
	b.ReportAllocs()
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c)
		}
	})
}

func BenchmarkAtomicGet(b *testing.B) {
	var config = new(atomic.Value)
	config.Store(Config{endpoint: "api.example.com"})
	b.ReportAllocs()
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load().(Config)
		}
	})
}
