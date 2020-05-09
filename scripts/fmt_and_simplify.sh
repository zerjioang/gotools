#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Formatting source code with gofmt (and simplification)"

packageName="github.com/zerjioang/gotools"

#get all files excluding vendors
filelist=$(find ./ -name "*.go" | grep -vendor)
toreplace="./"
toreplaceBy="/"
for file in ${filelist}
do
	echo "Formatting file $file"
	gofmt -s -w ${file}
done

echo "(fmt and simplify) Code formatting done!"
