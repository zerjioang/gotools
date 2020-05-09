// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package model

// Object - The transaction call object
type EthRequestParams struct {

	// DATA, 20 UnsafeBytes - (optional) The address the transaction is sent from.
	From string `json:"from,omitempty"`
	// DATA, 20 UnsafeBytes - The address the transaction is directed to.
	To string `json:"to,omitempty"`
	// QUANTITY - (optional) Integer of the gas provided for the transaction execution.
	// eth_call consumes zero gas, but this parameter may be needed by some executions.
	Gas string `json:"gas,omitempty"`
	// QUANTITY - (optional) Integer of the gasPrice used for each paid gas
	GasPrice string `json:"gasPrice,omitempty"`
	// QUANTITY - (optional) Integer of the value sent with this transaction
	Value string `json:"value,omitempty"`
	// DATA - (optional) Hash of the method signature and encoded parameters.
	Data string `json:"data,omitempty"`
	// DATA target block name or hexadecimal value
	Tag string
}

func (params EthRequestParams) String() string {
	return ""
}

type EthRequest struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  string `json:"params"`
}

func (request *EthRequest) SetBlockPeriod(period string) {
	// todo append block period
}
