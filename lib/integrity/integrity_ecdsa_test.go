// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package integrity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"testing"
)

// Test for ECDSA
func TestECDSAIntegrity(t *testing.T) {

	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := priv.PublicKey

	//t.Log(encode(priv, &pub))

	//create test message
	str := []byte("Lorem Ipsum dolor sit Amet")
	// hash test message
	h := sha256.New()
	h.Write(str)
	signhash := h.Sum(nil)

	r, s, err := ecdsaSign(signhash, priv)
	if err != nil {
		t.Error("signing hash stack", err)
	}
	//now verify
	verify := ecdsaVerify(signhash, r, s, &pub)
	t.Log("signature verification result", verify)
	if !verify {
		t.Error("ecsa sign and verify process failed")
	}

	t.Log(r, s)
	raw := PointsToDER(r, s)
	t.Log(raw)

	verify = ecdsaVerify([]byte("different source message"), r, s, &pub)
	t.Log("signature verification result", verify)
	if verify {
		t.Error("ecsa sign and verify process failed")
	}
}
