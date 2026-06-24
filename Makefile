GOLANGCI_LINT_VERSION := v2.12.2
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint
BINARY_PATH := bin/gendiff
COVERAGE_PROFILE := coverage.out

.PHONY: build lint lint-fix test test-coverage install-lint require-lint

install-lint:
	@if test -x $(GOLANGCI_LINT); then \
		echo "Golangci-lint is already installed at $(GOLANGCI_LINT)"; \
		$(GOLANGCI_LINT) --version; \
	else \
		echo "Installing version $(GOLANGCI_LINT_VERSION)..."; \
		go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION); \
		echo "Installation completed"; \
	fi

require-lint:
	@test -x $(GOLANGCI_LINT) || (echo "golangci-lint not found. Run 'make install-lint' first."; exit 1)

lint: require-lint
	@echo "Running linters..."
	@test -z "$$(gofmt -l .)" || (echo "Code is not formatted"; exit 1)
	@$(GOLANGCI_LINT) run
	@echo "Linting completed."

lint-fix: require-lint
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

test-coverage:
	@go test ./... -coverprofile=$(COVERAGE_PROFILE)
	@go tool cover -func=$(COVERAGE_PROFILE)
