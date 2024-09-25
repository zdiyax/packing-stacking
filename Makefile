#include .env

GOPATH := $(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install go-micro.dev/v4/cmd/protoc-gen-micro@v4

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go-out=:. proto/auth.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy
	@go mod vendor

.PHONY: build
build:
	@go mod build

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: lint
lint:
	$(V)golangci-lint run --config .golangci.yml

.PHONY: swag
swag:
	@swag init