// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package base64

var (
	ssse3 = false
)

func init() {
	ssse3 = _haveSSSE3()
}

// haveSSSE3 returns whether the CPU supports SSSE3 instructions (i.e. PSHUFB).
// Note that this is SSSE3, not SSE3.
// note: this is an assembly instruction
// go:noescape
// BenchmarkFastEncode/has-ssse3-4
// 14011170   86.1 ns/op   11.61 MB/s   0 B/op   0 allocs/op
func _haveSSSE3() bool

// BenchmarkFastEncode/has-ssse3-4
// 1000000000   0.404 ns/op	2476.74 MB/s   0 B/op   0 allocs/op
func haveSSSE3() bool {
	return ssse3
}

// note: this is an assembly instruction
// go:noescape
func _encode12ByteGroups(lookup []int8, dst, src []byte) (di int, si int)

func encode12ByteGroups(lookup []int8, dst, src []byte) (di int, si int) {
	return _encode12ByteGroups(lookup, dst, src)
}

var lookupStd = []int8{
	65, 71, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -19, -16, 0, 0,
}

var lookupURL = []int8{
	65, 71, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -17, 32, 0, 0,
}

func encodeAccelerated(enc *Encoding, dst, src []byte) (int, int) {
	// if no ssse3 available
	if !ssse3 {
		return 0, 0
	}
	// If the source slice is less than 12 bytes fallback to the standard
	// go encoder.
	if len(src) < 12 {
		return 0, 0
	}

	// If our SIMD map is too small or not set fallback to go encoder.
	// This will happen if a non-standard encoding is being used.
	if len(enc.accEncode) < 16 {
		return 0, 0
	}
	// The destination slice is too small.  As the assembly code doesn't
	// check slice bounds we're going to fallback to the go code, which
	// will panic appropriately.
	if len(dst) < enc.EncodedLen(len(src)) {
		return 0, 0
	}
	return encode12ByteGroups(enc.accEncode, dst, src)
}

func accelerateEncodeMap(encoder string) []int8 {
	switch encoder {
	case encodeStd:
		return lookupStd
	case encodeURL:
		return lookupURL
	}
	return nil
}
