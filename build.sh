#!/bin/sh
# Adapted from https://github.com/matrix-org/dendrite/blob/main/build.sh (Apache 2.0)

# Put installed packages into ./bin
export GOBIN=$PWD/`dirname $0`/bin

mkdir -p bin

CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -v -o "bin/" ./cmd/...
