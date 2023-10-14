.PHONY: run, build, run-build, test

run-server:
	@echo "running the program..."
	@go run main.go

build:
	@echo "building the go binary..."
	@go build -o bin/api

run: build
	@echo "running the built binary..."
	@./bin/api

test:
	@echo "running all the tests..."
	@go test -v ./...