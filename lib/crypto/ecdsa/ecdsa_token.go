package ecdsa

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	jwt "github.com/zerjioang/gotools/thirdparty/jwt-go"
)

func GenEcdsaToken() (*ecdsa.PrivateKey, string, error) {
	// create a valid ecdsa key pair
	priv, _, err := GenerateECDSAKit()
	if err != nil {
		return nil, "", err
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		//issuer
		"iss": "issuer",
		// subject or user of the token
		"sub": "anonymous",
		// audience of the token
		"aud": "all",
		// expiration time
		"exp": now.Add(30 * time.Minute).Unix(),
		// not before
		"nbf": now.Unix(),
		// issued at
		"iat": now.Unix(),
		// unique token id
		"jti": "",
		// extra data
		"type": "register",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString(priv)
	if err != nil {
		return nil, "", err
	}
	return priv, tokenStr, nil
}

func ECDSAPrivateToPem(key *ecdsa.PrivateKey) (string, error) {
	if key != nil {
		raw, err := x509.MarshalECPrivateKey(key)
		if err != nil {
			return "", err
		}
		var pemPrivateBlock = &pem.Block{
			Type:  "ECDSA PRIVATE KEY",
			Bytes: raw,
		}
		rawPem := pem.EncodeToMemory(pemPrivateBlock)
		return string(rawPem), nil
	}
	return "", errors.New("private key missing")
}

func ECDSAPublicToPem(key *ecdsa.PublicKey) (string, error) {
	if key != nil {
		raw, err := x509.MarshalPKIXPublicKey(key)
		if err != nil {
			return "", err
		}
		var pemPublicBlock = &pem.Block{
			Type:  "ECDSA PUBLIC KEY",
			Bytes: raw,
		}
		rawPem := pem.EncodeToMemory(pemPublicBlock)
		return string(rawPem), nil
	}
	return "", errors.New("private key missing")
}
