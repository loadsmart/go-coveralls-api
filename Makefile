GOBIN?=$(shell go env GOPATH)/bin

.PHONY: install
install: ## Install tools used by the project
	fgrep '_' tools.go | cut -f2 -d' ' | xargs go install

.PHONY: test
test: ## Run tests
	go test -v ./...
