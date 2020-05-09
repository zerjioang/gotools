// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package erc721

import (
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/logger"
)

//What is ERC-721?
//
//ERC-721 is a free, open standard that describes how to build
//non-fungible or unique tokens on the Ethereum blockchain.
//While most tokens are fungible (every token is the same as
//every other token), ERC-721 tokens are all unique.
//
//Think of them like rare, one-of-a-kind collectables.

//pragma solidity ^0.4.20;
//
///// @title ERC-721 Non-Fungible Token Standard
///// @dev See https://github.com/ethereum/EIPs/blob/master/EIPS/eip-721.md
/////  Note: the ERC-165 identifier for this interface is 0x80ac58cd
//interface ERC721 /* is ERC165 */ {
///// @dev This emits when ownership of any NFT changes by any mechanism.
/////  This event emits when NFTs are created (`from` == 0) and destroyed
/////  (`to` == 0). Exception: during contract creation, any number of NFTs
/////  may be created and assigned without emitting transfer. At the time of
/////  any transfer, the approved address for that NFT (if any) is reset to none.
//event transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
//
///// @dev This emits when the approved address for an NFT is changed or
/////  reaffirmed. The zero address indicates there is no approved address.
/////  When a transfer event emits, this also indicates that the approved
/////  address for that NFT (if any) is reset to none.
//event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId);
//
///// @dev This emits when an operator is enabled or disabled for an owner.
/////  The operator can manage all NFTs of the owner.
//event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved);
//
///// @notice Count all NFTs assigned to an owner
///// @dev NFTs assigned to the zero address are considered invalid, and this
/////  function throws for queries about the zero address.
///// @param _owner An address for whom to query the balance
///// @return The number of NFTs owned by `_owner`, possibly zero
//function balanceOf(address _owner) external view returns (uint256);
//
///// @notice Find the owner of an NFT
///// @dev NFTs assigned to zero address are considered invalid, and queries
/////  about them do throw.
///// @param _tokenId The identifier for an NFT
///// @return The address of the owner of the NFT
//function ownerOf(uint256 _tokenId) external view returns (address);
//
///// @notice Transfers the ownership of an NFT from one address to another address
///// @dev Throws unless `msg.sender` is the current owner, an authorized
/////  operator, or the approved address for this NFT. Throws if `_from` is
/////  not the current owner. Throws if `_to` is the zero address. Throws if
/////  `_tokenId` is not a valid NFT. When transfer is complete, this function
/////  checks if `_to` is a smart contract (code size > 0). If so, it calls
/////  `onERC721Received` on `_to` and throws if the return value is not
/////  `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`.
///// @param _from The current owner of the NFT
///// @param _to The new owner
///// @param _tokenId The NFT to transfer
///// @param data Additional data with no specified format, sent in call to `_to`
//function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data) external payable;
//
///// @notice Transfers the ownership of an NFT from one address to another address
///// @dev This works identically to the other function with an extra data parameter,
/////  except this function just sets data to ""
///// @param _from The current owner of the NFT
///// @param _to The new owner
///// @param _tokenId The NFT to transfer
//function safeTransferFrom(address _from, address _to, uint256 _tokenId) external payable;
//
///// @notice transfer ownership of an NFT -- THE CALLER IS RESPONSIBLE
/////  TO CONFIRM THAT `_to` IS CAPABLE OF RECEIVING NFTS OR ELSE
/////  THEY MAY BE PERMANENTLY LOST
///// @dev Throws unless `msg.sender` is the current owner, an authorized
/////  operator, or the approved address for this NFT. Throws if `_from` is
/////  not the current owner. Throws if `_to` is the zero address. Throws if
/////  `_tokenId` is not a valid NFT.
///// @param _from The current owner of the NFT
///// @param _to The new owner
///// @param _tokenId The NFT to transfer
//function transferFrom(address _from, address _to, uint256 _tokenId) external payable;
//
///// @notice Set or reaffirm the approved address for an NFT
///// @dev The zero address indicates there is no approved address.
///// @dev Throws unless `msg.sender` is the current NFT owner, or an authorized
/////  operator of the current owner.
///// @param _approved The new approved NFT controller
///// @param _tokenId The NFT to approve
//function approve(address _approved, uint256 _tokenId) external payable;
//
///// @notice Enable or disable approval for a third party ("operator") to manage
/////  all of `msg.sender`'s assets.
///// @dev Emits the ApprovalForAll event. The contract MUST allow
/////  multiple operators per owner.
///// @param _operator Address to add to the set of authorized operators.
///// @param _approved True if the operator is approved, false to revoke approval
//function setApprovalForAll(address _operator, bool _approved) external;
//
///// @notice Get the approved address for a single NFT
///// @dev Throws if `_tokenId` is not a valid NFT
///// @param _tokenId The NFT to find the approved address for
///// @return The approved address for this NFT, or the zero address if there is none
//function getApproved(uint256 _tokenId) external view returns (address);
//
///// @notice Query if an address is an authorized operator for another address
///// @param _owner The address that owns the NFTs
///// @param _operator The address that acts on behalf of the owner
///// @return True if `_operator` is an approved operator for `_owner`, false otherwise
//function isApprovedForAll(address _owner, address _operator) external view returns (bool);
//}
//
//interface ERC165 {
///// @notice Query if a contract implements an interface
///// @param interfaceID The interface identifier, as specified in ERC-165
///// @dev Interface identification is specified in ERC-165. This function
/////  uses less than 30,000 gas.
///// @return `true` if the contract implements `interfaceID` and
/////  `interfaceID` is not 0xffffffff, `false` otherwise
//function supportsInterface(bytes4 interfaceID) external view returns (bool);
//}
//
//interface ERC721TokenReceiver {
///// @notice Handle the receipt of an NFT
///// @dev The ERC721 smart contract calls this function on the
///// recipient after a `transfer`. This function MAY throw to revert and reject the transfer. Return
///// of other than the magic value MUST result in the transaction being reverted.
///// @notice The contract address is always the message sender.
///// @param _operator The address which called `safeTransferFrom` function
///// @param _from The address which previously owned the token
///// @param _tokenId The NFT identifier which is being transferred
///// @param _data Additional data with no specified format
///// @return `bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`
///// unless throwing
//function onERC721Received(address _operator, address _from, uint256 _tokenId, bytes _data) external returns(bytes4);
//}

