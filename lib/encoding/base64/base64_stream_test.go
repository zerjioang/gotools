package base64

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/encoding/generic"
)

const (
	plainText      = "Hello, world from b64 benchmark"
	exampleEncoded = "SGVsbG8sIHdvcmxkIGZyb20gYjY0IGJlbmNobWFyaw=="
	// sonrie, sonreir, sonreido
)

type testWriter struct {
	generic.StreamInterface
}

func (w *testWriter) Write(b byte) {
}

func TestBase64(t *testing.T) {
	t.Run("native-go", func(t *testing.T) {
		raw := []byte(plainText)
		encoded := base64.StdEncoding.EncodeToString(raw)
		decodedRaw, err := base64.StdEncoding.DecodeString(encoded)
		assert.Nil(t, err)
		assert.Equal(t, raw, decodedRaw)
		t.Log(string(encoded))
	})
	t.Run("custom", func(t *testing.T) {
		t.Run("encode-string", func(t *testing.T) {
			raw := []byte(plainText)
			encoded := EncodeToString(raw)
			assert.Equal(t, exampleEncoded, encoded)
			t.Log(string(encoded))
		})
		t.Run("encode-stream", func(t *testing.T) {
			raw := []byte(plainText)
			EncodeToStream(raw, new(testWriter).Write)
		})
	})
}
