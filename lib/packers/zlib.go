// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package packers

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

func ZlibUnpack(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	unpackedData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	closeErr := r.Close()
	if closeErr != nil {
		return nil, closeErr
	}
	return unpackedData, nil
}

func ZlibPack(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	closeErr := w.Close()
	if closeErr != nil {
		return nil, closeErr
	}
	return b.Bytes(), nil
}
