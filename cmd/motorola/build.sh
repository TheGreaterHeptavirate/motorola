#!/bin/bash
#
# Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
# All Rights Reserved
#
# All copies of this software (if not stated otherwise) are dedicated
# ONLY to personal, non-commercial use.
#

#
echo "building for linux"
go build -o motorola.bin 

echo "build for windows"
GOOS=windows GOARCH=amd64 \
CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
go build \
-ldflags="-s -w -extldflags -static" \
-o motorola.exe .
