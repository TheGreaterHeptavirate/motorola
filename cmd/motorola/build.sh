#!/bin/bash
#
echo "building for linux"
go build -o motorola.bin 

echo "build for windows"
GOOS=windows GOARCH=amd64 \
        CGO_ENABLED=1 \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
        go build \
        -ldflags="-s -w -linkmode external -extldflags -static" \
        -o motorola.exe .
