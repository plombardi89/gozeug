SHELL := /usr/bin/env bash

GIT_COMMIT=$(shell git rev-parse --verify HEAD)

GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GOBUILD = go build -o bin/$(BINARY_BASENAME)-$(GOOS)-$(GOARCH)

all: clean fmt test

clean:
	rm -rf bin

fmt:
	go fmt ./...

test:
	go test -v ./...
