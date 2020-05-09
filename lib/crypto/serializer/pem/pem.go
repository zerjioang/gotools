package pem

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
)

func RsaPublicToPEM(pubkey *rsa.PublicKey) ([]byte, error) {
	var buf bytes.Buffer
	asn1Bytes, err := asn1.Marshal(*pubkey)
	if err != nil {
		return nil, err
	}
	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	encErr := pem.Encode(&buf, pemkey)
	return buf.Bytes(), encErr
}

func RsaPrivateToPEM(key *rsa.PrivateKey) ([]byte, error) {
	var buf bytes.Buffer
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err := pem.Encode(&buf, privateKey)
	return buf.Bytes(), err
}
