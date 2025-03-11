.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'




.PHONY: static-check
static-check: ## Run static code checks
	golangci-lint run -c .config/golangci.yaml --timeout 5m

.PHONY: test
test: static-check test-unit ## Run all tests
	go tool cover -func cover.out | grep total:
	go tool cover -html=cover.out -o cover.html; xdg-open cover.html
	go-cover-treemap -coverprofile cover.out > cover.svg; firefox cover.svg #xdg-open cover.svg

.PHONY: test-unit
test-unit:
	go test -race ./... -coverprofile cover.out




.PHONY: upgrade
upgrade:
	go get -t -u ./...
	go mod tidy

.PHONY: install-tools
install-tools: ## Initialise this machine with development dependencies
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v1.60.3
	go install github.com/nikolaydubina/go-cover-treemap@latest

	@# enable git hooks
	git config --global core.hooksPath .config/githooks