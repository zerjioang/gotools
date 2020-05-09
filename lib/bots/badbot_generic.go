// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bots

import (
	"github.com/zerjioang/gotools/common"
	"github.com/zerjioang/gotools/lib/hashset"
	"github.com/zerjioang/gotools/lib/logger"
)

var (
	badBotsList hashset.HashSetWORM
)

/*
this method should one be executed once!
*/
func Init(filepath string, debug bool) {
	logger.Debug("[module] loading anti-bots policy data")
	badBotsList = hashset.NewHashSetWORM()
	badBotsList.LoadFromRaw(filepath, common.NewLine)
	if debug {
		// allow calls made with apachebench and curl tools
		badBotsList.Remove("apachebench")
		badBotsList.Remove("curl")
	}
}

func GetBadBotsList() hashset.HashSetWORM {
	return badBotsList
}

func MatchAny(userAgent string) bool {
	return badBotsList.MatchAny(userAgent)
}

func Size() int {
	return badBotsList.Size()
}
