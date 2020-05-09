package jsonboost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssembly(t *testing.T) {
	t.Run("lookup-json-error-code-1", func(t *testing.T) {
		result, err := Lookup("", "")
		assert.NotNil(t, err)
		assert.Equal(t, result, nil)
	})
	t.Run("lookup-json-error-code-2", func(t *testing.T) {
		result, err := Lookup("{}", "")
		assert.NotNil(t, err)
		assert.Equal(t, result, nil)
	})
	t.Run("lookup-json-error-3", func(t *testing.T) {
		// this test must return a corrupted json error
		result, err := Lookup(`{"a": 23`, "a")
		assert.NotNil(t, err)
		assert.Equal(t, result, nil)
	})
	t.Run("lookup-json-extract-value-a", func(t *testing.T) {
		result, err := Lookup(`{"a": 23}`, "a")
		assert.Nil(t, err)
		assert.Equal(t, result, 23)
	})
}
