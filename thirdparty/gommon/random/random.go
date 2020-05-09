// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package random

import (
	"bytes"
	"math/rand"
	"strings"
	"unsafe"

	"github.com/zerjioang/gotools/lib/fastime"
)

type (
	Random struct {
	}
)

// Charsets
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	Hex          = Numeric + "abcdef"
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	global = New()
	src    = rand.NewSource(fastime.UnixNano())
)

func New() *Random {
	rand.Seed(fastime.UnixNano())
	return new(Random)
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}

func String(length uint8, charsets ...string) string {
	return global.String(length, charsets...)
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	b [32]byte
)

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// concurrent safe 32 bytes UIUD generation function
func RandomUUID32() string {
	var b bytes.Buffer
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	var i, remain uint8
	var cache int64
	i, cache, remain = 31, src.Int63(), letterIdxMax
	for i != 255 {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	//return b.String()
	raw := b.Bytes()
	return *(*string)(unsafe.Pointer(&raw))
}

// concurrent not safe 32 bytes UIUD generation function
func RandomUUID32Shared() string {
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	var i, remain uint8
	var cache int64
	i, cache, remain = 31, src.Int63(), letterIdxMax
	for i != 255 {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	//return b.String()
	raw := b[:]
	return *(*string)(unsafe.Pointer(&raw))
}
