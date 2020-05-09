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

echo "Downloading & Installing required dependencies"
# for gogo/protobuf
# Install the protoc-gen-gofast binary
echo "downloading from https://github.com/gogo/protobuf"
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto
go get github.com/gogo/protobuf/protoc-gen-gogofaster
go get github.com/gogo/protobuf/protoc-gen-gofast
echo "use it to generate faster marshaling and unmarshaling go code for your protocol buffers."
echo "protoc --gofast_out=. myproto.proto"

echo "installing messagepack"
go get -u github.com/vmihailenco/msgpack