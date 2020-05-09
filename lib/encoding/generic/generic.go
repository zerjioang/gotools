// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package generic

import (
	"bytes"
	"strings"
)

type Writer func(byte)
type StreamInterface interface {
	Write(byte)
}

// encode buffer to given common
func NbaseEncode(nb uint64, buf *bytes.Buffer, base string) {
	l := uint64(len(base))
	if nb/l != 0 {
		NbaseEncode(nb/l, buf, base)
	}
	buf.WriteByte(base[nb%l])
}

//decode string to given common
func NbaseDecode(enc, base string) uint64 {
	var nb uint64
	lbase := len(base)
	le := len(enc)
	for i := 0; i < le; i++ {
		mult := 1
		for j := 0; j < le-i-1; j++ {
			mult *= lbase
		}
		nb += uint64(strings.IndexByte(base, enc[i]) * mult)
	}
	return nb
}
