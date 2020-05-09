// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package counter32

import "testing"

func TestCounter(t *testing.T) {
	t.Run("atomic-uint32-instantiate", func(t *testing.T) {
		_ = NewCounter32()
	})
	t.Run("atomic-uint32-get", func(t *testing.T) {
		var c1 Count32
		value := c1.Get()
		t.Log(value)
	})
	t.Run("atomic-uint32-add", func(t *testing.T) {
		var c2 Count32
		value := c2.Increment()
		t.Log(value)
	})
	t.Run("atomic-uint32-get-add", func(t *testing.T) {
		var c3 Count32
		v1 := c3.Get()
		if v1 == 0 {

		}
		v2 := c3.Increment()
		if v2 == 1 {

		}
	})
	t.Run("atomic-uint32-bytes", func(t *testing.T) {
		var c3 Count32
		v1 := c3.Get()
		if v1 == 0 {

		}
		c3.Increment()
		_ = c3.UnsafeBytes()
	})
	t.Run("atomic-uint32-bytes-fixed", func(t *testing.T) {
		var c3 Count32
		v1 := c3.Get()
		if v1 == 0 {

		}
		c3.Increment()
		_ = c3.UnsafeBytesFixed()
	})
	t.Run("atomic-uint32-bytes-safe", func(t *testing.T) {
		var c3 Count32
		v1 := c3.Get()
		if v1 == 0 {

		}
		c3.Increment()
		_ = c3.SafeBytes()
	})
}
