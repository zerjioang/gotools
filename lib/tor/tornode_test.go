// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package tor

import (
	"testing"

	"github.com/zerjioang/gotools/lib/radix"
)

func TestCreateRadix(t *testing.T) {
	// Create a tree
	r := radix.New()
	/*
		1.163.34.119
		1.172.104.133
		1.41.132.176
		100.1.197.216
	*/
	r.Insert("1.163.34.119", nil)
	r.Insert("1.172.104.133", nil)
	r.Insert("1.41.132.176", nil)

	// Find the longest prefix match
	_, _ = r.Get("1.41.132.176")
}
