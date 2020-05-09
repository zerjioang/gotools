// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package str

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

var (
	empty    []byte
	jsonfast = jsoniter.ConfigFastest
	jsonstd  = jsoniter.ConfigCompatibleWithStandardLibrary
)

func ReadFileAsString(path string) string {
	data, _ := ReadFileData(path)
	return data
}

func ReadFileData(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func UnsafeBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// A hack until issue golang/go#2632 is fixed.
// See: https://github.com/golang/go/issues/2632
func UnsafeString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// converts ascii chars of a given string in lowercase
// this function is at least
// twice as fast as standard to lower function of go standard library
func ToLowerAscii(src string) string {
	if src == "" {
		return src
	}
	rawBytes := []byte(src)
	start := uintptr(unsafe.Pointer(&rawBytes[0]))
	s := len(rawBytes)
	for i := 0; i < s; i++ {
		// get char at current index
		c := *(*byte)((unsafe.Pointer)(start + uintptr(i)))
		if c >= 'A' && c <= 'Z' {
			*(*byte)((unsafe.Pointer)(start + uintptr(i))) = c + 32
		}
	}
	return *(*string)(unsafe.Pointer(&rawBytes))
}

// A slightly faster lowercase function.
func toLower(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

func strLen(s string) int {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
}

func IntToByte(v int) []byte {
	b := []byte(strconv.FormatInt(int64(v), 10))
	return b
}

func StreamToByte(stream io.Reader) []byte {
	buf := bytes.Buffer{}
	_, _ = buf.ReadFrom(stream)
	return buf.Bytes()
}
