// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package randomuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/util/str"
)

func TestGenerateUUID(t *testing.T) {
	t.Run("uuid-entropy", func(t *testing.T) {
		value := GenerateUUIDFromEntropy()
		t.Log("uuid value:", value)
		// example: 07dc4b26-caef-43c9-b068-54fff6222653
		if value == "" || len(value) != 36 {
			t.Error("failed to create uuid from entropy")
		}
	})

	t.Run("id-entropy", func(t *testing.T) {
		value := GenerateIDString()
		t.Log("uuid value:", value.String())
	})

	t.Run("random-str-charset", func(t *testing.T) {
		idstr := RandomStr(32)
		assert.NotNil(t, idstr)
		assert.Equal(t, len(idstr), 32)
		t.Log(str.UnsafeString(idstr))
	})
}
