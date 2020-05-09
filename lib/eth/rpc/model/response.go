// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package model

import (
	"encoding/json"
	"errors"
	"strconv"
)

type EthResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	//Data string    `json:"result"`
	Error *EthError `json:"error"`
}

func (response EthResponse) Errored() error {
	return errors.New(response.Error.Message + ". Error code: " + strconv.Itoa(response.Error.Code))
}
