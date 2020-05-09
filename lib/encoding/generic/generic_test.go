// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package generic

import (
	"bytes"
	"testing"
)

func TestNbaseEncode(t *testing.T) {
	type args struct {
		nb   uint64
		buf  *bytes.Buffer
		base string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NbaseEncode(tt.args.nb, tt.args.buf, tt.args.base)
		})
	}
}

func TestNbaseDecode(t *testing.T) {
	type args struct {
		enc  string
		base string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbaseDecode(tt.args.enc, tt.args.base); got != tt.want {
				t.Errorf("NbaseDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
