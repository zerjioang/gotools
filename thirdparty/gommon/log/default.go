// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

// +build appengine

package log

import (
	"io"
	"os"
)

func output() io.Writer {
	return os.Stdout
}
