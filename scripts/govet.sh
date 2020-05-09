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

log "Checking source code with Go vet"

#get all files excluding vendors
filelist=$(go list ./... | grep -vendor)
for file in ${filelist}
do
	log "static analysis of package $file"
	go vet "$@" ${file}
done

ok "Code checking done!"
