// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// +build ignore

package integrity

import (
	"crypto/ecdsa"
	"reflect"
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		privateKey *ecdsa.PrivateKey
		publicKey  *ecdsa.PublicKey
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := encode(tt.args.privateKey, tt.args.publicKey)
			if got != tt.want {
				t.Errorf("encode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("encode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	type args struct {
		pemEncoded    []byte
		pemEncodedPub []byte
	}
	tests := []struct {
		name  string
		args  args
		want  *ecdsa.PrivateKey
		want1 *ecdsa.PublicKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decode(tt.args.pemEncoded, tt.args.pemEncodedPub)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("decode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
