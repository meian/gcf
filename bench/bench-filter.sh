#!/bin/bash

DIR=$(cd $(dirname $BASH_SOURCE); pwd)

gotest() {
    cpu=${3:-$(nproc)}
    go test -bench "$1" -benchmem -benchtime 5s -cpu $cpu . \
        | tee "$2"
}

(
    cd $DIR
    gotest BenchmarkFilter_Volumes filter-volumes.txt
    gotest BenchmarkFilter_Compare filter-compare-8.txt
    gotest BenchmarkFilter_Compare filter-compare-1.txt 1
)
