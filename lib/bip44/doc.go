// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip44

/*
Path levels

We define the following 5 levels in BIP32 path:

m / purpose' / coin_type' / account' / change / address_index

Apostrophe in the path indicates that BIP32 hardened derivation is used.

Each level has a special meaning, described in the chapters below.
Purpose

Purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation. It indicates that the subtree of this node is used according to this specification.

Hardened derivation is used at this level.

## Package performance

* initial analysis results:

BenchmarkBIP44/NewKeyFromMnemonic-4         	     300	   4077059 ns/op	   0.00 MB/s	   16340 B/op	     233 allocs/op
BenchmarkBIP44/NewKeyFromMasterKey-4        	     300	   4096654 ns/op	   0.00 MB/s	   18100 B/op	     261 allocs/op

*/
