#!/bin/bash

export GOPATH=$(pwd)
go get github.com/samuel/go-zookeeper/zk
go get github.com/gorilla/mux

go build -o instance-tracker *.go

sudo docker build .