// more info at http://erc721.org/

/*
{
    "095ea7b3": "approve(address,uint256)",
    "70a08231": "balanceOf(address)",
    "081812fc": "getApproved(uint256)",
    "e985e9c5": "isApprovedForAll(address,address)",
    "150b7a02": "onERC721Received(address,address,uint256,bytes)",
    "6352211e": "ownerOf(uint256)",
    "42842e0e": "safeTransferFrom(address,address,uint256)",
    "b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
    "a22cb465": "setApprovalForAll(address,bool)",
    "23b872dd": "transferFrom(address,address,uint256)"
}

*/

const (
	defaultErc721Abi = `[
	{
		"constant": true,
		"inputs": [
			{
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "getApproved",
		"outputs": [
			{
				"name": "",
				"type": "address"
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
				"name": "_approved",
				"type": "address"
			},
			{
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [],
		"payable": true,
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_operator",
				"type": "address"
			},
			{
				"name": "_from",
				"type": "address"
			},
			{
				"name": "_tokenId",
				"type": "uint256"
			},
			{
				"name": "_data",
				"type": "bytes"
			}
		],
		"name": "onERC721Received",
		"outputs": [
			{
				"name": "",
				"type": "bytes4"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_from",
				"type": "address"
			},
			{
				"name": "_to",
				"type": "address"
			},
			{
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [],
		"payable": true,
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_from",
				"type": "address"
			},
			{
				"name": "_to",
				"type": "address"
			},
			{
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"payable": true,
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "ownerOf",
		"outputs": [
			{
				"name": "",
				"type": "address"
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
				"name": "_owner",
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
				"name": "_operator",
				"type": "address"
			},
			{
				"name": "_approved",
				"type": "bool"
			}
		],
		"name": "setApprovalForAll",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_from",
				"type": "address"
			},
			{
				"name": "_to",
				"type": "address"
			},
			{
				"name": "_tokenId",
				"type": "uint256"
			},
			{
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "safeTransferFrom",
		"outputs": [],
		"payable": true,
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "_owner",
				"type": "address"
			},
			{
				"name": "_operator",
				"type": "address"
			}
		],
		"name": "isApprovedForAll",
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
				"name": "_from",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "_to",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "_owner",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "_approved",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "_tokenId",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"name": "_owner",
				"type": "address"
			},
			{
				"indexed": true,
				"name": "_operator",
				"type": "address"
			},
			{
				"indexed": false,
				"name": "_approved",
				"type": "bool"
			}
		],
		"name": "ApprovalForAll",
		"type": "event"
	}
]`
)

const (
	ApproveParams                   string = "095ea7b3" // approve(address,uint256)
	BalanceOfParams                 string = "70a08231" // balanceOf(address)
	GetApprovedParams               string = "081812fc" // getApproved(uint256)
	IsApprovedForAllParams          string = "e985e9c5" // isApprovedForAll(address,address)
	OnERC721ReceivedParams          string = "150b7a02" // onERC721Received(address,address,uint256,bytes)
	OwnerOfParams                   string = "6352211e" // ownerOf(uint256)
	SafeTransferFromParams          string = "42842e0e" // safeTransferFrom(address,address,uint256)
	SafeTransferFromParamsWithBytes string = "b88d4fde" // safeTransferFrom(address,address,uint256,bytes)
	SetApprovalForAllParams         string = "a22cb465" // setApprovalForAll(address,bool)
	TransferFromParams              string = "23b872dd" // transferFrom(address,address,uint256)
)

var (
	// read only variables for ERC721
	erc721AbiModel *abi.ABI
)

func init() {
	erc721AbiModel = new(abi.ABI)
	unmErr := erc721AbiModel.UnmarshalJSON([]byte(defaultErc721Abi))
	if unmErr != nil {
		logger.Error("failed to load ERC721 interaction model internals")
		return
	}
}

func LoadErc721Abi() *abi.ABI {
	return erc721AbiModel
}
