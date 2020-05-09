package browser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasGraphicInterface(t *testing.T) {
	t.Run("detect-ui", func(t *testing.T) {
		status := detectUI()
		assert.NotNil(t, status)
	})
	t.Run("has-ui", func(t *testing.T) {
		status := HasGraphicInterface()
		assert.NotNil(t, status)
		t.Log("has ui: ", status)
	})
}
