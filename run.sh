#!/bin/bash
#
#    Starts debug version of the server
#
# 
#


if [[ $(env | grep GOPATH | wc -l) = *0* ]]; then
    export GOPATH=`pwd`
fi


go run src/app/app.go