// Package atomicvar provides atomic Boolean type for
// cleaner code and better performance.
package atomicvar

import "sync/atomic"

// AtomicBoolInt32 is an atomic Boolean
// Its methods are all atomic, thus safe to be called by
// multiple goroutines simultaneously
// Note: When embedding into a struct, one should always use
// *AtomicBoolInt32 to avoid copy
type AtomicBoolInt32 int32

// New creates an AtomicBoolInt32 with default to false
func New() *AtomicBoolInt32 {
	return new(AtomicBoolInt32)
}

// NewBool creates an AtomicBoolInt32 with given default value
func NewBool(ok bool) *AtomicBoolInt32 {
	ab := New()
	if ok {
		ab.Set()
	}
	return ab
}

// Set sets the Boolean to true
func (ab *AtomicBoolInt32) Set() {
	atomic.StoreInt32((*int32)(ab), 1)
}

// UnSet sets the Boolean to false
func (ab *AtomicBoolInt32) UnSet() {
	atomic.StoreInt32((*int32)(ab), 0)
}

// IsTrue returns whether the Boolean is true
func (ab *AtomicBoolInt32) IsSet() bool {
	return atomic.LoadInt32((*int32)(ab)) == 1
}

// SetTo sets the boolean with given Boolean
func (ab *AtomicBoolInt32) SetTo(yes bool) {
	if yes {
		atomic.StoreInt32((*int32)(ab), 1)
	} else {
		atomic.StoreInt32((*int32)(ab), 0)
	}
}

// SetToIf sets the Boolean to new only if the Boolean matches the old
// Returns whether the set was done
func (ab *AtomicBoolInt32) SetToIf(old, new bool) (set bool) {
	var o, n int32
	if old {
		o = 1
	}
	if new {
		n = 1
	}
	return atomic.CompareAndSwapInt32((*int32)(ab), o, n)
}
