package multihash_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/multihash"
)

func TestMultihash(t *testing.T) {
	t.Run("encode-md2", func(t *testing.T) {
		plain := "hello-world"
		raw, err := multihash.Encode(multihash.Md2, []byte(plain))
		assert.NoError(t, err)
		assert.NotNil(t, raw)
		assert.Equal(t, hex.EncodeToString(raw), "d36f37b1678c4b9e82380bf2c95d3de3")
	})
}
