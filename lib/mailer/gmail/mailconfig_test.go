// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package gmail

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultMailServerInstantiation(t *testing.T) {
	t.Run("init-instance", func(t *testing.T) {
		i := GetGmailServerConfigInstanceInit()
		assert.NotNil(t, i, "could not be instantiated")
	})
	t.Run("singleton-unsafe-instance", func(t *testing.T) {
		i := GetGmailServerConfigInstance("", "", "")
		assert.NotNil(t, i, "could not be instantiated")
	})
	t.Run("singleton-safe-instance", func(t *testing.T) {
		i := GetGmailServerConfigInstanceThreadSafe("", "", "")
		assert.NotNil(t, i, "could not be instantiated")
	})
}
