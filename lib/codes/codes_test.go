package codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	t.Run("get-code-standard", func(t *testing.T) {
		msg := StatusText(100)
		assert.Equal(t, msg, "Continue")
	})
	t.Run("get-code-optimized", func(t *testing.T) {
		msg := StatusTextOptimized(100)
		assert.Equal(t, msg, "Continue")
	})
	t.Run("is-informational", func(t *testing.T) {
		result := IsInformational(StatusProcessing)
		assert.True(t, result)
	})
	t.Run("is-successful", func(t *testing.T) {
		result := IsSuccessful(StatusOK)
		assert.True(t, result)
	})
}
