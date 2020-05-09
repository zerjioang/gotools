// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package solc

import "testing"

func Test_slurpFiles(t *testing.T) {
	type args struct {
		files []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slurpFiles(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("slurpFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("slurpFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
