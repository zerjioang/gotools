// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package integrity

import (
	"crypto/ecdsa"
	"math/big"
	"reflect"
	"testing"
)

func Test_newBig(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBig(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignMsgWithIntegrity(t *testing.T) {
	type args struct {
		message string
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
			got, got1 := SignMsgWithIntegrity(tt.args.message)
			if got != tt.want {
				t.Errorf("SignMsgWithIntegrity() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SignMsgWithIntegrity() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_ecdsaSign(t *testing.T) {
	type args struct {
		message []byte
		priv    *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		wantR   *big.Int
		wantS   *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotS, err := ecdsaSign(tt.args.message, tt.args.priv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ecdsaSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("ecdsaSign() gotR = %v, want %v", gotR, tt.wantR)
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("ecdsaSign() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_ecdsaVerify(t *testing.T) {
	type args struct {
		hash []byte
		r    *big.Int
		s    *big.Int
		pub  *ecdsa.PublicKey
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ecdsaVerify(tt.args.hash, tt.args.r, tt.args.s, tt.args.pub); got != tt.want {
				t.Errorf("ecdsaVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointsToDER(t *testing.T) {
	type args struct {
		r *big.Int
		s *big.Int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PointsToDER(tt.args.r, tt.args.s); got != tt.want {
				t.Errorf("PointsToDER() = %v, want %v", got, tt.want)
			}
		})
	}
}
