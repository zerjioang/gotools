// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package aeshash_test

import (
	"testing"

	"github.com/zerjioang/gotools/lib/hash/aeshash"
)

func TestAESHash(t *testing.T) {
	t.Run("aeshash", func(t *testing.T) {
		val := aeshash.Hash("cheese")
		if val != 1315767268 {
			t.Errorf("Expected 1315767268, got %d", val)
		}
	})
}
