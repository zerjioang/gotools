// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/lib/encoding/hex"
)

const (
	exampleCode = `0x60806040526004361061005c576000357c0100000000000000000000000000000000000000000000000000000000900480630121b93f146100615780632d35a8a21461009c5780633477ee2e146100c7578063a3ec138d14610189575b600080fd5b34801561006d57600080fd5b5061009a6004803603602081101561008457600080fd5b81019080803590602001909291905050506101f2565b005b3480156100a857600080fd5b506100b161040c565b6040518082815260200191505060405180910390f35b3480156100d357600080fd5b50610100600480360360208110156100ea57600080fd5b8101908080359060200190929190505050610412565b6040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561014c578082015181840152602081019050610131565b50505050905090810190601f1680156101795780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561019557600080fd5b506101d8600480360360208110156101ac57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104d4565b604051808215151515815260200191505060405180910390f35b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515156102b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4572726f21204573746520656c6569746f72206ac3a120766f746f750000000081525060200191505060405180910390fd5b6000811180156102c557506002548111155b151561035f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f4572726f212045737465206ec3a36f20c3a920756d2063616e64696461746f2081526020017f76c3a16c69646f0000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600082815260200190815260200160002060020160008154809291906001019190505550807ffff3c900d938d21d0990d786e819f29b8d05c1ef587b462b939609625b684b1660405160405180910390a250565b60025481565b6001602052806000526040600020600091509050806000015490806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104c45780601f10610499576101008083540402835291602001916104c4565b820191906000526020600020905b8154815290600101906020018083116104a757829003601f168201915b5050505050908060020154905083565b60006020528060005260406000206000915054906101000a900460ff168156fea165627a7a723058206996e059ba9036b30a3acd5655270d1e14e578a3eb5fa18152b6bb7b98a9cbe80029`
)

func TestReverse(t *testing.T) {
	t.Run("hex-decode", func(t *testing.T) {
		data, err := hex.FromEthHex(exampleCode)
		assert.Nil(t, err)
		assert.NotNil(t, data)
		t.Log(data)
		t.Log(string(data))
	})
}
