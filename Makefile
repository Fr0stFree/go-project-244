GOLANGCI_LINT_VERSION := v2.12.2
GOLANGCI_LINT := $(shell go env GOPATH)/bin/golangci-lint
BINARY_PATH := bin/gendiff
COVERAGE_PROFILE := coverage.out
COVERAGE_BADGE := coverage-badge.json

.PHONY: build lint lint-fix test test-coverage install-lint

install-lint:
	@if test -x $(GOLANGCI_LINT); then \
		echo "Using existing golangci-lint: $(GOLANGCI_LINT)"; \
	else \
		echo "golangci-lint not found"; \
		echo "Installing version $(GOLANGCI_LINT_VERSION)..."; \
		go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION); \
		echo "Installation completed"; \
	fi

lint: install-lint
	@echo "Running linters..."
	@test -z "$$(gofmt -l .)" || (echo "Code is not formatted"; exit 1)
	@$(GOLANGCI_LINT) run
	@echo "Linting completed."

lint-fix: install-lint
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
	@coverage=$$(go tool cover -func=$(COVERAGE_PROFILE) | awk '/^total:/ {print substr($$3, 1, length($$3)-1)}'); \
	color=$$(awk -v coverage="$$coverage" 'BEGIN { if (coverage >= 80) print "brightgreen"; else if (coverage >= 60) print "yellow"; else print "red" }'); \
	printf '{\n  "schemaVersion": 1,\n  "label": "coverage",\n  "message": "%s%%",\n  "color": "%s"\n}\n' "$$coverage" "$$color" > $(COVERAGE_BADGE)
