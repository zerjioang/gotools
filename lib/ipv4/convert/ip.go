// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package convert

import (
	"encoding/binary"
	"net"
	"regexp"
	"strconv"
)

const (
	asciiDot  uint8 = 46
	asciiZero uint8 = 48
)

var (
	ipRegex, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
)

// converts an IP address to uint32 value
func Ip2int(ip string) uint32 {
	rawBytes := net.ParseIP(ip).To4()
	return binary.BigEndian.Uint32(rawBytes)
}

func Ip2intLow(ip string) uint32 {
	var octets [4][4]byte
	var currentOctect uint8 = 0
	var currentOctectPos uint8 = 0
	s := len(ip)
	_ = ip[s-1]
	for i := 0; i < s; i++ {
		ipVal := ip[i]
		if ipVal == asciiDot {
			octets[currentOctect][3] = currentOctectPos
			//move to the next octect
			currentOctect++
			currentOctectPos = 0
		} else {
			// assign value to current octect
			octets[currentOctect][currentOctectPos] = ipVal
			currentOctectPos++
		}
	}
	// save last octet information
	octets[currentOctect][3] = currentOctectPos

	// convert octects string bytes to decimal
	var octectsDecimal [4]byte
	var l uint8 = 4
	var i uint8 = 0
	for i < l {
		//process each octect data
		// convert octects to uint32
		// octets[0]*256³ + octets[1]*256² + octets[2]*256¹ + octets[1]*256⁰
		_ = octets[i][2]
		switch octets[i][3] {
		case 0:
			octectsDecimal[i] = 0
		case 1:
			octectsDecimal[i] = octets[i][0] - asciiZero
		case 2:
			octectsDecimal[i] = (octets[i][0]-asciiZero)*10 + (octets[i][1] - asciiZero)
		case 3:
			octectsDecimal[i] = (octets[i][0]-asciiZero)*100 + (octets[i][1]-asciiZero)*10 + (octets[i][2] - asciiZero)
		}
		i++
	}
	var intIp uint32
	// intIp = uint32(octectsDecimal[0])*16777216 + uint32(octectsDecimal[1])*65536 + uint32(octectsDecimal[2])*256 + uint32(octectsDecimal[3])
	intIp = uint32(octectsDecimal[3]) | uint32(octectsDecimal[2])<<8 | uint32(octectsDecimal[1])<<16 | uint32(octectsDecimal[0])<<24
	return intIp
}

// converts an uint32 to IP address
func Int2ip(ipInt int64) string {
	// need to do two bit shifting and "0xff" masking

	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt(ipInt&0xff, 10)

	return b0 + "." + b1 + "." + b2 + "." + b3
}
