// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package detect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIpv4Methods(t *testing.T) {
	t.Run("isIpv4-net-valid", func(t *testing.T) {
		v := IsIpv4Net("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-net-invalid", func(t *testing.T) {
		v := IsIpv4Net("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("isIpv4-regex-valid", func(t *testing.T) {
		v := IsIpv4Regex("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-regex-invalid", func(t *testing.T) {
		v := IsIpv4Regex("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("isIpv4-simple-valid", func(t *testing.T) {
		v := IsIpv4Simple("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-simple-invalid", func(t *testing.T) {
		v := IsIpv4Simple("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("isIpv4-valid", func(t *testing.T) {
		v := IsIpv4("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-invalid", func(t *testing.T) {
		v := IsIpv4("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("custom-atoi", func(t *testing.T) {
		v, err := integerAtoi("89789")
		assert.Nil(t, err)
		assert.Equal(t, v, 89789, "failed to detect a invalid ipv4")
	})
}
