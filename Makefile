BUILD_DATE := $(shell date +"%Y-%m-%d")
BUILD_VER := "0.0.1"

default: help

.PHONY: help
help:
	@echo "devapp makefile usage: make [target]"
	@echo "  Targets:"
	@echo "  » clean           Remove build artifacts and clean up the project"
	@echo "  » bin             Build the binary and output to _bin/ directory"
	@echo "  » test            Run all unit tests and generate coverage report"
	@echo "  » run             Run the main.go file to start the server"
	@echo "  » image-dev       Build the docker development image to testing, i.e., non-production"
	@echo "  » image-rel       Build the docker image that is a production release candidate"

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) > /dev/null 2>&1

.PHONY: bin
bin: clean
	go build \
	-ldflags "-X 'token_go_module/internal/conf.buildDate=$(BUILD_DATE)' -s -w" -o ./bin/token_go_bin cmd/main.go

.PHONY: test
test:
	go test ./internal/conf ./internal/server -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: run
run: bin
	./bin/token_go_bin

.PHONY:image-dev
image-dev:
	BUILD_VER="$(BUILD_VER)" ./build/build-develop.sh

.PHONY:image-rel
image-rel:
	BUILD_VER="$(BUILD_VER)" ./build/build-release.sh
