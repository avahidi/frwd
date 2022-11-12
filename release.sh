#!/bin/bash

rm -rf build
mkdir -p build

export CMD="go build -ldflags='-s -w'"

for GOOS in linux darwin windows
do
    for GOARCH in 386 amd64 arm arm64
    do
        GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o build/frwd_${GOOS}_${GOARCH}
    done
done

ls -l