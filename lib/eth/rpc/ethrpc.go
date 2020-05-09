// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
	"syscall"

	"github.com/zerjioang/gotools/lib/cache"
	"github.com/zerjioang/gotools/lib/encoding/hex"
	"github.com/zerjioang/gotools/lib/eth/fixtures"
	"github.com/zerjioang/gotools/lib/eth/fixtures/abi"
	"github.com/zerjioang/gotools/lib/eth/paramencoder/erc20"
	"github.com/zerjioang/gotools/lib/eth/rpc/client"
	"github.com/zerjioang/gotools/lib/eth/rpc/model"
	"github.com/zerjioang/gotools/lib/httpclient"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/lib/stack"
	"github.com/zerjioang/gotools/lib/worker"
	"github.com/zerjioang/gotools/util/str"
)

// https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J
// https://github.com/ethereum/go-ethereum/blob/master/ethclient/ethclient.go
type contractFunction func(string) (string, error)
type ParamsCallback func() string
type Web3Resolver func() (interface{}, error)

var (
	instance         = new(EthRPC)
	summaryFunctions = []contractFunction{
		instance.Erc20Name,
		instance.Erc20Symbol,
		instance.Erc20Decimals,
		instance.Erc20TotalSupply,
	}
	summaryFunctionsNames = []string{
		"name", "symbol", "decimals", "totalsupply",
	}
	defaultRpcHeader = http.Header{
		"Content-Type": []string{httpclient.ApplicationJSON},
	}
	defaultGraphQLHeader = http.Header{
		"Content-Type": []string{httpclient.ApplicationJSON},
	}
)

var (
	oneEth        = big.NewInt(1000000000000000000)
	oneEthInt64   = oneEth.Int64()
	defaultBigInt = new(big.Int)
)

var (
	netVersion = []byte("net_version")
)

type ConnectionMode uint8

const (
	HttpMode ConnectionMode = iota
	UnixMode
)

// EthRPC - Ethereum rpc client
type EthRPC struct {
	//connection mode: http or unix
	mode ConnectionMode
	//ethereum or quorum node endpoint
	url string
	//ethereum interaction cache
	cache *cache.MemoryCache
	// http client
	client *client.EthClient
	// debug flag
	Debug           bool
	connectionCache model.ConnectionCache
}

// New create new rpc client with given url
func NewDefaultRPCPtr(url string, debug bool, client *client.EthClient) *EthRPC {
	c := NewDefaultRPC(url, debug, client)
	return &c
}

// New create new rpc client with given url
func NewDefaultRPC(url string, debug bool, client *client.EthClient) EthRPC {
	rpc := EthRPC{
		url:    url,
		cache:  cache.NewMemoryCache(),
		client: client,
		Debug:  debug,
	}
	if strings.LastIndex(url, ".ipc") == 0 {
		//unix mode detected
		rpc.mode = UnixMode
	}
	return rpc
}

func (rpc *EthRPC) post(method string, target interface{}, params ParamsCallback) error {
	paramsStr := ""
	if params != nil {
		paramsStr = params()
	}
	result, err := rpc.makePostWithMethodParams(method, paramsStr)
	if err != nil {
		return err
	}

	if target == nil || result == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// URL returns client url
func (rpc *EthRPC) URL() string {
	return rpc.url
}

// makePostWithMethodParams returns raw response of method post
/*

eth_call

Executes a new message post immediately without creating a transaction on the block chain.
Parameters

    Object - The transaction post object

    from: DATA, 20 UnsafeBytes - (optional) The address the transaction is sent from.
    to: DATA, 20 UnsafeBytes - The address the transaction is directed to.
    gas: QUANTITY - (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
    gasPrice: QUANTITY - (optional) Integer of the gasPrice used for each paid gas
    value: QUANTITY - (optional) Integer of the value sent with this transaction
    data: DATA - (optional) Hash of the method signature and encoded parameters. For details see Ethereum Contract ABI

    QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter

Returns

DATA - the return value of executed contract.
*/
func (rpc *EthRPC) makePostWithMethodParams(method string, params string) (json.RawMessage, error) {
	if params == "" {
		request := `{"id": 1,"jsonrpc": "2.0","method": "` + method + `"}`
		return rpc.makePostString(request)
	} else {
		request := `{"id": 1,"jsonrpc": "2.0","method": "` + method + `","params":` + params + `}`
		return rpc.makePostString(request)
	}
}

func (rpc *EthRPC) makePostString(data string) (json.RawMessage, error) {
	return rpc.makePostBytes(str.UnsafeBytes(data))
}

func (rpc *EthRPC) makePostBytes(data []byte) (json.RawMessage, error) {

	responseData, postErr := httpclient.MakePost(rpc.client, rpc.url, defaultRpcHeader, data)
	if postErr != nil {
		return nil, postErr
	}

	resp := model.EthResponse{}
	unmErr := json.Unmarshal(responseData, &resp)
	if unmErr != nil {
		return nil, unmErr
	}

	if resp.Error != nil {
		return nil, resp.Errored()
	}

	return resp.Result, nil
}

func (rpc *EthRPC) proxyRequest(data []byte) (json.RawMessage, error) {
	return httpclient.MakePost(rpc.client, rpc.url, defaultRpcHeader, data)
}

func unixSocketReader(r io.Reader, notifier chan worker.GoroutineResponse) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			notifier <- worker.GoroutineResponse{Err: err, Data: nil}
		}
		notifier <- worker.GoroutineResponse{Err: nil, Data: buf[0:n]}
	}
}

