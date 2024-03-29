//
// Copyright Helix Distributed Ledger. All Rights Reserved.
// SPDX-License-Identifier: GNU GPL v3
//

// Package monotime provides a fast monotonic clock source.
package monotonic

import (
	_ "unsafe" // required to use //go:linkname

	"github.com/zerjioang/gotools/lib/fastime"
)

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

// Now returns the current time in nanoseconds from a monotonic clock.
// The time returned is based on some arbitrary platform-specific point in the
// past.  The time returned is guaranteed to increase monotonically at a
// constant rate, unlike time.Now() from the Go standard library, which may
// slow down, speed up, jump forward or backward, due to NTP activity or leap
// seconds.
func Now() uint64 {
	return uint64(nanotime())
}

// Since returns the amount of time that has elapsed since t. t should be
// the result of a call to Now() on the same machine.
func Since(t uint64) fastime.Duration {
	return fastime.Duration(uint64(nanotime()) - t)
}
