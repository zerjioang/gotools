// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip44

import (
	"crypto/sha512"
	"testing"

	"github.com/zerjioang/gotools/lib/bip32"
	"github.com/zerjioang/gotools/lib/bip39"
)

func BenchmarkBIP44(b *testing.B) {
	b.Run("NewKeyFromMnemonic", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"
			_, _ = NewKeyFromMnemonic(mnemonic, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
		}
	})

	b.Run("NewKeyFromMasterKey", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			mnemonic := "fork mutual wise slush quality ripple purse shiver case whisper derive ball duty fabric autumn drill account kidney case clean hollow food either soul"
			seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
			if err.Occur() {
				b.Error(err)
			}

			masterKey, mErr := bip32.NewMasterKey(seed, "Bitcoin seed", sha512.New)
			if mErr != nil {
				b.Error(err)
			}

			fKey, mErr := NewKeyFromMasterKey(masterKey, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
			if mErr != nil {
				b.Error(err)
			}
			if fKey == nil {
				b.Error("key is null")
			}
		}
	})
}
