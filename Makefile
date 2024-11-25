.PHONY: build test

VERSION=`git describe --tags --dirty --always`

.PHONY: build
build:
	go build -ldflags "-X 'github.com/conduitio-labs/conduit-connector-algolia.version=${VERSION}'" -o conduit-connector-algolia cmd/connector/main.go

.PHONY: test
test:
	go test $(GOTEST_FLAGS) -race ./...

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt
fmt:
	gofumpt -l -w .

.PHONY: install-tools
install-tools:
	@echo Installing tools from tools.go
	@go list -e -f '{{ join .Imports "\n" }}' tools.go | xargs -I % go list -f "%@{{.Module.Version}}" % | xargs -tI % go install %
	@go mod tidy
