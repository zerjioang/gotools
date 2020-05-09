// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package ip

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpToUint32(t *testing.T) {

	t.Run("ip-to-int", func(t *testing.T) {
		t.Run("10.41.132.6", func(t *testing.T) {
			intVal := Ip2int("10.41.132.6")
			t.Log("str ip:", strconv.Itoa(int(intVal)))
			assert.Equal(t, int(intVal), 170492934, "failed to convert ip to numeric")
		})
		t.Run("218.255.173.218", func(t *testing.T) {
			intVal := Ip2int("218.255.173.218")
			t.Log("str ip:", strconv.Itoa(int(intVal)))
			assert.Equal(t, int(intVal), 3674189274, "failed to convert ip to numeric")
		})
	})
	t.Run("convert-uint32-low", func(t *testing.T) {
		t.Run("convert-uint32-low-10.41.132.6", func(t *testing.T) {
			intVal := Ip2intLow("10.41.132.6")
			t.Log("str ip:", strconv.Itoa(int(intVal)))
			assert.Equal(t, int(intVal), 170492934, "failed to convert ip to numeric")
		})
		t.Run("convert-uint32-low-218.255.173.218", func(t *testing.T) {
			intVal := Ip2intLow("218.255.173.218")
			t.Log("str ip:", strconv.Itoa(int(intVal)))
			assert.Equal(t, int(intVal), 3674189274, "failed to convert ip to numeric")
		})
	})
	t.Run("assembly_amd64", func(t *testing.T) {
		v := IpToInt2([]byte("10.41.132.6"))
		t.Log("str ip:", strconv.Itoa(int(v)))
		assert.Equal(t, int(v), 170492934, "failed to convert ip to numeric")
	})
	t.Run("isIpv4-valid", func(t *testing.T) {
		v := IsIpv4("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-invalid", func(t *testing.T) {
		v := IsIpv4("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("isIpv4-regex-valid", func(t *testing.T) {
		v := IsIpv4Regex("10.41.132.6")
		assert.True(t, v, "failed to detect a valid ipv4")
	})
	t.Run("isIpv4-regex-invalid", func(t *testing.T) {
		v := IsIpv4Regex("10.280.132.6")
		assert.False(t, v, "failed to detect a invalid ipv4")
	})
	t.Run("custom-atoi", func(t *testing.T) {
		v, err := integerAtoi("89789")
		assert.Nil(t, err)
		assert.Equal(t, v, 89789, "failed to detect a invalid ipv4")
	})
}
