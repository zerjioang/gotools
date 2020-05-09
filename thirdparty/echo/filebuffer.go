// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

package echo

import (
	"bytes"
	"errors"
	"time"
)

var (
	errNotSupported = errors.New("unsupported seek method")
)

const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)

type FileBuffer struct {
	//metadata values for http
	name string
	time time.Time
	// content data
	Buffer bytes.Buffer
	Index  int64
}

func NewFileBuffer() FileBuffer {
	return FileBuffer{}
}

func (fbuffer *FileBuffer) Bytes() []byte {
	return fbuffer.Buffer.Bytes()
}

func (fbuffer *FileBuffer) Read(p []byte) (int, error) {
	n, err := bytes.NewBuffer(fbuffer.Buffer.Bytes()[fbuffer.Index:]).Read(p)

	if err == nil {
		if fbuffer.Index+int64(len(p)) < int64(fbuffer.Buffer.Len()) {
			fbuffer.Index += int64(len(p))
		} else {
			fbuffer.Index = int64(fbuffer.Buffer.Len())
		}
	}

	return n, err
}

func (fbuffer *FileBuffer) Write(p []byte) (int, error) {
	n, err := fbuffer.Buffer.Write(p)

	if err == nil {
		fbuffer.Index = int64(fbuffer.Buffer.Len())
	}

	return n, err
}

func (fbuffer *FileBuffer) Seek(offset int64, whence int) (int64, error) {
	var err error
	var Index int64 = 0

	switch whence {
	case SeekStart:
		if offset >= int64(fbuffer.Buffer.Len()) || offset < 0 {
			err = errors.New("invalid offset")
		} else {
			fbuffer.Index = offset
			Index = 0
		}
	case SeekCurrent:
		if offset >= int64(fbuffer.Buffer.Len()) || offset < 0 {
			err = errors.New("invalid offset")
		} else {
			fbuffer.Index = offset
			Index = offset
		}
	case SeekEnd:
		if offset >= int64(fbuffer.Buffer.Len()) || offset < 0 {
			err = errors.New("invalid offset")
		} else {
			fbuffer.Index = offset
			Index = int64(fbuffer.Buffer.Len())
		}
	default:
		err = errNotSupported
	}

	return Index, err
}
