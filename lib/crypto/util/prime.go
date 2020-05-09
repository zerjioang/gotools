package util

import (
	"crypto/rand"
	"math/big"
)

// generate a random prime of length `bits`
func RandomPrime(bits int) (p *big.Int, err error) {
	return rand.Prime(rand.Reader, bits)
}
