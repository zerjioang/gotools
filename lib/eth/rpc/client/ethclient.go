package client

import (
	"time"

	"github.com/valyala/fasthttp"
)

type EthClient struct {
	*fasthttp.Client
}

func NewEthClient() *EthClient {
	c := EthClient{}
	c.Client = &fasthttp.Client{
		ReadTimeout:     time.Second * 3,
		WriteTimeout:    time.Second * 3,
		WriteBufferSize: 2048,
		ReadBufferSize:  2048,
	}
	return &c
}

func (cli *EthClient) HttpClient() *fasthttp.Client {
	return cli.Client
}
