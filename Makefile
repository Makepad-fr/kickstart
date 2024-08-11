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

test: clean build skaffold-test-project
	@echo "Running tests..."

	go test ./...

.PHONY: skaffold-test-project
skaffold-test-project:
	@echo "Testing the skaffold file format by creating a project"
	@echo "Creating a new project"
	./out/kickstart init-project test-project
	@echo "Navigating to the created project"
	cd ./test-project && skaffold diagnose
	@echo "Running skaffold diagnose"

install: build
	@echo "Installing the binary..."
	sudo install -m 0755 $(BINARY_PATH) $(INSTALL_PATH)
	@echo "Kickstart installed to $(INSTALL_PATH)"

