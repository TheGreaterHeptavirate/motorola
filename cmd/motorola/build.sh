#!/bin/bash
#
echo "building for linux"
go build -o motorola.bin 

echo "build for windows"
LD_LIBRARY_PATH="/usr/x86_64-w64-mingw32/sys-root/mingw/lib" PKG_CONFIG_PATH="/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig/" GOOS=windows GOARCH=amd64 \
        CGO_ENABLED=1 \
        CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ \
        go build \
        -ldflags "-s -w -extldflags=-static" \
        -o motorola.exe .
