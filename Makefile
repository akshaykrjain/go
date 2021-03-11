SHELL := /bin/bash

build:
	go build -o hello -v ./... 

testlocal:
	go vet ./...
	go test ./... -v

