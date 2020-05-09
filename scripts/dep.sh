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

if ! [ -x "$(command -v dep)" ]; then
  echo 'error: dep is not installed.'
  echo 'installing dep...'
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi

if [ -z "$GOBIN" ]; then
	echo "\$GOBIN is empty"
	export GOBIN=${GOPATH}/bin
else
	echo "\$GOBIN is NOT empty"
fi

echo "Downloading dependencies using go dep"
$GOBIN/dep ensure -v
echo "all dependencies downloaded"