// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
	})
}

func TestToLowerAscii(t *testing.T) {
	t.Run("ToLowerAscii-ua-1", func(t *testing.T) {
		val := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		assert.Equal(t, converted, "mozilla/5.0 (macintosh; intel mac os x 10_11_6) applewebkit/601.7.7 (khtml, like gecko) version/9.1.2 safari/601.7.7")
	})
	t.Run("ToLowerAscii-ua-2", func(t *testing.T) {
		val := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:61.0) Gecko/20100101 Firefox/61.0"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "mozilla/5.0 (x11; ubuntu; linux x86_64; rv:61.0) gecko/20100101 firefox/61.0" {
			t.Error("failed to lowercase")
		}
	})
}
