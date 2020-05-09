// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hex

import (
	"errors"
	"fmt"
	"unsafe"
)

// ErrLength reports an attempt to decode an odd-length input
// using Decode or DecodeString.
// The stream-based Decoder returns ioproto.ErrUnexpectedEOF instead of ErrLength.
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError values describe errors resulting from an invalid byte in a hex string.
type InvalidByteError byte

func init() {
}

// UnsafeDecodeString returns the bytes represented by the hexadecimal string s.
//
// UnsafeDecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func UnsafeDecodeString(s string) ([]byte, error) {
	//make string mutable
	src := []byte(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := decodeUnsafe(src, src)
	return src[:n], err
}

// DecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func DecodeString(s string) ([]byte, error) {
	//make string mutable
	src := []byte(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := decode(src, src)
	return src[:n], err
}

// todo: WIP
func decodeUnsafe(dst, src []byte) (int, error) {

	start := uintptr(unsafe.Pointer(&dst[0]))
	step := unsafe.Sizeof(dst[0])

	var i int
	for i = 0; i < len(src)/2; i++ {
		c := src[i*2]
		a, ok1 := fromHexChar(src[i*2])
		b, ok2 := fromHexChar(src[i*2+1])
		if !ok1 || !ok2 {
			return i, InvalidByteError(c)
		}
		*(*byte)((unsafe.Pointer)(start + step*uintptr(i))) = (a << 4) | b
	}
	if len(src)%2 == 1 {
		// Check for invalid char before reporting bad length,
		// since the invalid char (if present) is an earlier problem.
		c := src[i*2]
		if _, ok := fromHexChar(c); !ok {
			return i, InvalidByteError(c)
		}
		return i, ErrLength
	}
	return i, nil
}

// Decode decodes src into DecodedLen(len(src)) bytes,
// returning the actual number of bytes written to dst.
//
// Decode expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, Decode returns the number
// of bytes decoded before the error.
func decode(dst, src []byte) (int, error) {
	i, j := 0, 1
	for ; j < len(src); j += 2 {
		a, ok := fromHexChar(src[j-1])
		if !ok {
			return i, InvalidByteError(src[j-1])
		}
		b, ok := fromHexChar(src[j])
		if !ok {
			return i, InvalidByteError(src[j])
		}
		dst[i] = (a << 4) | b
		i++
	}
	if len(src)%2 == 1 {
		// Check for invalid char before reporting bad length,
		// since the invalid char (if present) is an earlier problem.
		if _, ok := fromHexChar(src[j-1]); !ok {
			return i, InvalidByteError(src[j-1])
		}
		return i, ErrLength
	}
	return i, nil
}

// fromHexChar converts a hex character into its value and a success flag.
func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}
	return 0, false
}
func (e InvalidByteError) Error() string {
	return fmt.Sprintf("encoding/hex: invalid byte: %#U", rune(e))
}
