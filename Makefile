.PHONY: build test lint run clean help

build: ## Build the binary
	go build -o ollama-proxy

test: ## Run tests
	go test -v ./...

lint: ## Run linter
	golangci-lint run

run: ## Run the application (requires OP_AUTH_TOKEN env var)
	go run .

clean: ## Remove build artifacts
	rm -f ollama-proxy

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
