.PHONY: build test

build:
	go build -o conduit-connector-algolia cmd/algolia/main.go

test:
	go test $(GOTEST_FLAGS) -race ./...

