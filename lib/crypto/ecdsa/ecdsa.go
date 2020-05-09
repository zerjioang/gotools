package ecdsa

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	//see http://golang.org/pkg/crypto/elliptic/#P256
	p256          = elliptic.P256()
	errNotInCurve = errors.New("key points are not in curve")
)

// this generates a public & private key pair
func GenerateECDSAKey(c elliptic.Curve) (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		return nil, err
	}
	if !c.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
		return nil, errNotInCurve
	}
	return priv, err
}

func GenerateECDSAKit() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	pk, err := GenerateECDSAKey(p256)
	if err == nil {
		return pk, &pk.PublicKey, nil
	}
	return nil, nil, err
}

func ECDSALoadPrivateFromPem(keyPem string) (*ecdsa.PrivateKey, error) {
	var priv *ecdsa.PrivateKey
	block, _ := pem.Decode([]byte(keyPem))
	if block == nil {
		return nil, errors.New("failed to decode pem block")
	}
	if block.Type == "ECDSA PRIVATE KEY" {
		var err error
		var pk crypto.PrivateKey
		pk, err = parsePrivateKey(block.Bytes)
		if err != nil {
			return nil, errors.New("failure parsing private key bytes")
		}
		priv = pk.(*ecdsa.PrivateKey)
		return priv, nil
	}
	return nil, errors.New("invalid pem block found. does not belong to ecdsa private key")
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, fmt.Errorf("found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, fmt.Errorf("failed to parse private key")
}
