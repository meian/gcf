#!/bin/bash

DIR=$(cd $(dirname $BASH_SOURCE); pwd)

gotest() {
    cpu=${3:-$(nproc)}
    go test -bench "$1" -benchmem -benchtime 2s -cpu $cpu . \
        | tee "$2"
}

(
    cd $DIR
    gotest BenchmarkRange range.txt
)
