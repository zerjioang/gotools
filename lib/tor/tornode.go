// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package tor

import (
	"io/ioutil"
	"strings"

	"github.com/zerjioang/gotools/lib/hashset"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/util/ip"
	"github.com/zerjioang/gotools/util/str"
)

type TorList struct {
	hashset.HashSetAtomic
}

func (l *TorList) LoadIps(path string) {
	logger.Debug("loading tor node list")
	if path != "" {
		logger.Debug("loading tor list with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		itemList := strings.Split(str.UnsafeString(data), "\n")
		if itemList != nil {
			for _, v := range itemList {
				//convert string ip to uint
				ipvalue := ip.Ip2intLow(v)
				l.UnsafeAddUint32(ipvalue)
			}
		}
	}
}
