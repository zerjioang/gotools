package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dataBytes         = []byte{45, 78, 125, 4, 80, 50, 50, 50}
	dataUint64 uint64 = 0x2d4e7d0450323232
)

func TestEncoding(t *testing.T) {
	t.Run("BytesToUint64", func(t *testing.T) {
		n := BytesToUint64(dataBytes)
		assert.Equal(t, n, dataUint64)
	})
	t.Run("Uint64ToBytes", func(t *testing.T) {
		b := Uint64ToBytes(dataUint64)
		assert.Equal(t, b, dataBytes)
	})
}
