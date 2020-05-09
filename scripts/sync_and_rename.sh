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

echo "Syncing go.ethereum with latest version of github.com"

ethPath="core/modules/ethfork"
packageName="gotools"

if [[ ! -f ${ethPath} ]] ; then
	echo "etherem secp256k1 C files missing...downloading..."
	if [[ ! -d ${ethPath} ]] ; then
		echo "ethereum c files were not previously downloaded"
		mkdir -p $ethPath
		cd $ethPath
		# files are not downloaded
		echo "cloning eth files..."
		git clone https://github.com/ethereum/go-ethereum .
		# remove vendor folder
		echo "removing vendor folder"
		rm -rf ./vendor

		echo "renaming package name"
		# renaming package names of the files to include them in the project
		find . -type f -exec sed -i "s/github\.com\/ethereum\/go-ethereum/github\.com\/zerjioang\/$packageName\/core\/modules\/ethfork/g" {} +

		echo "renaming internal package"
		# rename locked internal (to internals) package to be usable from external
		find . -type f -exec sed -i 's/\/internal\//\/internals\//g' {} +

		echo "renaming internal dirs"
		# rename all dirs named internal to internals
		# sudo apt-install rename
		rename "s/internal/internals/"  *
	else
		# files were already downloaded previously
		echo "previously downloaded ethereum c files found. skipping"
	fi
	echo "go.ethereum c files copied"
fi

echo "sync done"