package xor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	t.Run("example-string", func(t *testing.T) {
		key := "JOKER"
		data := "hello world from xor"

		encrypted := EncryptDecrypt(data, key)
		fmt.Println("Encrypted:", encrypted)
		assert.Equal(t, encrypted, `"*')=j8$7>.o-7='o3* `)

		decrypted := EncryptDecrypt(encrypted, key)
		fmt.Println("Decrypted:", decrypted)
		assert.Equal(t, decrypted, data)
	})
	t.Run("example-bytes", func(t *testing.T) {
		key := "JOKER"
		data := "hello world from xor"

		encrypted := EncryptDecryptBytes([]byte(data), []byte(key))
		fmt.Println("Encrypted:", string(encrypted))
		assert.Equal(t, string(encrypted), `"*')=j8$7>.o-7='o3* `)

		decrypted := EncryptDecryptBytes([]byte(encrypted), []byte(key))
		fmt.Println("Decrypted:", string(decrypted))
		assert.Equal(t, string(decrypted), data)
	})
}
