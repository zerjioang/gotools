// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package spinlock

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
)

// Reentrant allowable spin locks
type ReentrableSpinLock struct {
	owner uint32
	count int
}

func (sl *ReentrableSpinLock) Lock() {
	me := GetGoroutineId()
	if sl.owner == me { /// If the current thread has acquired the lock, the number of threads increases by one, and then returns
		sl.count++
		return
	}
	// If the lock is not acquired, spin through CAS
	for !atomic.CompareAndSwapUint32(&sl.owner, 0, 1) {
		runtime.Gosched()
	}
}
func (sl *ReentrableSpinLock) Unlock() {
	if sl.owner != GetGoroutineId() {
		panic("illegalMonitorStateError")
	}
	if sl.count > 0 { // if greater than 0, it means that the current thread has acquired the lock many times, and the release lock is simulated by subtracting count from one.
		sl.count--
	} else {
		// If count== 0, the lock can be released, which ensures that the number of acquisitions of the lock is the same as the number of releases of the lock.
		atomic.StoreUint32(&sl.owner, 0)
	}
}

func GetGoroutineId() uint32 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return uint32(id)
}

func NewReentrableSpinLock() *ReentrableSpinLock {
	return new(ReentrableSpinLock)
}
