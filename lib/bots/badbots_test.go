package bots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBadBots(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		Init("./bots", true)
		assert.True(t, Size() > 0)
	})
}
