package rsa

import (
	"testing"

	"github.com/zerjioang/gotools/lib/crypto/serializer/pem"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRSA(t *testing.T) {
	t.Run("generate-1024", func(t *testing.T) {
		pk, err := GenerateRSA(1024)
		assert.NoError(t, err)
		assert.NotNil(t, pk)
		assert.Equal(t, pk.Size(), 1024/8)
	})
	t.Run("encode-private-key", func(t *testing.T) {
		t.Run("to-pem", func(t *testing.T) {
			pk, err := GenerateRSA(1024)
			assert.NoError(t, err)
			assert.NotNil(t, pk)
			assert.Equal(t, pk.Size(), 1024/8)
			raw, err := pem.RsaPrivateToPEM(pk)
			assert.NoError(t, err)
			assert.NotNil(t, raw)
			t.Log(string(raw))
		})
	})
	t.Run("encode-public-key", func(t *testing.T) {
		t.Run("to-pem", func(t *testing.T) {
			pk, err := GenerateRSA(1024)
			assert.NoError(t, err)
			assert.NotNil(t, pk)
			assert.Equal(t, pk.Size(), 1024/8)
			raw, err := pem.RsaPublicToPEM(&pk.PublicKey)
			assert.NoError(t, err)
			assert.NotNil(t, raw)
			t.Log(string(raw))
		})
	})
}