// out := make(chan GoroutineResponse, 1)
func (rpc *EthRPC) makePostUnix(data string, notifier chan worker.GoroutineResponse) {
	unixPath := "/tmp/echo.sock"
	linKErr := syscall.Unlink(unixPath)
	if linKErr != nil {
		logger.Error(linKErr)
	}
	c, err := net.Dial("unix", unixPath)
	if err != nil {
		notifier <- worker.GoroutineResponse{Err: err, Data: nil}
	} else {
		go unixSocketReader(c, notifier)

		//send the message
		_, err = c.Write(str.UnsafeBytes(data))
		if err != nil {
			notifier <- worker.GoroutineResponse{Err: err, Data: nil}
		}
		//close connection
		err = c.Close()
		notifier <- worker.GoroutineResponse{Err: err, Data: nil}
	}
}

// RawCall returns raw response of method post (Deprecated)
func (rpc *EthRPC) RawCall(method string, params string) (json.RawMessage, error) {
	return rpc.makePostWithMethodParams(method, params)
}

func (rpc *EthRPC) EthMethodNoParams(methodName string) (interface{}, error) {
	var response interface{}
	err := rpc.post(methodName, &response, nil)
	return response, err
}

// returns ethereum node information
func (rpc *EthRPC) EthNodeInfo() (string, error) {
	var response string

	err := rpc.post("eth_info", &response, nil)
	return response, err
}

// Web3ClientVersion returns the current client version.
func (rpc *EthRPC) Web3ClientVersion() (string, error) {
	var clientVersion string

	err := rpc.post("web3_clientVersion", &clientVersion, nil)
	return clientVersion, err
}

func (rpc *EthRPC) IsGanache() (bool, error) {
	data, err := rpc.Web3ClientVersion()
	if err != nil {
		return false, err
	} else {
		// check if response data is similar to ganache response
		// by default ganache node simulator starts with: EthereumJS TestRPC
		isGanache := strings.Contains(data, "ethereum-js") || strings.Contains(data, "TestRPC")
		return isGanache, nil
	}
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (rpc *EthRPC) Web3Sha3(data []byte) (string, error) {
	var hash string
	//prepare the params of the sha3 function
	params := func() string {
		return "[" + rpc.doubleQuote(fixtures.Encode(data)) + "]"
	}
	err := rpc.post("web3_sha3", &hash, params)
	return hash, err
}

// web3 request resolver that includes cache support for all request given a known key
func (rpc *EthRPC) resolve(key []byte, resolver Web3Resolver) (interface{}, error) {
	cachedValue, found := rpc.cache.Get(key)
	if found && cachedValue != nil {
		//cache value found, so return it
		return cachedValue, nil
	} else {
		// try to resolve the value
		result, err := resolver()
		if err != nil {
			//error occurred while executing
			return nil, err
		} else {
			//vale the result value in the cache
			rpc.cache.Set(key, result)
			return result, nil
		}
	}
}

// NetVersion returns the current network protocol version.
func (rpc *EthRPC) NetVersion() (string, error) {
	response, err := rpc.resolve(netVersion, func() (interface{}, error) {
		var version string
		err := rpc.post("net_version", &version, nil)
		return version, err
	})
	result, _ := response.(string)
	return result, err
}

// NetListening returns true if client is actively listening for network connections.
func (rpc *EthRPC) NetListening() (bool, error) {
	var listening bool

	err := rpc.post("net_listening", &listening, nil)
	return listening, err
}

// NetPeerCount returns number of peers currently connected to the client.
func (rpc *EthRPC) NetPeerCount() (int, error) {
	var response int
	if err := rpc.post("net_peerCount", &response, nil); err != nil {
		return 0, err
	}

	return response, nil
}

// EthProtocolVersion returns the current ethereum protocol version.
func (rpc *EthRPC) EthProtocolVersion() (string, error) {
	var protocolVersion string

	err := rpc.post("eth_protocolVersion", &protocolVersion, nil)
	return protocolVersion, err
}

// EthSyncing returns an object with data about the sync status or false.
func (rpc *EthRPC) EthSyncing() (*Syncing, error) {
	result, err := rpc.makePostWithMethodParams("eth_syncing", "")
	if err != nil {
		return nil, err
	}
	syncing := new(Syncing)
	if bytes.Equal(result, []byte("false")) {
		//no syncing, return false
		return syncing, nil
	}
	err = json.Unmarshal(result, syncing)
	return syncing, err
}

// returns ethereum node information
func (rpc *EthRPC) EthNetVersion() (string, error) {
	var response string

	err := rpc.post("net_version", &response, nil)
	return response, err
}

// EthCoinbase returns the client coinbase address
func (rpc *EthRPC) EthCoinbase() (string, error) {
	var address string

	err := rpc.post("eth_coinbase", &address, nil)
	return address, err
}

// EthMining returns true if client is actively mining new blocks.
func (rpc *EthRPC) EthMining() (bool, error) {
	var mining bool

	err := rpc.post("eth_mining", &mining, nil)
	return mining, err
}

// EthHashRate returns the number of hashes per second that the node is mining with.
func (rpc *EthRPC) EthHashRate() (int64, error) {
	var response string
	err := rpc.post("eth_hashrate", &response, nil)
	if err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGasPrice returns the current price per gas in wei.
func (rpc *EthRPC) EthGasPrice() (int64, error) {
	var response string
	err := rpc.post("eth_gasPrice", &response, nil)
	if err != nil {
		return 0, err
	}
	// example 0x4a817c800
	// fast check if string starts with 0x
	if len(response) > 2 && response[1] == 'x' && response[0] == '0' {
		response = response[2:]
	}
	return ParseHexToInt(response)
}

// EthAccounts returns a list of addresses owned by client.
func (rpc *EthRPC) EthAccounts() ([]string, error) {
	var accounts []string
	err := rpc.post("eth_accounts", &accounts, nil)
	return accounts, err
}

// EthBlockNumber returns the number of most recent block.
func (rpc *EthRPC) EthBlockNumber() (int64, error) {
	var response string
	if err := rpc.post("eth_blockNumber", &response, nil); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBalance returns the balance of the account of given address in wei.
func (rpc *EthRPC) EthGetBalance(address string, block string) (*big.Int, string, error) {
	var response string
	//prepare the params of the get balance function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	if err := rpc.post("eth_getBalance", &response, params); err != nil {
		return defaultBigInt, response, err
	}
	return ParseBigInt(response)
}

// EthGetStorageAt returns the value from a storage position at a given address.
func (rpc *EthRPC) EthGetStorageAt(data string, key string, block string) (string, error) {
	var result string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(data) + "," + rpc.doubleQuote(key) + "," + rpc.doubleQuote(block) + "]"
	}
	err := rpc.post("eth_getStorageAt", &result, params)
	return result, err
}

// EthGetTransactionCount returns the number of transactions sent from an address.
func (rpc *EthRPC) EthGetTransactionCount(address string, block string) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	if err := rpc.post("eth_getTransactionCount", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetBlockTransactionCountByHash(hash string) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	if err := rpc.post("eth_getBlockTransactionCountByHash", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
func (rpc *EthRPC) EthGetBlockTransactionCountByNumber(blockNumber string) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(blockNumber) + "]"
	}
	if err := rpc.post("eth_getBlockTransactionCountByNumber", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetUncleCountByBlockHash(hash string) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	if err := rpc.post("eth_getUncleCountByBlockHash", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (rpc *EthRPC) EthGetUncleCountByBlockNumber(number int) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(IntToHex(number)) + "]"
	}
	if err := rpc.post("eth_getUncleCountByBlockNumber", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetCode returns code at a given address.
func (rpc *EthRPC) EthGetCode(address string, block string) (string, error) {
	var code string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	err := rpc.post("eth_getCode", &code, params)
	return code, err
}

// EthSign signs data with a given address.
// Calculates an Ethereum specific signature
// with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func (rpc *EthRPC) EthSign(address, data string) (string, error) {
	var signature string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(data) + "]"
	}
	err := rpc.post("eth_sign", &signature, params)
	return signature, err
}

// EthSendTransaction creates new message post transaction
// or a contract creation, if the data field contains code.
func (rpc *EthRPC) EthSendTransaction(transaction TransactionData) (string, error) {
	return rpc.EthSendTransactionPtr(&transaction)
}

// EthSendTransaction creates new message post transaction
// or a contract creation, if the data field contains code.
func (rpc *EthRPC) EthSendTransactionPtr(transaction *TransactionData) (string, error) {
	var hash string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw)
	}
	err := rpc.post("eth_sendTransaction", &hash, params)
	return hash, err
}

