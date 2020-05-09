// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package integrity

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
	"testing"
)

// Test for DSA(Digital Signature Algorithm) package functions
func TestIntegrity(t *testing.T) {

	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params
	genErr := dsa.GenerateKey(privatekey, rand.Reader) // this generates a public & private key pair
	if genErr != nil {
		t.Error(genErr)
	}
	var pubkey dsa.PublicKey
	pubkey = privatekey.PublicKey

	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privatekey)

	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubkey)

	// Sign
	var h hash.Hash
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)

	io.WriteString(h, "This is the message to be signed and verified!")
	signhash := h.Sum(nil)

	r, s, err := dsa.Sign(rand.Reader, privatekey, signhash)
	if err != nil {
		fmt.Println(err)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)

	// Verify
	verifystatus := dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be true

	// we add additional data to change the signhash
	io.WriteString(h, "This message is NOT to be signed and verified!")
	signhash = h.Sum(nil)

	verifystatus = dsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus) // should be false
}
