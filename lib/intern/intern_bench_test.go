// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package intern_test

import (
	"strconv"
	"testing"

	"github.com/zerjioang/gotools/lib/intern"
)

func BenchmarkIntern(b *testing.B) {
	s := make([]string, b.N)
	for i := range s {
		s[i] = strconv.Itoa(i)
	}

	i := intern.New(16)

	b.ReportAllocs()
	b.ResetTimer()

	var dedupe string
	for _, v := range s {
		dedupe = i.Deduplicate(v)
	}

	if dedupe != strconv.Itoa(b.N-1) {
		b.Errorf("last dedupe not as expected. Have %s expected %d", dedupe, b.N-1)
	}
}

func BenchmarkInternBasic(b *testing.B) {
	s := make([]string, b.N)
	for i := range s {
		s[i] = strconv.Itoa(i)
	}

	i := make(map[string]string)

	b.ReportAllocs()
	b.ResetTimer()

	var dedupe string
	var ok bool
	for _, v := range s {
		dedupe, ok = i[v]
		if !ok {
			i[v] = v
			dedupe = v
		}
	}

	if dedupe != strconv.Itoa(b.N-1) {
		b.Errorf("last dedupe not as expected. Have %s expected %d", dedupe, b.N-1)
	}
}
