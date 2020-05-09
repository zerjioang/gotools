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
start

function start(){
	echo "Ofuscating current code before compilation"
	echo "simple code ofuscation done"
}

function rename_path(){
	src=$1
	dst=$2
	echo "renaming path from $src $dst"
	# renaming package names of the files to include them in the project
	echo "replacing references with regex: s/$src/$dst/g"
	# find . -type f -exec sed -i 's/$src/$dst/g' {} +

	echo "renaming dirs..."
	# rename all dirs named internal to internals
	# sudo apt-install rename
	echo "renaming dirs with: s/$src/$dst/ *"
	# rename "s/$src/$dst/" *
}