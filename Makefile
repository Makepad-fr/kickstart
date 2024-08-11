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

test: clean build e2e
	@echo "Running tests..."

	go test ./...

.PHONY: e2e
e2e: create-test-project skaffold-test-project run-helm-lint-on-created-project test-application

.PHONY: create-test-project
create-test-project:
	@echo "Testing the skaffold file format by creating a project"
	@echo "Creating a new project"
	./out/kickstart init-project test-project
	
.PHONY: run-helm-lint-on-created-project
run-helm-lint-on-created-project:
	@echo "Adding a test chart to the project"
	cd ./test-project && ../out/kickstart add-chart server
	@echo "Running helm lint on the created chart"
	helm lint ./test-project/charts/server

.PHONY: skaffold-test-project
skaffold-test-project:
	@echo "Running skaffold diagnose"
	cd ./test-project && skaffold diagnose


.PHONY: test-add-application
test-application:
	@echo "Adding test application to the project"
	cd ./test-project && ../out/kickstart add-app server
	@echo "Building the project with skaffold"
	cd ./test-project && skaffold build

.PHONY: insall
install: build
	@echo "Installing the binary..."
	sudo install -m 0755 $(BINARY_PATH) $(INSTALL_PATH)
	@echo "Kickstart installed to $(INSTALL_PATH)"

