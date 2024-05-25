# Get the operating system 

SHELL := /bin/bash    # by default make uses /bin/sh
BIN_DIR := $(CURDIR)/bin	
PROJECTNAME=$(shell basename "$(PWD)")
PWD_PROJECT=$(shell pwd)
LDFLAGS="-w -X 'main.buildTime=$(shell date)' -X 'main.lastCommit=$(shell git rev-parse HEAD)' -X 'main.semanticVersion=$(shell git describe --tags --dirty=-dev)'"
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

export GOBIN := $(BIN_DIR)


## help: info about what can be done
help: Makefile
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help


build:
	@echo "--> Building gofetch binary for $(GOOS):$(GOARCH)"
	  env go build -ldflags $(LDFLAGS) -o gofetch ./cmd/;\
	@echo "--> gofetch for $(GOOS):$(GOARCH) built at $(PWD_PROJECT)"

.PHONY: build

## run linter
linter: setup-linter
	@echo "Checking code"
	$(BIN_DIR)/golangci-lint run $(CURDIR)/...

.PHONY: linter

setup-linter:
	@echo "Installing golanglint dependency"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

test: linter
	@echo "Running all tests"
	go test -v ./...

.PHONY: setup-linter
