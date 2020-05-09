// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package fastime

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFastTime(t *testing.T) {

	t.Run("comparison", func(t *testing.T) {
		unixStdTime := time.Now().UnixNano()
		fastTime := Unix()
		t.Log(unixStdTime, fastTime)
		diff := fastTime - unixStdTime
		t.Log(diff)
	})
	t.Run("duration", func(t *testing.T) {
		var d Duration
		d = Nanosecond * 200
		if d.Nanoseconds() != 200 {
			t.Error("failed to get nanoseconds")
		}
	})

	t.Run("now", func(t *testing.T) {
		tm1 := Now()
		if tm1.sec > 0 {

		}
	})
	t.Run("struct-now", func(t *testing.T) {
		tt := NewFastTime()
		tt.now()
		t.Log(tt)
	})
	t.Run("add", func(t *testing.T) {
		tm2 := Now()
		u := tm2.Add(Nanosecond * 200)
		t.Log(u.Unix())
		t.Log(tm2.Unix())
		if u.Unix()-tm2.Unix() != 0 {
			t.Error("failed to add time")
		}
	})
	t.Run("struct-unix", func(t *testing.T) {
		tm2 := Now()
		u := tm2.Unix()
		t.Log(u)
	})
	t.Run("global-unix", func(t *testing.T) {
		timenow := Unix()
		t.Log(timenow)
	})
	t.Run("safe-bytes", func(t *testing.T) {
		tm2 := Now()
		raw := tm2.SafeBytes()
		t.Log(raw)
	})
	t.Run("unsafe-bytes", func(t *testing.T) {
		tm2 := Now()
		raw := tm2.UnsafeBytes()
		t.Log(raw)
	})
	t.Run("format", func(t *testing.T) {
		// get current date time
		millis := Now().Unix()
		timeStr := time.Unix(millis, 0).Format(time.RFC3339)
		millisStr := strconv.FormatInt(millis, 10)
		t.Log(millis)
		t.Log(timeStr)
		t.Log(millisStr)
		assert.Equal(t, len(millisStr), 10)
	})
}

func TestStandardTime(t *testing.T) {

	t.Run("standard-now", func(t *testing.T) {
		tm3 := time.Now()
		t.Log(tm3)
	})
	t.Run("standard-now-unix", func(t *testing.T) {
		tm4 := time.Now()
		u := tm4.Unix()
		t.Log(u)
	})
}
