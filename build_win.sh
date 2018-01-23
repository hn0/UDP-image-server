#!/bin/bash
#
#
#  Build example for windows platform
#
#

if [[ $(env | grep GOPATH | wc -l) = *0* ]]; then
    export GOPATH=`pwd`
fi

export GOOS=windows
export GOARCH=amd64

go build -v -x -o bin/server.exe src/app/app.go