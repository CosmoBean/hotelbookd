.PHONY: run, build, run-build, test, start-basic-containers, stop-basic-containers, clear-cache

start-basic-containers:
	@echo "booting up the basic containers..."
	@docker-compose -f docker-compose.yml up mongo -d

stop-basic-containers:
	@echo "shutting down the basic containers..."
	@docker-compose -f docker-compose.yml down -v --remove-orphans

run-server:
	@echo "running the program..."
	@go run main.go

build:
	@echo "building the go binary..."
	@go build -o bin/api

run: build
	@echo "executing the binary..."
	@./bin/api

test:
	@echo "running all the tests..."
	@go test -v ./... -count 1

clear-cache:
	@echo "clearing the cached test results.."
	@go clean -testcache

