// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package erc20

import (
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/eth/paramencoder"
	"github.com/zerjioang/gotools/lib/logger"
)

/*
7

To get token balance with
eth_call you need 'to' and 'data' parameter.
'to' is contract address, here we need to generate the 'data' parameter.
As the doc eth_call says, data: DATA - (optional) Hash of the method signature and encoded parameters.
address example: 0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0

We also need to know function specification. In the case of ERC20 the specs are:

contract ERC20 {
    function totalSupply() constant returns (uint supply);
    function balanceOf( address who ) constant returns (uint value);
    function allowance( address owner, address spender ) constant returns (uint _allowance);

    function transfer( address to, uint value) returns (bool ok);
    function transferFrom( address from, address to, uint value) returns (bool ok);
    function approve( address spender, uint value ) returns (bool ok);

    event transfer( address indexed from, address indexed to, uint value);
    event Approval( address indexed owner, address indexed spender, uint value);
}

// from each function, we need to hash its header as:
1. Web3.sha3("balanceOf(address)")
2: generate hex result

Encoded in hex: 0x70a08231b98ef4ca268c9cc3f6b4590e4bfec28280db06bb5d45e689f2a360be

after that we need to take first 4 bytes: 70a08231

Now ew need to encode our arguments:

Argument Encoding:

* address: equivalent to uint160, except for the assumed interpretation and language typing.
* int: enc(X) is the big-endian two's complement encoding of X,
padded on the higher-order (left) side with 0xff for negative X
and with zero bytes for positive X such that the length is a multiple of 32 bytes.

Padding the 20 bytes token address to 32 bytes with 0 to token holder address:

0000000000000000000000000b88516a6d22bf8e0d3657effbd41577c5fd4cb7
Then concat the function selector and encoded parameter, we get data parameter:

0x70a082310000000000000000000000000b88516a6d22bf8e0d3657effbd41577c5fd4cb7

curl -X POST --data '{"jsonrpc":"2.0","method":"eth_call","params":[{"to": "0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0", "data":"0x70a082310000000000000000000000000b88516a6d22bf8e0d3657effbd41577c5fd4cb7"}, "latest"],"id":67}' -H "Content-Type: application/json" http://127.0.0.1:8545/
*/

const (
	erc20Abi = `[
   {
      "constant":true,
      "inputs":[

      ],
      "name":"name",
      "outputs":[
         {
            "name":"",
            "type":"string"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":false,
      "inputs":[
         {
            "name":"_spender",
            "type":"address"
         },
         {
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"approve",
      "outputs":[
         {
            "name":"success",
            "type":"bool"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":true,
      "inputs":[

      ],
      "name":"totalSupply",
      "outputs":[
         {
            "name":"",
            "type":"uint256"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":false,
      "inputs":[
         {
            "name":"_from",
            "type":"address"
         },
         {
            "name":"_to",
            "type":"address"
         },
         {
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"transferFrom",
      "outputs":[
         {
            "name":"success",
            "type":"bool"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":true,
      "inputs":[

      ],
      "name":"decimals",
      "outputs":[
         {
            "name":"",
            "type":"uint256"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":true,
      "inputs":[
         {
            "name":"_owner",
            "type":"address"
         }
      ],
      "name":"balanceOf",
      "outputs":[
         {
            "name":"balance",
            "type":"uint256"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":true,
      "inputs":[

      ],
      "name":"symbol",
      "outputs":[
         {
            "name":"",
            "type":"string"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":false,
      "inputs":[
         {
            "name":"_to",
            "type":"address"
         },
         {
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"showMeTheMoney",
      "outputs":[

      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":false,
      "inputs":[
         {
            "name":"_to",
            "type":"address"
         },
         {
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"transfer",
      "outputs":[
         {
            "name":"success",
            "type":"bool"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "constant":true,
      "inputs":[
         {
            "name":"_owner",
            "type":"address"
         },
         {
            "name":"_spender",
            "type":"address"
         }
      ],
      "name":"allowance",
      "outputs":[
         {
            "name":"remaining",
            "type":"uint256"
         }
      ],
      "payable":false,
      "type":"function"
   },
   {
      "anonymous":false,
      "inputs":[
         {
            "indexed":true,
            "name":"_from",
            "type":"address"
         },
         {
            "indexed":true,
            "name":"_to",
            "type":"address"
         },
         {
            "indexed":false,
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"transfer",
      "type":"event"
   },
   {
      "anonymous":false,
      "inputs":[
         {
            "indexed":true,
            "name":"_owner",
            "type":"address"
         },
         {
            "indexed":true,
            "name":"_spender",
            "type":"address"
         },
         {
            "indexed":false,
            "name":"_value",
            "type":"uint256"
         }
      ],
      "name":"Approval",
      "type":"event"
   }
]`
)

var (
	// read only variables for ERC20
	erc20AbiModel     *abi.ABI
	NameParams        string
	DecimalsParams    string
	TotalSupplyParams string
	SymbolParams      string
)

/*
You can use some tool to calculate the function identifier which
are the first 4 bytes of the keccak hash of the function identifier as shown below.
These identifiers are for ERC20:

18160ddd -> totalSupply()
70a08231 -> balanceOf(address)
dd62ed3e -> allowance(address,address)
a9059cbb -> transfer(address,uint256)
095ea7b3 -> approve(address,uint256)
23b872dd -> transferFrom(address,address,uint256)

When calling json-rpc, params needs to be encoded in 32 bytes long.
so the full data field should be identifier + "000000000000000000000000" + payload

for example: in the case of total supply (0 params needed), would be:
0x18160ddd0000000000000000000000000000000000000000000000000000000000000000

# The method selector (4 bytes)
0xee919d5
# The 1st argument (32 bytes)
00000000000000000000000000000000000000000000000000000000000000001

*/
func init() {
	erc20AbiModel = new(abi.ABI)
	unmErr := erc20AbiModel.UnmarshalJSON([]byte(erc20Abi))
	if unmErr != nil {
		logger.Error("failed to load ERC20 interaction model internals")
		return
	}
	// preload erc20 functions abi encoded data field
	TotalSupplyParams, _ = paramencoder.GeneratePayload(erc20AbiModel, "totalSupply")
	DecimalsParams, _ = paramencoder.GeneratePayload(erc20AbiModel, "decimals")
	NameParams, _ = paramencoder.GeneratePayload(erc20AbiModel, "name")
	SymbolParams, _ = paramencoder.GeneratePayload(erc20AbiModel, "symbol")
}

func LoadErc20Abi() *abi.ABI {
	return erc20AbiModel
}
