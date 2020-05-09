// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package stack

import "github.com/zerjioang/gotools/util/str"

type Error struct {
	cause []byte
}

var (
	nilErr = New("")
)

func Nil() Error {
	return nilErr
}

func Ret(e error) Error {
	if e == nil {
		return nilErr
	}
	return Error{str.UnsafeBytes(e.Error())}
}

func New(msg string) Error {
	return Error{str.UnsafeBytes(msg)}
}

func (stack Error) Error() string {
	return str.UnsafeString(stack.cause)
}

func (stack Error) Bytes() []byte {
	return stack.cause
}

func (stack Error) Occur() bool {
	return len(stack.cause) > 0
}
func (stack Error) None() bool {
	return len(stack.cause) == 0
}
