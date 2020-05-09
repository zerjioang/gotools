package torips

import (
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/lib/tor"
)

var (
	list tor.TorList
)

/*
this method should one be executed once!
*/
func Init(path string, debug bool) {
	logger.Info("loading tor nodes hash data")
	list.LoadIps(path)
}

func Size() int {
	return list.Size()
}

func IsBackListedIp(ip uint32) bool {
	return list.Contains(ip)
}
