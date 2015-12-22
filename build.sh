#!/bin/bash


PWD=`pwd`

export GOPATH=$PWD

go build goweather.go

echo "Done"
