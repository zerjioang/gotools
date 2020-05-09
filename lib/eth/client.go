// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package eth

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"

	"github.com/zerjioang/gotools/lib/eth/fixtures"
	"github.com/zerjioang/gotools/lib/eth/fixtures/crypto"
	"github.com/zerjioang/gotools/lib/eth/fixtures/crypto/secp256k1"
)

// generates new ethereum ECDSA key
// This is the private key which is used for signing transactions and is to be treated
// like a password and never be shared, since who ever is in possesion
// of it will have access to all your funds.
func GenerateNewKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
}

// get the bytes of private key
func GetPrivateKeyBytes(privateKey *ecdsa.PrivateKey) []byte {
	return crypto.FromECDSA(privateKey)
}

// get the private key encoded as 0x..... string
func GetPrivateKeyAsEthString(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privateKey.D.Bytes())
}

// get the bytes of private key
func GetPublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	return privateKey.Public().(*ecdsa.PublicKey)
}

// get ethereum address from its private key
func GetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) fixtures.Address {
	return fixtures.PubkeyToAddress(privateKey.PublicKey)
}

// get the bytes of public key
func GetPublicKeyBytes(pub *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(pub)
}

// The public address is simply the Keccak-256 hash of the public key
// and then we take the last 40 characters (20 bytes) and prefix it with 0x
func GetPublicKeyAsEthString(pub *ecdsa.PublicKey) string {
	b := GetPublicKeyBytes(pub)
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}
