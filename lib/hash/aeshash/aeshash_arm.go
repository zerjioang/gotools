// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// +build arm arm64 arm64be

package aeshash

import (
	"hash/fnv"

	"github.com/zerjioang/gotools/util/str"
)

// Hash hashes the given string using the algorithm used by Go's hash tables
// God knows what it really is.
func Hash(key string) uint32 {
	h := fnv.New32a()
	h.Write(str.UnsafeBytes(key))
	return h.Sum32()
}
