build:
	go build ./...

lint:
	golangci-lint run ./...

.PHONY: build lint
