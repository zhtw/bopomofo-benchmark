#!/bin/bash
pushd `dirname $0` > /dev/null
SCRIPTPATH=`pwd`
popd > /dev/null

GOPATH=$SCRIPTPATH
PROJECT=github.com/zhtw/bopomofo-benchmark
go fmt $PROJECT
go test $PROJECT
go build $PROJECT
