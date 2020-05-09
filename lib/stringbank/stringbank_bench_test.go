// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package stringbank

import (
	"strconv"
	"testing"
)

func BenchmarkStringbank(b *testing.B) {
	s := make([]string, b.N)
	for i := range s {
		s[i] = strconv.Itoa(i)
	}

	index := make([]int, b.N)

	b.ReportAllocs()
	b.ResetTimer()
	sb := Stringbank{}
	for i, v := range s {
		index[i] = sb.Save(v)
	}

	var out string
	for _, i := range index {
		out = sb.Get(i)
	}
	if out != s[b.N-1] {
		b.Fatalf("final string should be %s, is %s", s[b.N-1], out)
	}
}
