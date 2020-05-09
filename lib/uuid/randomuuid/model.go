// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package randomuuid

import (
	"encoding/hex"

	"github.com/zerjioang/gotools/util/str"
)

type UniqueIdHex [32]byte

func (uid UniqueIdHex) String() string {
	return string(uid[:])
}

func (uid UniqueIdHex) UnsafeString() string {
	return str.UnsafeString(uid[:])
}

func (uid UniqueIdHex) Slice() []byte {
	return uid[:]
}
func (uid UniqueIdHex) Bytes() [32]byte {
	return uid
}

type UniqueIdRaw [16]byte

func (uid UniqueIdRaw) String() string {
	return hex.EncodeToString(uid[:])
}

func (uid UniqueIdRaw) Slice() []byte {
	return uid[:]
}

func (uid UniqueIdRaw) Bytes() [16]byte {
	return uid
}