// EthSendRawTransaction creates new message post transaction
// or a contract creation for signed transactions.
func (rpc *EthRPC) EthSendRawTransaction(data string) (string, error) {
	var hash string
	//prepare the params of the function
	params := func() string {
		return data
	}
	err := rpc.post("eth_sendRawTransaction", &hash, params)
	return hash, err
}

// EthCall executes a new message post immediately without
// creating a transaction on the block chain.
func (rpc *EthRPC) EthCall(transaction TransactionData, tag string) (string, error) {
	var data string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw) + "," + tag
	}
	err := rpc.post("eth_call", &data, params)
	return data, err
}

// EthEstimateGas makes a post or transaction, which won't be
// added to the blockchain and returns the used gas, which can
// be used for estimating the used gas.
func (rpc *EthRPC) EthEstimateGas(transaction TransactionData) (int64, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw)
	}
	err := rpc.post("eth_estimateGas", &response, params)
	if err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// getBlock gets current block information
func (rpc *EthRPC) getBlock(method string, withTransactions bool, params string) (*Block, error) {
	result, err := rpc.makePostWithMethodParams(method, params)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var response proxyBlock
	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}

	block := response.toBlock()
	return &block, nil
}

// EthGetBlockByHash returns information about a block by hash.
func (rpc *EthRPC) EthGetBlockByHash(hash string, withTransactions bool) (*Block, error) {
	params := hash
	return rpc.getBlock("eth_getBlockByHash", withTransactions, params)
}

// EthGetBlockByNumber returns information about a block by block number.
func (rpc *EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	params := IntToHex(number)
	return rpc.getBlock("eth_getBlockByNumber", withTransactions, params)
}

func (rpc *EthRPC) getTransaction(method string, params ParamsCallback) (*Transaction, error) {
	transaction := new(Transaction)
	err := rpc.post(method, transaction, params)
	return transaction, err
}

// EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
func (rpc *EthRPC) EthGetTransactionByHash(hash string) (*Transaction, error) {
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	return rpc.getTransaction("eth_getTransactionByHash", params)
}

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	params := func() string {
		return blockHash + "," + IntToHex(transactionIndex)
	}
	return rpc.getTransaction("eth_getTransactionByBlockHashAndIndex", params)
}

// EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	params := func() string {
		return IntToHex(blockNumber) + "," + IntToHex(transactionIndex)
	}
	return rpc.getTransaction("eth_getTransactionByBlockNumberAndIndex", params)
}

// EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note That the receipt is not available for pending transactions.
func (rpc *EthRPC) EthGetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return rpc.doubleQuote(hash)
	}
	err := rpc.post("eth_getTransactionReceipt", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// TODO implement
// EthGetPendingTransactions returns the list of pending transactions
func (rpc *EthRPC) EthGetPendingTransactions(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return hash
	}
	err := rpc.post("eth_pendingTransactions", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// eth_getUncleByBlockHashAndIndex
// Returns information about a uncle of a block by hash and uncle index position.
// params: [
//   '0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b',
//   '0x0' // 0
// ]
// Note: An uncle doesn't contain individual transactions.
func (rpc *EthRPC) EthGetUncleByBlockHashAndIndex(hash string, uncleIndex int) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "," + rpc.doubleQuote(IntToHex(uncleIndex)) + "]"
	}
	err := rpc.post("eth_getUncleByBlockHashAndIndex", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// eth_getUncleByBlockNumberAndIndex
// Returns information about a uncle of a block by number and uncle index position.
// params: [
//   '0x29c', // 668
//   '0x0' // 0
// ]
// Note: An uncle doesn't contain individual transactions.
func (rpc *EthRPC) EthGetUncleByBlockNumberAndIndex(blockNumber string, uncleIndex int) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return "[" + rpc.doubleQuote(blockNumber) + "," + rpc.doubleQuote(IntToHex(uncleIndex)) + "]"
	}
	err := rpc.post("eth_getUncleByBlockNumberAndIndex", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// EthGetCompilers returns a list of available compilers in the client.
// @deprecated
// returns Array - Array of available compilers.
func (rpc *EthRPC) EthGetCompilers() ([]string, error) {
	var compilers []string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_getCompilers", &compilers, params)
	return compilers, err
}

// eth_compileSolidity
// @deprecated
// Returns compiled solidity code.
//   DATA - The compiled source code.
func (rpc *EthRPC) EthCompileSolidity(code string) ([]string, error) {
	var compilers []string
	params := func() string {
		return rpc.doubleQuote(code)
	}
	err := rpc.post("eth_compileSolidity", &compilers, params)
	return compilers, err
}

// eth_compileLLL
// @deprecated
// Returns compiled LLL code.
//   DATA - The compiled source code.
func (rpc *EthRPC) EthCompileLLL(code string) ([]string, error) {
	var compilers []string
	params := func() string {
		return rpc.doubleQuote(code)
	}
	err := rpc.post("eth_compileLLL", &compilers, params)
	return compilers, err
}

// eth_compileSerpent
// @deprecated
// Returns compiled Serpent code.
//   DATA - The compiled source code.
func (rpc *EthRPC) EthCompileSerpent(code string) ([]string, error) {
	var compilers []string
	params := func() string {
		return rpc.doubleQuote(code)
	}
	err := rpc.post("eth_compileSerpent", &compilers, params)
	return compilers, err
}

// EthNewFilter creates a new filter object.
func (rpc *EthRPC) EthNewFilter(filter FilterParams) (string, error) {
	var filterID string
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_newFilter", &filterID, params)
	return filterID, err
}

// EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
// To check if the state has changed, post EthGetFilterChanges.
func (rpc *EthRPC) EthNewBlockFilter() (string, error) {
	var filterID string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_newBlockFilter", &filterID, params)
	return filterID, err
}

// EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, post EthGetFilterChanges.
func (rpc *EthRPC) EthNewPendingTransactionFilter() (string, error) {
	var filterID string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_newPendingTransactionFilter", &filterID, params)
	return filterID, err
}

// EthUninstallFilter uninstalls a filter with given id.
func (rpc *EthRPC) EthUninstallFilter(filterID string) (bool, error) {
	var res bool
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_uninstallFilter", &res, params)
	return res, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
func (rpc *EthRPC) EthGetFilterChanges(filterID string) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_getFilterChanges", &logs, params)
	return logs, err
}

// EthGetFilterLogs returns an array of all logs matching filter with given id.
func (rpc *EthRPC) EthGetFilterLogs(filterID string) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_getFilterLogs", &logs, params)
	return logs, err
}

// EthGetLogs returns an array of all logs matching a given filter object.
func (rpc *EthRPC) EthGetLogs(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_getLogs", &logs, params)
	return logs, err
}

// eth_getWork
// Returns the hash of the current block, the seedHash, and the boundary condition to be met ("target").
// EthGetWork returns an array of all logs matching a given filter object.
// Returns
// Array - Array with the following properties:
// DATA, 32 Bytes - current block header pow-hash
// DATA, 32 Bytes - the seed hash used for the DAG.
// DATA, 32 Bytes - the boundary condition ("target"), 2^256 / difficulty.
func (rpc *EthRPC) EthGetWork() ([]string, error) {
	var logs []string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_getWork", &logs, params)
	return logs, err
}

// eth_submitWork
// EthSubmitWork
// Used for submitting a proof-of-work solution.
// DATA, 8 Bytes - The nonce found (64 bits)
// DATA, 32 Bytes - The header's pow-hash (256 bits)
// DATA, 32 Bytes - The mix digest (256 bits)
// params: [
//  "0x0000000000000001",
//  "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
//  "0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"
// ]
// Boolean - returns true if the provided solution is valid, otherwise false.
func (rpc *EthRPC) EthSubmitWork(nonce, header, digest string) (bool, error) {
	var result bool
	params := func() string {
		return "[" + rpc.doubleQuote(nonce) + "," + rpc.doubleQuote(header) + "," + rpc.doubleQuote(digest) + "]"
	}
	err := rpc.post("eth_submitWork", &result, params)
	return result, err
}

// eth_submitHashrate
// Used for submitting mining hashrate.
// Parameters
//    Hashrate, a hexadecimal string representation (32 bytes) of the hash rate
//    ID, String - A random hexadecimal(32 bytes) ID identifying the client
// params: [
//  "0x0000000000000000000000000000000000000000000000000000000000500000",
//  "0x59daa26581d0acd1fce254fb7e85952f4c09d0915afd33d3886cd914bc7d283c"
// ]
// returns
// Boolean - returns true if submitting went through succesfully and false otherwise.
func (rpc *EthRPC) EthSubmitHashrate(hashrate string, clientid string) (bool, error) {
	var result bool
	params := func() string {
		return "[" + rpc.doubleQuote(hashrate) + "," + rpc.doubleQuote(clientid) + "]"
	}
	err := rpc.post("eth_submitHashrate", &result, params)
	return result, err
}

// eth_getProof
// Returns the account- and storage-values of the specified account including the Merkle-proof.
// getProof-Parameters
//
//    DATA, 20 bytes - address of the account or contract
//    ARRAY, 32 Bytes - array of storage-keys which should be proofed and included. See eth_getStorageAt
//    QUANTITY|TAG - integer block number, or the string "latest" or "earliest", see the default block parameter
//
//Example Parameters
//
//params: ["0x1234567890123456789012345678901234567890",["0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000000000000000000000000000000000000000000001"],"latest"]
//
//getProof-Returns

//Returns Object - A account object:
//balance: QUANTITY - the balance of the account. See eth_getBalance
//codeHash: DATA, 32 Bytes - hash of the code of the account. For a simple Account without code it will return "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
//nonce: QUANTITY, - nonce of the account. See eth_getTransactionCount
//storageHash: DATA, 32 Bytes - SHA3 of the StorageRoot. All storage will deliver a MerkleProof starting with this rootHash.
//accountProof: ARRAY - Array of rlp-serialized MerkleTree-Nodes, starting with the stateRoot-Node, following the path of the SHA3 (address) as key.
//storageProof: ARRAY - Array of storage-entries as requested. Each entry is a object with these properties:
//key: QUANTITY - the requested storage key value: QUANTITY - the storage value proof: ARRAY - Array of rlp-serialized MerkleTree-Nodes, starting with the storageHash-Node, following the path of the SHA3 (key) as path.
func (rpc *EthRPC) EthGetProof(address string, keys []string, blocktime string) (bool, error) {
	var response bool
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "]"
	}
	err := rpc.post("eth_getProof", &response, params)
	return response, err
}

// db_putString
// Stores a string in the local database.
//Note this function is deprecated and will be removed in the future.
// Parameters
//
//    String - Database name.
//    String - Key name.
//    String - String to store.
//
//Example Parameters
//
//params: [
//  "testDB",
//  "myKey",
//  "myString"
//]
//
//Returns
//
//Boolean - returns true if the value was stored, otherwise false.
func (rpc *EthRPC) DbPutString(dbname string, key string, value string) (bool, error) {
	var response bool
	params := func() string {
		return "[" + rpc.doubleQuote(dbname) + "," + rpc.doubleQuote(key) + "," + rpc.doubleQuote(value) + "]"
	}
	err := rpc.post("db_putString", &response, params)
	return response, err
}

// db_getString
// Returns string from the local database.
// Note this function is deprecated and will be removed in the future.
// Parameters
// String - Database name.
// String - Key name.
// Example Parameters
//
// params: [
//  "testDB",
//  "myKey",
// ]
// Returns
//
// String - The previously stored string.
func (rpc *EthRPC) DbGetString(dbname string, key string) (string, error) {
	var value string
	params := func() string {
		return "[" + rpc.doubleQuote(dbname) + "," + rpc.doubleQuote(key) + "]"
	}
	err := rpc.post("db_getString", &value, params)
	return value, err
}

// db_putHex
// Stores binary data in the local database.
//Note this function is deprecated and will be removed in the future.
// Parameters
//
//    String - Database name.
//    String - Key name.
//    DATA - The data to store.
//
//Example Parameters
//
//params: [
//  "testDB",
//  "myKey",
//  "0x68656c6c6f20776f726c64"
//]
//
//Returns
//
//Boolean - returns true if the value was stored, otherwise false.
func (rpc *EthRPC) DbPutHex(dbname string, key string, value string) (bool, error) {
	var response bool
	params := func() string {
		return "[" + rpc.doubleQuote(dbname) + "," + rpc.doubleQuote(key) + "," + rpc.doubleQuote(value) + "]"
	}
	err := rpc.post("db_putHex", &response, params)
	return response, err
}

// db_getHex
// Returns binary data from the local database.
// Note this function is deprecated and will be removed in the future.
// Parameters
//
//    String - Database name.
//    String - Key name.
// Example Parameters
//
// params: [
//  "testDB",
//  "myKey",
// ]
//
// Returns
//
// DATA - The previously stored data.
func (rpc *EthRPC) DbGetHex(dbname string, key string) (string, error) {
	var value string
	params := func() string {
		return "[" + rpc.doubleQuote(dbname) + "," + rpc.doubleQuote(key) + "]"
	}
	err := rpc.post("db_getHex", &value, params)
	return value, err
}

// shh_version
// Returns the current whisper protocol version.
// Parameters
// 		none
// Returns
// 		String - The current whisper protocol version
func (rpc *EthRPC) ShhVersion() (string, error) {
	var sshVersion string
	err := rpc.post("shh_version", &sshVersion, nil)
	return sshVersion, err
}

// shh_post
// Sends a whisper message.
// Parameters
//
//     Object - The whisper post object:
//
//     from: DATA, 60 Bytes - (optional) The identity of the sender.
//     to: DATA, 60 Bytes - (optional) The identity of the receiver. When present whisper will encrypt the message so that only the receiver can decrypt it.
//     topics: Array of DATA - Array of DATA topics, for the receiver to identify messages.
//     payload: DATA - The payload of the message.
//     priority: QUANTITY - The integer of the priority in a range from ... (?).
//     ttl: QUANTITY - integer of the time to live in seconds.
//
// Example Parameters
//
// params: [{
//   from: "0x04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a03e245533f97284d442460f2998cd41858798ddfd4d661997d3940272b717b1",
//   to: "0x3e245533f97284d442460f2998cd41858798ddf04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a0d4d661997d3940272b717b1",
//   topics: ["0x776869737065722d636861742d636c69656e74", "0x4d5a695276454c39425154466b61693532"],
//   payload: "0x7b2274797065223a226d6",
//   priority: "0x64",
//   ttl: "0x64",
// }]
//
// Returns
//
// Boolean - returns true if the message was send, otherwise false.
// TODO develop this option
func (rpc *EthRPC) ShhPost(data model.WhisperParams) (string, error) {
	var sshVersion string
	err := rpc.post("shh_post", &sshVersion, nil)
	return sshVersion, err
}

// shh_newIdentity
// Creates new whisper identity in the client.
// Parameters: none
// Returns
//		DATA, 60 Bytes - the address of the new identiy.
func (rpc *EthRPC) ShhNewIdentity() (string, error) {
	var newIdAddress string
	err := rpc.post("shh_newIdentity", &newIdAddress, nil)
	return newIdAddress, err
}

// shh_hasIdentity
// Checks if the client hold the private keys for a given identity.
// Parameters
//    DATA, 60 Bytes - The identity address to check.
//Example Parameters
//params: [  "0x04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a03e245533f97284d442460f2998cd41858798ddfd4d661997d3940272b717b1"
//]
// Returns
// Boolean - returns true if the client holds the privatekey for that identity, otherwise false.
func (rpc *EthRPC) SshHasIdentity(identityKey string) (bool, error) {
	var response bool
	params := func() string {
		return "[" + rpc.doubleQuote(identityKey) + "]"
	}
	err := rpc.post("shh_hasIdentity", &response, params)
	return response, err
}

// shh_newGroup
// Creates a new group.
// Parameters: none
//Returns
//	DATA, 60 Bytes - the address of the new group.
func (rpc *EthRPC) ShhNewGroup() (string, error) {
	var groupId string
	err := rpc.post("shh_newGroup", &groupId, nil)
	return groupId, err
}

// shh_addToGroup
// Adds a whisper identity to the group.
//Parameters
//    DATA, 60 Bytes - The identity address to add to a group.
// Example Parameters
// params: [ "0x04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a03e245533f97284d442460f2998cd41858798ddfd4d661997d3940272b717b1"
// ]
// Returns
// 		Boolean - returns true if the identity was successfully added to the group, otherwise false.
func (rpc *EthRPC) ShhAddToGroup(identityKey string) (bool, error) {
	var response bool
	params := func() string {
		return "[" + rpc.doubleQuote(identityKey) + "]"
	}
	err := rpc.post("shh_addToGroup", &response, params)
	return response, err
}

// shh_newFilter
// Creates filter to notify, when client receives whisper message matching the filter options.
// Parameters
//
//     Object - The filter options:
//
//     to: DATA, 60 Bytes - (optional) Identity of the receiver. When present it will try to decrypt any incoming message if the client holds the private key to this identity.
//     topics: Array of DATA - Array of DATA topics which the incoming message's topics should match. You can use the following combinations:
//         [A, B] = A && B
//         [A, [B, C]] = A && (B || C)
//         [null, A, B] = ANYTHING && A && B null works as a wildcard
//
// Example Parameters
//
// params: [{
//    "topics": ['0x12341234bf4b564f'],
//    "to": "0x04f96a5e25610293e42a73908e93ccc8c4d4dc0edcfa9fa872f50cb214e08ebf61a03e245533f97284d442460f2998cd41858798ddfd4d661997d3940272b717b1"
// }]
//
// Returns
//
// QUANTITY - The newly created filter.
// TODO develop
func (rpc *EthRPC) ShhNewFilter() (string, error) {
	var groupId string
	err := rpc.post("shh_newGroup", &groupId, nil)
	return groupId, err
}

// shh_uninstallFilter
// Uninstalls a filter with given id. Should always be called when
// watch is no longer needed. Additonally Filters timeout when they
// aren't requested with shh_getFilterChanges for a period of time.
// Parameters
//    QUANTITY - The filter id.
// Example Parameters
// params: [
// "0x7" // 7
// ]
// Returns
// Boolean - true if the filter was successfully uninstalled, otherwise false.
func (rpc *EthRPC) ShhUninstallFilter() (bool, error) {
	var response bool
	err := rpc.post("shh_uninstallFilter", &response, nil)
	return response, err
}

// shh_getFilterChanges
// Polling method for whisper filters. Returns new messages since the last call of this method.
// Note calling the shh_getMessages method, will reset the buffer for this method, so that you won't receive duplicate messages.
// Parameters
//    QUANTITY - The filter id.
// Example Parameters
// params: [
// "0x7" // 7
// ]
// Returns
// Array - Array of messages received since last poll:
//    hash: DATA, 32 Bytes (?) - The hash of the message.
//    from: DATA, 60 Bytes - The sender of the message, if a sender was specified.
//    to: DATA, 60 Bytes - The receiver of the message, if a receiver was specified.
//    expiry: QUANTITY - Integer of the time in seconds when this message should expire (?).
//    ttl: QUANTITY - Integer of the time the message should float in the system in seconds (?).
//    sent: QUANTITY - Integer of the unix timestamp when the message was sent.
//    topics: Array of DATA - Array of DATA topics the message contained.
//    payload: DATA - The payload of the message.
//    workProved: QUANTITY - Integer of the work this message required before it was send (?).
func (rpc *EthRPC) ShhGetFilterChanges() (string, error) {
	var groupId string
	err := rpc.post("shh_getFilterChanges", &groupId, nil)
	return groupId, err
}

// shh_getMessages
// Get all messages matching a filter. Unlike shh_getFilterChanges this returns all messages.
// Parameters
//    QUANTITY - The filter id.
// params: [
//  "0x7" // 7
// ]
// Returns: see shh_getFilterChanges
func (rpc *EthRPC) SshGetMessages(filterId string) ([]model.FilterChanges, error) {
	var response []model.FilterChanges
	params := func() string {
		return "[" + rpc.doubleQuote(filterId) + "]"
	}
	err := rpc.post("shh_getMessages", &response, params)
	return response, err
}
func (rpc *EthRPC) generateTransactionPayload(contract string, data string, block string, gas string, gasprice string, params *model.EthRequestParams) string {
	requestParams := map[string]interface{}{
		"to":   contract,
		"data": data,
		/*
			"gas":      "0xaae60", //700000,
			"gasPrice": "0x15f90", //90000,
		*/
	}
	if gas != "" {
		requestParams["gas"] = gas
	}
	if gasprice != "" {
		requestParams["gasPrice"] = gasprice
	}
	raw, _ := json.Marshal(requestParams)
	paramsStr := str.UnsafeString(raw)
	request := `{"id":1, "jsonrpc":"2.0","method":"eth_call","params":[` + paramsStr + `]}`
	return request
}

// todo add from field
func (rpc *EthRPC) generateCallPayload(contract string, data string, block string) string {
	if block != model.NoPeriod {
		request := `{"id": 1,"jsonrpc": "2.0","method": "eth_call",
"params":[{
"to": ` + rpc.doubleQuote(contract) + `,
"data": ` + rpc.doubleQuote(data) + `},
` + rpc.doubleQuote(block) + `]}`
		return request
	} else {
		request := `{"id": 1,"jsonrpc": "2.0","method": "eth_call",
"params":[{
"to": ` + rpc.doubleQuote(contract) + `,
"data": ` + rpc.doubleQuote(data) + `}]}`
		return request
	}
}

// this method converts standard contract params to abi encoded params given a
// contract address, method name and abi model
func (rpc *EthRPC) convertParamsToAbi(contract string, method string, args interface{}) ([]byte, error) {
	var abiModel abi.ABI
	//try to fetch the abi model linked to given contract address
	return abiModel.Pack(method, args)
}

// post ethereum network contract with no parameters
func (rpc *EthRPC) ContractCall(contract string, methodName string, params string, block string, gas string, gasprice string) (string, error) {
	abiparams, abiEncErr := rpc.convertParamsToAbi(contract, methodName, params)
	if abiEncErr != nil {
		//failed to encode post abi data
		logger.Error("failed to encode contract post abi parameters: ", abiEncErr)
		return "", abiEncErr
	} else {
		paramsStr := str.UnsafeString(abiparams)
		data := paramsStr + "," + "," + gas + "," + gasprice
		payload := rpc.generateCallPayload(contract, data, block)
		raw, err := rpc.makePostString(payload)
		if err == nil {
			var data string
			unErr := json.Unmarshal(raw, &data)
			return data, unErr
		}
		return "", err
	}
}

// post ethereum network contract with no parameters
func (rpc *EthRPC) contractCallAbiParams(contract string, data string, block string) (string, error) {
	payload := rpc.generateCallPayload(contract, data, block)
	raw, err := rpc.makePostString(payload)
	if err == nil {
		var data string
		unErr := json.Unmarshal(raw, &data)
		return data, unErr
	}
	return "", err
}

func (rpc *EthRPC) Erc20Summary(contract string) (map[string]string, error) {
	var response = map[string]string{
		summaryFunctionsNames[0]: "",
		summaryFunctionsNames[1]: "",
		summaryFunctionsNames[2]: "",
		summaryFunctionsNames[3]: "",
	}
	// copy target url to summary functions
	instance.url = rpc.url
	// execute erc20 summary functions
	for i := 0; i < len(summaryFunctions); i++ {
		raw, err := summaryFunctions[i](contract)
		if err == nil && raw != "" {
			response[summaryFunctionsNames[i]] = raw
		}
	}
	return response, nil
}

func (rpc *EthRPC) Erc20TotalSupply(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, erc20.TotalSupplyParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Symbol(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, erc20.SymbolParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Name(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, erc20.NameParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Decimals(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, erc20.DecimalsParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20BalanceOf(contract string, tokenOwner string) (json.RawMessage, error) {
	tokenOwnerAddress, decodeErr := fromStringToAddress(tokenOwner)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := erc20.LoadErc20Abi().Pack("balanceOf", tokenOwnerAddress)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 balanceof method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	req := rpc.generateCallPayload(contract, dataContent, model.LatestBlockNumber)
	return rpc.makePostString(req)
}

func (rpc *EthRPC) Erc20Allowance(contract string, tokenOwner string, spender string) (json.RawMessage, error) {
	tokenOwnerAddress, decodeErr := fromStringToAddress(tokenOwner)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	spenderAddress, decodeErr := fromStringToAddress(spender)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := erc20.LoadErc20Abi().Pack("allowance", tokenOwnerAddress, spenderAddress)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 allowance method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	params := model.EthRequestParams{
		To:   contract,
		Data: dataContent,
		Tag:  model.LatestBlockNumber,
	}
	return rpc.makePostWithMethodParams("eth_call", params.String())
}

func (rpc *EthRPC) Erc20Transfer(contract string, address string, amount int) (json.RawMessage, error) {
	senderAddress, decodeErr := fromStringToAddress(address)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := erc20.LoadErc20Abi().Pack("transfer", senderAddress, amount)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 transfer method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	params := model.EthRequestParams{
		To:   contract,
		Data: dataContent,
		Tag:  model.LatestBlockNumber,
	}
	return rpc.makePostWithMethodParams("eth_sendTransaction", params.String())
}

// Chain id returns target network chain id
func (rpc *EthRPC) ChainId() (string, error) {
	var result string
	err := rpc.post("eth_chainId", &result, nil)
	if err != nil {
		return "", err
	}
	return result, nil
}

// NetworkID returns the network ID (also known as the chain ID) for this chain.
func (rpc *EthRPC) NetworkId() (*big.Int, error) {
	version := new(big.Int)
	var ver string
	ver, err := rpc.NetVersion()
	if err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (rpc *EthRPC) PendingNonceAt(address string) (uint64, error) {
	result, err := rpc.EthGetTransactionCount(address, "pending")
	return uint64(result), err
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (rpc *EthRPC) SuggestGasPrice() (*big.Int, error) {
	var response string
	err := rpc.post("eth_gasPrice", &response, nil)
	if err != nil {
		return nil, err
	}
	v, _, err := ParseBigInt(response)
	return v, err
}

// curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x8aff0a12f3e8d55cc718d36f84e002c335df2f4a", "data": "606060405260728060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480636ffa1caa146037576035565b005b604b60048080359060200190919050506061565b6040518082815260200191505060405180910390f35b6000816002029050606d565b91905056"}],"id":1}
func (rpc *EthRPC) DeployContract(fromAddress string, bytecode string, gas string, gasPrice string) (json.RawMessage, error) {
	payload := `[{"from":` + rpc.doubleQuote(fromAddress) + `,
"data":` + rpc.doubleQuote(bytecode) + `,
"gasprice":` + rpc.doubleQuote(gasPrice) + `,
"gas":` + rpc.doubleQuote(gas) + `}]`
	return rpc.makePostWithMethodParams("eth_sendTransaction", payload)
}

func (rpc *EthRPC) IsSmartContractAddress(addr string) (bool, error) {
	bytecode, err := rpc.EthGetCode(addr, model.LatestBlockNumber)
	// if the address has valid bytecode, is a contract
	// if is not code addres 0x is returned
	return len(bytecode) > 2, err
}

// helper methods

func fromStringToAddress(addr string) (fixtures.Address, error) {
	var a fixtures.Address
	raw, decodeErr := hex.FromEthHex(addr)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return a, decodeErr
	}
	a.SetBytes(raw)
	return a, nil
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return oneEth
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1Int64() int64 {
	return oneEthInt64
}

// double quotes given string
func (rpc *EthRPC) doubleQuote(data string) string {
	return `"` + data + `"`
}

func (rpc *EthRPC) Proxy(requestContent []byte) ([]byte, stack.Error) {
	result, err := rpc.proxyRequest(requestContent)
	return result, stack.Ret(err)
}

// method that calls graph ql api of ethereum nodes (supported on geth +1.9.0)
func (rpc *EthRPC) GraphQL(graphQLendpointURI string, qlQuery []byte) ([]byte, stack.Error) {
	result, err := httpclient.MakePost(rpc.client, graphQLendpointURI, defaultGraphQLHeader, qlQuery)
	return result, stack.Ret(err)
}

func ResolveNetworkId(id string) string {
	if id == "0" {
		return "Olympic, Ethereum public pre-release PoW testnet"
	} else if id == "1" {
		return "Ethereum Mainnet"
	} else if id == "2" {
		return "Morden Testnet"
	} else if id == "3" {
		return "Ropsten Testnet"
	} else if id == "4" {
		return "Rinkeby Testnet"
	} else if id == "5" {
		return "Goerli, the public cross-client PoA testnet"
	} else if id == "6" {
		return "Kotti Classic, the public cross-client PoA testnet for Classic"
	} else if id == "8" {
		return "Ubiq, the public Gubiq main network with flux difficulty"
	} else if id == "42" {
		return "Kovan, the public Parity-only PoA testnet"
	} else if id == "60" {
		return "GoChain, the GoChain networks mainnet"
	} else if id == "61" {
		return "Ethereum Classic PoW main network"
	} else if id == "77" {
		return "Sokol, the public POA Network testnet"
	} else if id == "99" {
		return "Core, the public POA Network main network"
	} else if id == "100" {
		return "xDai, the public MakerDAO/POA Network main network"
	} else if id == "5777" {
		return "Ganache Local Testnet"
	} else if id == "31337" {
		return "GoChain testnet, the GoChain networks public testnet"
	} else if id == "401697" {
		return "Tobalaba, the public Energy Web Foundation testnet"
	} else if id == "7762959" {
		return "Musicoin, the music blockchain"
	} else if id == "61717561" {
		return "Aquachain, ASIC resistant chain"
	}
	return "unknown"
}

var (
	networkLookupMap = map[int]string{
		0:        "Olympic, Ethereum public pre-release PoW testnet",
		1:        "Ethereum Mainnet",
		2:        "Morden Testnet",
		3:        "Ropsten Testnet",
		4:        "Rinkeby Testnet",
		5:        "Goerli, the public cross-client PoA testnet",
		6:        "Kotti Classic, the public cross-client PoA testnet for Classic",
		8:        "Ubiq, the public Gubiq main network with flux difficulty chain ID 8",
		42:       "Kovan, the public Parity-only PoA testnet",
		60:       "GoChain, the GoChain networks mainnet",
		61:       "Ethereum Classic PoW main network",
		77:       "Sokol, the public POA Network testnet",
		99:       "Core, the public POA Network main network",
		100:      "xDai, the public MakerDAO/POA Network main network",
		5777:     "Ganache Local Testnet",
		31337:    "GoChain testnet, the GoChain networks public testnet",
		401697:   "Tobalaba, the public Energy Web Foundation testnet",
		7762959:  "Musicoin, the music blockchain",
		61717561: "Aquachain, ASIC resistant chain",
	}
)

// @deprecated
func ResolveNetworkId2(id string) string {
	idx, err := strconv.Atoi(id)
	if err == nil {
		name, ok := networkLookupMap[idx]
		if ok {
			return name
		}
	}
	return ""
}
