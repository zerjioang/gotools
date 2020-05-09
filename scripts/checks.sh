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

project="github.com/zerjioang/gotools"

function install(){
	name=$1
	package=$2
	log "checking if $name is installed in $GOPATH"
	if [[ ! -f ${GOPATH}/bin/$name ]]; then
		#statements
		log "$name not found. Downloading via go get"
		go get -u $package
	fi
}

function gotools(){
	log "Transforms an unkeyed struct literal into a keyed one."
	# keyify github.com/zerjioang/gotools

	log "Go static analysis, detecting bugs, performance issues, and much more."
	# staticcheck -tags dev oss $(go list ./...)

	log "Reorders struct fields to minimize the amount of padding."
	# structlayout-optimize github.com/zerjioang/gotools
}

function run(){
	log "golinting code.."

	# A set of utilities for checking Go sources.
	install "aligncheck" "gitlab.com/opennota/check/cmd/aligncheck"
	install "structcheck" "gitlab.com/opennota/check/cmd/structcheck"
	install "varcheck" "gitlab.com/opennota/check/cmd/varcheck"
	# tool to detect Go structs that would take less memory if their fields were sorted.
	install "maligned" "github.com/mdempsky/maligned"
	# prealloc is a Go static analysis tool to find slice declarations that could potentially be preallocated.
	install "prealloc" "github.com/alexkohler/prealloc"

	# keyify 	Transforms an unkeyed struct literal into a keyed one.
	# rdeps 	Find all reverse dependencies of a set of packages
	# staticcheck 	Go static analysis, detecting bugs, performance issues, and much more.
	# structlayout 	Displays the layout (field sizes and padding) of structs.
	# structlayout-optimize 	Reorders struct fields to minimize the amount of padding.
	# structlayout-pretty 	Formats the output of structlayout with ASCII art.
	install "go-tools" "honnef.co/go/tools/cmd/..."

	# install golant ci
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin latest

	log "aligncheck of ${project}"
	${GOPATH}/bin/aligncheck ${project}

	#get all files excluding vendors
	filelist=$(find . -type f -name "*.go" | grep -vendor)
	for file in ${filelist}
	do
		log "checking file $file"
		#${GOPATH}/bin/goimports -v -w ${file}
	done

	ok "Code checks done!"
}

gotools
run