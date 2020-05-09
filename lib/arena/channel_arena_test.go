package arena

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArena(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		a := NewChannelArena(1000, 1024)
		assert.NotNil(t, a)
	})
	t.Run("pop-push", func(t *testing.T) {
		a := NewChannelArena(1000, 1024)
		assert.NotNil(t, a)
		a.Push(a.Pop())
	})
	t.Run("insert", func(t *testing.T) {
		a := NewChannelArena(1000, 1024)
		assert.NotNil(t, a)
		l := a.Pop()
		l[0] = 'a'
		l[1] = 'b'
		l[2] = 'c'
		a.Push(l)
	})
}
