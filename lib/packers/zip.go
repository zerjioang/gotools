// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package packers

import (
	"archive/zip"
	"bytes"
	"io/ioutil"

	"github.com/zerjioang/gotools/lib/logger"
)

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(data []byte) ([]string, error) {

	var filenames []string

	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return filenames, err
	}
	for _, f := range r.File {
		// read file name
		filename := f.Name
		// read file contents
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		unpackedData, err := ioutil.ReadAll(rc)
		if err != nil {
			return filenames, err
		}
		// close file stream
		_ = rc.Close()
		logger.Debug("zipped file found: ", filename)
		logger.Debug("zipped file data bytes: ", len(unpackedData))
	}
	return filenames, nil
}
