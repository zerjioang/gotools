// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package bip39

import (
	"encoding/hex"
	"testing"
)

func TestE2ENewMnemonic(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		// the entropy can be any byte slice, generated how pleased,
		// as long its bit size is a multiple of 32 and is within
		// the inclusive range of {128,256}
		entropy, _ := hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")

		// generate a mnemomic
		/*
			128 bits -> 12 words
			160 bits -> 15 words
			192 bits -> 18 words
			224 bits -> 21 words
			256 bits -> 24 words
		*/
		mnemomic, serr := NewMnemonic(entropy)
		t.Log(serr)
		t.Log(mnemomic)
		// output:
		// all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform
	})
}
