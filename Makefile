TARGET    := go-server # TODO: replace this with your target/app name
VERSION   := v0.1.0    # TODO: replace this with your target/app version
BUILD     := $(shell date -u +%Y-%m-%d.%H:%M)

CLI_PKG   := github.com/akaahmedkamal/go-cli

SRC_DIR   := .
BUILD_DIR := build
EXE       := $(BUILD_DIR)/$(TARGET)

GO        ?= go
LDFLAGS   += -X $(CLI_PKG)/v1.AppName=$(TARGET)
LDFLAGS   += -X $(CLI_PKG)/v1.AppVersion=$(VERSION)
LDFLAGS   += -X $(CLI_PKG)/v1.AppBuild=$(BUILD)

all: clean build

.PHONY: build

build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(EXE) $(SRC_DIR)

.PHONY: build/debug

build/debug:
	$(GO) build -tags debug -ldflags "$(LDFLAGS)" -o $(EXE) $(SRC_DIR)

.PHONY: clean

clean:
	$(RM) -rf $(BUILD_DIR)
