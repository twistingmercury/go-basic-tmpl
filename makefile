BIN_DIR := "./_bin/"
BIN_NAME := "{{bin_name}}"
BUILD_DATE := $(shell date +"%Y-%m-%d")
BUILD_VER := "n/a"
GIT_COMMIT := "n/a"
MODULE_NAME :=  "{{module_name}}"
ALPINE_VERSION := "3.19"
GO_VERSION := "1.21.9"
DESCRIPTION:= "{{description}}"
VENDOR:= "{{vendor_name}}"
TARGET:= "main.go"


ifeq ($(shell git rev-parse --is-inside-work-tree 2>/dev/null),true)
  TAG := $(shell git describe --tags --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2>/dev/null)
  ifdef TAG
    BUILD_VER := $(shell echo $(TAG) | sed 's/v//')
  else
    BUILD_VER := "0.0.0"
  endif
  GIT_COMMIT := $(shell git rev-parse --short HEAD)pwd
endif

default: help

.PHONY: help
help:
	@echo "\{{bin_name}} makefile usage: make [target]"
	@echo "  Targets:"
	@echo "  » clean           Remove build artifacts and clean up the project"
	@echo "  » bin             Build the binary and output to _bin/ directory"
	@echo "  » test            Run all unit tests and generate coverage report"
	@echo "  » image           Build the docker image with using a multi-stage build"
	@echo "  » run             Run the main.go file to start the server"
	@echo "  » test-heartbeat  Test the heartbeat endpoint using cURL\n"

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) > /dev/null 2>&1

.PHONY: bin
bin: clean
	go build \
	-ldflags "-X '$(MODULE_NAME)/cmd/conf.buildDate=$(BUILD_DATE)' \
	-X '$(MODULE_NAME)/cmd/conf.buildVer=$(BUILD_VER)' \
	-X '$(MODULE_NAME)/cmd/conf.buildCommit=$(GIT_COMMIT)' -s -w" \
	-o $(BIN_DIR)$(BIN_NAME) $(TARGET)

.PHONY: test
test:
	go test -v ./conf ./server -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: run
run:
	go run main.go

.PHONY: test-heartbeat
test-heartbeat:
	curl -X GET -H "Content-Type: application/json" http://localhost:8081/heartbeat

.PHONY: image
image:
	BUILD_DATE="$(BUILD_DATE)" \
	BUILD_VER="$(BUILD_VER)" \
	GIT_COMMIT="$(GIT_COMMIT)" \
	DOCKERFILE_DIR="$(PWD)" \
	ALPINE_VERSION="$(ALPINE_VERSION)" \
	GO_VERSION="$(GO_VERSION)" \
	TARGET="$(TARGET)" \
	./_build/build.sh
