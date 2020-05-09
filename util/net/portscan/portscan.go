package portscan

import (
	"net"
	"strconv"
	"time"

	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/lib/stack"
)

var (
	errTimeOut = stack.New("timeout error checking port liveness")
	// known json-rpc ports used for common apps such as geth, parity
	knownPorts = []int{8545, 7545, 4000, 3000}
)

// generic function that checks if a remote tcp port is opened or not
func IsOpenPort(host string, port int, timeoutSecs int64) (bool, stack.Error) {
	// return data
	isOpen := false
	stErr := stack.Nil()

	// host -> the remote host
	// timeoutSecs -> the timeout value
	endpoint := host + ":" + strconv.Itoa(port)
	logger.Debug("checking if port is open for endpoint: ", endpoint)
	conn, err := net.DialTimeout("tcp", endpoint, time.Duration(timeoutSecs)*time.Second)

	if terr, ok := err.(*net.OpError); ok && terr.Timeout() {
		stErr = errTimeOut
	} else if err != nil {
		stErr = stack.Ret(err)
	} else {
		isOpen = true
		_ = conn.Close()
	}
	return isOpen, stErr
}

// tries to find a common json-rpc port
func FindJsonRpcPort(endpoint string) int {
	logger.Debug("finding json-rpc port available")
	for _, port := range knownPorts {
		op, err := IsOpenPort(endpoint, port, 3)
		if err.None() && op {
			return port
		}
	}
	return -1
}
