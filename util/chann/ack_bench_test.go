// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// Go encourages us to organize our code using goroutines and to use
// channels of channels to implement request-response semantics [1].
//
// I have encountered far more instances that require acknowledgment
// than fully-fledged respones so I became curious whether channels
// of channels were indeed the best implementation strategy.
//
// In summary, yes, they are.  These benchmarks demonstrate that
// channels perform better than mutexes, that condition variables are
// still clumsy, and that preallocation is a huge win when and if you
// can manage it.
//
// This result makes sense because a mutex is implemented in terms of
// a semaphore [2] while a channel is implemented by a different
// primitive which I'll return to research later.
//
// [1] <http://golang.org/doc/effective_go.html#chan_of_chan>
// [2] <http://swtch.com/semaphore.pdf>
// [3] http://www.golangpatterns.info/home/updates
package chann

import (
	"sync"
	"testing"
)

func BenchmarkChannelBool(b *testing.B) {
	b.StopTimer()
	ch := make(chan chan bool)
	go benchmarkChannelBool(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chB := make(chan bool)
		ch <- chB
		<-chB
	}
}

func BenchmarkChannelBoolPreallocated(b *testing.B) {
	b.StopTimer()
	ch := make(chan chan bool)
	go benchmarkChannelBool(ch)
	chB := make(chan bool)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ch <- chB
		<-chB
	}
}

func BenchmarkChannelStruct(b *testing.B) {
	b.StopTimer()
	ch := make(chan chan sentinel)
	go benchmarkChannelStruct(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		chS := make(chan sentinel)
		ch <- chS
		<-chS
	}
}

func BenchmarkChannelStructPreallocated(b *testing.B) {
	b.StopTimer()
	ch := make(chan chan sentinel)
	go benchmarkChannelStruct(ch)
	chS := make(chan sentinel)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ch <- chS
		<-chS
	}
}

// This will deadlock when `go benchmarkCond(ch)` calls `c.Signal()`
// before the benchmark loop calls `c.Wait()`.  By inspection it should
// be slower than `BenchmarkMutex`, anyway.
/*
func BenchmarkCond(b *testing.B) {
        b.StopTimer()
        ch := make(chan *sync.Cond)
        go benchmarkCond(ch)
        b.StartTimer()
        for i := 0; i < b.N; i++ {
                c := sync.NewCond(&sync.Mutex{})
                c.L.Lock()
                ch <- c
                c.Wait()
        }
}
*/

// This will deadlock when `go benchmarkCond(ch)` calls `c.Signal()`
// before the benchmark loop calls `c.Wait()`.  By inspection it should
// be slower than `BenchmarkMutex`, anyway.
/*
func BenchmarkCondPreallocated(b *testing.B) {
        b.StopTimer()
        ch := make(chan *sync.Cond)
        go benchmarkCond(ch)
        c := sync.NewCond(&sync.Mutex{})
        b.StartTimer()
        for i := 0; i < b.N; i++ {
                c.L.Lock()
                ch <- c
                c.Wait()
                c.L.Unlock()
        }
}
*/

func BenchmarkMutex(b *testing.B) {
	b.StopTimer()
	ch := make(chan *sync.Mutex)
	go benchmarkMutex(ch)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m := &sync.Mutex{}
		m.Lock()
		ch <- m
		m.Lock()
	}
}

func BenchmarkMutexPreallocated(b *testing.B) {
	b.StopTimer()
	ch := make(chan *sync.Mutex)
	go benchmarkMutex(ch)
	m := &sync.Mutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Lock()
		ch <- m
		m.Lock()
		m.Unlock()
	}
}

func TestNoWarning(t *testing.T) {}

func benchmarkChannelBool(ch chan chan bool) {
	for {
		<-ch <- true
	}
}

func benchmarkChannelStruct(ch chan chan sentinel) {
	for {
		<-ch <- sentinel{}
	}
}

func benchmarkCond(ch chan *sync.Cond) {
	for {
		(<-ch).Signal()
	}
}

func benchmarkMutex(ch chan *sync.Mutex) {
	for {
		(<-ch).Unlock()
	}
}

type sentinel struct{}
