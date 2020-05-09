package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	plainText      = "hello-world"
	exampleEncoded = "AAwf93s01ikS26e"
	// sonrie, sonreir, sonreido
)

func TestBase62(t *testing.T) {
	t.Run("encode-test", func(t *testing.T) {
		encodedValue := StdEncoding.EncodeToString([]byte(plainText))
		assert.Equal(t, encodedValue, exampleEncoded)
	})
	t.Run("e2e", func(t *testing.T) {
		urlVal := "example-content-1234"
		encodedValue := StdEncoding.EncodeToString([]byte(urlVal))
		assert.NotEqual(t, "", encodedValue)
		rtVal, err := StdEncoding.DecodeString(encodedValue)
		assert.Equal(t, nil, err)
		assert.Equal(t, string(rtVal), urlVal)
	})
}
