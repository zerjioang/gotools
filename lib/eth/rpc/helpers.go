// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"bytes"
	"math/big"
	"strconv"
	"strings"

	"github.com/zerjioang/gotools/util/str"

	"github.com/zerjioang/gotools/lib/eth/fixtures/crypto"
	"github.com/zerjioang/gotools/lib/logger"
	"golang.org/x/crypto/sha3"

	"github.com/zerjioang/gotools/lib/eth/fixtures"
)

// ParseInt parse hex string value to int
func ParseInt(value string) (int64, error) {
	if value == "0x0" {
		return 0, nil
	}
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (*big.Int, string, error) {
	i := new(big.Int)
	if value == "0x0" {
		return i, value, nil
	}
	//check if the value starts with 0x. is so, remove that content
	// 48 decimal is '0' in ascii
	// 120 decimal is 'x' is ascii
	if len(value) > 2 && value[0] == 48 && value[1] == 120 {
		value = value[2:]
	}
	i, _ = i.SetString(value, 16)
	return i, value, nil
}

// ParseBigInt parse hex string value to big.Int
func ParseHexToInt(value string) (int64, error) {
	return strconv.ParseInt(value, 16, 64)
}

// Int64ToHex convert int64 to hexadecimal representation
func Int64ToHex(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

// IntToHex convert int to hexadecimal representation
func IntToHex(i int) string {
	return Int64ToHex(int64(i))
}

// BigToHex covert big.Int to hexadecimal representation
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return fixtures.Encode(bigInt.Bytes())
}

func HexToBigInt(hex string) *big.Int {
	i := new(big.Int)
	i, _ = i.SetString(hex, 16)
	return i
}

// TextAndHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func TextAndHash(data []byte) ([]byte, string) {
	var b bytes.Buffer
	b.WriteString("\x19Ethereum Signed Message:\n")
	b.Write(str.IntToByte(len(data)))
	b.Write(data)
	msg := b.Bytes()
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(msg)
	return hasher.Sum(nil), string(msg)
}

func LocalSigning(addrHex string, privHex string, plainMessage string) ([]byte, error) {
	key, cErr := crypto.HexToECDSA(privHex)
	if cErr != nil {
		logger.Error("private key load error: ", cErr)
		return nil, cErr
	}

	hashed, message := TextAndHash([]byte("foo-bar"))
	hashHexMsg := fixtures.Encode(hashed)
	logger.Info(string(hashed))
	logger.Info(hashHexMsg)
	logger.Info(message)

	sig, err := crypto.Sign(hashed, key)
	if err != nil {
		logger.Error("signing error: ", err)
		return nil, err
	}
	return sig, err
	//addr := fixtures.HexToAddress(addrHex)
	/*recoveredPub, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		logger.Error("eCRecover error: %s", err)
		return err
	}
	pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	recoveredAddr := fixtures.PubkeyToAddress(*pubKey)
	if addr != recoveredAddr {
		logger.Error("address mismatch: want: %x have: %x", addr, recoveredAddr)
		return err
	}

	// should be equal to SigToPub
	recoveredPub2, err := crypto.SigToPub(msg, sig)
	if err != nil {
		logger.Error("eCRecover error: %s", err)
		return err
	}
	recoveredAddr2 := fixtures.PubkeyToAddress(*recoveredPub2)
	if addr != recoveredAddr2 {
		logger.Error("address mismatch: want: %x have: %x", addr, recoveredAddr2)
		return err
	}*/
}
