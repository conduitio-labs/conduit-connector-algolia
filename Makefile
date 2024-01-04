.PHONY: build test

VERSION=`git describe --tags --dirty --always`

build:
	go build -ldflags "-X 'github.com/conduitio-labs/conduit-connector-algolia.version=${VERSION}'" -o conduit-connector-algolia cmd/connector/main.go

test:
	go test $(GOTEST_FLAGS) -race ./...

test-integration:
	# run required docker containers, execute integration tests, stop containers after tests
	docker compose -f test/docker-compose.yml up -d
	go test $(GOTEST_FLAGS) -v -race ./...; ret=$$?; \
		docker compose -f test/docker-compose.yml down; \
		exit $$ret

lint:
	golangci-lint run -v
.PHONY: install-tools

install-tools:
	@echo Installing tools from tools.go
	@go list -e -f '{{ join .Imports "\n" }}' tools.go | xargs -tI % go install %
	@go mod tidy
