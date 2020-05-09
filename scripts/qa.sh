#!/bin/bash

#
# Copyright gotools
# SPDX-License-Identifier: GNU GPL v3
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Checking code quality with linters..."

# go fmt
./scripts/fmt.sh
# go fmt simplify
./scripts/fmt_and_simplify.sh
# go imports
./scripts/goimports.sh

# go vet
./scripts/govet.sh -tags "dev oss"
./scripts/govet.sh -tags "pre oss"
./scripts/govet.sh -tags "prod oss"

./scripts/govet.sh -tags "dev pro"
./scripts/govet.sh -tags "pre pro"
./scripts/govet.sh -tags "prod pro"


# add license header to files
# ./scripts/license_header.sh

echo "qa scripts finished"
