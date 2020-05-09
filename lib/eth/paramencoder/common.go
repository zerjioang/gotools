// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package paramencoder

import (
	"github.com/zerjioang/gotools/lib/encoding/hex"
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/eth/fixtures/common"
	"github.com/zerjioang/gotools/lib/logger"
)

func GeneratePayload(aibmodel *abi.ABI, functionName string) (string, error) {
	//preload symbol function params
	temp, err := aibmodel.Pack(functionName)
	if err != nil {
		logger.Error("failed to load ERC20 '", functionName, "' interaction model")
		return "", err
	} else {
		// add 32 byte padding to set that this function has no parameters
		data := hex.ToEthHex(common.RightPadBytes(temp, 32+len(temp)))
		return data, err
	}
}
