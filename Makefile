.DEFAULT_GOAL := build
.PHONY: dev-init fmt vet test testv lint clean build

APP_NAME := ratiogoblin
BUILD_DIR := bin
MAIN_FILE := ./cmd/$(APP_NAME)

dev-init:
	@go run $(MAIN_FILE) init

dev-run:
	@go run $(MAIN_FILE) run

fmt: # Format the code
	go fmt ./...

vet: # Vet the code
	go vet ./... 

test: 
	@go test ./...

testv: 
	@go test -v -cover ./...

lint: # Lint the code
	@golangci-lint run --timeout 5m

clean: # Clean the code
	@rm -rf ./$(BUILD_DIR)

build: clean # Build the code 
	@go build -ldflags "-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
