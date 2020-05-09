// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex

import (
	"encoding/hex"
	"unsafe"
)

const (
	hextable     = "0123456789abcdef"
	doubleQuotes = 34
)

func ToEthHex(raw []byte) string {
	dst := make([]byte, 2+len(raw)*2)
	dst[0] = 48  // 0
	dst[1] = 120 // x
	for i, v := range raw {
		dst[2+i*2] = hextable[v>>4]
		dst[2+i*2+1] = hextable[v&0x0f]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func ToHex(raw []byte) string {
	dst := make([]byte, len(raw)*2)
	for i, v := range raw {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

// decode an standard hex string
func FromHex(raw string) ([]byte, error) {
	return hex.DecodeString(raw)
}

// decode ethereum 0x hex string
func FromEthHex(content string) ([]byte, error) {
	raw := []byte(content)
	idx := 0
	if raw[0] == 48 && raw[1] == 120 {
		//starts with 0x
		idx = 2
	}
	end := len(raw)
	if raw[0] == doubleQuotes {
		idx = 3
		end -= 1
	}
	data := raw[idx:end]
	n, err := hex.Decode(data, data)
	return data[:n], err
}
