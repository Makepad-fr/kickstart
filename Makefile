# Makefile for Kickstart CLI tool

BINARY_PATH=./out/kickstart
INSTALL_PATH=/usr/local/bin/kickstart

.PHONY: all build clean run test create-out-directory install

all: create-out-directory build

create-out-directory:
	mkdir -p ./out

build: create-out-directory
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

install: build
	@echo "Installing the binary..."
	sudo install -m 0755 $(BINARY_PATH) $(INSTALL_PATH)
	@echo "Kickstart installed to $(INSTALL_PATH)"

