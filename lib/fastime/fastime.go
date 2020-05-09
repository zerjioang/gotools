// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package fastime

import (
	"encoding/binary"
	"unsafe"
)

// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
type Duration int64

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 { return int64(d) }

func (d Duration) Seconds() float64 {
	return float64(d.Nanoseconds() / 1000000)
}

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
//
// To count the number of units in a Duration, divide:
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // prints 1000
//
// To convert an integer number of units to a Duration, multiply:
//	seconds := 10
//	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
//
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

// fast time struct stored on stack
type FastTime struct {
	nsec uint32
	sec  int64
}

func (t FastTime) Raw() int64 {
	return t.sec
}

func (t FastTime) Unix() int64 {
	return t.sec / 1e9
}

func (t FastTime) Nanos() uint32 {
	return t.nsec
}

func (t FastTime) SafeBytes() []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, t.sec)
	return buf[:n]
}

func (t FastTime) UnsafeBytes() []byte {
	return (*[8]byte)(unsafe.Pointer(&t.sec))[:]
}

func (t FastTime) Add(duration Duration) FastTime {
	ns := duration.Nanoseconds()
	t.nsec += uint32(ns)
	t.sec += ns / 10e9
	return t
}

func Unix() int64 {
	_, t := internalNow()
	return t
}

func UnixNano() int64 {
	return Unix()
}

func Now() FastTime {
	t := FastTime{}
	t.now()
	return t
}

func FromTime(nanos int, unix int64) FastTime {
	return FastTime{
		nsec: uint32(nanos),
		sec:  unix,
	}
}

func NewFastTime() FastTime {
	return FastTime{}
}
