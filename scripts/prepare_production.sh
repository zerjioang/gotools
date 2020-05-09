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

echo "Preparing code for production..."

# remove logging lines with just text strings
coreFiles=$(find . -type f -name "*.go" | grep -vendor)
regex="(log|logger)\.(Debug|Info|Error|Warn|Critical){1}\(\"(\w|\s)*\"\)"
for file in ${coreFiles}
do
	echo "optimizing file: $file"
	sed -ri "s/$regex//g" ${file}
done

# after removing logs, we may need to remove some orphaned imports
./scripts/goimports.sh
./scripts/fmt.sh
./scripts/fmt_and_simplify.sh
# go vet
# ./scripts/govet.sh -tags "dev oss"
# ./scripts/govet.sh -tags "pre oss"
# ./scripts/govet.sh -tags "prod oss"

# ./scripts/govet.sh -tags "dev pro"
# ./scripts/govet.sh -tags "pre pro"
# ./scripts/govet.sh -tags "prod pro"