// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	gohex "github.com/tmthrgd/go-hex"
	hex2 "github.com/zerjioang/gotools/util/hex"
)

func TestEncode(t *testing.T) {
	t.Run("default-encode", func(t *testing.T) {
		result := hex.EncodeToString([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
	t.Run("fast-encode", func(t *testing.T) {
		result := hex2.UnsafeEncodeToString([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
	t.Run("fast-encode-pooled", func(t *testing.T) {
		result := hex2.UnsafeEncodeToStringPooled([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
	t.Run("gohex-encode", func(t *testing.T) {
		result := gohex.EncodeToString([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
}

func TestDecode(t *testing.T) {
	t.Run("default-decode", func(t *testing.T) {
		result, err := hex.DecodeString("746869732d69732d612d74657374")
		assert.Nil(t, err)
		assert.Equal(t, result, []byte("this-is-a-test"))
	})
	t.Run("fast-decode", func(t *testing.T) {
		result, err := hex2.UnsafeDecodeString("746869732d69732d612d74657374")
		assert.Nil(t, err)
		assert.Equal(t, string(result), "this-is-a-test")
	})
}
