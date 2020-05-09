// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3
package bip32

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_hashSha256(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashSha256(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashSha256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashSha256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashDoubleSha256(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashDoubleSha256(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashDoubleSha256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashDoubleSha256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashRipeMD160(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashRipeMD160(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashRipeMD160() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashRipeMD160() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hash160(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hash160(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("hash160() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hash160() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checksum(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("checksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addChecksumToBytes(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addChecksumToBytes(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("addChecksumToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addChecksumToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base58Encode(t *testing.T) {
	type args struct {
		data []byte
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
			if got := base58Encode(tt.args.data); got != tt.want {
				t.Errorf("base58Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base58Decode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base58Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("base58Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base58Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_publicKeyForPrivateKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := publicKeyForPrivateKey(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("publicKeyForPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addPublicKeys(t *testing.T) {
	type args struct {
		key1 []byte
		key2 []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addPublicKeys(tt.args.key1, tt.args.key2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addPublicKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addPrivateKeys(t *testing.T) {
	type args struct {
		key1 []byte
		key2 []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addPrivateKeys(tt.args.key1, tt.args.key2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addPrivateKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressPublicKey(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressPublicKey(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compressPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandPublicKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := expandPublicKey(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandPublicKey() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("expandPublicKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_validatePrivateKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validatePrivateKey(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("validatePrivateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateChildPublicKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChildPublicKey(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("validateChildPublicKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_uint32Bytes(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uint32Bytes(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uint32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
