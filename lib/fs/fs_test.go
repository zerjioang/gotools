package fs_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/fs"
)

func TestFs(t *testing.T) {
	t.Run("pagesize", func(t *testing.T) {
		ps := fs.PageSize()
		t.Log("pagesize: ", ps)
		assert.True(t, ps > 0)
	})
	t.Run("exists", func(t *testing.T) {
		status := fs.Exists("/tmp")
		assert.True(t, status)
	})
	t.Run("read-entropy", func(t *testing.T) {
		data, err := fs.ReadEntropy(rand.Reader, 16)
		assert.Nil(t, err)
		assert.NotNil(t, data)
	})
	t.Run("read-entropy-16", func(t *testing.T) {
		data, err := fs.ReadEntropy16()
		assert.Nil(t, err)
		assert.NotNil(t, data)
		t.Log(data)
	})
	t.Run("buffered-reader", func(t *testing.T) {
		data, size, err := fs.ReadFile(rand.Reader)
		assert.Nil(t, err)
		assert.NotNil(t, size)
		assert.NotNil(t, data)
	})
}
