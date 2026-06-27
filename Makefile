GOLANGCI_LINT_VERSION := v2.12.2
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint
BINARY_PATH := bin/gendiff
COVERAGE_PROFILE := coverage.out

.PHONY: build lint lint-fix test test-coverage install-lint require-lint

install-lint:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

require-lint:
	test -x $(GOLANGCI_LINT) || (echo "golangci-lint not found. Run 'make install-lint' first."; exit 1)

lint: require-lint
	test -z "$$(gofmt -l .)" || (echo "Code is not formatted"; exit 1)
	$(GOLANGCI_LINT) run

lint-fix: require-lint
	gofmt -w .
	$(GOLANGCI_LINT) run --fix

build:
	go build -o $(BINARY_PATH) ./cmd/gendiff

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=$(COVERAGE_PROFILE)
	go tool cover -func=$(COVERAGE_PROFILE)