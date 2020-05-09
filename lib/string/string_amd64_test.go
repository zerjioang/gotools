package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsAssembly(t *testing.T) {
	t.Run("is-digit", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			r := IsDigit('0')
			t.Log(r)
			assert.True(t, r)
		})
		t.Run("false", func(t *testing.T) {
			r := IsDigit('h')
			t.Log(r)
			assert.False(t, r)
		})
	})
	t.Run("is-numeric-array", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			example := []byte("1485545485")
			r := IsNumericArray(example)
			assert.True(t, r)
		})
		t.Run("false", func(t *testing.T) {
			example := []byte("foo")
			r := IsNumericArray(example)
			assert.False(t, r)
		})
	})
	t.Run("lowercase", func(t *testing.T) {
		example := []byte("HELLO WORLD")
		LowerCase(example)
		t.Log(string(example))
	})
}
