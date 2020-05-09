// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package eth

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	ethrpc "github.com/zerjioang/gotools/lib/eth/rpc"
)

const (
	ganacheUITestEndpoint   = "HTTP://127.0.0.1:9545"
	ganacheUIAddress0       = "0xcD1C300209FeE0dd6C68c69115C9148f9FDF3102"
	ganacheCliTestEndpoint  = "HTTP://127.0.0.1:9545"
	ganacheCLIAddress0      = "0xa156Cf364ff355c5727AC79e5363377b6F726129"
	ganacheCLIAddressoEIP55 = "0xa156Cf364Ff355c5727aC79E5363377b6f726129"
	// run tests using ganache cli
	address0            = ganacheCLIAddress0
	ganacheTestEndpoint = ganacheCliTestEndpoint
)

func TestConvertAddress(t *testing.T) {
	addr := ConvertAddress(address0)
	t.Log("address converted", addr.Hex())
	assert.Equal(t, addr.Hex(), ganacheCLIAddressoEIP55, "failed to convert account")
}

func TestIsValidAddress(t *testing.T) {
	result := IsValidAddress(address0)
	assert.True(t, result)
}

func TestIsValidBlockHash(t *testing.T) {
	t.Run("empty-string", func(t *testing.T) {
		result := IsValidBlockHash("")
		assert.False(t, result)
	})
	t.Run("latest", func(t *testing.T) {
		result := IsValidBlockHash("315615145")
		assert.False(t, result)
	})
	t.Run("valid", func(t *testing.T) {
		result := IsValidBlockHash("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")
		assert.True(t, result)
	})
}

func TestIsValidBlock(t *testing.T) {
	t.Run("empty-string", func(t *testing.T) {
		result := IsValidBlockNumber("")
		assert.False(t, result)
	})
	t.Run("latest", func(t *testing.T) {
		result := IsValidBlockNumber("latest")
		assert.True(t, result)
	})
	t.Run("earliest", func(t *testing.T) {
		result := IsValidBlockNumber("earliest")
		assert.True(t, result)
	})
	t.Run("earliest-fail", func(t *testing.T) {
		result := IsValidBlockNumber("earliestt")
		assert.False(t, result)
	})
	t.Run("pending", func(t *testing.T) {
		result := IsValidBlockNumber("pending")
		assert.True(t, result)
	})
	t.Run("pending-fail", func(t *testing.T) {
		result := IsValidBlockNumber("pendiiing")
		assert.False(t, result)
	})
	t.Run("hex-string", func(t *testing.T) {
		result := IsValidBlockNumber("0xff")
		assert.True(t, result)
	})
}

func TestIsValidHexSignature(t *testing.T) {
	t.Run("empty-string", func(t *testing.T) {
		result := IsValidHexSignature("")
		assert.False(t, result)
	})
	t.Run("latest", func(t *testing.T) {
		result := IsValidHexSignature("315615145")
		assert.False(t, result)
	})
	t.Run("earliest", func(t *testing.T) {
		result := IsValidBlockNumber("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238")
		assert.True(t, result)
	})
}

func TestIsValidHexPrivateKey(t *testing.T) {
	t.Run("empty-string", func(t *testing.T) {
		result := IsValidHexPrivateKey("")
		assert.False(t, result)
	})
	t.Run("latest", func(t *testing.T) {
		result := IsValidHexPrivateKey("315615145")
		assert.False(t, result)
	})
	t.Run("valid", func(t *testing.T) {
		result := IsValidHexPrivateKey("cdfa05e62455fae56b1fea15607691975db23b6bef5342f9f50505769529d")
		assert.True(t, result)
	})
	t.Run("0x-valid", func(t *testing.T) {
		result := IsValidHexPrivateKey("0xcdfa05e62455fae56b1fea15607691975db23b6bef5342f9f50505769529d")
		assert.True(t, result)
	})
}

func TestGetAccountBalance(t *testing.T) {
	// define the client
	cli := ethrpc.NewDefaultRPC(ganacheTestEndpoint, false, nil)
	expected := big.NewInt(0)
	expected, _ = expected.SetString("100000000000000000000", 10)
	balance, raw, err := cli.EthGetBalance(address0, "latest")
	if err != nil && raw != "" {
		t.Error("failed to get the client", err)
	} else {
		t.Log("readed account balance", balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}

func TestGetAccountBalanceAtBlock(t *testing.T) {
	// define the client
	cli := ethrpc.NewDefaultRPC(ganacheTestEndpoint, true, nil)
	expected := big.NewInt(0)
	expected, _ = expected.SetString("100000000000000000000", 10)
	balance, raw, err := cli.EthGetBalance(address0, "0")
	if err != nil && raw != "" {
		t.Error("failed to get the client", err)
	} else {
		t.Log("readed account balance", balance)
		if balance.Cmp(expected) != 0 {
			t.Error("failed to get balance for ganache account[0]")
		}
	}
}
