# Makefile for Kickstart CLI tool

BINARY_PATH=./out/kickstart

.PHONY: 
create-out-directory:
	mkdir -p ./out

all: create-out-directory build

build:
	@echo "Building the binary..."
	go build -o $(BINARY_PATH) main.go

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_PATH)

run:
	@echo "Running the CLI..."
	./$(BINARY_PATH)

test:
	@echo "Running tests..."
	go test ./...

.PHONY: build clean run test create-out-directory
