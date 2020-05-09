// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip44

import (
	"crypto/sha512"

	"github.com/zerjioang/gotools/lib/bip32"
	"github.com/zerjioang/gotools/lib/bip39"
)

// path schema: we define the following 5 levels in BIP32:
// Apostrophe in the path indicates that BIP32 hardened derivation is used.
// m / purpose' / coin_type' / account' / change / address_index

// Purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation.
// It indicates that the subtree of this node is used according to this specification.
const Purpose uint32 = 0x8000002C

//https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
//https://github.com/satoshilabs/slips/blob/master/slip-0044.md
//https://github.com/FactomProject/FactomDocs/blob/master/wallet_info/wallet_test_vectors.md

func NewKeyFromMnemonic(mnemonic string, coin, account, chain, address uint32) (*bip32.Key, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err.Occur() {
		return nil, err
	}
	masterKey, mErr := bip32.NewMasterKey(seed, "Bitcoin seed", sha512.New)
	if mErr != nil {
		return nil, err
	}
	return NewKeyFromMasterKey(masterKey, coin, account, chain, address)
}

func NewKeyFromMasterKey(masterKey *bip32.Key, coin, account, chain, address uint32) (*bip32.Key, error) {
	child, err := masterKey.NewChildKey(Purpose)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(coin)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(account)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(chain)
	if err != nil {
		return nil, err
	}

	child, err = child.NewChildKey(address)
	if err != nil {
		return nil, err
	}

	return child, nil
}
