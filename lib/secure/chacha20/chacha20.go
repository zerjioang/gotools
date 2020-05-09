package chacha20

import (
	"crypto/rand"

	"github.com/zerjioang/gotools/lib/logger"
	"golang.org/x/crypto/chacha20poly1305"
)

// these method are safe to be used concurrently

func Encrypt(key []byte, message []byte) ([]byte, []byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		logger.Error("Failed to instantiate XChaCha20-Poly1305:", err)
		return nil, nil, err
	}

	// Encryption: 1 read random nonce
	nonce, err := Nonce()
	if err != nil {
		logger.Error("failed to read random nonce", err)
		return nil, nil, err
	}

	// Encryption: 2 make encryption seal
	ciphertext := aead.Seal(nil, nonce, message, nil)
	return ciphertext, nonce, nil
}

func Decrypt(key []byte, nonce []byte, ciphertext []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		logger.Error("Failed to instantiate XChaCha20-Poly1305:", err)
		return nil, err
	}
	// Decryption.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logger.Error("Failed to decrypt or authenticate message:", err)
		return nil, err
	}
	return plaintext, nil
}

func Nonce() ([]byte, error) {
	logger.Debug("generating nonce for chacha20 encoder")
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		logger.Error("failed to read random nonce", err)
		return nil, err
	}
	return nonce, nil
}
