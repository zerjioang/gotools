package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsAsm(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		r := Add(2, 4)
		t.Log(r)
		assert.Equal(t, r, 6)
	})
}
