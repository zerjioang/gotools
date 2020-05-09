// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"encoding/hex"
	"testing"

	"github.com/zerjioang/gotools/lib/eth/fixtures"
)

const (
	localGanache           = "http://127.0.0.1:8545"
	ganacheExpectedVersion = "EthereumJS TestRPC/v2.5.3/ethereum-js"
	ganacheNetVersion      = "1552315719982"
	address0               = "0x428a40d2452976feb126bdd272faaab0652f89b3"
	address0key            = "030cde5208870c9d15f0d598c853a10fa39c6fcba2220cd161370b8328563337"
)

func TestEthRPC_Call(t *testing.T) {
	t.Run("web3_clientVersion", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		// Returns the current client version.
		result, err := client.Web3ClientVersion()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != ganacheExpectedVersion {
				t.Error("failed to get client version. wrong response")
			}
		}
	})
	t.Run("web3_sha3", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		// Returns Keccak-256 (not the standardized SHA3-256) of the given data.
		result, err := client.Web3Sha3([]byte("hello-world"))
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			// test with https://emn178.github.io/online-tools/keccak_256.html
			if result != "0xd41bad2284cfa351467b5db9418bbe3a5c02162c02ee585f07e5553d823ebad9" {
				t.Error("failed to get sha3. wrong response")
			}
		}
	})
	t.Run("net_version", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.NetVersion()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != ganacheNetVersion {
				t.Error("failed to get net version. wrong response")
			}
		}
	})
	t.Run("net_version", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.NetListening()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			// test with https://emn178.github.io/online-tools/keccak_256.html
			if !result {
				t.Error("ganache test server is not listening")
			}
		}
	})
	t.Run("net_peerCount", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.NetPeerCount()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != 0 {
				t.Error("ganache test server is not responding to peer count correctly. should be 0")
			}
		}
	})
	t.Run("eth_protocolVersion", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthProtocolVersion()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != "62" && result != "63" {
				t.Error("ganache is not responsing with expected protocol version")
			}
		}
	})
	t.Run("eth_syncing", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthSyncing()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result == nil || result.IsSyncing == true {
				t.Error("ganache wrong response")
			}
		}
	})
	t.Run("eth_coinbase", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthCoinbase()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if len(result) != 42 {
				t.Error("failed to get coinbase address")
			}
		}
	})
	t.Run("eth_mining", func(t *testing.T) {
	})
	t.Run("eth_hashrate", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthHashRate()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result > 0 {
				t.Error("failed to get coinbase address")
			}
		}
	})
	t.Run("eth_gasprice", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthGasPrice()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != 20000000000 {
				t.Error("failed to get gas price")
			}
		}
	})
	t.Run("eth_accounts", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		result, err := client.EthAccounts()
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if len(result) != 10 {
				t.Error("failed to get accounts")
			}
		}
	})
	t.Run("ethereum-node-sign", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		// send sign nrequest of plain text
		result, err := client.EthSign(address0, "foo-bar")
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != "0x0cbef474e48d6429dbc27f356cdd3e92086972415b7329cdbf2b7c1d4d314889279db69d702530620b2b46578483114b977035eaa39058b6f006208e1704465801" {
				t.Error("failed to node sign")
			}
		}
	})
	t.Run("hash-msg", func(t *testing.T) {
		hashed, message := TextAndHash([]byte("foo-bar"))
		hexhash := fixtures.Encode(hashed)
		// print hex encoded hash string
		t.Log(hexhash)
		//print plain raw message
		t.Log(message)
		if hexhash != "0x1378cc2b893327ad6b4c3a65929b405ac3373027e2a5eb9a498ea6353fb9e34d" {
			t.Error("failed to keccack hash message")
		}
	})
	t.Run("ethereum-node-sign-2", func(t *testing.T) {
		client := NewDefaultRPC(localGanache, true, nil)
		// send sign request of hashed message
		hashed, message := TextAndHash([]byte("foo-bar"))
		t.Log(string(hashed))
		t.Log(fixtures.Encode(hashed))
		t.Log(message)
		result, err := client.EthSign(address0, fixtures.Encode(hashed))
		if err != nil {
			t.Error("failed to post", err)
		} else {
			t.Log(result)
			if result != "0x2398cc14454bf83a74ab0dba8a34e3cedc640b3c1380751363bbbd69742b87b52512d2dff71baf4a27a234dfd6f26a2e22b008cba0af08bba65694f2f81524a701" {
				t.Error("failed to node sign")
			}
		}
	})
	t.Run("ethereum-offline-sign", func(t *testing.T) {
		// The sign method calculates an Ethereum specific signature with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))).
		signature, err := LocalSigning(address0, address0key, "foo-bar")
		if err != nil {
			t.Error(err)
		} else {
			signatureStr := hex.EncodeToString(signature)
			t.Log(signatureStr)
			if signatureStr != "0x2398cc14454bf83a74ab0dba8a34e3cedc640b3c1380751363bbbd69742b87b52512d2dff71baf4a27a234dfd6f26a2e22b008cba0af08bba65694f2f81524a701" {
				t.Error("failed to local sign")
			}
		}
	})
	t.Run("install_contract", func(t *testing.T) {
		/*
			curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x428a40d2452976feb126bdd272faaab0652f89b3", "data": "606060405260728060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480636ffa1caa146037576035565b005b604b60048080359060200190919050506061565b6040518082815260200191505060405180910390f35b6000816002029050606d565b91905056"}],"id":1}'

			The response should be similar to:

			{
			  "jsonrpc": "2.0",
			  "id": 1,
			  "result": "0x6755e3fa3a927b0b0c43764fc014334fd00cc121ec645566776ec2c8c2ae41c2"
			}

			curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0x8290c22bd9b4d61bc57222698799edd7bbc8df5214be44e239a95f679249c59c"],"id":1}'

			Then it returns

			{
			  "jsonrpc": "2.0",
			  "id": 1,
			  "result": {
				"transactionHash": "0x6755e3fa3a927b0b0c43764fc014334fd00cc121ec645566776ec2c8c2ae41c2",
				"transactionIndex": "0x0",
				"blockHash": "0xc17d6960b338ef21be44d48208ead2f6d8e8a63d42c89ef61d28ddc4ee7c5d87",
				"blockNumber": "0x897",
				"cumulativeGasUsed": "0x015f90",
				"gasUsed": "0x015f90",
				"contractAddress": "0x73ec81da0c72dd112e06c09a6ec03b5544d26f05",
				"logs": [],
				"from": "0xcd2a3d9f938e13cd947ec05abc7fe734df8dd826",
				"to": "0x00",
				"root": "0x9182b4ffc855052ce52ad12d38eb2de42f068f3b6df9c218e7dba46fc404be21"
			  }
			}

		*/
	})
}
