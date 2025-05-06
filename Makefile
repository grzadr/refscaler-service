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
SERVICE_IMAGE=refscaler-service:$(VERSION)
NAMESPACE=refscaler
FRONTEND_BIN=frontend
FRONTEND_IMAGE=refscaler-frontend:$(VERSION)

all: setup test build-images

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

setup:
	go mod download

build-frontend: $(BIN_DIR) setup
	$(GOBUILD) -o $(BIN_DIR)/$(FRONTEND_BIN) \
	-v \
	-ldflags "-X main.Version=$(VERSION)" \
	$(CMD_DIR)/$(FRONTEND_BIN)/main.go

run-frontend: build-frontend
	./$(BIN_DIR)/$(FRONTEND_BIN)

swag-service:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/service/main.go -o cmd/service/docs

build-service: $(BIN_DIR) setup swag-service
	$(GOBUILD) -o $(BIN_DIR)/$(SERVICE_BIN) \
	-v \
	-ldflags "-X main.Version=$(VERSION)" \
	$(CMD_DIR)/$(SERVICE_BIN)/main.go

build-image-service: lint
	docker build -t $(SERVICE_IMAGE) --file Service.Dockerfile .

build-image-frontend: lint
	docker build -t $(FRONTEND_IMAGE) --file Frontend.Dockerfile .

build-images: build-image-service

kind-upload: build-images
	kind load docker-image $(SERVICE_IMAGE)
	kind load docker-image $(FRONTEND_IMAGE)

fmt:
	golangci-lint fmt $(SRC_DIR)

lint: fmt
	golangci-lint run

test:
	$(GOTEST) $(SRC_DIR)

test-cover: lint test
	mkdir -p $(COVER_DIR)
	$(GOTEST) -coverprofile=$(COVER_DIR)/coverage.out $(SRC_DIR)
	go tool cover -html=$(COVER_DIR)/coverage.out

clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

run-service: build-service
	./$(BIN_DIR)/$(SERVICE_BIN)

version:
	@echo $(VERSION)

kind-install: kind-upload
	kubectl get ns $(NAMESPACE) &> /dev/null || kubectl create ns $(NAMESPACE)
	kubectl label ns $(NAMESPACE) momate-gateway-access=true
	helm upgrade -i -n $(NAMESPACE) refscaler ./refscaler

.PHONY: \
	all \
	build-frontend \
	build-image-service \
	build-images \
	build-service \
	clean \
	fmt \
	kind-upload \
	lint \
	run-frontend \
	run-service \
	setup \
	swa-service \
	test \
	test-cover \
	version
