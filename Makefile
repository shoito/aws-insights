.PHONY: build generate clean test

export GO111MODULE=on

all: build

build: generate
	@go build -o bin/iaws -ldflags "-s -w" .

generate:
	@go generate
	@go mod tidy
	@go mod verify

test: generate
	@go test -v ./...

clean:
	@rm -rf bin/