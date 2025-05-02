# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
SERVICE_BIN=service
BIN_DIR=bin
CMD_DIR=cmd
COVER_DIR=coverage
SRC_DIR=./...
VERSION=$(shell cat VERSION)

all: test build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

swag:
	swag init -g cmd/service/main.go -o cmd/service/docs

build-service: $(BIN_DIR) swag
	$(GOBUILD) -o $(BIN_DIR)/$(SERVICE_BIN) -v -ldflags "-X main.Version=$(VERSION)" $(CMD_DIR)/$(SERVICE_BIN)/main.go

fmt:
	golangci-lint fmt $(SRC_DIR)

lint: fmt
	golangci-lint run

test: lint
	$(GOTEST) $(SRC_DIR)

test-cover: test
	mkdir -p $(COVER_DIR)
	$(GOTEST) -coverprofile=$(COVER_DIR)/coverage.out $(SRC_DIR)
	go tool cover -html=$(COVER_DIR)/coverage.out

clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

run-service: build
	./$(BIN_DIR)/$(SERVICE_BIN)

# Add a target to display the version
version:
	@echo $(VERSION)

.PHONY: all build-service lint test clean run-service test-cover fmt version swag
