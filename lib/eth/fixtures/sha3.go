// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package fixtures

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

// NewKeccak256 creates a new Keccak-256 hash.
func NewKeccak256() hash.Hash {
	return sha3.New256()
}

// NewKeccak512 creates a new Keccak-512 hash.
func NewKeccak512() hash.Hash {
	return sha3.New512()
}
