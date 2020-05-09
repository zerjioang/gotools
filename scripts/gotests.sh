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

# load colored logs
source ./scripts/colors.sh

log "generating test functions base with gotests"
log "checking if gotests is installed in $GOPATH"
if [[ ! -f ${GOPATH}/bin/gotests ]]; then
	#statements
	log "gotests not found. Downloading via go get"
	go get -v github.com/cweill/gotests
	cd ${GOPATH}/src/github.com/cweill/gotests/gotests
	go build && go install
fi

if [[ ! -f ${GOPATH}/bin/gotests ]]; then
	fail "failed to install gotests in ${GOPATH}"
	return -1
fi

#get all files excluding vendors
filelist=$(find . -type f -name "*.go" | grep -vendor)
for file in ${filelist}
do
	log "generating gotests for file $file"
	${GOPATH}/bin/gotests -excl Benchmark.* -w ${file}
done

ok "(gotests) Code formatting done!"