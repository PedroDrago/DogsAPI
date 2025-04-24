GOPATH=$(shell go env GOPATH)
run: build
	@./bin/api

build:
	@go build -o bin/api ./cmd/api/main.go

compose_build:
	@docker compose build

compose: compose_build
	@docker compose up

test:
	@go test ./...

lint:
	$(GOPATH)/bin/golangci-lint run
