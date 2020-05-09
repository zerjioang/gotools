// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"encoding/json"
	"math/big"
	"testing"
)

func BenchmarkHexIntUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		test := struct {
			ID hexInt `json:"id"`
		}{}

		data := []byte(`{"id": "0x1cc348"}`)
		_ = json.Unmarshal(data, &test)
	}
}

func BenchmarkHexBigUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		test := struct {
			ID hexBig `json:"id"`
		}{}

		data := []byte(`{"id": "0x51248487c7466b7062d"}`)
		_ = json.Unmarshal(data, &test)

		b := big.Int{}
		b.SetString("23949082357483433297453", 10)
	}

}

func BenchmarkSyncingUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		syncing := new(Syncing)
		_ = json.Unmarshal([]byte("0"), syncing)

		data := []byte(`{
		"startingBlock": "0x384",
		"currentBlock": "0x386",
		"highestBlock": "0x454"
	}`)
		_ = json.Unmarshal(data, syncing)
	}
}

func BenchmarkTransactionUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tx := new(Transaction)
		_ = json.Unmarshal([]byte("111"), tx)

		data := []byte(`{
        "blockHash": "0x3003694478c108eaec173afcb55eafbb754a0b204567329f623438727ffa90d8",
        "blockNumber": "0x83319",
        "from": "0x201354729f8d0f8b64e9a0c353c672c6a66b3857",
        "gas": "0x15f90",
        "gasPrice": "0x4a817c800",
        "hash": "0xfc7dcd42eb0b7898af2f52f7c5af3bd03cdf71ab8b3ed5b3d3a3ff0d91343cbe",
        "input": "0xe1fa8e8425f1af44eb895e4900b8be35d9fdc28744a6ef491c46ec8601990e12a58af0ed",
        "nonce": "0x6ba1",
        "to": "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf",
        "transactionIndex": "0x3",
        "value": "0x0"
    }`)
		_ = json.Unmarshal(data, tx)
	}
}

func BenchmarkLogUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		log := new(Log)
		_ = json.Unmarshal([]byte("111"), log)

		data := []byte(`{
        "address": "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf",
        "topics": ["0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"],
        "data": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "blockNumber": "0x7f2cd",
        "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
        "transactionIndex": "0x1",
        "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
        "logIndex": "0x6",
        "removed": false
    }`)
		_ = json.Unmarshal(data, log)
	}
}

func BenchmarkTransactionReceiptUnmarshal(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		receipt := new(TransactionReceipt)
		_ = json.Unmarshal([]byte("[1]"), receipt)

		data := []byte(`{
        "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
        "blockNumber": "0x7f2cd",
        "contractAddress": null,
        "cumulativeGasUsed": "0x13356",
        "gasUsed": "0x6384",
        "logs": [{
            "address": "0xd10e3be2bc8f959bc8c41cf65f60de721cf89adf",
            "topics": ["0x78e4fc71ff7e525b3b4660a76336a2046232fd9bba9c65abb22fa3d07d6e7066"],
            "data": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "blockNumber": "0x7f2cd",
            "blockHash": "0x3757b6efd7f82e3a832f0ec229b2fa36e622033ae7bad76b95763055a69374f7",
            "transactionIndex": "0x1",
            "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
            "logIndex": "0x6",
            "removed": false
        }],
        "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000020000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000",
        "root": "0xe367ea197d629892e7b25ea246fba93cd8ae053d468cc5997a816cc85d660321",
        "transactionHash": "0xecd8a21609fa852c08249f6c767b7097481da34b9f8d2aae70067918955b4e69",
        "transactionIndex": "0x1"
    }`)
		_ = json.Unmarshal(data, receipt)
	}
}
