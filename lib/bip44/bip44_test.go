// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip44

import (
	"crypto/sha512"
	"testing"

	"github.com/zerjioang/gotools/lib/bip32"
	"github.com/zerjioang/gotools/lib/bip39"
)

func TestNewKeyFromMnemonic(t *testing.T) {
	mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"
	fKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if fKey.String() != "xprvA2vH8KdcBBKhMxhENJpJdbwiU5cUXSkaHR7QVTpBmusgYMR8NsZ4BFTNyRLUiaPHg7UYP8u92FJkSEAmmgu3PDQCoY7gBsdvpB7msWGCpXG" {
		t.Errorf("Invalid Factoid key - %v", fKey.String())
	}

	ecKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if ecKey.String() != "xprvA2ziNegvZRfAAUtDsjeS9LvCP1TFXfR3hUzMcJw7oYAhdPqZyiJTMf1ByyLRxvQmGvgbPcG6Q569m26ixWsmgTR3d3PwicrG7hGD7C7seJA" {
		t.Errorf("Invalid EC key - %v", ecKey.String())
	}
}

func TestNewKeyFromMasterKey(t *testing.T) {
	mnemonic := "fork mutual wise slush quality ripple purse shiver case whisper derive ball duty fabric autumn drill account kidney case clean hollow food either soul"

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err.Occur() {
		t.Error(err)
	}

	masterKey, mErr := bip32.NewMasterKey(seed, "Bitcoin seed", sha512.New)
	if mErr != nil {
		t.Error(err)
	}

	fKey, mErr := NewKeyFromMasterKey(masterKey, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if mErr != nil {
		t.Error(err)
	}
	if fKey.String() != "xprvA2ozvmqhejZwf53z9jMcYEMqar7rW4v8fFhkiQgMWSiRwvU1t5v5GEWZdk9hYf4MEWDBFktwTMuE6D2uzCkeBvSB327Dfbha41B18Vn54LK" {
		t.Errorf("Invalid Factoid key - %v", fKey.String())
	}

	ecKey, mErr := NewKeyFromMasterKey(masterKey, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if mErr != nil {
		t.Error(err)
	}
	if ecKey.String() != "xprvA47GVbRXBW8tCq5NHa6ZmWU3ReDmySFjddVH2951oHErgteBAnJzALPiUw6HeChQmajqzAN7nW5j53o569ehM4QAb7qQs7MsJGyJSnAqMn7" {
		t.Errorf("Invalid EC key - %v", ecKey.String())
	}
}

func TestBitLength(t *testing.T) {
	child, err := NewKeyFromMnemonic(
		"element fence situate special wrap snack method volcano busy ribbon neck sphere",
		TypeFactomFactoids,
		2147483648,
		0,
		19,
	)

	if err != nil {
		t.Errorf("%v", err)
	}
	if len(child.Key) != 32 {
		t.Errorf("len: %d, child.Key:%x\n", len(child.Key), child.Key)
		t.Errorf("%v", child.String())
	}

	child, err = NewKeyFromMnemonic(
		"element fence situate special wrap snack method volcano busy ribbon neck sphere",
		TypeFactomFactoids,
		2147483648,
		1,
		19,
	)

	if err != nil {
		t.Errorf("%v", err)
	}
	if len(child.Key) != 32 {
		t.Errorf("len: %d, child.Key:%x\n", len(child.Key), child.Key)
		t.Errorf("%v", child.String())
	}
}
