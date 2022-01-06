APP_NAME=go-server

BIN_DIR=bin

build:
	go build -o $(BIN_DIR)/$(APP_NAME) .
.PHONY: build
