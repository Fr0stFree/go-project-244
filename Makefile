GOLANGCI_LINT_VERSION := v2.12.2
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint
BINARY_PATH := bin/gendiff

.PHONY: build lint lint-fix test install-lint

install-lint:
	@if test -x $(GOLANGCI_LINT); then \
		echo "Using existing golangci-lint: $(GOLANGCI_LINT)"; \
	else \
		echo "golangci-lint not found"; \
		echo "Installing version $(GOLANGCI_LINT_VERSION)..."; \
		go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION); \
		echo "Installation completed"; \
	fi

lint:
	@echo "Running linters..."
	@test -z "$$(gofmt -l .)" || (echo "Code is not formatted"; exit 1)
	@$(GOLANGCI_LINT) run
	@echo "Linting completed."

lint-fix:
	@echo "Running linters with auto-fix..."
	@gofmt -w .
	@$(GOLANGCI_LINT) run --fix
	@echo "Linting and auto-fix completed."

build:
	@echo "Building $(BINARY_PATH)..."
	@go build -o $(BINARY_PATH) ./cmd/gendiff
	@echo "Binary created at: $$(pwd)/$(BINARY_PATH)"

test:
	@go test ./... -v
