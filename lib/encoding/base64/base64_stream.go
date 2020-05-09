package base64

import (
	"unsafe"

	"github.com/zerjioang/gotools/lib/encoding/generic"
	"github.com/zerjioang/gotools/lib/logger"
)

const (
	empty = ""
)

var (
	charMap        = []byte(encodeStd)
	stdPaddingByte = byte(StdPadding)
)

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		logger.Debug("little endian")
	case [2]byte{0xAB, 0xCD}:
		logger.Debug("big endian")
	default:
		logger.Error("could not determine native endianness.")
	}
}

/*
 The Base64 encoding process is to:

    Divid the input bytes stream into blocks of 3 bytes.
    Divid 24 bits of each 3-byte block into 4 groups of 6 bits.
    Map each group of 6 bits to 1 printable character, based on the 6-bit value using the Base64 character set map.
    If the last 3-byte block has only 1 byte of input data, pad 2 bytes of zero (\x0000). After encoding it as a normal block, override the last 2 characters with 2 equal signs (==), so the decoding process knows 2 bytes of zero were padded.
    If the last 3-byte block has only 2 bytes of input data, pad 1 byte of zero (\x00). After encoding it as a normal block, override the last 1 character with 1 equal signs (=), so the decoding process knows 1 byte of zero was padded.
    Carriage return (\r) and new line (\n) are inserted into the output character stream. They will be ignored by the decoding process.

This solution requires following operations:

    3 memory loads;
    3 bit-ands;
    2 bit-ors
    5 bit-shifts.

*/
func EncodeToString(src []byte) string {
	//calculate group count
	sizeof := len(src)
	if sizeof == 0 {
		return empty
	}
	// calculate encoded base64 text size
	resultSize := (sizeof + 2) / 3 * 4
	// allocate space for storing result value
	dst := make([]byte, resultSize)
	// destination index counter
	// source data index counter
	di, si := 0, 0
	n := (sizeof / 3) * 3
	for si < n {
		//Divid the input bytes stream into blocks of 3 bytes.
		b3 := src[si+2]
		b2 := src[si+1]
		b1 := src[si+0]

		// Convert 3x 8bit source bytes into 4 bytes of 6 bits.
		// Map each group of 6 bits to 1 printable character, based on the 6-bit value using the Base64 character set map.

		// uint32
		// these three 8-bit (ASCII) characters become one 24-bit number
		mapIdx := uint(b1)<<16 | uint(b2)<<8 | uint(b3)
		// this 24-bit number gets separated into four 6-bit numbers
		// those four 6-bit numbers are used as indices into the base64 character list
		dst[di+0] = resolve(mapIdx >> 18 & 0x3F)
		dst[di+1] = resolve(mapIdx >> 12 & 0x3F)
		dst[di+2] = resolve(mapIdx >> 6 & 0x3F)
		dst[di+3] = resolve(mapIdx & 0x3F)

		//increase source data counter
		si += 3
		// increase destination index to the next 4 bytes group
		di += 4
	}
	remain := sizeof - si

	// Add the remaining small block
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	dst[di+0] = charMap[val>>18&0x3F]
	dst[di+1] = charMap[val>>12&0x3F]

	switch remain {
	case 2:
		dst[di+2] = charMap[val>>6&0x3F]
		dst[di+3] = stdPaddingByte
	case 1:
		dst[di+2] = stdPaddingByte
		dst[di+3] = stdPaddingByte
	}
	return string(dst)
}

func resolve(v uint) byte {
	return charMap[v]
}

func EncodeToStream(src []byte, writer generic.Writer) {
	//calculate group count
	sizeof := len(src)
	if sizeof == 0 {
		return
	}
	// destination index counter
	// source data index counter
	di, si := 0, 0
	n := (sizeof / 3) * 3
	for si < n {
		//Divid the input bytes stream into blocks of 3 bytes.
		b3 := src[si+2]
		b2 := src[si+1]
		b1 := src[si+0]

		// Convert 3x 8bit source bytes into 4 bytes of 6 bits.
		// Map each group of 6 bits to 1 printable character, based on the 6-bit value using the Base64 character set map.

		// these three 8-bit (ASCII) characters become one 24-bit number
		mapIdx := uint(b1)<<16 | uint(b2)<<8 | uint(b3)
		// this 24-bit number gets separated into four 6-bit numbers
		// those four 6-bit numbers are used as indices into the base64 character list
		writer(charMap[mapIdx>>18&0x3F])
		writer(charMap[mapIdx>>12&0x3F])
		writer(charMap[mapIdx>>6&0x3F])
		writer(charMap[mapIdx&0x3F])

		//increase source data counter
		si += 3
		// increase destination index to the next 4 bytes group
		di += 4
	}
	remain := sizeof - si

	// Add the remaining small block
	val := uint(src[si+0]) << 16
	if remain == 2 {
		val |= uint(src[si+1]) << 8
	}

	writer(charMap[val>>18&0x3F])
	writer(charMap[val>>12&0x3F])

	switch remain {
	case 2:
		writer(charMap[val>>6&0x3F])
		writer(stdPaddingByte)
	case 1:
		writer(stdPaddingByte)
	}
}
