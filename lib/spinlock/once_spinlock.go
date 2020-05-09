// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package spinlock

import (
	"runtime"
	"sync/atomic"
)

// SpinLock implements a simple atomic spin lock, the zero value for a SpinLock is an unlocked spinlock.
type SpinLock uint32

// Lock locks sl. If the lock is already in use, the caller blocks until Unlock is called
func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		// Gosched yields the processor, allowing other goroutines to run. It does not
		// suspend the current goroutine, so execution resumes automatically.
		runtime.Gosched()
	}
}

// Unlock unlocks sl, unlike [Mutex.Unlock](http://golang.org/pkg/sync/#Mutex.Unlock),
// there's no harm calling it on an unlocked SpinLock
func (sl *SpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// TryLock will try to lock sl and return whether it succeed or not without blocking.
func (sl *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1)
}

func (sl *SpinLock) String() string {
	if atomic.LoadUint32((*uint32)(sl)) == 1 {
		return "Locked"
	}
	return "Unlocked"
}

// NewSpinLock instantiates a spin-lock.
func NewSpinLock() *SpinLock {
	return new(SpinLock)
}
