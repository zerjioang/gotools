// +build !cgo

package cgo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Run("call-nocgo-version", func(t *testing.T) {
		vstr := CgoVersion()
		assert.NotNil(t, vstr)
		assert.True(t, strings.Index(vstr, "unknown") != -1)
		t.Log(vstr)
	})
	t.Run("call-nocgo-version-flush", func(t *testing.T) {
		vstr := CgoVersion()
		assert.NotNil(t, vstr)
		assert.True(t, strings.Index(vstr, "unknown") != -1)
		t.Log(vstr)
		FlushCache()
		vstr2 := CgoVersion()
		assert.NotNil(t, vstr2)
		assert.True(t, strings.Index(vstr2, "unknown") != -1)
		t.Log(vstr2)
	})
}
