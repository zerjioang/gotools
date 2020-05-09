// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package badips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadIps(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		Init("./list_any_3", true)
		assert.True(t, Size() > 0)
	})

	t.Run("get-list", func(t *testing.T) {
		Init("./list_any_3", true)
		assert.True(t, Size() > 0)

		l := GetBadIPList()
		assert.NotNil(t, l)
	})
	t.Run("contains-true", func(t *testing.T) {
		Init("./list_any_3", true)
		assert.True(t, Size() > 0)

		l := GetBadIPList()
		result := l.Contains("31.6.220.31")
		assert.True(t, result)
	})
	t.Run("contains-false", func(t *testing.T) {
		Init("./list_any_3", true)
		assert.True(t, Size() > 0)

		l := GetBadIPList()
		result := l.Contains("127.0.0.1")
		assert.False(t, result)
	})
}
