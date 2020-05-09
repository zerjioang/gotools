package rsa

import (
	"crypto/rand"
	"crypto/rsa"
)

var reader = rand.Reader

func GenerateRSA(bitSize uint32) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(reader, int(bitSize))
}
