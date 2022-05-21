TARGET    := static
VERSION   := v0.1.0

SRC_DIR   := .
BUILD_DIR := bin
EXE       := $(BUILD_DIR)/$(TARGET)

all: clean build

.PHONY: build

build:
	go build -o $(EXE) $(SRC_DIR)

.PHONY: build/debug

build/debug:
	go build -tags debug -o $(EXE) $(SRC_DIR)

.PHONY: clean

clean:
	$(RM) -rf $(BUILD_DIR)
