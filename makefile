SHELL := /bin/bash

.PHONY: all build deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: deps build

${BUILD_DIR}:
	mkdir -p $(BUILD_DIR)

build: ${BUILD_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd

location: ## Start function location
	$(GOCMD) run ./cmd location $(arg)

message: ## Start function message
	$(GOCMD) run ./cmd message $(arg)

server: ## Start server
	$(GOCMD) run ./cmd server

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache
