// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/zerjioang/gotools/lib/eth/rpc/client"

	"github.com/valyala/fasthttp"
	"github.com/zerjioang/gotools/lib/logger"

	"github.com/zerjioang/gotools/util/str"
)

const (
	postMethod = "POST"
)

var (
	ApplicationJSON = "application/json"
	fallbackClient  = client.NewEthClient()
	br              = bytes.NewReader(nil)
)

func MakePost(client *client.EthClient, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	return MakeCall(client, postMethod, url, headers, content)
}

func MakeCall(client *client.EthClient, method string, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	br.Reset(content)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)                    //set URL
	req.Header.SetMethodBytes([]byte(method)) //set method mode
	req.SetBody(content)                      //set body

	res := fasthttp.AcquireResponse()

	//prepare the client to be used for requests
	var err error
	if client == nil {
		err = fallbackClient.Do(req, res)
	} else {
		err = client.Do(req, res)
	}
	if err != nil {
		return nil, err
	}
	responseData := res.Body()
	logger.Debug("response received: ", str.UnsafeString(responseData))
	return responseData, nil
}
