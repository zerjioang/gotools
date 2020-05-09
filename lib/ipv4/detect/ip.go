// Copyright https://github.com/zerjioang/gotools
// SPDX-License-Identifier: GPL-3.0-only

package detect

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

// IP address lengths (bytes).
const (
	iPv4len = 4
	// Bigger than we need, not too big to worry about overflow
	big = 0xFFFFFF
)

var (
	ipRegex, _ = regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
)

// detect if given string is an ipv4 using net package
func IsIpv4Net(host string) bool {
	return net.ParseIP(host) != nil
}

func IsIpv4Simple(host string) bool {
	parts := strings.Split(host, ".")
	if len(parts) < 4 {
		return false
	}
	for _, x := range parts {
		if i, err := integerAtoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func IsIpv4(s string) bool {
	if len(s) == 0 {
		// Missing octets.
		return false
	}
	for i := 0; i < iPv4len; i++ {
		if i > 0 {
			if s[0] != '.' {
				return false
			}
			s = s[1:]
		}
		var n int
		var i int
		var ok bool
		for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
			n = n*10 + int(s[i]-'0')
			if n >= big {
				n = big
				ok = false
			}
		}
		if i == 0 {
			n = 0
			i = 0
			ok = false
		} else {
			ok = true
		}
		if !ok || n > 0xFF {
			return false
		}
		s = s[i:]
	}
	return len(s) == 0
}

func IsIpv4Regex(ipAddress string) bool {
	ipAddress = strings.Trim(ipAddress, " ")
	return ipRegex.MatchString(ipAddress)
}

func atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func integerAtoi(str string) (int, error) {
	res := 0 // Initialize result

	// Iterate through all characters of input string and
	// update result
	for i := 0; i < len(str); i++ {
		res = res*10 + int(str[i]-'0')
	}

	// return result.
	return res, nil
}
