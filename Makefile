.PHONY: run, build, run-build, test, start-basic-containers, stop-basic-containers

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
	@echo "building and executing the binary..."
	@./bin/api

test:
	@echo "running all the tests..."
	@go test -v ./...


