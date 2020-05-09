// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package gmail

import (
	"testing"
)

func BenchmarkGetMailServerConfigInstance(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetGmailServerConfigInstance("", "", "")
	}
}

func BenchmarkGetMailServerConfigInstanceThreadSafe(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetGmailServerConfigInstanceThreadSafe("", "", "")
	}
}

func BenchmarkGetMailServerConfigInstanceInit(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetGmailServerConfigInstanceInit()
	}
}
