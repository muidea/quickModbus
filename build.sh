#!/bin/sh

#export GOOS=windows
export GOOS=linux
export GOARCH=amd64

go build github.com/muidea/quickModbus/cmd/quickModbus
