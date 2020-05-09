package ecdsa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	jwt "github.com/zerjioang/gotools/thirdparty/jwt-go"
)

func TestGenEcdsaToken(t *testing.T) {
	t.Run("generate-token", func(t *testing.T) {
		priv, token, err := GenEcdsaToken()
		assert.NoError(t, err)
		assert.NotNil(t, priv)
		assert.NotNil(t, token)
		t.Log(token)
	})
	t.Run("key-to-pem", func(t *testing.T) {
		priv, pub, err := GenerateECDSAKit()
		assert.NoError(t, err)
		assert.NotNil(t, priv)
		assert.NotNil(t, pub)

		strPub, err := ECDSAPublicToPem(pub)
		assert.NoError(t, err)
		assert.NotNil(t, strPub)
		t.Log(strPub)

		str, err := ECDSAPrivateToPem(priv)
		assert.NoError(t, err)
		assert.NotNil(t, str)
		t.Log(str)
	})
	t.Run("validate-ecdsa-token", func(t *testing.T) {
		// sample token string taken from the New example
		priv, tokenString, err := GenEcdsaToken()
		assert.NoError(t, err)
		assert.NotNil(t, priv)
		assert.NotNil(t, tokenString)
		t.Log(tokenString)

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return &priv.PublicKey, nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
	})
	t.Run("validate-ecdsa-token-pem", func(t *testing.T) {
		// sample token string taken from the New example
		tokenString := `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbGwiLCJleHAiOjE1NzA2MjA3NzMsImlhdCI6MTU3MDYxODk3MywiaXNzIjoiaXNzdWVyIiwianRpIjoiIiwibmJmIjoxNTcwNjE4OTczLCJzdWIiOiJhbm9ueW1vdXMiLCJ0eXBlIjoicmVnaXN0ZXIifQ.r88DBFwAqlbGOVgyXg2vG2lMbCeXkOZuNXgQ1FRrZe9qfOFhyrSP1RJhXFqfj81W5tWwRTd9JwmxfyNT1s-TQQ`
		assert.NotNil(t, tokenString)
		t.Log(tokenString)

		privateKeyPem := `-----BEGIN ECDSA PRIVATE KEY-----
MHcCAQEEICFlgU5pTVrZ/muIahj3IKk+3fX2wDQI63dg/Taa63iooAoGCCqGSM49
AwEHoUQDQgAE3QTuRWkvg63wFASYaoOUSY19hvgrJg1YIC4sQM2INE8lGZNERQPH
pVCir9FIziDDzlQcgo6XlRssTlZqAbfkwQ==
-----END ECDSA PRIVATE KEY-----
`

		pkey, err := ECDSALoadPrivateFromPem(privateKeyPem)
		assert.NoError(t, err)
		assert.NotNil(t, pkey)

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return &pkey.PublicKey, nil
		})
		assert.NoError(t, err)
		assert.NotNil(t, token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
	})
}
