// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ethrpc

import (
	"fmt"
	"math/big"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSizeOf(t *testing.T) {
	t.Run("size-of-big.Int", func(t *testing.T) {
		i := new(big.Int)
		i.SetString("644", 8) // octal
		fmt.Printf("size of %TransactionData:, %d\n", i, unsafe.Sizeof(i))
		fmt.Printf("align of %TransactionData:, %d\n", i, unsafe.Alignof(i))
	})
	t.Run("size-of-int64", func(t *testing.T) {
		var i int64
		i = 55465474
		fmt.Printf("size of %TransactionData:, %d\n", i, unsafe.Sizeof(i))
		fmt.Printf("align of %TransactionData:, %d\n", i, unsafe.Alignof(i))
	})
	t.Run("size-of-*int64", func(t *testing.T) {
		var i *int64
		fmt.Printf("size of %TransactionData:, %d\n", i, unsafe.Sizeof(i))
		fmt.Printf("align of %TransactionData:, %d\n", i, unsafe.Alignof(i))
	})
	t.Run("size-of-pointer", func(t *testing.T) {
		var i uintptr
		fmt.Printf("size of %TransactionData:, %d\n", i, unsafe.Sizeof(i))
		fmt.Printf("align of %TransactionData:, %d\n", i, unsafe.Alignof(i))
	})
	t.Run("size-of-pointer", func(t *testing.T) {
		const ptrSize = 32 + int(^uintptr(0)>>63<<5)
		fmt.Printf("pointer size is %d:, %d\n", ptrSize, unsafe.Sizeof(ptrSize))
	})
}

func TestParseInt(t *testing.T) {
	i, err := ParseInt("0x143")
	assert.Nil(t, err)
	assert.Equal(t, 323, i)

	i, err = ParseInt("143")
	assert.Nil(t, err)
	assert.Equal(t, 323, i)

	i, err = ParseInt("0xaaa")
	assert.Nil(t, err)
	assert.Equal(t, 2730, i)

	i, err = ParseInt("1*29")
	assert.NotNil(t, err)
	assert.Equal(t, 0, i)
}

func TestParseBigInt(t *testing.T) {
	i, raw, err := ParseBigInt("0xabc")
	assert.Nil(t, err)
	assert.Equal(t, raw, "0xabc")
	assert.Equal(t, int64(2748), i.Int64())

	i, raw, err = ParseBigInt("$%1")
	assert.Equal(t, raw, "$%1")
	assert.NotNil(t, err)
}

func TestIntToHex(t *testing.T) {
	assert.Equal(t, "0xde0b6b3a7640000", IntToHex(1000000000000000000))
	assert.Equal(t, "0x6f", IntToHex(111))
}

func TestBigToHex(t *testing.T) {
	i1, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	assert.Equal(t, "0xde0b6b3a7640000", BigToHex(*i1))

	i2, _ := big.NewInt(0).SetString("100000000000000000000", 10)
	assert.Equal(t, "0x56bc75e2d63100000", BigToHex(*i2))

	i3, _ := big.NewInt(0).SetString("0", 10)
	assert.Equal(t, "0x0", BigToHex(*i3))
}
