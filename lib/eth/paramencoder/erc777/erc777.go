// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package erc777

import (
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/logger"
)

const (
	defaultAbi = `[
	{
		"constant": true,
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "granularity",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "from",
				"type": "address"
			},
			{
				"name": "to",
				"type": "address"
			},
			{
				"name": "amount",
				"type": "uint256"
			},
			{
				"name": "userData",
				"type": "bytes"
			},
			{
				"name": "operatorData",
				"type": "bytes"
			}
		],
		"name": "operatorSend",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "owner",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "operator",
				"type": "address"
			}
		],
		"name": "authorizeOperator",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "to",
				"type": "address"
			},
			{
				"name": "amount",
				"type": "uint256"
			},
			{
				"name": "userData",
				"type": "bytes"
			}
		],
		"name": "send",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "to",
				"type": "address"
			},
			{
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "send",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "operator",
				"type": "address"
			},
			{
				"name": "tokenHolder",
				"type": "address"
			}
		],
		"name": "isOperatorFor",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "operator",
				"type": "address"
			}
		],
		"name": "revokeOperator",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "userData",
				"type": "bytes"
			},
			{
				"indexed": false,
				"name": "operatorData",
				"type": "bytes"
			}
		],
		"name": "Sent",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "operatorData",
				"type": "bytes"
			}
		],
		"name": "Minted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "from",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "amount",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "userData",
				"type": "bytes"
			},
			{
				"indexed": false,
				"name": "operatorData",
				"type": "bytes"
			}
		],
		"name": "Burned",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "tokenHolder",
				"type": "address"
			}
		],
		"name": "AuthorizedOperator",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "operator",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "tokenHolder",
				"type": "address"
			}
		],
		"name": "RevokedOperator",
		"type": "event"
	}
]`
)

const (
	//preload methods identifiers
	AuthorizeOperator = "959b8c3f" // "authorizeOperator(address)
	BalanceOf         = "70a08231" // "balanceOf(address)
	Granularity       = "556f0dc7" // "granularity()
	IsOperatorFor     = "d95b6371" // "isOperatorFor(address,address)
	Name              = "06fdde03" // "name()
	OperatorSend      = "62ad1b83" // "operatorSend(address,address,uint256,bytes,bytes)
	RevokeOperator    = "fad8b32a" // "revokeOperator(address)
	Send              = "d0679d34" // "send(address,uint256)
	SendWithBytes     = "9bd9bbc6" // "send(address,uint256,bytes)
	Symbol            = "95d89b41" // "symbol()
	TotalSupply       = "18160ddd" // "totalSupply()
)

var (
	// read only variables for ERC721
	erc777AbiModel *abi.ABI
)

func init() {
	erc777AbiModel = new(abi.ABI)
	unmErr := erc777AbiModel.UnmarshalJSON([]byte(defaultAbi))
	if unmErr != nil {
		logger.Error("failed to load ERC777 interaction model internals")
		return
	}
}
