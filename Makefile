.DEFAULT_GOAL := build

fmt:
	@go fmt ./...

lint: fmt
	@go lint ./...

vet: fmt
	@go vet ./...

build: vet
	@go build .

run: fmt
	@go run ./...

test: 
	@go test -v -cover ./...

.PHONY: fmt lint vet build run  test


