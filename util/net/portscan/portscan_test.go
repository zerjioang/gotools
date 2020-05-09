package portscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOpenPort(t *testing.T) {
	t.Run("negative-port", func(t *testing.T) {
		op, err := IsOpenPort("127.0.0.1", -20, 3)
		assert.Equal(t, op, false)
		assert.NotNil(t, err.Error())
	})
	t.Run("too-high-port", func(t *testing.T) {
		op, err := IsOpenPort("127.0.0.1", 8000000, 3)
		assert.Equal(t, op, false)
		assert.NotNil(t, err.Error())
	})
	t.Run("valid-port", func(t *testing.T) {
		// todo ganache cli should be running on port 7545
		op, err := IsOpenPort("127.0.0.1", 7545, 3)
		assert.True(t, op)
		assert.True(t, err.None())
	})
}

func TestFindJsonRpcPort(t *testing.T) {
	t.Run("find-127.0.0.1", func(t *testing.T) {
		// todo ganache cli should be running on port 7545
		port := FindJsonRpcPort("127.0.0.1")
		assert.True(t, port != -1)
		assert.Equal(t, port, 7545)
	})
}
