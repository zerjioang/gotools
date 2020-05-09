// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

// +build !appengine

package log

import (
	"io"

	"github.com/mattn/go-colorable"
)

var (
	colorableOut = colorable.NewColorableStdout()
)

func output() io.Writer {
	return colorableOut
}
