// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package erc223

import (
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/logger"
)

const (
	defaultAbi = `[
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
		"inputs": [
			{
				"name": "who",
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
				"name": "to",
				"type": "address"
			},
			{
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "transfer",
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
				"name": "value",
				"type": "uint256"
			},
			{
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "transfer",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
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
				"name": "value",
				"type": "uint256"
			},
			{
				"indexed": false,
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "Transfer",
		"type": "event"
	}
]`
	BalanceOfParams         = "70a08231" //balanceOf(address)
	TotalSupplyParams       = "18160ddd" //totalSupply()
	TransferParams          = "a9059cbb" //transfer(address,uint256)
	TransferParamsWithBytes = "be45fd62" //transfer(address,uint256,bytes)
)

var (
	// read only variables for ERC721
	erc223AbiModel *abi.ABI
)

func init() {
	erc223AbiModel = new(abi.ABI)
	unmErr := erc223AbiModel.UnmarshalJSON([]byte(defaultAbi))
	if unmErr != nil {
		logger.Error("failed to load ERC223 interaction model internals")
		return
	}
}
