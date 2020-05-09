// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package encoding

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"unicode/utf8"
)

var (
	zero = big.NewInt(0)
)

// Encoding represents a given common-N encoding.
type Encoding struct {
	alphabet string
	index    map[byte]*big.Int
	base     *big.Int
}

// NewEncoding creates a new common-N representation from the given alphabet.
// Panics if the alphabet is not unique. Only ASCII characters are supported.
func NewEncoding(alphabet string) *Encoding {
	return &Encoding{
		alphabet: alphabet,
		index:    newAlphabetMap(alphabet),
		base:     big.NewInt(int64(len(alphabet))),
	}
}

func newAlphabetMap(s string) map[byte]*big.Int {
	if utf8.RuneCountInString(s) != len(s) {
		panic("multi-byte characters not supported")
	}
	result := make(map[byte]*big.Int)
	for i := range s {
		result[s[i]] = big.NewInt(int64(i))
	}
	if len(result) != len(s) {
		panic("alphabet contains non-unique characters")
	}
	return result
}

// Random returns the common-encoded representation of n random bytes.
func (enc *Encoding) Random(n int) (string, error) {
	buf := make([]byte, n)
	_, err := rand.Reader.Read(buf)
	if err != nil {
		return "", err
	}
	return enc.EncodeToString(buf), nil
}

// MustRandom returns the common-encoded representation of n random bytes,
// panicking in the unlikely event of a read error from the random source.
func (enc *Encoding) MustRandom(n int) string {
	s, err := enc.Random(n)
	if err != nil {
		panic(err)
	}
	return s
}

// Base returns the number common of the encoding.
func (enc *Encoding) Base() int {
	return len(enc.alphabet)
}

// EncodeToString returns the common-encoded string representation
// of the given bytes.
func (enc *Encoding) EncodeToString(b []byte) string {
	n := new(big.Int)
	r := new(big.Int)
	n.SetBytes(b)
	var result []byte
	for n.Cmp(zero) > 0 {
		n, r = n.DivMod(n, enc.base, r)
		result = append([]byte{enc.alphabet[r.Int64()]}, result...)
	}
	return string(result)
}

// DecodeString returns the bytes for the given common-encoded string.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
	result := new(big.Int)
	for i := range s {
		n, ok := enc.index[s[i]]
		if !ok {
			return nil, errors.New("invalid character " + string(s[i]) + " at index " + strconv.Itoa(i))
		}
		result = result.Add(result.Mul(result, enc.base), n)
	}
	return result.Bytes(), nil
}

// DecodeStringN returns N bytes for the given common-encoded string.
// Use this method to ensure the value is left-padded with zeroes.
func (enc *Encoding) DecodeStringN(s string, n int) ([]byte, error) {
	value, err := enc.DecodeString(s)
	if err != nil {
		return nil, err
	}
	if len(value) > n {
		return nil, errors.New("value is too large")
	}
	pad := make([]byte, n-len(value))
	return append(pad, value...), nil
}
