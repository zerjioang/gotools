// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package badips

import (
	"github.com/zerjioang/gotools/lib/hashset"
	"github.com/zerjioang/gotools/lib/logger"
)

var (
	// https://www.howtoforge.com/tutorial/protect-your-server-computer-with-badips-and-fail2ban/
	// loaded from https://www.badips.com/get/list/any/5
	badIpList hashset.HashSetWORM
)

/*
this method should one be executed once!
*/
func Init(path string, debug bool) {
	logger.Debug("loading bad ip list policy data")
	badIpList = hashset.NewHashSetWORM()
	badIpList.LoadFromRaw(path, "\n")
}

func GetBadIPList() hashset.HashSetWORM {
	return badIpList
}

func Size() int {
	return badIpList.Size()
}

func IsBackListedIp(ip string) bool {
	return badIpList.Contains(ip)
}
