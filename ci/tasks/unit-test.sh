#!/usr/bin/env sh

mkdir -p ${GOPATH}/src/github.com/cappyzawa/pipelines
cp -r pr/* ${GOPATH}/src/github.com/cappyzawa/pipelines
cd ${GOPATH}/src/github.com/cappyzawa/pipelines

go test ./...
