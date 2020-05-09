// +build cgo

package cgo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Run("call-cgo-version", func(t *testing.T) {
		vstr := CgoVersion()
		assert.NotNil(t, vstr)
		assert.True(t, strings.Index(vstr, "gcc ") != -1)
		t.Log(vstr)
	})
	t.Run("call-cgo-version-flush", func(t *testing.T) {
		vstr := CgoVersion()
		assert.NotNil(t, vstr)
		assert.True(t, strings.Index(vstr, "gcc ") != -1)
		t.Log(vstr)
		FlushCache()
		vstr2 := CgoVersion()
		assert.NotNil(t, vstr2)
		assert.True(t, strings.Index(vstr2, "gcc ") != -1)
		t.Log(vstr2)
	})
}
