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

# global variables
GO_EXEC=go
BENCH_FILE=old.bench

function requirements(){
	log "checking if benchstat is installed"
	if ! [ -x "$(command -v benchstat)" ]; then
		#statements
		fail "benchstat tool not found"
		log "installing benchstat..."
		go get golang.org/x/perf/cmd/benchstat
	else
		ok "benchstat found"
	fi
}

function benchmark(){
	INITIAL_DIR=$(pwd)
	TEST_FILES=$1
	if [[ -z ${TEST_FILES} ]]; then
		fail "no go package specified"
		fail "usage: benchtest /absolute/path/to/package"
		cd ${INITIAL_DIR}
		return -1
	fi

	log "moving to test dir: ${TEST_FILES}"
	cd ${TEST_FILES}

	if [[ -f new.bench ]]; then
		log "benchmark cleanup"
		# clean up
		# remove old.bench
		# rename new.bench to old.bench
		rm old.bench
		mv new.bench old.bench
	fi

	if [[ -f old.bench ]]; then
		#statements
		log "old benchmarking results found"
		log "configuring benchmark process for comparison..."
		BENCH_FILE=new.bench
	else
		BENCH_FILE=old.bench
	fi

	log "executing benchmarking"
	${GO_EXEC} test -tags "dev oss" -test.bench=. -cpu=1,2,4 -benchtime=2s -test.count=1 | tee ${BENCH_FILE}

	if [[ -f new.bench ]]; then
		benchstat old.bench new.bench
	else
		fail "no previous results were found"
	fi

	ok "benchstat finished"
	ok "current results saved as ${BENCH_FILE}"
	cd ${INITIAL_DIR}
}

requirements
benchmark $@