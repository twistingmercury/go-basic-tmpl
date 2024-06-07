BIN_DIR := "./bin/"
OUT := "BIN_NAME"
BUILD_DATE := $(shell date +"%Y-%m-%d")
BUILD_VER := "0.0.1"
TARGET:= "main.go"

default: help

.PHONY: help
help:
	@echo "\devapp makefile usage: make [target]"
	@echo "  Targets:"
	@echo "  » clean           Remove build artifacts and clean up the project"
	@echo "  » bin             Build the binary and output to _bin/ directory"
	@echo "  » test            Run all unit tests and generate coverage report"
	@echo "  » image           Build the docker image with using a multi-stage build"
	@echo "  » run             Run the main.go file to start the server"

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) > /dev/null 2>&1

.PHONY: bin
bin: clean
	go build \
	-ldflags "-X 'MODULE_NAME/conf.buildDate=$(BUILD_DATE)' -X 'MODULE_NAME/conf.buildVer=$(BUILD_VER)' -s -w" \
	-o $(BIN_DIR)$(OUT) $(TARGET)

.PHONY: test
test:
	go test -v ./conf ./server -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: run
run: bin
	NAMESPACE=test ./$(BIN_DIR)$(OUT)

.PHONY: image
image:
	BUILD_DATE="$(BUILD_DATE)" \
	BUILD_VER="$(BUILD_VER)" \
	DOCKERFILE_DIR="$(PWD)" \
	./_build/build.sh
