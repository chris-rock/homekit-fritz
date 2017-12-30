#!/bin/bash

build/linux-arm:
	@GOOS=linux GOARCH=arm go build -o dist/linux-arm/hkfritz main.go

build/linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/hkfritz main.go 

build/macos-amd64:
	@GOOS=darwin GOARCH=amd64 go build -o dist/macos-amd64/hkfritz main.go 

run:
	@go run main.go serve

test/unit	:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

test/vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')

test/fmt:
	@go fmt $(shell go list ./... | grep -v '/vendor/')

test/lint:
	@for package in $(shell go list ./... | grep -v '/vendor/' | grep -v '/api' | grep -v '/server/internal'); do \
      golint -set_exit_status $$package $$i || exit 1; \
	done

install/golint:
	@go get -u github.com/golang/lint/golint

install/goreleaser:
	@go get github.com/goreleaser/goreleaser
  