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

all: test build

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

build: $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/$(SERVICE_BIN) -v $(CMD_DIR)/$(SERVICE_BIN)/main.go

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

run: build
	./$(BIN_DIR)/$(BINARY_NAME)

.PHONY: all build lint test clean run test-cover fmt