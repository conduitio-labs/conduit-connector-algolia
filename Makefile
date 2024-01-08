.PHONY: build test

VERSION=`git describe --tags --dirty --always`

build:
	go build -ldflags "-X 'github.com/conduitio-labs/conduit-connector-algolia.version=${VERSION}'" -o conduit-connector-algolia cmd/connector/main.go

test:
	go test $(GOTEST_FLAGS) -race ./...

lint:
	golangci-lint run -v

.PHONY: install-tools
install-tools:
	@echo Installing tools from tools.go
	@go list -e -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy
