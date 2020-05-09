package browser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenBrowser(t *testing.T) {
	t.Run("empty-url", func(t *testing.T) {
		err := OpenBrowser("")
		assert.Nil(t, err)
	})
	t.Run("google-url", func(t *testing.T) {
		err := OpenBrowser("http://google.com")
		assert.Nil(t, err)
	})
	t.Run("localhost-url", func(t *testing.T) {
		err := OpenBrowser("http://localhost:8080")
		assert.Nil(t, err)
	})
}
