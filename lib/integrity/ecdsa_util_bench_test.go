// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package integrity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"
)

var (
	p224    = elliptic.P224()
	rns     = rand.Reader
	priv, _ = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	str     = []byte("Lorem Ipsum dolor sit Amet")
	r, s    *big.Int
)

func BenchmarkECDSA(b *testing.B) {
	b.Run("create", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = ecdsa.GenerateKey(p224, rns)
		}
	})
	b.Run("sign", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			r, s, _ = ecdsaSign(str, priv)
		}
	})
	b.Run("verify", func(b *testing.B) {
		priv, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
		pub := priv.PublicKey

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			//now verify
			_ = ecdsaVerify(str, r, s, &pub)
		}
	})
	b.Run("sign-integrity-msg", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			SignMsgWithIntegrity("hello-world")
		}
	})
}
